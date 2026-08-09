package solana

import (
	"crypto/sha256"
	"errors"
	"math"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/require"
)

func createProgramAddressReference(seeds [][]byte, programID PublicKey) (PublicKey, error) {
	if len(seeds) > MaxSeeds {
		return PublicKey{}, ErrMaxSeedLengthExceeded
	}
	for _, seed := range seeds {
		if len(seed) > MaxSeedLength {
			return PublicKey{}, ErrMaxSeedLengthExceeded
		}
	}
	buffer := []byte{}
	for _, seed := range seeds {
		buffer = append(buffer, seed...)
	}
	buffer = append(buffer, programID[:]...)
	buffer = append(buffer, []byte(PDA_MARKER)...)
	hash := sha256.Sum256(buffer)
	if IsOnCurve(hash[:]) {
		return PublicKey{}, errors.New("invalid seeds; address must fall off the curve")
	}
	return PublicKeyFromBytes(hash[:]), nil
}

func findProgramAddressReference(seeds [][]byte, programID PublicKey) (PublicKey, uint8, error) {
	bump := uint8(math.MaxUint8)
	for bump != 0 {
		address, err := createProgramAddressReference(append(seeds, []byte{bump}), programID)
		if err == nil {
			return address, bump, nil
		}
		bump--
	}
	return PublicKey{}, bump, errors.New("unable to find a valid program address")
}

func TestProgramAddressOptimizedParity(t *testing.T) {
	rng := rand.New(rand.NewSource(1))
	for iteration := 0; iteration < 500; iteration++ {
		seedCount := rng.Intn(MaxSeeds + 2)
		seeds := make([][]byte, seedCount)
		for i := range seeds {
			seeds[i] = make([]byte, rng.Intn(MaxSeedLength+2))
			_, _ = rng.Read(seeds[i])
		}
		var programID PublicKey
		_, _ = rng.Read(programID[:])

		want, wantErr := createProgramAddressReference(seeds, programID)
		got, gotErr := CreateProgramAddress(seeds, programID)
		require.Equal(t, want, got, "iteration %d", iteration)
		if wantErr == nil {
			require.NoError(t, gotErr, "iteration %d", iteration)
		} else {
			require.EqualError(t, gotErr, wantErr.Error(), "iteration %d", iteration)
		}
	}

	for iteration := 0; iteration < 100; iteration++ {
		seedCount := rng.Intn(MaxSeeds)
		seeds := make([][]byte, seedCount)
		for i := range seeds {
			seeds[i] = make([]byte, rng.Intn(MaxSeedLength+1))
			_, _ = rng.Read(seeds[i])
		}
		var programID PublicKey
		_, _ = rng.Read(programID[:])

		wantAddress, wantBump, wantErr := findProgramAddressReference(seeds, programID)
		gotAddress, gotBump, gotErr := FindProgramAddress(seeds, programID)
		require.Equal(t, wantAddress, gotAddress, "iteration %d", iteration)
		require.Equal(t, wantBump, gotBump, "iteration %d", iteration)
		if wantErr == nil {
			require.NoError(t, gotErr, "iteration %d", iteration)
		} else {
			require.EqualError(t, gotErr, wantErr.Error(), "iteration %d", iteration)
		}
	}
}

func BenchmarkCreateProgramAddress(b *testing.B) {
	programID := MustPublicKeyFromBase58("BPFLoaderUpgradeab1e11111111111111111111111")
	seed := MustPublicKeyFromBase58("SeedPubey1111111111111111111111111111111111")
	seeds := [][]byte{seed[:], {1}}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		address, err := CreateProgramAddress(seeds, programID)
		if err != nil {
			b.Fatal(err)
		}
		_ = address
	}
}

func BenchmarkCreateProgramAddressReference(b *testing.B) {
	programID := MustPublicKeyFromBase58("BPFLoaderUpgradeab1e11111111111111111111111")
	seed := MustPublicKeyFromBase58("SeedPubey1111111111111111111111111111111111")
	seeds := [][]byte{seed[:], {1}}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		address, err := createProgramAddressReference(seeds, programID)
		if err != nil {
			b.Fatal(err)
		}
		_ = address
	}
}

func BenchmarkFindProgramAddress(b *testing.B) {
	programID := MustPublicKeyFromBase58("BPFLoader1111111111111111111111111111111111")
	seeds := [][]byte{[]byte("Lil'"), []byte("Bits")}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		address, bump, err := FindProgramAddress(seeds, programID)
		if err != nil {
			b.Fatal(err)
		}
		_, _ = address, bump
	}
}

func BenchmarkFindProgramAddressReference(b *testing.B) {
	programID := MustPublicKeyFromBase58("BPFLoader1111111111111111111111111111111111")
	seeds := [][]byte{[]byte("Lil'"), []byte("Bits")}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		address, bump, err := findProgramAddressReference(seeds, programID)
		if err != nil {
			b.Fatal(err)
		}
		_, _ = address, bump
	}
}

func BenchmarkFindAssociatedTokenAddress(b *testing.B) {
	wallet := MustPublicKeyFromBase58("SeedPubey1111111111111111111111111111111111")
	mint := MustPublicKeyFromBase58("So11111111111111111111111111111111111111112")
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		address, bump, err := FindAssociatedTokenAddress(wallet, mint)
		if err != nil {
			b.Fatal(err)
		}
		_, _ = address, bump
	}
}
