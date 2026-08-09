package solana

import (
	"bytes"
	"fmt"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/require"
)

type compilerParityCase struct {
	instructions []Instruction
	blockhash    Hash
	payer        PublicKey
	tables       map[PublicKey]PublicKeySlice
}

type countingCompilerInstruction struct {
	Instruction
	accountCalls int
}

func (instruction *countingCompilerInstruction) Accounts() []*AccountMeta {
	instruction.accountCalls++
	return instruction.Instruction.Accounts()
}

func compilerParityPublicKey(rng *rand.Rand) PublicKey {
	var key PublicKey
	_, _ = rng.Read(key[:])
	key[0] |= 1
	return key
}

func buildCompilerParityCase(seed int64) compilerParityCase {
	rng := rand.New(rand.NewSource(seed))
	pool := make(PublicKeySlice, 2+rng.Intn(40))
	for i := range pool {
		pool[i] = compilerParityPublicKey(rng)
	}

	instructionCount := 1 + rng.Intn(10)
	instructions := make([]Instruction, instructionCount)
	for instructionIndex := range instructions {
		accountCount := 1 + rng.Intn(20)
		accounts := make([]*AccountMeta, accountCount)
		for accountIndex := range accounts {
			key := pool[rng.Intn(len(pool))]
			accounts[accountIndex] = &AccountMeta{
				PublicKey:  key,
				IsSigner:   rng.Intn(7) == 0,
				IsWritable: rng.Intn(3) == 0,
			}
		}
		if instructionIndex == 0 {
			accounts[0].IsSigner = true
		}

		programID := compilerParityPublicKey(rng)
		if rng.Intn(5) == 0 {
			programID = pool[rng.Intn(len(pool))]
		}
		data := make([]byte, rng.Intn(24))
		_, _ = rng.Read(data)
		instructions[instructionIndex] = newTestInstruction(programID, accounts, data)
	}

	var payer PublicKey
	if rng.Intn(2) == 0 {
		if rng.Intn(3) == 0 {
			payer = compilerParityPublicKey(rng)
		} else {
			payer = pool[rng.Intn(len(pool))]
		}
	}

	var tables map[PublicKey]PublicKeySlice
	if rng.Intn(4) != 0 {
		tables = make(map[PublicKey]PublicKeySlice)
		for tableIndex := 0; tableIndex < 1+rng.Intn(4); tableIndex++ {
			tableKey := compilerParityPublicKey(rng)
			addresses := make(PublicKeySlice, 1+rng.Intn(30))
			for addressIndex := range addresses {
				if rng.Intn(5) == 0 {
					addresses[addressIndex] = compilerParityPublicKey(rng)
				} else {
					addresses[addressIndex] = pool[rng.Intn(len(pool))]
				}
			}
			tables[tableKey] = addresses
		}
	}

	var blockhash Hash
	_, _ = rng.Read(blockhash[:])
	return compilerParityCase{
		instructions: instructions,
		blockhash:    blockhash,
		payer:        payer,
		tables:       tables,
	}
}

func cloneCompilerParityInstructions(instructions []Instruction) []Instruction {
	cloned := make([]Instruction, len(instructions))
	for i, instruction := range instructions {
		accounts := instruction.Accounts()
		clonedAccounts := make([]*AccountMeta, len(accounts))
		for j, account := range accounts {
			copyOfAccount := *account
			clonedAccounts[j] = &copyOfAccount
		}
		data, err := instruction.Data()
		if err != nil {
			panic(err)
		}
		cloned[i] = newTestInstruction(
			instruction.ProgramID(),
			clonedAccounts,
			bytes.Clone(data),
		)
	}
	return cloned
}

func compilerParityOptions(testCase compilerParityCase) []TransactionOption {
	options := make([]TransactionOption, 0, 2)
	if !testCase.payer.IsZero() {
		options = append(options, TransactionPayer(testCase.payer))
	}
	if testCase.tables != nil {
		options = append(options, TransactionAddressTables(testCase.tables))
	}
	return options
}

func requireCompilerParity(t testing.TB, testCase compilerParityCase) {
	t.Helper()
	options := compilerParityOptions(testCase)
	want, wantErr := newTransactionReference(
		cloneCompilerParityInstructions(testCase.instructions),
		testCase.blockhash,
		options...,
	)
	got, gotErr := newTransactionOptimized(
		cloneCompilerParityInstructions(testCase.instructions),
		testCase.blockhash,
		options...,
	)
	require.Equal(t, wantErr == nil, gotErr == nil)
	if wantErr != nil {
		require.EqualError(t, gotErr, wantErr.Error())
		return
	}

	wantWire, err := want.MarshalBinary()
	require.NoError(t, err)
	gotWire, err := got.MarshalBinary()
	require.NoError(t, err)
	require.Equal(t, wantWire, gotWire)
	require.Equal(t, want.Message.Header, got.Message.Header)
	require.Equal(t, want.Message.AccountKeys, got.Message.AccountKeys)
	require.Equal(t, want.Message.Instructions, got.Message.Instructions)
	require.Equal(t, want.Message.AddressTableLookups, got.Message.AddressTableLookups)
}

func TestNewTransactionOptimizedParity(t *testing.T) {
	for seed := int64(1); seed <= 1_000; seed++ {
		t.Run(fmt.Sprintf("seed_%d", seed), func(t *testing.T) {
			requireCompilerParity(t, buildCompilerParityCase(seed))
		})
	}
}

func TestNewTransactionReadsInstructionAccountsOnce(t *testing.T) {
	payer := compilerParityPublicKey(rand.New(rand.NewSource(10_001)))
	program := compilerParityPublicKey(rand.New(rand.NewSource(10_002)))
	instruction := &countingCompilerInstruction{Instruction: newTestInstruction(
		program,
		[]*AccountMeta{{PublicKey: payer, IsSigner: true, IsWritable: true}},
		[]byte{1, 2, 3},
	)}

	_, err := NewTransaction([]Instruction{instruction}, Hash{})
	require.NoError(t, err)
	require.Equal(t, 1, instruction.accountCalls)
}

func TestNewTransactionDoesNotMutateCallerAccountMetas(t *testing.T) {
	rng := rand.New(rand.NewSource(10_003))
	payer := compilerParityPublicKey(rng)
	other := compilerParityPublicKey(rng)
	program := compilerParityPublicKey(rng)
	payerMeta := &AccountMeta{PublicKey: payer}
	otherMeta := &AccountMeta{PublicKey: other, IsSigner: true}
	instruction := newTestInstruction(program, []*AccountMeta{payerMeta, otherMeta}, nil)

	transaction, err := NewTransaction(
		[]Instruction{instruction},
		Hash{},
		TransactionPayer(payer),
	)
	require.NoError(t, err)
	require.Equal(t, payer, transaction.Message.AccountKeys[0])
	require.False(t, payerMeta.IsSigner)
	require.False(t, payerMeta.IsWritable)
	require.True(t, otherMeta.IsSigner)
	require.False(t, otherMeta.IsWritable)
}

func TestNewTransactionMergesCrossedPrivilegesWithReferenceOrdering(t *testing.T) {
	rng := rand.New(rand.NewSource(10_004))
	payer := compilerParityPublicKey(rng)
	shared := compilerParityPublicKey(rng)
	readonly := compilerParityPublicKey(rng)
	program := compilerParityPublicKey(rng)
	testCase := compilerParityCase{
		payer: payer,
		instructions: []Instruction{newTestInstruction(program, []*AccountMeta{
			{PublicKey: payer, IsSigner: true, IsWritable: true},
			{PublicKey: shared, IsSigner: true},
			{PublicKey: shared, IsWritable: true},
			{PublicKey: readonly},
		}, nil)},
	}

	requireCompilerParity(t, testCase)
}

func FuzzNewTransactionOptimizedParity(f *testing.F) {
	for _, seed := range []int64{1, 2, 3, 7, 31, 127, 1_024, 65_535} {
		f.Add(seed)
	}
	f.Fuzz(func(t *testing.T, seed int64) {
		requireCompilerParity(t, buildCompilerParityCase(seed))
	})
}

func BenchmarkNewTransactionReference(b *testing.B) {
	for _, testCase := range benchTxShapes {
		testCase := testCase
		b.Run(testCase.name, func(b *testing.B) {
			instructions, blockhash := buildBenchInstructions(testCase.numInstructions, testCase.accountsPerIx)
			b.ReportAllocs()
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				transaction, err := newTransactionReference(instructions, blockhash)
				if err != nil {
					b.Fatal(err)
				}
				_ = transaction
			}
		})
	}
}

func BenchmarkNewTransactionReferenceWithLookupTable(b *testing.B) {
	for _, testCase := range benchTxShapes {
		testCase := testCase
		b.Run(testCase.name, func(b *testing.B) {
			instructions, blockhash, tables := buildBenchInstructionsWithLookups(testCase.numInstructions, testCase.accountsPerIx)
			options := TransactionAddressTables(tables)
			b.ReportAllocs()
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				transaction, err := newTransactionReference(instructions, blockhash, options)
				if err != nil {
					b.Fatal(err)
				}
				_ = transaction
			}
		})
	}
}
