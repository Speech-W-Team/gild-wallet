package tron

import (
	"crypto/ecdsa"
	"crypto/sha256"
	"encoding/hex"
	gild "gild-wallet/internal/crypto"
	"github.com/ethereum/go-ethereum/crypto"
	"golang.org/x/crypto/sha3"
	"log"
)

func GenerateWallet() (string, string, error) {
	privKey, pubKey, err := gild.GenerateKeyPair()

	if err != nil {
		return "", "", err
	}

	address, err := generateAddress(pubKey)
	if err != nil {
		return "", "", err
	}

	//privateKey := encodePrivateKey(privKey)
	return address, hex.EncodeToString(append(privKey.D.Bytes())), nil
}

func RestoreWallet(userPrivateKey string) (string, string, error) {
	privKey, err := crypto.HexToECDSA(userPrivateKey)
	if err != nil {
		return "", "", err
	}
	privateKey := crypto.FromECDSA(privKey)
	publicKey := privKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	address, err := generateAddress(publicKeyECDSA)
	if err != nil {
		return "", "", err
	}
	return address, hex.EncodeToString(privateKey), nil
}

func generateAddress(pubKey *ecdsa.PublicKey) (string, error) {
	pubKeyBytes := append(pubKey.X.Bytes(), pubKey.Y.Bytes()...)
	hash3 := sha3.NewLegacyKeccak256()
	hash3.Write(pubKeyBytes)
	hash3256 := hash3.Sum(nil)
	initialAddress := hash3256[len(hash3256)-21:]
	initialAddress[0] = 0x41
	hash2 := sha256.New()
	hash2.Write(initialAddress)
	hash := hash2.Sum(nil)
	hash2 = sha256.New()
	hash2.Write(hash)
	hash = hash2.Sum(nil)
	verificationCode := hash[:4]
	initialAddress = append(initialAddress, verificationCode...)

	return gild.Base58Encode(initialAddress), nil
}
