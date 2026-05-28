package solana

import (
	"testing"
)

// Benchmarks for PublicKey/Signature base58 methods, which are the
// hot paths that most users hit via Stringer and JSON marshaling.

var (
	benchPubkey    = MustPublicKeyFromBase58("4cHoJNmLed5PBgFBezHmJkMJLEZrcTvr3aopjnYBRxUb")
	benchSignature = MustSignatureFromBase58("5YBLhMBLjhAHnEPnHKLLnVwHSfXGPJMCvKAfNsiaEw2T63edrYxVFHKUxRXfP6KA1HVo7c9JZ3LAJQR72giX7Cb")

	benchPubkeyStr = "4cHoJNmLed5PBgFBezHmJkMJLEZrcTvr3aopjnYBRxUb"
	benchSigStr    = "5YBLhMBLjhAHnEPnHKLLnVwHSfXGPJMCvKAfNsiaEw2T63edrYxVFHKUxRXfP6KA1HVo7c9JZ3LAJQR72giX7Cb"
)

func BenchmarkPublicKey_String(b *testing.B) {
	pk := benchPubkey
	for b.Loop() {
		_ = pk.String()
	}
}

func BenchmarkPublicKeyFromBase58(b *testing.B) {
	for b.Loop() {
		PublicKeyFromBase58(benchPubkeyStr)
	}
}

func BenchmarkSignature_String(b *testing.B) {
	sig := benchSignature
	for b.Loop() {
		_ = sig.String()
	}
}

func BenchmarkSignatureFromBase58(b *testing.B) {
	for b.Loop() {
		SignatureFromBase58(benchSigStr)
	}
}

func BenchmarkPublicKey_MarshalJSON(b *testing.B) {
	pk := benchPubkey
	for b.Loop() {
		pk.MarshalJSON()
	}
}

func BenchmarkSignature_MarshalJSON(b *testing.B) {
	sig := benchSignature
	for b.Loop() {
		sig.MarshalJSON()
	}
}
