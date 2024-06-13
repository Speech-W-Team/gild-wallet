package main

import (
	"fmt"
	"gild-wallet/api"
	"gild-wallet/wallets/blockchains"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil"
	"github.com/btcsuite/btcutil/base58"
)

func main() {
	//generateNewWallets()

	restoreWallets()

	//restoreWalletsViaMnemonic()
}

func generateNewWallets() {
	walletApi := api.WalletAPI{}
	cryptoBtcType := blockchains.BTC
	address, privateKey, err := walletApi.CreateWallet(cryptoBtcType)
	if err != nil {
		fmt.Println("Error creating wallet:", err)
	}
	fmt.Println(cryptoBtcType, "Address:", address)
	fmt.Println(cryptoBtcType, "Private Key:", privateKey)

	cryptoTronType := blockchains.TRON
	address, privateKey, err = walletApi.CreateWallet(blockchains.TRON)
	if err != nil {
		fmt.Println("Error creating wallet:", err)
	}
	fmt.Println(cryptoTronType, "TRON Address:", address)
	fmt.Println(cryptoTronType, "TRON Private Key:", privateKey)
	cryptoETHType := blockchains.ETH
	address, privateKey, err = walletApi.CreateWallet(blockchains.ETH)
	if err != nil {
		fmt.Println("Error creating wallet:", err)
	}
	fmt.Println(cryptoETHType, "Address:", address)
	fmt.Println(cryptoETHType, "Private Key:", privateKey)
}

func restoreWallets() {
	println("————————— Restore")
	btcAddress := "1Hoj1vV6TPZnMC4L8h5amdj7HASQsWUf6n"
	btcPrivateKey := "KxukPyzudG2eFUCJcnMeUBKsboLiVmrrALBgvQyNe1T4kRuXxxKz"
	privKey := base58.Decode(btcPrivateKey)
	wif, err := btcutil.DecodeWIF(btcPrivateKey)
	if err != nil {
		fmt.Println("Error decoding BTC wif private key:", err)
	}
	privKey = wif.PrivKey.D.Bytes()
	addressPubKey, err := btcutil.NewAddressPubKey(wif.SerializePubKey(), &chaincfg.MainNetParams)
	if err != nil {
		fmt.Println("Error decoding BTC public key:", err)
	}
	address := addressPubKey.EncodeAddress()
	walletApi := api.WalletAPI{}
	_, privateKey, err := walletApi.RestoreWallet(privKey, blockchains.BTC)
	if err != nil {
		fmt.Println("Error restoring wallet:", err)
	}
	if address != btcAddress {
		fmt.Println("Error restoring BTC wallet: address mismatch")
	}
	fmt.Println("BTC restored Address:", address)
	fmt.Println("BTC restored Private Key:", privateKey)

	//ethAddress := "0xAAd58655708a93afB3a5529cF8c53840321b8C9d"
	//ethPrivateKey := "1046d074171e5892869739868de43242409f6d2facad9917f63a0a01e60683c5"
	//privKey, err = hex.DecodeString(ethPrivateKey)
	//if err != nil {
	//	fmt.Println("Error decoding ETH private key:", err)
	//}
	//address, privateKey, err = walletApi.RestoreWallet(privKey, blockchains.ETH)
	//if err != nil {
	//	fmt.Println("Error restoring wallet:", err)
	//}
	//if address != ethAddress {
	//	fmt.Println("Error restoring eth wallet: address mismatch")
	//}
	//fmt.Println("ETH restored Address:", address)
	//fmt.Println("ETH restored Private Key:", privateKey)
	//
	//tronAddress := "THk8sMZwtwFmEWjTzmAvNNSBtXiqm7YknU"
	//tronPrivateKey := "6991f30d97035525aa143d2a85f51f833d6e465df0665f8971fd0626b69015de"
	//privKey, err = hex.DecodeString(tronPrivateKey)
	//if err != nil {
	//	fmt.Println("Error decoding TRON private key:", err)
	//}
	//address, privateKey, err = walletApi.RestoreWallet(privKey, blockchains.TRON)
	//if err != nil {
	//	fmt.Println("Error restoring tron wallet:", err)
	//}
	//if address != tronAddress {
	//	fmt.Println("Error restoring tron wallet: address mismatch")
	//}
	//fmt.Println("TRON restored Address:", address)
	//fmt.Println("TRON restored Private Key:", privateKey)
}

func restoreWalletsViaMnemonic() {
	println("————————— Restore Via Mnemonic")
	println("————————— Restore Bitcoin Via Mnemonic")
	btcAddress := "1EMJpg4FR2RYfGrDQxda1HVKmgGNyrnvUE"
	ethAddress := "0x107B57cDF9F2309d0e62c402AaDc863Fc97B0cB3"
	tronAddress := "TWxWU4brH7GkL1hiyXGjLbbs6YDvah7zmd"
	mnemonicPhrase := "other budget write video mimic captain cargo anger emerge chalk neck series"
	walletApi := api.WalletAPI{}
	//address, err := walletApi.RestorePrivateKeyViaMnemonic(mnemonicPhrase, "", blockchains.BTC)
	address, privateKey, err := walletApi.RestorePrivateKeyViaMnemonic(mnemonicPhrase, "", blockchains.BTC)
	if err != nil {
		fmt.Println("Error restoring wallet:", err)
	}
	if address != btcAddress {
		fmt.Println("Error restoring BTC wallet: address mismatch")
	}
	fmt.Println("BTC restored Address:", address)
	fmt.Println("BTC restored Private Key:", privateKey)

	println("————————— Restore Ethereum Via Mnemonic")
	address, privateKey, err = walletApi.RestorePrivateKeyViaMnemonic(mnemonicPhrase, "", blockchains.ETH)
	if err != nil {
		fmt.Println("Error restoring wallet:", err)
	}
	if address != ethAddress {
		fmt.Println("Error restoring eth wallet: address mismatch")
	}
	fmt.Println("ETH restored Address:", address)
	fmt.Println("ETH restored Private Key:", privateKey)

	println("————————— Restore TRON Via Mnemonic")
	address, privateKey, err = walletApi.RestorePrivateKeyViaMnemonic(mnemonicPhrase, "", blockchains.TRON)
	if err != nil {
		fmt.Println("Error restoring tron wallet:", err)
	}
	if address != tronAddress {
		fmt.Println("Error restoring tron wallet: address mismatch")
	}
	fmt.Println("TRON restored Address:", address)
	fmt.Println("TRON restored Private Key:", privateKey)
}
