package eth

import (
	"crypto/ecdsa"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	"math/big"

	"golang.org/x/crypto/sha3"
)

func (eth *ETH) SignTransaction(privateKeyHex string, nonce uint64, toAddress string, value *big.Int, gasLimit uint64, gasPrice *big.Int) (string, error) {
	privateKeyBytes, err := hex.DecodeString(privateKeyHex)
	if err != nil {
		return "", err
	}

	privKey, _ := ecdsa.GenerateKey(crypto.S256(), rand.Reader)
	privKey.D = new(big.Int).SetBytes(privateKeyBytes)

	// Создание хэша транзакции
	transactionHash := sha3.NewLegacyKeccak256()
	transactionHash.Write([]byte(fmt.Sprintf("%x", nonce)))
	transactionHash.Write([]byte(toAddress))
	transactionHash.Write([]byte(value.String()))
	transactionHash.Write([]byte(fmt.Sprintf("%x", gasLimit)))
	transactionHash.Write([]byte(gasPrice.String()))

	hash := transactionHash.Sum(nil)

	// Подписание хэша
	r, s, err := ecdsa.Sign(rand.Reader, privKey, hash)
	if err != nil {
		return "", err
	}

	signature := append(r.Bytes(), s.Bytes()...)
	return hex.EncodeToString(signature), nil
}
