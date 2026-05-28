package token2022

import (
	"bytes"
	"strconv"
	"testing"

	ag_gofuzz "github.com/gagliardetto/gofuzz"
	ag_require "github.com/stretchr/testify/require"
)

func TestEncodeDecode_InitializeMint(t *testing.T) {
	fu := ag_gofuzz.New().NilChance(0)
	for i := 0; i < 1; i++ {
		t.Run("InitializeMint"+strconv.Itoa(i), func(t *testing.T) {
			{
				params := new(InitializeMint)
				fu.Fuzz(params)
				params.AccountMetaSlice = nil
				buf := new(bytes.Buffer)
				err := encodeT(*params, buf)
				ag_require.NoError(t, err)
				got := new(InitializeMint)
				err = decodeT(got, buf.Bytes())
				got.AccountMetaSlice = nil
				ag_require.NoError(t, err)
				ag_require.Equal(t, params, got)
			}
		})
	}
}

func TestEncodeDecode_InitializeAccount(t *testing.T) {
	fu := ag_gofuzz.New().NilChance(0)
	for i := 0; i < 1; i++ {
		t.Run("InitializeAccount"+strconv.Itoa(i), func(t *testing.T) {
			{
				params := new(InitializeAccount)
				fu.Fuzz(params)
				params.AccountMetaSlice = nil
				buf := new(bytes.Buffer)
				err := encodeT(*params, buf)
				ag_require.NoError(t, err)
				got := new(InitializeAccount)
				err = decodeT(got, buf.Bytes())
				got.AccountMetaSlice = nil
				ag_require.NoError(t, err)
				ag_require.Equal(t, params, got)
			}
		})
	}
}

func TestEncodeDecode_InitializeMultisig(t *testing.T) {
	fu := ag_gofuzz.New().NilChance(0)
	for i := 0; i < 1; i++ {
		t.Run("InitializeMultisig"+strconv.Itoa(i), func(t *testing.T) {
			{
				params := new(InitializeMultisig)
				fu.Fuzz(params)
				params.Accounts = nil
				params.Signers = nil
				buf := new(bytes.Buffer)
				err := encodeT(*params, buf)
				ag_require.NoError(t, err)
				got := new(InitializeMultisig)
				err = decodeT(got, buf.Bytes())
				got.Accounts = nil
				got.Signers = nil
				ag_require.NoError(t, err)
				ag_require.Equal(t, params, got)
			}
		})
	}
}

func TestEncodeDecode_Transfer(t *testing.T) {
	fu := ag_gofuzz.New().NilChance(0)
	for i := 0; i < 1; i++ {
		t.Run("Transfer"+strconv.Itoa(i), func(t *testing.T) {
			{
				params := new(Transfer)
				fu.Fuzz(params)
				params.Accounts = nil
				params.Signers = nil
				buf := new(bytes.Buffer)
				err := encodeT(*params, buf)
				ag_require.NoError(t, err)
				got := new(Transfer)
				err = decodeT(got, buf.Bytes())
				got.Accounts = nil
				got.Signers = nil
				ag_require.NoError(t, err)
				ag_require.Equal(t, params, got)
			}
		})
	}
}

func TestEncodeDecode_Approve(t *testing.T) {
	fu := ag_gofuzz.New().NilChance(0)
	for i := 0; i < 1; i++ {
		t.Run("Approve"+strconv.Itoa(i), func(t *testing.T) {
			{
				params := new(Approve)
				fu.Fuzz(params)
				params.Accounts = nil
				params.Signers = nil
				buf := new(bytes.Buffer)
				err := encodeT(*params, buf)
				ag_require.NoError(t, err)
				got := new(Approve)
				err = decodeT(got, buf.Bytes())
				got.Accounts = nil
				got.Signers = nil
				ag_require.NoError(t, err)
				ag_require.Equal(t, params, got)
			}
		})
	}
}

func TestEncodeDecode_Revoke(t *testing.T) {
	fu := ag_gofuzz.New().NilChance(0)
	for i := 0; i < 1; i++ {
		t.Run("Revoke"+strconv.Itoa(i), func(t *testing.T) {
			{
				params := new(Revoke)
				fu.Fuzz(params)
				params.Accounts = nil
				params.Signers = nil
				buf := new(bytes.Buffer)
				err := encodeT(*params, buf)
				ag_require.NoError(t, err)
				got := new(Revoke)
				err = decodeT(got, buf.Bytes())
				got.Accounts = nil
				got.Signers = nil
				ag_require.NoError(t, err)
				ag_require.Equal(t, params, got)
			}
		})
	}
}

func TestEncodeDecode_SetAuthority(t *testing.T) {
	fu := ag_gofuzz.New().NilChance(0)
	for i := 0; i < 1; i++ {
		t.Run("SetAuthority"+strconv.Itoa(i), func(t *testing.T) {
			{
				params := new(SetAuthority)
				fu.Fuzz(params)
				params.Accounts = nil
				params.Signers = nil
				buf := new(bytes.Buffer)
				err := encodeT(*params, buf)
				ag_require.NoError(t, err)
				got := new(SetAuthority)
				err = decodeT(got, buf.Bytes())
				got.Accounts = nil
				got.Signers = nil
				ag_require.NoError(t, err)
				ag_require.Equal(t, params, got)
			}
		})
	}
}

func TestEncodeDecode_MintTo(t *testing.T) {
	fu := ag_gofuzz.New().NilChance(0)
	for i := 0; i < 1; i++ {
		t.Run("MintTo"+strconv.Itoa(i), func(t *testing.T) {
			{
				params := new(MintTo)
				fu.Fuzz(params)
				params.Accounts = nil
				params.Signers = nil
				buf := new(bytes.Buffer)
				err := encodeT(*params, buf)
				ag_require.NoError(t, err)
				got := new(MintTo)
				err = decodeT(got, buf.Bytes())
				got.Accounts = nil
				got.Signers = nil
				ag_require.NoError(t, err)
				ag_require.Equal(t, params, got)
			}
		})
	}
}

func TestEncodeDecode_Burn(t *testing.T) {
	fu := ag_gofuzz.New().NilChance(0)
	for i := 0; i < 1; i++ {
		t.Run("Burn"+strconv.Itoa(i), func(t *testing.T) {
			{
				params := new(Burn)
				fu.Fuzz(params)
				params.Accounts = nil
				params.Signers = nil
				buf := new(bytes.Buffer)
				err := encodeT(*params, buf)
				ag_require.NoError(t, err)
				got := new(Burn)
				err = decodeT(got, buf.Bytes())
				got.Accounts = nil
				got.Signers = nil
				ag_require.NoError(t, err)
				ag_require.Equal(t, params, got)
			}
		})
	}
}

func TestEncodeDecode_CloseAccount(t *testing.T) {
	fu := ag_gofuzz.New().NilChance(0)
	for i := 0; i < 1; i++ {
		t.Run("CloseAccount"+strconv.Itoa(i), func(t *testing.T) {
			{
				params := new(CloseAccount)
				fu.Fuzz(params)
				params.Accounts = nil
				params.Signers = nil
				buf := new(bytes.Buffer)
				err := encodeT(*params, buf)
				ag_require.NoError(t, err)
				got := new(CloseAccount)
				err = decodeT(got, buf.Bytes())
				got.Accounts = nil
				got.Signers = nil
				ag_require.NoError(t, err)
				ag_require.Equal(t, params, got)
			}
		})
	}
}

func TestEncodeDecode_FreezeAccount(t *testing.T) {
	fu := ag_gofuzz.New().NilChance(0)
	for i := 0; i < 1; i++ {
		t.Run("FreezeAccount"+strconv.Itoa(i), func(t *testing.T) {
			{
				params := new(FreezeAccount)
				fu.Fuzz(params)
				params.Accounts = nil
				params.Signers = nil
				buf := new(bytes.Buffer)
				err := encodeT(*params, buf)
				ag_require.NoError(t, err)
				got := new(FreezeAccount)
				err = decodeT(got, buf.Bytes())
				got.Accounts = nil
				got.Signers = nil
				ag_require.NoError(t, err)
				ag_require.Equal(t, params, got)
			}
		})
	}
}

func TestEncodeDecode_ThawAccount(t *testing.T) {
	fu := ag_gofuzz.New().NilChance(0)
	for i := 0; i < 1; i++ {
		t.Run("ThawAccount"+strconv.Itoa(i), func(t *testing.T) {
			{
				params := new(ThawAccount)
				fu.Fuzz(params)
				params.Accounts = nil
				params.Signers = nil
				buf := new(bytes.Buffer)
				err := encodeT(*params, buf)
				ag_require.NoError(t, err)
				got := new(ThawAccount)
				err = decodeT(got, buf.Bytes())
				got.Accounts = nil
				got.Signers = nil
				ag_require.NoError(t, err)
				ag_require.Equal(t, params, got)
			}
		})
	}
}

func TestEncodeDecode_TransferChecked(t *testing.T) {
	fu := ag_gofuzz.New().NilChance(0)
	for i := 0; i < 1; i++ {
		t.Run("TransferChecked"+strconv.Itoa(i), func(t *testing.T) {
			{
				params := new(TransferChecked)
				fu.Fuzz(params)
				params.Accounts = nil
				params.Signers = nil
				buf := new(bytes.Buffer)
				err := encodeT(*params, buf)
				ag_require.NoError(t, err)
				got := new(TransferChecked)
				err = decodeT(got, buf.Bytes())
				got.Accounts = nil
				got.Signers = nil
				ag_require.NoError(t, err)
				ag_require.Equal(t, params, got)
			}
		})
	}
}

func TestEncodeDecode_ApproveChecked(t *testing.T) {
	fu := ag_gofuzz.New().NilChance(0)
	for i := 0; i < 1; i++ {
		t.Run("ApproveChecked"+strconv.Itoa(i), func(t *testing.T) {
			{
				params := new(ApproveChecked)
				fu.Fuzz(params)
				params.Accounts = nil
				params.Signers = nil
				buf := new(bytes.Buffer)
				err := encodeT(*params, buf)
				ag_require.NoError(t, err)
				got := new(ApproveChecked)
				err = decodeT(got, buf.Bytes())
				got.Accounts = nil
				got.Signers = nil
				ag_require.NoError(t, err)
				ag_require.Equal(t, params, got)
			}
		})
	}
}

func TestEncodeDecode_MintToChecked(t *testing.T) {
	fu := ag_gofuzz.New().NilChance(0)
	for i := 0; i < 1; i++ {
		t.Run("MintToChecked"+strconv.Itoa(i), func(t *testing.T) {
			{
				params := new(MintToChecked)
				fu.Fuzz(params)
				params.Accounts = nil
				params.Signers = nil
				buf := new(bytes.Buffer)
				err := encodeT(*params, buf)
				ag_require.NoError(t, err)
				got := new(MintToChecked)
				err = decodeT(got, buf.Bytes())
				got.Accounts = nil
				got.Signers = nil
				ag_require.NoError(t, err)
				ag_require.Equal(t, params, got)
			}
		})
	}
}

func TestEncodeDecode_BurnChecked(t *testing.T) {
	fu := ag_gofuzz.New().NilChance(0)
	for i := 0; i < 1; i++ {
		t.Run("BurnChecked"+strconv.Itoa(i), func(t *testing.T) {
			{
				params := new(BurnChecked)
				fu.Fuzz(params)
				params.Accounts = nil
				params.Signers = nil
				buf := new(bytes.Buffer)
				err := encodeT(*params, buf)
				ag_require.NoError(t, err)
				got := new(BurnChecked)
				err = decodeT(got, buf.Bytes())
				got.Accounts = nil
				got.Signers = nil
				ag_require.NoError(t, err)
				ag_require.Equal(t, params, got)
			}
		})
	}
}

func TestEncodeDecode_InitializeAccount2(t *testing.T) {
	fu := ag_gofuzz.New().NilChance(0)
	for i := 0; i < 1; i++ {
		t.Run("InitializeAccount2"+strconv.Itoa(i), func(t *testing.T) {
			{
				params := new(InitializeAccount2)
				fu.Fuzz(params)
				params.AccountMetaSlice = nil
				buf := new(bytes.Buffer)
				err := encodeT(*params, buf)
				ag_require.NoError(t, err)
				got := new(InitializeAccount2)
				err = decodeT(got, buf.Bytes())
				got.AccountMetaSlice = nil
				ag_require.NoError(t, err)
				ag_require.Equal(t, params, got)
			}
		})
	}
}

func TestEncodeDecode_SyncNative(t *testing.T) {
	fu := ag_gofuzz.New().NilChance(0)
	for i := 0; i < 1; i++ {
		t.Run("SyncNative"+strconv.Itoa(i), func(t *testing.T) {
			{
				params := new(SyncNative)
				fu.Fuzz(params)
				params.AccountMetaSlice = nil
				buf := new(bytes.Buffer)
				err := encodeT(*params, buf)
				ag_require.NoError(t, err)
				got := new(SyncNative)
				err = decodeT(got, buf.Bytes())
				got.AccountMetaSlice = nil
				ag_require.NoError(t, err)
				ag_require.Equal(t, params, got)
			}
		})
	}
}

func TestEncodeDecode_InitializeAccount3(t *testing.T) {
	fu := ag_gofuzz.New().NilChance(0)
	for i := 0; i < 1; i++ {
		t.Run("InitializeAccount3"+strconv.Itoa(i), func(t *testing.T) {
			{
				params := new(InitializeAccount3)
				fu.Fuzz(params)
				params.AccountMetaSlice = nil
				buf := new(bytes.Buffer)
				err := encodeT(*params, buf)
				ag_require.NoError(t, err)
				got := new(InitializeAccount3)
				err = decodeT(got, buf.Bytes())
				got.AccountMetaSlice = nil
				ag_require.NoError(t, err)
				ag_require.Equal(t, params, got)
			}
		})
	}
}

func TestEncodeDecode_InitializeMultisig2(t *testing.T) {
	fu := ag_gofuzz.New().NilChance(0)
	for i := 0; i < 1; i++ {
		t.Run("InitializeMultisig2"+strconv.Itoa(i), func(t *testing.T) {
			{
				params := new(InitializeMultisig2)
				fu.Fuzz(params)
				params.Accounts = nil
				params.Signers = nil
				buf := new(bytes.Buffer)
				err := encodeT(*params, buf)
				ag_require.NoError(t, err)
				got := new(InitializeMultisig2)
				err = decodeT(got, buf.Bytes())
				got.Accounts = nil
				got.Signers = nil
				ag_require.NoError(t, err)
				ag_require.Equal(t, params, got)
			}
		})
	}
}

func TestEncodeDecode_InitializeMint2(t *testing.T) {
	fu := ag_gofuzz.New().NilChance(0)
	for i := 0; i < 1; i++ {
		t.Run("InitializeMint2"+strconv.Itoa(i), func(t *testing.T) {
			{
				params := new(InitializeMint2)
				fu.Fuzz(params)
				params.AccountMetaSlice = nil
				buf := new(bytes.Buffer)
				err := encodeT(*params, buf)
				ag_require.NoError(t, err)
				got := new(InitializeMint2)
				err = decodeT(got, buf.Bytes())
				got.AccountMetaSlice = nil
				ag_require.NoError(t, err)
				ag_require.Equal(t, params, got)
			}
		})
	}
}

// === Token-2022 specific instructions ===

func TestEncodeDecode_GetAccountDataSize(t *testing.T) {
	t.Run("GetAccountDataSize", func(t *testing.T) {
		params := &GetAccountDataSize{
			ExtensionTypes: []ExtensionType{
				ExtensionTransferFeeConfig,
				ExtensionMintCloseAuthority,
				ExtensionImmutableOwner,
			},
		}
		buf := new(bytes.Buffer)
		err := encodeT(*params, buf)
		ag_require.NoError(t, err)
		got := new(GetAccountDataSize)
		err = decodeT(got, buf.Bytes())
		got.AccountMetaSlice = nil
		params.AccountMetaSlice = nil
		ag_require.NoError(t, err)
		ag_require.Equal(t, params, got)
	})
}

func TestEncodeDecode_InitializeImmutableOwner(t *testing.T) {
	t.Run("InitializeImmutableOwner", func(t *testing.T) {
		params := new(InitializeImmutableOwner)
		params.AccountMetaSlice = nil
		buf := new(bytes.Buffer)
		err := encodeT(*params, buf)
		ag_require.NoError(t, err)
		got := new(InitializeImmutableOwner)
		err = decodeT(got, buf.Bytes())
		got.AccountMetaSlice = nil
		ag_require.NoError(t, err)
		ag_require.Equal(t, params, got)
	})
}

func TestEncodeDecode_AmountToUiAmount(t *testing.T) {
	fu := ag_gofuzz.New().NilChance(0)
	for i := 0; i < 1; i++ {
		t.Run("AmountToUiAmount"+strconv.Itoa(i), func(t *testing.T) {
			{
				params := new(AmountToUiAmount)
				fu.Fuzz(params)
				params.AccountMetaSlice = nil
				buf := new(bytes.Buffer)
				err := encodeT(*params, buf)
				ag_require.NoError(t, err)
				got := new(AmountToUiAmount)
				err = decodeT(got, buf.Bytes())
				got.AccountMetaSlice = nil
				ag_require.NoError(t, err)
				ag_require.Equal(t, params, got)
			}
		})
	}
}

func TestEncodeDecode_UiAmountToAmount(t *testing.T) {
	t.Run("UiAmountToAmount", func(t *testing.T) {
		uiAmount := "123.456"
		params := &UiAmountToAmount{
			UiAmount: &uiAmount,
		}
		buf := new(bytes.Buffer)
		err := encodeT(*params, buf)
		ag_require.NoError(t, err)
		got := new(UiAmountToAmount)
		err = decodeT(got, buf.Bytes())
		got.AccountMetaSlice = nil
		params.AccountMetaSlice = nil
		ag_require.NoError(t, err)
		ag_require.Equal(t, params, got)
	})
}

func TestEncodeDecode_InitializeMintCloseAuthority(t *testing.T) {
	fu := ag_gofuzz.New().NilChance(0)
	for i := 0; i < 1; i++ {
		t.Run("InitializeMintCloseAuthority"+strconv.Itoa(i), func(t *testing.T) {
			{
				params := new(InitializeMintCloseAuthority)
				fu.Fuzz(params)
				params.AccountMetaSlice = nil
				buf := new(bytes.Buffer)
				err := encodeT(*params, buf)
				ag_require.NoError(t, err)
				got := new(InitializeMintCloseAuthority)
				err = decodeT(got, buf.Bytes())
				got.AccountMetaSlice = nil
				ag_require.NoError(t, err)
				ag_require.Equal(t, params, got)
			}
		})
	}
}

// Test full instruction encode/decode round-trip through the Instruction wrapper.
func TestInstructionRoundTrip(t *testing.T) {
	mint := ag_gofuzz.New().NilChance(0)
	var key [32]byte
	mint.Fuzz(&key)

	t.Run("InitializeMint2_FullRoundTrip", func(t *testing.T) {
		var decimals uint8 = 9
		inst := NewInitializeMint2InstructionBuilder().
			SetDecimals(decimals).
			SetMintAuthority(key).
			SetMintAccount(key)
		built := inst.Build()
		data, err := built.Data()
		ag_require.NoError(t, err)
		ag_require.Equal(t, byte(Instruction_InitializeMint2), data[0])

		decoded, err := DecodeInstruction(nil, data)
		ag_require.NoError(t, err)
		ag_require.Equal(t, "InitializeMint2", InstructionIDToName(decoded.TypeID.Uint8()))
	})

	t.Run("InitializeImmutableOwner_FullRoundTrip", func(t *testing.T) {
		inst := NewInitializeImmutableOwnerInstructionBuilder().
			SetAccount(key)
		built := inst.Build()
		data, err := built.Data()
		ag_require.NoError(t, err)
		ag_require.Equal(t, byte(Instruction_InitializeImmutableOwner), data[0])

		decoded, err := DecodeInstruction(nil, data)
		ag_require.NoError(t, err)
		ag_require.Equal(t, "InitializeImmutableOwner", InstructionIDToName(decoded.TypeID.Uint8()))
	})

	t.Run("InitializeMintCloseAuthority_FullRoundTrip", func(t *testing.T) {
		inst := NewInitializeMintCloseAuthorityInstructionBuilder().
			SetCloseAuthority(key).
			SetMintAccount(key)
		built := inst.Build()
		data, err := built.Data()
		ag_require.NoError(t, err)
		ag_require.Equal(t, byte(Instruction_InitializeMintCloseAuthority), data[0])

		decoded, err := DecodeInstruction(nil, data)
		ag_require.NoError(t, err)
		ag_require.Equal(t, "InitializeMintCloseAuthority", InstructionIDToName(decoded.TypeID.Uint8()))
	})
}
