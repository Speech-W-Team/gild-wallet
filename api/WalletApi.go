package api

import (
	"fmt"
	"gild-wallet/blockchain/btc"
	"gild-wallet/blockchain/eth"
	"gild-wallet/blockchain/trx"
	"gild-wallet/core"
)

//export GenerateWalletWithMnemonic
func GenerateWalletWithMnemonic(blockchain string, mnemonic string, password string, network string, bipType int) (*core.Wallet, error) {
	networkType := core.NetworkType(network)

	config, err := getBIPConfig(bipType)
	if err != nil {
		return nil, err
	}

	walletManager, err := getWalletManager(blockchain, networkType)
	if err != nil {
		return nil, err
	}

	return walletManager.RestoreWalletFromMnemonic(mnemonic, password, config)
}

//export GenerateWalletWithPrivateKey
func GenerateWalletWithPrivateKey(blockchain string, privateKey string, network string, bipType int) (*core.Wallet, error) {
	networkType := core.NetworkType(network)

	config, err := getBIPConfig(bipType)
	if err != nil {
		return nil, err
	}

	walletManager, err := getWalletManager(blockchain, networkType)
	if err != nil {
		return nil, err
	}

	return walletManager.RestoreWalletFromString(privateKey, config)
}

//export GenerateWallet
func GenerateWallet(blockchain string, network string, bipType int) (*core.Wallet, string, error) {
	networkType := core.NetworkType(network)

	config, err := getBIPConfig(bipType)
	if err != nil {
		return nil, "", err
	}

	walletManager, err := getWalletManager(blockchain, networkType)
	if err != nil {
		return nil, "", err
	}

	return walletManager.GenerateWallet(config)
}

func getBIPConfig(bipType int) (core.BIPConfig, error) {
	var config core.BIPConfig
	switch bipType {
	case 44:
		config = core.BIP44
	case 49:
		config = core.BIP49
	case 84:
		config = core.BIP84
	default:
		return 0, fmt.Errorf("invalid BIP type")
	}
	return config, nil
}

func getWalletManager(blockchain string, networkType core.NetworkType) (core.WalletManager, error) {
	var walletManager core.WalletManager
	switch blockchain {
	case "BTC", "Bitcoin":
		walletManager = &btc.BitcoinWalletManager{NetworkType: networkType}
	case "ETH", "Ethereum":
		walletManager = &eth.EthereumWalletManager{NetworkType: networkType}
	case "TRX", "Tron":
		walletManager = &trx.TronWalletManager{NetworkType: networkType}
	default:
		return nil, fmt.Errorf("unsupported blockchain type %s", blockchain)
	}
	return walletManager, nil
}
