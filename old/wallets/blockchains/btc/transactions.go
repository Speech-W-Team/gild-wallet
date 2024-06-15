package btc

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
)

func SignTransaction(message string, privateKey *ecdsa.PrivateKey) ([]byte, error) {
	hash := sha256.Sum256([]byte(message))
	r, s, err := ecdsa.Sign(rand.Reader, privateKey, hash[:])
	if err != nil {
		return nil, err
	}

	// Преобразование подписи в компактный формат
	signature := append(r.Bytes(), s.Bytes()...)

	return signature, nil
}
