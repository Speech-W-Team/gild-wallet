package eth

import (
	"crypto/ecdsa"
	"encoding/hex"
	_ "errors"
	"gild-wallet/wallets/blockchains"
	"log"

	"github.com/ethereum/go-ethereum/crypto"
)

type ETH struct{}

func NewETH() *ETH {
	return &ETH{}
}

func (eth *ETH) GenerateWallet(networkType blockchains.NetworkType) (blockchains.Wallet, error) {
	privKey, err := crypto.GenerateKey()
	if err != nil {
		return blockchains.Wallet{}, err
	}

	privateKey := crypto.FromECDSA(privKey)
	publicKey := privKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	return blockchains.Wallet{
		Address:    address,
		PrivateKey: hex.EncodeToString(privateKey),
		CryptoType: blockchains.ETH,
	}, err
}
