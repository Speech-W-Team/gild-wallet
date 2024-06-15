package tron

import (
	"crypto/sha256"
	gild "gild-wallet/cryptography"
	"gild-wallet/wallets/blockchains"
	"golang.org/x/crypto/sha3"
)

func (tron *TRON) GenerateAddress(pubKey []byte, networkType blockchains.NetworkType) (string, error) {
	hash3 := sha3.NewLegacyKeccak256()
	hash3.Write(pubKey)
	hashed256PubKey := hash3.Sum(nil)

	initialAddress := hashed256PubKey[len(hashed256PubKey)-21:]
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
