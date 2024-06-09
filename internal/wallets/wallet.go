package wallets

import (
	"errors"
	_ "errors"
	"gild-wallet/internal/wallets/btc"
	"gild-wallet/internal/wallets/eth"
	"gild-wallet/internal/wallets/tron"
)

type CryptoType int

const (
	BTC CryptoType = iota
	ETH
	TRON
	SOL
	BSC
	ARB
	AVAX
	MATIC
	SUI
	OP
)

type Wallet struct {
	Address    string
	PrivateKey string
	CryptoType CryptoType
}

func NewWallet(cryptoType CryptoType) (*Wallet, error) {
	var address, privateKey string
	var err error

	switch cryptoType {
	case BTC:
		address, privateKey, err = btc.GenerateWallet()
	case ETH:
		address, privateKey, err = eth.GenerateWallet()
	case TRON:
		address, privateKey, err = tron.GenerateWallet()
	default:
		return nil, errors.New("unsupported wallets type")
	}

	if err != nil {
		return nil, err
	}

	return &Wallet{
		Address:    address,
		PrivateKey: privateKey,
		CryptoType: cryptoType,
	}, nil
}

func RestoreWallet(privateKey string, cryptoType CryptoType) (*Wallet, error) {
	var address string
	var err error

	switch cryptoType {
	//case BTC:
	//address, privateKey, err = btc.GenerateWallet()
	case ETH:
		address, privateKey, err = eth.RestoreWallet(privateKey)
	case TRON:
		address, privateKey, err = tron.RestoreWallet(privateKey)
	default:
		return nil, errors.New("unsupported wallets type")
	}

	if err != nil {
		return nil, err
	}

	return &Wallet{
		Address:    address,
		PrivateKey: privateKey,
		CryptoType: cryptoType,
	}, nil
}
