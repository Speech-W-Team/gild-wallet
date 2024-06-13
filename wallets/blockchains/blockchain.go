package blockchains

type NetworkType string

const (
	Mainnet NetworkType = "mainnet"
	Testnet NetworkType = "testnet"
)

type CryptoType string

const (
	BTC   CryptoType = "BTC"
	ETH              = "ETH"
	TRON             = "TRON"
	SOL              = "SOL"
	BSC              = "BSC"
	ARB              = "ARB"
	AVAX             = "AVAX"
	MATIC            = "MATIC"
	SUI              = "SUI"
	OP               = "OP"
)

// Wallet represents a cryptocurrency wallet.
type Wallet struct {
	Address    string
	PrivateKey string
	CryptoType CryptoType
}

// Cryptocurrency defines the interface for a cryptocurrency.
type Cryptocurrency interface {
	GenerateWallet(networkType NetworkType) (Wallet, error)
	RestoreWallet(privateKey []byte, networkType NetworkType) (Wallet, error)
	GenerateAddress(pubKey []byte, networkType NetworkType) (string, error)
}

// Transaction represents a cryptocurrency transaction.
type Transaction struct {
	From   string
	To     string
	Amount float64
	Fee    float64
	Data   string
}
