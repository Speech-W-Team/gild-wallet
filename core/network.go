package core

// NetworkConfig represents the network configuration for a blockchain
type NetworkConfig struct {
	Name    string
	RPCURL  string
	ChainID int
}

// NetworkManager interface for network management
type NetworkManager interface {
	SetNetwork(config NetworkConfig) error
	GetNetwork() NetworkConfig
}
