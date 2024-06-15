package btc

import (
	"crypto/ecdsa"
	"gild-wallet/wallets"
	"gild-wallet/wallets/blockchains"
	"github.com/btcsuite/btcutil/base58"
	"github.com/ethereum/go-ethereum/crypto"
	"log"
)

func (btc *BTC) RestoreWallet(userPrivateKey []byte, networkType blockchains.NetworkType) (*blockchains.Wallet, error) {
	privKey, err := crypto.ToECDSA(userPrivateKey)
	if err != nil {
		return &blockchains.Wallet{}, err
	}
	privateKey := crypto.FromECDSA(privKey)
	publicKey := privKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	address, err := btc.GenerateAddress(append(publicKeyECDSA.X.Bytes(), publicKeyECDSA.Y.Bytes()...), networkType)

	return &blockchains.Wallet{
		Address:    address,
		PrivateKey: base58.Encode(privateKey),
		CryptoType: blockchains.BTC,
	}, nil
}

func (btc *BTC) RestoreWalletFromString(privateKey string, networkType blockchains.NetworkType) (*blockchains.Wallet, error) {
	return nil, nil
}

func (btc *BTC) RestoreWalletFromMnemonic(mnemonic string, password string, networkType blockchains.NetworkType) (*blockchains.Wallet, error) {
	paths := []wallets.DerivationPathItem{
		wallets.NewDerivationPath(44, true),
		wallets.NewDerivationPath(0, true),
		wallets.NewDerivationPath(0, true),
		{Hardened: false, Path: 0},
	}
	temp, _, err := wallets.HdWallet(mnemonic, password, paths)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

}
