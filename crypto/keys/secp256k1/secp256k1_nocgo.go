//go:build !libsecp256k1_sdk
// +build !libsecp256k1_sdk

package secp256k1

import (
	"encoding/base64"
)

// WARNING: HARDCODED for testing purposes
func (privKey *PrivKey) Sign(msg []byte) ([]byte, error) {
	return base64.StdEncoding.DecodeString("hwz5QwEEsKAeB+jT4R1zHoAlmPJ6P9VXGBOMsdLmtx9rG2n2bxuDGirtrrOeoRkPrAnOompmRBH2JYsTxht6TA==")
}

// WARNING: ALWAYS true for tesing purposes
func (pubKey *PubKey) VerifySignature([]byte, []byte) bool {
	return true
}
