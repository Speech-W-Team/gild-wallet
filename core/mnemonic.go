package core

import (
	"errors"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil/hdkeychain"
	"github.com/tyler-smith/go-bip39"
)

func SeedFromMnemonic(mnemonic string, password string) []byte {
	return bip39.NewSeed(mnemonic, password)
}

func MasterKeyFromSeed(seed []byte) (*hdkeychain.ExtendedKey, error) {
	return hdkeychain.NewMaster(seed, &chaincfg.MainNetParams)
}

func HDWallet(masterKey *hdkeychain.ExtendedKey, path DerivationPathItem) (*hdkeychain.ExtendedKey, error) {
	var derivationKey *hdkeychain.ExtendedKey
	var err error
	var derivationPath = &path

	if path.Hardened {
		derivationKey, err = masterKey.Child(hdkeychain.HardenedKeyStart + derivationPath.Path)
	} else {
		derivationKey, err = masterKey.Child(path.Path)
	}
	derivationPath = path.Child
	if err != nil {
		return nil, err
	}

	for derivationPath != nil {
		if derivationKey != nil {
			if derivationPath.Hardened {
				derivationKey, err = derivationKey.Child(hdkeychain.HardenedKeyStart + derivationPath.Path)
			} else {
				derivationKey, err = derivationKey.Child(derivationPath.Path)
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
