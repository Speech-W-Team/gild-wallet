package api

import (
	"fmt"
	"gild-wallet/blockchain/btc"
	"gild-wallet/blockchain/eth"
	"gild-wallet/blockchain/trx"
	"gild-wallet/core"
)

//export GenerateWalletWithMnemonic
func GenerateWalletWithMnemonic(blockchain string, mnemonic string, password string, network string, bipType uint32) (*core.Wallet, error) {
	networkType := core.NetworkType(network)

	walletManager, err := getWalletManager(blockchain, networkType)
	if err != nil {
		return nil, err
	}

	return walletManager.RestoreWalletFromMnemonic(mnemonic, password, core.BIPConfig(bipType))
}

//export GenerateWalletWithPrivateKey
func GenerateWalletWithPrivateKey(blockchain string, privateKey string, network string, bipType uint32) (*core.Wallet, error) {
	networkType := core.NetworkType(network)

	walletManager, err := getWalletManager(blockchain, networkType)
	if err != nil {
		return nil, err
	}

	return walletManager.RestoreWalletFromString(privateKey, core.BIPConfig(bipType))
}

//export GenerateWallet
func GenerateWallet(blockchain string, network string, bipType uint32) (*core.Wallet, string, error) {
	networkType := core.NetworkType(network)

	walletManager, err := getWalletManager(blockchain, networkType)
	if err != nil {
		return nil, "", err
	}

	return walletManager.GenerateWallet(core.BIPConfig(bipType))
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
