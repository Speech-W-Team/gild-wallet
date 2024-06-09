package main

import (
	"fmt"
	"gild-wallet/internal/api"
	"gild-wallet/internal/wallets"
)

func main() {
	generateNewWallets()

	restoreWallets()
}

func generateNewWallets() {
	walletApi := api.WalletAPI{}
	cryptoBtcType := wallets.BTC
	address, privateKey, err := walletApi.CreateWallet(cryptoBtcType)
	if err != nil {
		fmt.Println("Error creating wallet:", err)
	}
	fmt.Println("BTC Address:", address)
	fmt.Println("BTC Private Key:", privateKey)

	cryptoTronType := wallets.TRON
	address, privateKey, err = walletApi.CreateWallet(cryptoTronType)
	if err != nil {
		fmt.Println("Error creating wallet:", err)
	}
	fmt.Println("TRON Address:", address)
	fmt.Println("TRON Private Key:", privateKey)
	cryptoETHType := wallets.ETH
	address, privateKey, err = walletApi.CreateWallet(cryptoETHType)
	if err != nil {
		fmt.Println("Error creating wallet:", err)
	}
	fmt.Println("ETH Address:", address)
	fmt.Println("ETH Private Key:", privateKey)
}

func restoreWallets() {
	ethAddress := "0xAAd58655708a93afB3a5529cF8c53840321b8C9d"
	ethPrivateKey := "1046d074171e5892869739868de43242409f6d2facad9917f63a0a01e60683c5"
	walletApi := api.WalletAPI{}
	address, privateKey, err := walletApi.RestoreWallet(ethPrivateKey, wallets.ETH)
	if err != nil {
		fmt.Println("Error restoring wallet:", err)
	}
	if address != ethAddress {
		fmt.Println("Error restoring eth wallet: address mismatch")
	}
	fmt.Println("ETH restored Address:", address)
	fmt.Println("ETH restored Private Key:", privateKey)

	tronAddress := "TC2KjLhLhfRN4xkMWTgLvAMTVYdSHCdyFD"
	tronPrivateKey := "464c36f56f94caf321fefe5b5fc153dd593fb95be76ee9c28e1bcccdeded5ddb"
	address, privateKey, err = walletApi.RestoreWallet(tronPrivateKey, wallets.TRON)
	if err != nil {
		fmt.Println("Error restoring tron wallet:", err)
	}
	if address != tronAddress {
		fmt.Println("Error restoring tron wallet: address mismatch")
	}
	fmt.Println("TRON restored Address:", address)
	fmt.Println("TRON restored Private Key:", privateKey)
}

// ETH Private key: 856072316d308a5adc12dd9ad000fba6c6757a21e48e434fee7d34974bb6079e
