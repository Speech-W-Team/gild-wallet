package core

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
