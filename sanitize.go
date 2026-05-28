package solana

import (
	"errors"
	"fmt"
)

// sanitizeError represents a message or transaction validation error.
type sanitizeError struct {
	msg string
}

func (e *sanitizeError) Error() string {
	return e.msg
}

func newSanitizeError(format string, args ...any) error {
	return &sanitizeError{msg: fmt.Sprintf(format, args...)}
}

// IsSanitizeError reports whether err is a sanitization validation error.
func IsSanitizeError(err error) bool {
	var se *sanitizeError
	return errors.As(err, &se)
}

// maxAccountKeys is the maximum number of accounts a message can reference.
// Account indices are encoded as u8, so the limit is 256.
const maxAccountKeys = 256

// Sanitize validates the structural integrity of a Message.
// Ported from solana-sdk/message: legacy.rs sanitize() and v0/mod.rs sanitize().
func (m *Message) Sanitize() error {
	if m.IsVersioned() {
		return m.sanitizeV0()
	}
	return m.sanitizeLegacy()
}

func (m *Message) sanitizeLegacy() error {
	numKeys := len(m.AccountKeys)

	// Signing area and read-only non-signing area should not overlap.
	if int(m.Header.NumRequiredSignatures)+int(m.Header.NumReadonlyUnsignedAccounts) > numKeys {
		return newSanitizeError("header references more accounts than available: required_signatures(%d) + readonly_unsigned(%d) > account_keys(%d)",
			m.Header.NumRequiredSignatures, m.Header.NumReadonlyUnsignedAccounts, numKeys)
	}

	// There should be at least 1 RW fee-payer account.
	if m.Header.NumReadonlySignedAccounts >= m.Header.NumRequiredSignatures {
		return newSanitizeError("no writable signer: readonly_signed(%d) >= required_signatures(%d)",
			m.Header.NumReadonlySignedAccounts, m.Header.NumRequiredSignatures)
	}

	for i, ci := range m.Instructions {
		if int(ci.ProgramIDIndex) >= numKeys {
			return newSanitizeError("instruction %d: program_id_index %d out of bounds (account_keys len %d)", i, ci.ProgramIDIndex, numKeys)
		}
		// A program cannot be the payer.
		if ci.ProgramIDIndex == 0 {
			return newSanitizeError("instruction %d: program_id_index cannot be 0 (fee payer)", i)
		}
		for _, ai := range ci.Accounts {
			if int(ai) >= numKeys {
				return newSanitizeError("instruction %d: account index %d out of bounds (account_keys len %d)", i, ai, numKeys)
			}
		}
	}

	return nil
}

func (m *Message) sanitizeV0() error {
	numStaticKeys := len(m.AccountKeys)

	// Signing area and read-only non-signing area should not overlap.
	if int(m.Header.NumRequiredSignatures)+int(m.Header.NumReadonlyUnsignedAccounts) > numStaticKeys {
		return newSanitizeError("header references more accounts than available: required_signatures(%d) + readonly_unsigned(%d) > static_keys(%d)",
			m.Header.NumRequiredSignatures, m.Header.NumReadonlyUnsignedAccounts, numStaticKeys)
	}

	// There should be at least 1 RW fee-payer account.
	if m.Header.NumReadonlySignedAccounts >= m.Header.NumRequiredSignatures {
		return newSanitizeError("no writable signer: readonly_signed(%d) >= required_signatures(%d)",
			m.Header.NumReadonlySignedAccounts, m.Header.NumRequiredSignatures)
	}

	// Count dynamic keys from address table lookups.
	numDynamicKeys := 0
	for _, lookup := range m.AddressTableLookups {
		numLookupIndexes := len(lookup.WritableIndexes) + len(lookup.ReadonlyIndexes)
		// Each lookup table must be used to load at least one account.
		if numLookupIndexes == 0 {
			return newSanitizeError("address table lookup for %s loads no accounts", lookup.AccountKey)
		}
		numDynamicKeys += numLookupIndexes
	}

	if numStaticKeys == 0 {
		return newSanitizeError("message has no account keys")
	}

	// The combined number of static and dynamic account keys must be <= 256
	// since account indices are encoded as u8.
	totalKeys := numStaticKeys + numDynamicKeys
	if totalKeys > maxAccountKeys {
		return newSanitizeError("total account keys %d exceeds maximum %d", totalKeys, maxAccountKeys)
	}

	maxAccountIdx := totalKeys - 1
	// Program IDs must be in static keys only (not from lookup tables).
	maxProgramIdx := numStaticKeys - 1

	for i, ci := range m.Instructions {
		if int(ci.ProgramIDIndex) > maxProgramIdx {
			return newSanitizeError("instruction %d: program_id_index %d exceeds static keys (max %d)", i, ci.ProgramIDIndex, maxProgramIdx)
		}
		// A program cannot be the payer.
		if ci.ProgramIDIndex == 0 {
			return newSanitizeError("instruction %d: program_id_index cannot be 0 (fee payer)", i)
		}
		for _, ai := range ci.Accounts {
			if int(ai) > maxAccountIdx {
				return newSanitizeError("instruction %d: account index %d out of bounds (max %d)", i, ai, maxAccountIdx)
			}
		}
	}

	return nil
}

// HasDuplicates checks if the message has duplicate account keys.
// Uses O(n^2) comparison but requires no heap allocation, which is faster
// for the typically small number of accounts in a message.
// Ported from solana-sdk/message/legacy.rs has_duplicates().
func (m *Message) HasDuplicates() bool {
	keys := m.AccountKeys
	for i := 1; i < len(keys); i++ {
		for j := i; j < len(keys); j++ {
			if keys[i-1].Equals(keys[j]) {
				return true
			}
		}
	}
	return false
}

// Sanitize validates the structural integrity of a Transaction.
// It checks that the signature count matches the message header and
// that the message itself is valid.
// Ported from solana-sdk/transaction: lib.rs and versioned/mod.rs sanitize().
func (tx *Transaction) Sanitize() error {
	numSigs := len(tx.Signatures)
	numRequired := int(tx.Message.Header.NumRequiredSignatures)
	numStaticKeys := len(tx.Message.AccountKeys)

	// Signature count must exactly match num_required_signatures.
	if numRequired > numSigs {
		return newSanitizeError("not enough signatures: required %d, got %d", numRequired, numSigs)
	}
	if numRequired < numSigs {
		return newSanitizeError("too many signatures: required %d, got %d", numRequired, numSigs)
	}

	// Signatures must not exceed static account keys count
	// (signatures are verified before lookup keys are loaded).
	if numSigs > numStaticKeys {
		return newSanitizeError("more signatures (%d) than static account keys (%d)", numSigs, numStaticKeys)
	}

	return tx.Message.Sanitize()
}

// VerifyWithResults verifies each signature independently and returns
// a per-signature boolean result.
// Ported from solana-sdk/transaction/lib.rs verify_with_results().
func (tx *Transaction) VerifyWithResults() ([]bool, error) {
	msg, err := tx.Message.MarshalBinary()
	if err != nil {
		return nil, err
	}

	results := make([]bool, len(tx.Signatures))
	for i, sig := range tx.Signatures {
		if i < len(tx.Message.AccountKeys) {
			results[i] = sig.Verify(tx.Message.AccountKeys[i], msg)
		}
	}
	return results, nil
}

// isAdvanceNonceInstructionData checks if the instruction data starts with
// the AdvanceNonceAccount discriminant (u32 LE value 4).
func isAdvanceNonceInstructionData(data []byte) bool {
	return len(data) >= 4 && data[0] == 4 && data[1] == 0 && data[2] == 0 && data[3] == 0
}

// nonceAdvanceInstruction returns the first instruction if it is a
// System Program AdvanceNonceAccount instruction, or nil otherwise.
func (tx *Transaction) nonceAdvanceInstruction() *CompiledInstruction {
	if len(tx.Message.Instructions) == 0 {
		return nil
	}
	ix := &tx.Message.Instructions[0]

	// Check that the program is the System Program.
	if int(ix.ProgramIDIndex) >= len(tx.Message.AccountKeys) {
		return nil
	}
	if !tx.Message.AccountKeys[ix.ProgramIDIndex].Equals(SystemProgramID) {
		return nil
	}
	if !isAdvanceNonceInstructionData(ix.Data) {
		return nil
	}
	return ix
}

// UsesDurableNonce checks whether this transaction uses a durable nonce
// by inspecting the first instruction. Returns true if the first instruction
// is a System Program AdvanceNonceAccount instruction.
// Ported from solana-sdk/transaction: uses_durable_nonce().
func (tx *Transaction) UsesDurableNonce() bool {
	return tx.nonceAdvanceInstruction() != nil
}

// GetNonceAccount returns the public key of the nonce account if this
// transaction uses a durable nonce. The nonce account is the first account
// of the first instruction (the AdvanceNonceAccount instruction).
// Returns the zero PublicKey and false if this is not a nonce transaction.
func (tx *Transaction) GetNonceAccount() (PublicKey, bool) {
	ix := tx.nonceAdvanceInstruction()
	if ix == nil {
		return PublicKey{}, false
	}
	if len(ix.Accounts) == 0 {
		return PublicKey{}, false
	}
	nonceAccountIdx := ix.Accounts[0]
	if int(nonceAccountIdx) >= len(tx.Message.AccountKeys) {
		return PublicKey{}, false
	}
	return tx.Message.AccountKeys[nonceAccountIdx], true
}
