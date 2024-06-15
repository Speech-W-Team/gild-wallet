package tron

import (
	"crypto/ecdsa"
	"encoding/hex"
	"gild-wallet/wallets/blockchains"
	"github.com/ethereum/go-ethereum/crypto"
	"log"
)

func (tron *TRON) RestoreWallet(userPrivateKey []byte, networkType blockchains.NetworkType) (blockchains.Wallet, error) {
	privKey, err := crypto.ToECDSA(userPrivateKey)
	if err != nil {
		return blockchains.Wallet{}, err
	}

	privateKey := crypto.FromECDSA(privKey)
	publicKey := privKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	address, err := tron.GenerateAddress(append(publicKeyECDSA.X.Bytes(), publicKeyECDSA.Y.Bytes()...), networkType)
	if err != nil {
		return blockchains.Wallet{}, err
	}
	return blockchains.Wallet{Address: address, PrivateKey: hex.EncodeToString(privateKey), CryptoType: blockchains.TRON}, nil
}

func (tron *TRON) RestoreWalletFromString(privateKey string, networkType blockchains.NetworkType) (blockchains.Wallet, error) {

}

func (tron *TRON) RestoreWalletFromMnemonic(mnemonic string, password string, networkType blockchains.NetworkType) (blockchains.Wallet, error) {

}
