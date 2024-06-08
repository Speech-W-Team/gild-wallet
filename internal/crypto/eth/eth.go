package eth

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/hex"
	_ "errors"

	"golang.org/x/crypto/sha3"
)

func GenerateWallet() (string, string, error) {
	privKey, pubKey, err := generateKeyPair()
	if err != nil {
		return "", "", err
	}

	address := generateAddress(pubKey)
	privateKey := encodePrivateKey(privKey)
	return address, privateKey, nil
}

func generateKeyPair() (*ecdsa.PrivateKey, *ecdsa.PublicKey, error) {
	privKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return nil, nil, err
	}

	pubKey := &privKey.PublicKey
	return privKey, pubKey, nil
}

func generateAddress(pubKey *ecdsa.PublicKey) string {
	pubKeyBytes := append(pubKey.X.Bytes(), pubKey.Y.Bytes()...)
	hash := sha3.NewLegacyKeccak256()
	hash.Write(pubKeyBytes[1:])
	address := hash.Sum(nil)[12:]
	return "0x" + hex.EncodeToString(address)
}

func encodePrivateKey(privKey *ecdsa.PrivateKey) string {
	return hex.EncodeToString(privKey.D.Bytes())
}
