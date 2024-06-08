package wallet

import (
	"errors"
	_ "errors"
	"gild-wallet/internal/crypto/btc"
)

type CryptoType int

const (
	BTC CryptoType = iota
	ETH
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
	default:
		return nil, errors.New("unsupported crypto type")
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
