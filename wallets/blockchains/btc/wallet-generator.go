package btc

import (
	"crypto/ecdsa"
	_ "errors"
	gild "gild-wallet/cryptography"
	"gild-wallet/wallets/blockchains"
)

type BTC struct{}

func NewBTC() *BTC {
	return &BTC{}
}

func (btc *BTC) GenerateWallet(networkType blockchains.NetworkType) (blockchains.Wallet, error) {
	privKey, pubKey, err := gild.GenerateKeyPair()
	if err != nil {
		return blockchains.Wallet{}, err
	}

	address, err := btc.GenerateAddress(append(pubKey.X.Bytes(), pubKey.Y.Bytes()...), networkType)
	if err != nil {
		return blockchains.Wallet{}, err
	}

	privateKey := encodePrivateKey(privKey)
	return blockchains.Wallet{
		Address:    address,
		PrivateKey: privateKey,
		CryptoType: blockchains.BTC,
	}, nil
}

func encodePrivateKey(privKey *ecdsa.PrivateKey) string {
	return gild.Base58Encode(privKey.D.Bytes())
}
