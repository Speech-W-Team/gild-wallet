package api

import (
	"gild-wallet/wallets"
	"gild-wallet/wallets/blockchains"
)

type WalletAPI struct{}

func (api *WalletAPI) CreateWallet(cryptoType blockchains.CryptoType) (string, string, error) {
	w, err := wallets.NewWallet(cryptoType)
	if err != nil {
		return "", "", err
	}

	return w.Address, w.PrivateKey, nil
}

func (api *WalletAPI) RestoreWallet(privateKey []byte, cryptoType blockchains.CryptoType) (string, string, error) {
	w, err := wallets.RestoreWallet(privateKey, cryptoType)
	if err != nil {
		return "", "", err
	}

	return w.Address, w.PrivateKey, nil
}

func (api *WalletAPI) RestorePrivateKeyViaMnemonic(mnemonicPhrase string, password string, cryptoType blockchains.CryptoType) (string, string, error) {
	wallet, err := wallets.RestoreWalletViaMnemonicPhrase(mnemonicPhrase, password, cryptoType)
	if err != nil {
		return "", "", err
	}
	return wallet.Address, wallet.PrivateKey, err
}
