package btc

import (
	"crypto/sha256"
	"gild-wallet/cryptography"
	"gild-wallet/wallets/blockchains"
	"golang.org/x/crypto/ripemd160"
)

func (btc *BTC) GenerateAddress(pubKey []byte, networkType blockchains.NetworkType) (string, error) {
	var networkBytes []byte
	switch networkType {
	case blockchains.Mainnet:
		networkBytes = []byte("80")
	case blockchains.Testnet:
		networkBytes = []byte("ef")
	}
	publicKey := append(networkBytes, pubKey...)
	//publicKey = append(publicKey, []byte("01")...)
	hash256 := sha256.New()
	hash256.Write(publicKey)
	hashedPubKey := hash256.Sum(nil)
	hash256 = sha256.New()
	hash256.Write(hashedPubKey)
	hashedPubKey = hash256.Sum(nil)[:4]

	ripemd160Hasher := ripemd160.New()
	ripemd160Hasher.Write(hashedPubKey)
	hashedPubKeyRipemd160 := ripemd160Hasher.Sum(nil)

	//versionedPayload := append([]byte{0x80}, hashedPubKeyRipemd160...)

	//hash256 = sha256.New()
	//hash256.Write(versionedPayload)
	//checksum := hash256.Sum(nil)
	//
	//hash256 = sha256.New()
	//hash256.Write(checksum)
	//checksum = hash256.Sum(nil)[:4]

	fullPayload := append(hashedPubKeyRipemd160, hashedPubKey...)

	address := cryptography.Base58Encode(fullPayload)

	return address, nil
}
