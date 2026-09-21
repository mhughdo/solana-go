// Package encryption provides the encryption primitives needed for confidential transfers:
// ElGamal keys and ciphertexts, Pedersen commitments and openings,
// grouped ElGamal ciphertexts, and authenticated encryption.
//
// # Key material is not zeroized
//
// ElGamalSecretKey, PedersenOpening and AeKey are fixed-size value arrays.
// Their Rust counterparts implement Zeroize and wipe themselves on drop; the Go
// types cannot.
//
// Treat secret key material as readable for the lifetime of the process:
// keep it out of logs, dumps and long-lived structures, and prefer rederiving a key
// over caching it.
package encryption
