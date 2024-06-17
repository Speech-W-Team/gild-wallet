package core

import (
	"errors"
	"github.com/btcsuite/btcd/btcutil/hdkeychain"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/tyler-smith/go-bip39"
)

func SeedFromMnemonic(mnemonic string, password string) []byte {
	return bip39.NewSeed(mnemonic, password)
}

func MasterKeyFromSeed(seed []byte, networkType NetworkType) (*hdkeychain.ExtendedKey, error) {
	var params *chaincfg.Params
	switch networkType {
	case Mainnet:
		params = &chaincfg.MainNetParams
	case Testnet:
		params = &chaincfg.TestNet3Params
	}

	return hdkeychain.NewMaster(seed, params)
}

func HDWallet(masterKey *hdkeychain.ExtendedKey, path DerivationPathItem) (*hdkeychain.ExtendedKey, error) {
	var derivationKey *hdkeychain.ExtendedKey
	var err error
	var derivationPath = &path

	if path.Hardened {
		derivationKey, err = masterKey.Derive(hdkeychain.HardenedKeyStart + derivationPath.Path)
	} else {
		derivationKey, err = masterKey.Derive(path.Path)
	}
	derivationPath = path.Child
	if err != nil {
		return nil, err
	}

	for derivationPath != nil {
		if derivationKey != nil {
			if derivationPath.Hardened {
				derivationKey, err = derivationKey.Derive(hdkeychain.HardenedKeyStart + derivationPath.Path)
			} else {
				derivationKey, err = derivationKey.Derive(derivationPath.Path)
			}
		}
		derivationPath = derivationPath.Child
	}

	if derivationKey == nil {
		return nil, ErrDerivationKeyNil
	}

	return derivationKey, nil
}

var ErrDerivationKeyNil = errors.New("derivation key is nil")
