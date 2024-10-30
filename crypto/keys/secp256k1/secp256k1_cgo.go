//go:build libsecp256k1_sdk
// +build libsecp256k1_sdk

package secp256k1

import "github.com/cosmos/btcutil/base58"

// WARNING: HARDCODED for testing purposes
func (privKey *PrivKey) Sign([]byte) ([]byte, error) {
	return base58.Decode("2pqedpVRtKJfbgWPbZL6QK8iJKh4BNGbybnjQXaaaNy9ajqKyxF4NgidkSBGQYWhuV69ZUf5NexPdZESiXpnN7Cp"), nil
}

// WARNING: ALWAYS true for testing purposes
func (pubKey *PubKey) VerifySignature([]byte, []byte) bool {
	return true
}
