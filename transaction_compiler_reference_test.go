package solana

import (
	"bytes"
	"fmt"
	"slices"

	"go.uber.org/zap"
)

type addressTablePubkeyWithIndex struct {
	addressTable PublicKey
	index        uint8
}

// newTransactionReference is the pre-optimization compiler kept test-only
// as a differential oracle for exact message and wire parity.
func newTransactionReference(instructions []Instruction, recentBlockHash Hash, opts ...TransactionOption) (*Transaction, error) {
	if len(instructions) == 0 {
		return nil, fmt.Errorf("requires at-least one instruction to create a transaction")
	}

	options := transactionOptions{}
	for _, opt := range opts {
		opt.apply(&options)
	}

	feePayer := options.payer
	if feePayer.IsZero() {
		found := false
		for _, act := range instructions[0].Accounts() {
			if act.IsSigner {
				feePayer = act.PublicKey
				found = true
				break
			}
		}
		if !found {
			return nil, fmt.Errorf("cannot determine fee payer. You can either pass the fee payer via the 'TransactionWithInstructions' option parameter or it falls back to the first instruction's first signer")
		}
	}

	totalTableEntries := 0
	for _, t := range options.addressTables {
		totalTableEntries += len(t)
	}
	addressLookupKeysMap := make(map[PublicKey]addressTablePubkeyWithIndex, totalTableEntries) // all accounts from tables as map
	sortedTableKeys := make(PublicKeySlice, 0, len(options.addressTables))
	for k := range options.addressTables {
		sortedTableKeys = append(sortedTableKeys, k)
	}
	slices.SortFunc(sortedTableKeys, func(a, b PublicKey) int {
		return bytes.Compare(a[:], b[:])
	})
	for _, addressTablePubKey := range sortedTableKeys {
		addressTable := options.addressTables[addressTablePubKey]
		if len(addressTable) > 256 {
			return nil, fmt.Errorf("max lookup table index exceeded for %s table", addressTablePubKey)
		}

		for i, address := range addressTable {
			_, ok := addressLookupKeysMap[address]
			if ok {
				continue
			}

			addressLookupKeysMap[address] = addressTablePubkeyWithIndex{
				addressTable: addressTablePubKey,
				index:        uint8(i),
			}
		}
	}

	totalAccounts := 0
	for _, instruction := range instructions {
		totalAccounts += len(instruction.Accounts())
	}
	programIDs := make(PublicKeySlice, 0, len(instructions))
	accounts := make([]*AccountMeta, 0, totalAccounts+len(instructions))
	for _, instruction := range instructions {
		accounts = append(accounts, instruction.Accounts()...)
		programIDs.UniqueAppend(instruction.ProgramID())
	}

	// for IsInvoked check
	programIDsMap := make(map[PublicKey]struct{}, len(programIDs))
	// Add programID to the account list
	for _, programID := range programIDs {
		accounts = append(accounts, &AccountMeta{
			PublicKey:  programID,
			IsSigner:   false,
			IsWritable: false,
		})

		programIDsMap[programID] = struct{}{}
	}

	// Sort. Prioritizing first by signer, then by writable
	slices.SortStableFunc(accounts, func(a, b *AccountMeta) int {
		if a.less(b) {
			return -1
		}
		if b.less(a) {
			return 1
		}
		return 0
	})

	// Hint the map only above a threshold: for small txs, an empty map is
	// cheaper than a single pre-allocated bucket (~640B for PublicKey keys).
	// For larger txs, the hint avoids several bucket-grow operations.
	var uniqAccountsMap map[PublicKey]uint64
	if len(accounts) > 16 {
		uniqAccountsMap = make(map[PublicKey]uint64, len(accounts))
	} else {
		uniqAccountsMap = map[PublicKey]uint64{}
	}
	uniqAccounts := make([]*AccountMeta, 0, len(accounts))
	for _, acc := range accounts {
		if index, found := uniqAccountsMap[acc.PublicKey]; found {
			uniqAccounts[index].IsWritable = uniqAccounts[index].IsWritable || acc.IsWritable
			continue
		}
		uniqAccounts = append(uniqAccounts, acc)
		uniqAccountsMap[acc.PublicKey] = uint64(len(uniqAccounts) - 1)
	}

	if debugNewTransaction {
		zlog.Debug("unique account sorted", zap.Int("account_count", len(uniqAccounts)))
	}
	// Move fee payer to the front
	feePayerIndex := -1
	for idx, acc := range uniqAccounts {
		if acc.PublicKey.Equals(feePayer) {
			feePayerIndex = idx
		}
	}
	if debugNewTransaction {
		zlog.Debug("current fee payer index", zap.Int("fee_payer_index", feePayerIndex))
	}

	accountCount := len(uniqAccounts)
	if feePayerIndex < 0 {
		// fee payer is not part of accounts we want to add it
		accountCount++
	}
	allKeys := make([]*AccountMeta, accountCount)

	itr := 1
	for idx, uniqAccount := range uniqAccounts {
		if idx == feePayerIndex {
			uniqAccount.IsSigner = true
			uniqAccount.IsWritable = true
			allKeys[0] = uniqAccount
			continue
		}
		allKeys[itr] = uniqAccount
		itr++
	}

	if feePayerIndex < 0 {
		// fee payer is not part of accounts we want to add it
		feePayerAccount := &AccountMeta{
			PublicKey:  feePayer,
			IsSigner:   true,
			IsWritable: true,
		}
		allKeys[0] = feePayerAccount
	}

	message := Message{
		RecentBlockhash: recentBlockHash,
	}
	lookupsMap := make(map[PublicKey]struct { // extended MessageAddressTableLookup
		AccountKey      PublicKey // The account key of the address table.
		WritableIndexes []uint8
		Writable        []PublicKey
		ReadonlyIndexes []uint8
		Readonly        []PublicKey
	})
	for idx, acc := range allKeys {

		if debugNewTransaction {
			zlog.Debug("transaction account",
				zap.Int("account_index", idx),
				zap.Stringer("account_pub_key", acc.PublicKey),
			)
		}

		addressLookupKeyEntry, isPresentedInTables := addressLookupKeysMap[acc.PublicKey]
		_, isInvoked := programIDsMap[acc.PublicKey]
		// skip fee payer
		if isPresentedInTables && idx != 0 && !acc.IsSigner && !isInvoked {
			lookup := lookupsMap[addressLookupKeyEntry.addressTable]
			if acc.IsWritable {
				lookup.WritableIndexes = append(lookup.WritableIndexes, addressLookupKeyEntry.index)
				lookup.Writable = append(lookup.Writable, acc.PublicKey)
			} else {
				lookup.ReadonlyIndexes = append(lookup.ReadonlyIndexes, addressLookupKeyEntry.index)
				lookup.Readonly = append(lookup.Readonly, acc.PublicKey)
			}

			lookupsMap[addressLookupKeyEntry.addressTable] = lookup
			continue // prevent changing message.Header properties
		}

		message.AccountKeys = append(message.AccountKeys, acc.PublicKey)

		if acc.IsSigner {
			message.Header.NumRequiredSignatures++
			if !acc.IsWritable {
				message.Header.NumReadonlySignedAccounts++
			}
			continue
		}

		if !acc.IsWritable {
			message.Header.NumReadonlyUnsignedAccounts++
		}
	}

	var lookupsWritableKeys []PublicKey
	var lookupsReadOnlyKeys []PublicKey
	if len(lookupsMap) > 0 {
		lookups := make([]MessageAddressTableLookup, 0, len(lookupsMap))

		sortedLookupKeys := make(PublicKeySlice, 0, len(lookupsMap))
		var totalWritable, totalReadonly int
		for k, l := range lookupsMap {
			sortedLookupKeys = append(sortedLookupKeys, k)
			totalWritable += len(l.Writable)
			totalReadonly += len(l.Readonly)
		}
		lookupsWritableKeys = make([]PublicKey, 0, totalWritable)
		lookupsReadOnlyKeys = make([]PublicKey, 0, totalReadonly)
		slices.SortFunc(sortedLookupKeys, func(a, b PublicKey) int {
			return bytes.Compare(a[:], b[:])
		})
		for _, tablePubKey := range sortedLookupKeys {
			l := lookupsMap[tablePubKey]
			lookupsWritableKeys = append(lookupsWritableKeys, l.Writable...)
			lookupsReadOnlyKeys = append(lookupsReadOnlyKeys, l.Readonly...)

			lookups = append(lookups, MessageAddressTableLookup{
				AccountKey:      tablePubKey,
				WritableIndexes: l.WritableIndexes,
				ReadonlyIndexes: l.ReadonlyIndexes,
			})
		}

		// prevent error created in ResolveLookups
		err := message.SetAddressTables(options.addressTables)
		if err != nil {
			return nil, fmt.Errorf("SetAddressTables: %w", err)
		}
		message.SetAddressTableLookups(lookups)
	}

	var idx uint16
	accountKeyIndex := make(map[PublicKey]uint16, len(message.AccountKeys)+len(lookupsWritableKeys)+len(lookupsReadOnlyKeys))
	for _, acc := range message.AccountKeys {
		accountKeyIndex[acc] = idx
		idx++
	}
	for _, acc := range lookupsWritableKeys {
		accountKeyIndex[acc] = idx
		idx++
	}
	for _, acc := range lookupsReadOnlyKeys {
		accountKeyIndex[acc] = idx
		idx++
	}

	if debugNewTransaction {
		zlog.Debug("message header compiled",
			zap.Uint8("num_required_signatures", message.Header.NumRequiredSignatures),
			zap.Uint8("num_readonly_signed_accounts", message.Header.NumReadonlySignedAccounts),
			zap.Uint8("num_readonly_unsigned_accounts", message.Header.NumReadonlyUnsignedAccounts),
		)
	}

	message.Instructions = make([]CompiledInstruction, 0, len(instructions))
	for txIdx, instruction := range instructions {
		accounts = instruction.Accounts()
		accountIndex := make([]uint16, len(accounts))
		for idx, acc := range accounts {
			accountIndex[idx] = accountKeyIndex[acc.PublicKey]
		}
		data, err := instruction.Data()
		if err != nil {
			return nil, fmt.Errorf("unable to encode instructions [%d]: %w", txIdx, err)
		}
		message.Instructions = append(message.Instructions, CompiledInstruction{
			ProgramIDIndex: accountKeyIndex[instruction.ProgramID()],
			Accounts:       accountIndex,
			Data:           data,
		})
	}

	return &Transaction{
		Message: message,
	}, nil
}
