package core

import (
	"github.com/btcsuite/btcutil/hdkeychain"
)

type WalletManager interface {
	GenerateWallet() (*Wallet, string, error)

	RestoreWallet(privateKey []byte) (*Wallet, error)
	RestoreWalletFromString(privateKey string) (*Wallet, error)
	RestoreWalletFromMnemonic(mnemonic string, password string) (*Wallet, error)

	GenerateAddress(pubKey []byte) (string, error)
	GenerateAddressFromString(pubKey string) (string, error)
}

type DerivationPathItem struct {
	Path     uint32
	Hardened bool
	Child    *DerivationPathItem
}

type Wallet struct {
	Address    string
	PrivateKey string
	Derivation string
}

type WalletBuilder struct {
	MasterKey *hdkeychain.ExtendedKey
	Mnemonic  string
	Password  string
}

func NewWalletBuilder(mnemonic string, password string) *WalletBuilder {
	seed := SeedFromMnemonic(mnemonic, password)
	masterKey, err := MasterKeyFromSeed(seed)
	if err != nil {
		return nil
	}
	return &WalletBuilder{
		MasterKey: masterKey,
		Mnemonic:  mnemonic,
		Password:  password,
	}
}

func (builder *WalletBuilder) BuildWalletForDerivationPath(walletManager WalletManager, path DerivationPathItem) (*Wallet, error) {
	walletKey, err := HDWallet(builder.MasterKey, path)
	if err != nil {
		return nil, err
	}

	ecdsaKey, err := walletKey.ECPrivKey()
	if err != nil {
		return nil, err
	}

	privateKey := ecdsaKey.ToECDSA().D.Bytes()

	return walletManager.RestoreWallet(privateKey)
}
