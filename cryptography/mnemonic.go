package cryptography

import (
	"errors"
	"github.com/tyler-smith/go-bip39"
)

type MnemonicApi struct{}

var ErrMnemonicRequired = errors.New("mnemonic is required")
var ErrMnemonicInvalid = errors.New("invalid mnemonic")

func (mnemonic *MnemonicApi) NewSeedFromMnemonic(mnemonicPhrase string, passphrase string) ([]byte, error) {
	if mnemonicPhrase == "" {
		return nil, ErrMnemonicRequired
	}
	if !bip39.IsMnemonicValid(mnemonicPhrase) {
		return nil, ErrMnemonicInvalid
	}

	seed, err := bip39.NewSeedWithErrorChecking(mnemonicPhrase, passphrase)
	if err != nil {
		return nil, err
	}

	return seed, nil
}

func (mnemonic *MnemonicApi) NewEntropy(bitsSize int) ([]byte, error) {
	return bip39.NewEntropy(bitsSize)
}

func (mnemonic *MnemonicApi) NewMnemonicFromEntropy(entropy []byte) (string, error) {
	return bip39.NewMnemonic(entropy)
}

func (mnemonic *MnemonicApi) NewMnemonic(bitsSize int) (string, error) {
	entropy, err := mnemonic.NewEntropy(bitsSize)
	if err != nil {
		return "", err
	}
	return mnemonic.NewMnemonicFromEntropy(entropy)
}
