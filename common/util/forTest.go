package util

import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/crypto"
)

func GetKeyByText(s string) *ecdsa.PrivateKey{
	aHash := Hash256(s);
	aKey, _ := crypto.ToECDSA(aHash)
	return aKey;
}