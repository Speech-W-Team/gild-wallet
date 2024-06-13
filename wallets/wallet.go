package wallets

import (
	"crypto/ecdsa"
	"errors"
	"gild-wallet/wallets/blockchains"
	"gild-wallet/wallets/blockchains/btc"
	"gild-wallet/wallets/blockchains/eth"
	"gild-wallet/wallets/blockchains/tron"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil/hdkeychain"
	"github.com/tyler-smith/go-bip39"
)

func NewWallet(cryptoType blockchains.CryptoType) (*blockchains.Wallet, error) {
	cryptocurrency, err := NewCryptocurrency(cryptoType)
	if err != nil {
		return nil, err
	}
	wallet, err := cryptocurrency.GenerateWallet(blockchains.Mainnet)
	return &wallet, err
}

func RestoreWallet(privateKey []byte, cryptoType blockchains.CryptoType) (*blockchains.Wallet, error) {
	cryptocurrency, err := NewCryptocurrency(cryptoType)
	if err != nil {
		return nil, err
	}
	wallet, err := cryptocurrency.RestoreWallet(privateKey, blockchains.Mainnet)
	return &wallet, err
}

func RestoreWalletViaMnemonicPhrase(mnemonicPhrase string, password string, cryptoType blockchains.CryptoType) (*blockchains.Wallet, error) {
	var path []DerivationPathItem
	switch cryptoType {
	case blockchains.BTC:
		path = []DerivationPathItem{NewDerivationPath(44, true), NewDerivationPath(0, true), NewDerivationPath(0, true), {Hardened: false, Path: 0}}
	case blockchains.ETH:
		path = []DerivationPathItem{NewDerivationPath(44, true), NewDerivationPath(60, true), NewDerivationPath(0, true), {Hardened: false, Path: 0}, {Hardened: false, Path: 0}}
	case blockchains.TRON:
		path = []DerivationPathItem{NewDerivationPath(44, true), NewDerivationPath(195, true), NewDerivationPath(0, true), {Hardened: false, Path: 0}, {Hardened: false, Path: 0}}
	}
	temp, _, _ := hdWallet(mnemonicPhrase, password, path)
	cryptocurrency, err := NewCryptocurrency(cryptoType)
	if err != nil {
		return nil, err
	}
	wallet, err := cryptocurrency.RestoreWallet(temp.D.Bytes(), blockchains.Mainnet)
	//address, err := cryptocurrency.GenerateAddress()
	return &wallet, err
}

func NewCryptocurrency(CryptoType blockchains.CryptoType) (blockchains.Cryptocurrency, error) {
	var cryptocurrency blockchains.Cryptocurrency
	var err error

	switch CryptoType {
	case blockchains.BTC:
		cryptocurrency = btc.NewBTC()
	case blockchains.ETH:
		cryptocurrency = eth.NewETH()
	case blockchains.TRON:
		cryptocurrency = tron.NewTRON()
	default:
		err = ErrUnsupportedCryptocurrency
	}

	return cryptocurrency, err
}

type DerivationPathItem struct {
	Path     uint32
	Hardened bool
}

func NewDerivationPath(path uint32, hardened bool) DerivationPathItem {
	return DerivationPathItem{
		Path:     path,
		Hardened: hardened,
	}
}

func hdWallet(mnemonic string, password string, paths []DerivationPathItem) (*ecdsa.PrivateKey, *string, error) {
	seed := bip39.NewSeed(mnemonic, password)
	// Generate a wallets master node using the seed.
	masterKey, err := hdkeychain.NewMaster(seed, &chaincfg.MainNetParams)
	if err != nil {
		return nil, nil, err
	}
	var derivationKey *hdkeychain.ExtendedKey
	for _, path := range paths {
		if derivationKey == nil {
			if path.Hardened {
				derivationKey, err = masterKey.Child(hdkeychain.HardenedKeyStart + path.Path)
			} else {
				derivationKey, err = masterKey.Child(path.Path)
			}
		} else {
			if path.Hardened {
				derivationKey, err = derivationKey.Child(hdkeychain.HardenedKeyStart + path.Path)
			} else {
				derivationKey, err = derivationKey.Child(path.Path)
			}
		}
	}

	//acc44H, err := masterKey.Child(hdkeychain.HardenedKeyStart + 44)
	//if err != nil {
	//	return nil, nil, err
	//}
	//// This gives the path: m/44H/60H
	//acc44H60H, err := acc44H.Child(hdkeychain.HardenedKeyStart + 195)
	//if err != nil {
	//	return nil, nil, err
	//}
	//// This gives the path: m/44H/60H/0H
	//acc44H60H0H, err := acc44H60H.Child(hdkeychain.HardenedKeyStart + 0)
	//if err != nil {
	//	return nil, nil, err
	//}
	//// This gives the path: m/44H/60H/0H/0
	//acc44H60H0H0, err := acc44H60H0H.Child(0)
	//if err != nil {
	//	return nil, nil, err
	//}
	//// This gives the path: m/44H/60H/0H/0/0
	//derivationKey, err := acc44H60H0H0.Child(0)
	//if err != nil {
	//	return nil, nil, err
	//}
	btcecPrivKey, err := derivationKey.ECPrivKey()
	if err != nil {
		return nil, nil, err
	}
	privateKey := btcecPrivKey.ToECDSA()
	path := "m/44H/60H/0H/0/0"
	return privateKey, &path, nil
}

var ErrUnsupportedCryptocurrency = errors.New("unsupported cryptocurrency type")
