package wallet

import (
	"errors"
)

type CryptoType int

const (
	BTC CryptoType = iota
	ETH
)

type Wallet struct {
	Address string
	PrivateKey string
	CryptoType CryptoType
}

func NewWallet(cryptoType CryptoType) (*Wallet, error) {
	var address, privateKey string
	var err error

	switch cryptoType {
	case BTC: 
	}
}
