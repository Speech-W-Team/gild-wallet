package eth

import (
	"crypto/ecdsa"
	"encoding/hex"
	_ "errors"
	"log"

	"github.com/ethereum/go-ethereum/crypto"
	"golang.org/x/crypto/sha3"
)

func GenerateWallet() (string, string, error) {
	privKey, err := crypto.GenerateKey()
	if err != nil {
		return "", "", err
	}
	privateKey := crypto.FromECDSA(privKey)
	publicKey := privKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	return address, hex.EncodeToString(privateKey), nil
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

	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	return address, hex.EncodeToString(privateKey), nil
}

func generateAddress(pubKey *ecdsa.PublicKey) string {
	pubKeyBytes := append(pubKey.X.Bytes(), pubKey.Y.Bytes()...)
	hash := sha3.NewLegacyKeccak256()
	hash.Write(pubKeyBytes[1:])
	address := hash.Sum(nil)[12:]
	return "0x" + hex.EncodeToString(address)
}
