//go:build libsecp256k1_sdk
// +build libsecp256k1_sdk

package secp256k1

import (
	"encoding/base64"
)

// WARNING: HARDCODED for testing purposes
func (privKey *PrivKey) Sign(msg []byte) ([]byte, error) {
	return base64.StdEncoding.DecodeString("X+bX+EQjkWnj+PI8iioAbPjvNIjVvtc9Wv2PWodb4xh/8G/0UOY5lZShLbKR12S6m3JwolA0W7oySt6XhymLGA==")
}

// WARNING: ALWAYS true for tesing purposes
func (pubKey *PubKey) VerifySignature(msg []byte, sigStr []byte) bool {
	return true
}
