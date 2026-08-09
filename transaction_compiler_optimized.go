package solana

import (
	"bytes"
	"fmt"
	"slices"

	"go.uber.org/zap"
)

const (
	compiledAccountSigner uint8 = 1 << iota
	compiledAccountWritable
	compiledAccountInvoked
)

type compiledAccountInfo struct {
	index    uint16
	flags    uint8
	sortRank uint8
}

type compiledAccount struct {
	PublicKey PublicKey
	flags     uint8
	sortRank  uint8
}

type compactAddressTableIndex struct {
	table uint32
	index uint8
}

type compiledAddressTableLookup struct {
	writableCount int
	readonlyCount int
	writable      []uint8
	readonly      []uint8
}

func accountCompileFlags(account *AccountMeta) (flags, sortRank uint8) {
	if account.IsSigner {
		flags |= compiledAccountSigner
		sortRank |= 2
	}
	if account.IsWritable {
		flags |= compiledAccountWritable
		sortRank |= 1
	}
	return flags, sortRank
}

// newTransactionOptimized compiles instructions into a transaction message
// while keeping only one account-index map and one cached Accounts result per
// instruction.
func newTransactionOptimized(instructions []Instruction, recentBlockHash Hash, opts ...TransactionOption) (*Transaction, error) {
	if len(instructions) == 0 {
		return nil, fmt.Errorf("requires at-least one instruction to create a transaction")
	}

	options := transactionOptions{}
	for _, opt := range opts {
		opt.apply(&options)
	}

	// Accounts() is allowed to materialize a slice. Cache it so generated
	// instructions do that work once rather than once per compiler pass.
	instructionAccounts := make([][]*AccountMeta, len(instructions))
	totalAccounts := 0
	for i, instruction := range instructions {
		accounts := instruction.Accounts()
		instructionAccounts[i] = accounts
		totalAccounts += len(accounts)
	}

	feePayer := options.payer
	if feePayer.IsZero() {
		found := false
		for _, account := range instructionAccounts[0] {
			if account.IsSigner {
				feePayer = account.PublicKey
				found = true
				break
			}
		}
		if !found {
			return nil, fmt.Errorf("cannot determine fee payer. You can either pass the fee payer via the 'TransactionWithInstructions' option parameter or it falls back to the first instruction's first signer")
		}
	}

	accountCapacity := totalAccounts + len(instructions)
	var accountInfo map[PublicKey]compiledAccountInfo
	if accountCapacity > 16 {
		accountInfo = make(map[PublicKey]compiledAccountInfo, accountCapacity)
	} else {
		accountInfo = map[PublicKey]compiledAccountInfo{}
	}
	for i, instruction := range instructions {
		for _, account := range instructionAccounts[i] {
			flags, rank := accountCompileFlags(account)
			info := accountInfo[account.PublicKey]
			info.flags |= flags
			// Preserve the reference compiler's sort-before-merge behavior.
			// A signer-readonly occurrence combined with a nonsigner-writable
			// occurrence becomes signer-writable, but remains in signer-readonly
			// sort order because that was the strongest original occurrence.
			if rank > info.sortRank {
				info.sortRank = rank
			}
			accountInfo[account.PublicKey] = info
		}

		programID := instruction.ProgramID()
		info := accountInfo[programID]
		info.flags |= compiledAccountInvoked
		accountInfo[programID] = info
	}

	// A payer supplied through an option need not occur in any instruction.
	// Add it before materializing the sortable account slice.
	if _, ok := accountInfo[feePayer]; !ok {
		accountInfo[feePayer] = compiledAccountInfo{}
	}

	accounts := make([]compiledAccount, 0, len(accountInfo))
	for publicKey, info := range accountInfo {
		accounts = append(accounts, compiledAccount{
			PublicKey: publicKey,
			flags:     info.flags,
			sortRank:  info.sortRank,
		})
	}
	slices.SortFunc(accounts, func(a, b compiledAccount) int {
		if a.sortRank != b.sortRank {
			return int(b.sortRank) - int(a.sortRank)
		}
		return bytes.Compare(a.PublicKey[:], b.PublicKey[:])
	})

	feePayerIndex := -1
	for i := range accounts {
		if accounts[i].PublicKey == feePayer {
			feePayerIndex = i
			break
		}
	}
	if feePayerIndex < 0 {
		panic("solana: transaction compiler lost fee payer")
	}
	feePayerAccount := accounts[feePayerIndex]
	copy(accounts[1:feePayerIndex+1], accounts[:feePayerIndex])
	feePayerAccount.flags |= compiledAccountSigner | compiledAccountWritable
	accounts[0] = feePayerAccount

	totalTableEntries := 0
	for _, table := range options.addressTables {
		totalTableEntries += len(table)
	}
	sortedTableKeys := make(PublicKeySlice, 0, len(options.addressTables))
	for key := range options.addressTables {
		sortedTableKeys = append(sortedTableKeys, key)
	}
	slices.SortFunc(sortedTableKeys, func(a, b PublicKey) int {
		return bytes.Compare(a[:], b[:])
	})

	var addressLookupKeys map[PublicKey]compactAddressTableIndex
	var lookupBuilders []compiledAddressTableLookup
	if totalTableEntries > 0 {
		addressLookupKeys = make(map[PublicKey]compactAddressTableIndex, totalTableEntries)
		lookupBuilders = make([]compiledAddressTableLookup, len(sortedTableKeys))
	}
	for tableIndex, tableKey := range sortedTableKeys {
		table := options.addressTables[tableKey]
		if len(table) > 256 {
			return nil, fmt.Errorf("max lookup table index exceeded for %s table", tableKey)
		}
		for addressIndex, address := range table {
			if _, exists := addressLookupKeys[address]; exists {
				continue
			}
			addressLookupKeys[address] = compactAddressTableIndex{
				table: uint32(tableIndex),
				index: uint8(addressIndex),
			}
		}
	}

	staticAccountCount := 0
	for i := range accounts {
		account := &accounts[i]
		lookup, inTable := addressLookupKeys[account.PublicKey]
		useLookup := inTable && i != 0 && account.flags&(compiledAccountSigner|compiledAccountInvoked) == 0
		if !useLookup {
			staticAccountCount++
			continue
		}
		builder := &lookupBuilders[lookup.table]
		if account.flags&compiledAccountWritable != 0 {
			builder.writableCount++
		} else {
			builder.readonlyCount++
		}
	}
	for i := range lookupBuilders {
		builder := &lookupBuilders[i]
		if builder.writableCount > 0 {
			builder.writable = make([]uint8, 0, builder.writableCount)
		}
		if builder.readonlyCount > 0 {
			builder.readonly = make([]uint8, 0, builder.readonlyCount)
		}
	}

	message := Message{
		RecentBlockhash: recentBlockHash,
		AccountKeys:     make(PublicKeySlice, 0, staticAccountCount),
	}
	for i := range accounts {
		account := &accounts[i]
		lookup, inTable := addressLookupKeys[account.PublicKey]
		useLookup := inTable && i != 0 && account.flags&(compiledAccountSigner|compiledAccountInvoked) == 0
		if useLookup {
			builder := &lookupBuilders[lookup.table]
			if account.flags&compiledAccountWritable != 0 {
				builder.writable = append(builder.writable, lookup.index)
			} else {
				builder.readonly = append(builder.readonly, lookup.index)
			}
			continue
		}

		info := accountInfo[account.PublicKey]
		info.index = uint16(len(message.AccountKeys))
		accountInfo[account.PublicKey] = info
		message.AccountKeys = append(message.AccountKeys, account.PublicKey)

		if account.flags&compiledAccountSigner != 0 {
			message.Header.NumRequiredSignatures++
			if account.flags&compiledAccountWritable == 0 {
				message.Header.NumReadonlySignedAccounts++
			}
		} else if account.flags&compiledAccountWritable == 0 {
			message.Header.NumReadonlyUnsignedAccounts++
		}
	}

	lookupCount := 0
	for i := range lookupBuilders {
		if lookupBuilders[i].writableCount+lookupBuilders[i].readonlyCount > 0 {
			lookupCount++
		}
	}
	if lookupCount > 0 {
		lookups := make([]MessageAddressTableLookup, 0, lookupCount)
		for i, builder := range lookupBuilders {
			if builder.writableCount+builder.readonlyCount == 0 {
				continue
			}
			lookups = append(lookups, MessageAddressTableLookup{
				AccountKey:      sortedTableKeys[i],
				WritableIndexes: builder.writable,
				ReadonlyIndexes: builder.readonly,
			})
		}

		nextIndex := uint16(len(message.AccountKeys))
		for tableIndex, builder := range lookupBuilders {
			table := options.addressTables[sortedTableKeys[tableIndex]]
			for _, addressIndex := range builder.writable {
				publicKey := table[addressIndex]
				info := accountInfo[publicKey]
				info.index = nextIndex
				accountInfo[publicKey] = info
				nextIndex++
			}
		}
		for tableIndex, builder := range lookupBuilders {
			table := options.addressTables[sortedTableKeys[tableIndex]]
			for _, addressIndex := range builder.readonly {
				publicKey := table[addressIndex]
				info := accountInfo[publicKey]
				info.index = nextIndex
				accountInfo[publicKey] = info
				nextIndex++
			}
		}

		if err := message.SetAddressTables(options.addressTables); err != nil {
			return nil, fmt.Errorf("SetAddressTables: %w", err)
		}
		message.SetAddressTableLookups(lookups)
	}

	if debugNewTransaction {
		zlog.Debug("message header compiled",
			zap.Uint8("num_required_signatures", message.Header.NumRequiredSignatures),
			zap.Uint8("num_readonly_signed_accounts", message.Header.NumReadonlySignedAccounts),
			zap.Uint8("num_readonly_unsigned_accounts", message.Header.NumReadonlyUnsignedAccounts),
		)
	}

	message.Instructions = make([]CompiledInstruction, len(instructions))
	for instructionIndex, instruction := range instructions {
		instructionAccountMetas := instructionAccounts[instructionIndex]
		accountIndexes := make([]uint16, len(instructionAccountMetas))
		for accountIndex, account := range instructionAccountMetas {
			accountIndexes[accountIndex] = accountInfo[account.PublicKey].index
		}
		data, err := instruction.Data()
		if err != nil {
			return nil, fmt.Errorf("unable to encode instructions [%d]: %w", instructionIndex, err)
		}
		message.Instructions[instructionIndex] = CompiledInstruction{
			ProgramIDIndex: accountInfo[instruction.ProgramID()].index,
			Accounts:       accountIndexes,
			Data:           data,
		}
	}

	return &Transaction{Message: message}, nil
}
