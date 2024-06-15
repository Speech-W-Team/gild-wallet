package btc

import (
	"gild-wallet/core"
)

type BTCWalletManager struct{}

func (manager *BTCWalletManager) GenerateWallet() (*core.Wallet, []byte, error) {
	panic("implement me")
}

func (manager *BTCWalletManager) RestoreWallet(privateKey []byte) (*core.Wallet, error) {
	//TODO implement me
	panic("implement me")
}

func (manager *BTCWalletManager) RestoreWalletFromString(privateKey string) (*core.Wallet, error) {
	//TODO implement me
	panic("implement me")
}

func (manager *BTCWalletManager) RestoreWalletFromMnemonic(mnemonic string, password string) (*core.Wallet, error) {
	//TODO implement me
	panic("implement me")
}

func (manager *BTCWalletManager) GenerateAddress(pubKey []byte) (string, error) {
	panic("implement me")
	pubKeyEcdsa = btc
}

func (manager *BTCWalletManager) GenerateAddressFromString(pubKey string) (string, error) {
	//TODO implement me
	panic("implement me")
}
