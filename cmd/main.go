package main

import (
	"gild-wallet/internal/wallet"
)

func main() {
	btc, err := wallet.NewWallet(wallet.BTC)

	if err != nil {
		panic(err)
	}

	println(btc.Address)

	eth, err := wallet.NewWallet(wallet.ETH)

	if err != nil {
		panic(err)
	}

	println(eth.Address)
}
