package core

import (
	"github.com/btcsuite/btcutil/hdkeychain"
)

type BIPConfig int

const (
	BIP44 BIPConfig = iota
	BIP49 BIPConfig = iota
	BIP84 BIPConfig = iota
)

type NetworkType string

const (
	Mainnet NetworkType = "mainnet"
	Testnet NetworkType = "testnet"
)

type WalletManager interface {
	GenerateWallet(config BIPConfig) (*Wallet, string, error)

	RestoreWallet(privateKey []byte, config BIPConfig) (*Wallet, error)
	RestoreWalletFromString(privateKey string, config BIPConfig) (*Wallet, error)
	RestoreWalletFromMnemonic(mnemonic string, password string, config BIPConfig) (*Wallet, error)

	GenerateAddress(pubKey []byte, config BIPConfig) (string, error)
	GenerateAddressFromString(pubKey string, config BIPConfig) (string, error)
}

var WalletZeroPath = DerivationPathItem{}

func PathBip(rootPath uint32, coinPath uint32, walletPath *DerivationPathItem) DerivationPathItem {
	return DerivationPathItem{
		Path:     rootPath,
		Hardened: true,
		Child: &DerivationPathItem{
			Path:     coinPath,
			Hardened: true,
			Child: &DerivationPathItem{
				Path:     0,
				Hardened: true,
				Child: &DerivationPathItem{
					Path:     0,
					Hardened: false,
					Child:    walletPath,
				},
			},
		},
	}
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
	NetworkType
	MasterKey *hdkeychain.ExtendedKey
	Mnemonic  string
	Password  string
}

func NewWalletBuilder(mnemonic string, password string, networkType NetworkType) *WalletBuilder {
	seed := SeedFromMnemonic(mnemonic, password)
	masterKey, err := MasterKeyFromSeed(seed, networkType)
	if err != nil {
		return nil
	}
	return &WalletBuilder{
		NetworkType: networkType,
		MasterKey:   masterKey,
		Mnemonic:    mnemonic,
		Password:    password,
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
	return walletManager.RestoreWallet(privateKey, BIP84)
}
