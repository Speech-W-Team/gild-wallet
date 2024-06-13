package tron

import (
	"encoding/hex"
	gild "gild-wallet/cryptography"
	"gild-wallet/wallets/blockchains"
)

type Tron struct{}

func NewTRON() *Tron {
	return &Tron{}
}

func (tron *Tron) GenerateWallet(networkType blockchains.NetworkType) (blockchains.Wallet, error) {
	privKey, pubKey, err := gild.GenerateKeyPair()

	if err != nil {
		return blockchains.Wallet{}, err
	}

	address, err := tron.GenerateAddress(append(pubKey.X.Bytes(), pubKey.Y.Bytes()...), networkType)
	if err != nil {
		return blockchains.Wallet{}, err
	}

	return blockchains.Wallet{
		Address:    address,
		PrivateKey: hex.EncodeToString(privKey.D.Bytes()),
		CryptoType: blockchains.TRON,
	}, nil
}
