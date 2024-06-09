package btc

import (
	"crypto/ecdsa"
	"crypto/sha256"
	_ "errors"
	"gild-wallet/internal/crypto"
	"golang.org/x/crypto/ripemd160"
)

func GenerateWallet() (string, string, error) {
	privKey, pubKey, err := crypto.GenerateKeyPair()
	if err != nil {
		return "", "", err
	}

	address, err := generateAddress(pubKey)
	if err != nil {
		return "", "", err
	}

	privateKey := encodePrivateKey(privKey)
	return address, privateKey, nil
}

func generateAddress(pubKey *ecdsa.PublicKey) (string, error) {
	pubKeyBytes := append(pubKey.X.Bytes(), pubKey.Y.Bytes()...)
	hash256 := sha256.New()
	hash256.Write(pubKeyBytes)
	hashedPubKey := hash256.Sum(nil)

	ripemd160Hasher := ripemd160.New()
	ripemd160Hasher.Write(hashedPubKey)
	hashedPubKeyRipemd160 := ripemd160Hasher.Sum(nil)

	versionedPayload := append([]byte{0x00}, hashedPubKeyRipemd160...)

	hash256 = sha256.New()
	hash256.Write(versionedPayload)
	checksum := hash256.Sum(nil)

	hash256 = sha256.New()
	hash256.Write(checksum)
	checksum = hash256.Sum(nil)[:4]

	fullPayload := append(versionedPayload, checksum...)

	address := crypto.Base58Encode(fullPayload)

	return address, nil
}

func encodePrivateKey(privKey *ecdsa.PrivateKey) string {
	return crypto.Base58Encode(privKey.D.Bytes())
}
