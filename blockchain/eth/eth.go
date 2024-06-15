package btc

import "gild-wallet/core"

type ETHWalletManager struct{}

func (manager *ETHWalletManager) GenerateWallet() (*core.Wallet, []byte, error) {
	//TODO implement me
	panic("implement me")
}

func (manager *ETHWalletManager) RestoreWallet(privateKey []byte) (*core.Wallet, error) {
	//TODO implement me
	panic("implement me")
}

func (manager *ETHWalletManager) RestoreWalletFromString(privateKey string) (*core.Wallet, error) {
	//TODO implement me
	panic("implement me")
}

func (manager *ETHWalletManager) RestoreWalletFromMnemonic(mnemonic string, password string) (*core.Wallet, error) {
	//TODO implement me
	panic("implement me")
}

func (manager *ETHWalletManager) GenerateAddress(pubKey []byte) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (manager *ETHWalletManager) GenerateAddressFromString(pubKey string) (string, error) {
	//TODO implement me
	panic("implement me")
}
