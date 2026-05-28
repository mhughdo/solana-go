package solana

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRegisterInstructionDecoder(t *testing.T) {

	decoder := func(instructionAccounts []*AccountMeta, data []byte) (any, error) {
		return nil, nil
	}
	decoderAnother := func(instructionAccounts []*AccountMeta, data []byte) (any, error) {
		return nil, nil
	}

	// First registration succeeds.
	require.NoError(t, RegisterInstructionDecoder(BPFLoaderProgramID, decoder))

	// Re-registering the same decoder is a no-op.
	require.NoError(t, RegisterInstructionDecoder(BPFLoaderProgramID, decoder))

	// Registering a different decoder for the same programID returns an error.
	assert.Error(t, RegisterInstructionDecoder(BPFLoaderProgramID, decoderAnother))
}
