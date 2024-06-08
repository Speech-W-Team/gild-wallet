package main

import (
	"gild-wallet/internal/wallet"
)

func main() {
	wallet, err := wallet.NewWallet(wallet.BTC)

	if err != nil {
		panic(err)
	}

	println(wallet.Address)
}
