package api

import (
	"gild-wallet/internal/wallets"
)

type WalletAPI struct{}

func (api *WalletAPI) CreateWallet(cryptoType wallets.CryptoType) (string, string, error) {
	w, err := wallets.NewWallet(cryptoType)
	if err != nil {
		return "", "", err
	}

	return w.Address, w.PrivateKey, nil
}

func (api *WalletAPI) RestoreWallet(privateKey string, cryptoType wallets.CryptoType) (string, string, error) {
	w, err := wallets.RestoreWallet(privateKey, cryptoType)
	if err != nil {
		return "", "", err
	}

	return w.Address, w.PrivateKey, nil
}
