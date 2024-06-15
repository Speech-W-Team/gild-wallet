package eth

import (
	"crypto/ecdsa"
	"encoding/hex"
	"gild-wallet/wallets/blockchains"
	"github.com/ethereum/go-ethereum/crypto"
	"log"
)

func (eth *ETH) RestoreWallet(userPrivateKey []byte, networkType blockchains.NetworkType) (blockchains.Wallet, error) {
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

	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	return blockchains.Wallet{
		Address:    address,
		PrivateKey: hex.EncodeToString(privateKey),
		CryptoType: blockchains.ETH,
	}, err
}

func (eth *ETH) RestoreWalletFromString(privateKey string, networkType blockchains.NetworkType) (blockchains.Wallet, error) {

}

func (eth *ETH) RestoreWalletFromMnemonic(mnemonic string, password string, networkType blockchains.NetworkType) (blockchains.Wallet, error) {

}
