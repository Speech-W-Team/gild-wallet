package btc

import (
	"crypto/sha256"
	"encoding/hex"
	"gild-wallet/core"
	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcutil/base58"
	"github.com/tyler-smith/go-bip39"
	"golang.org/x/crypto/ripemd160"
)

type BitcoinWalletManager struct {
	core.NetworkType
}

func (manager *BitcoinWalletManager) GenerateWallet(config core.BIPConfig) (*core.Wallet, string, error) {
	entropy, err := bip39.NewEntropy(128)
	if err != nil {
		return nil, "", err
	}
	mnemonic, err := bip39.NewMnemonic(entropy)
	if err != nil {
		return nil, "", err
	}

	wallet, err := manager.RestoreWalletFromMnemonic(mnemonic, "", config)
	return wallet, mnemonic, err
}

func (manager *BitcoinWalletManager) RestoreWallet(privateKey []byte, config core.BIPConfig) (*core.Wallet, error) {
	_, pubKey := btcec.PrivKeyFromBytes(btcec.S256(), privateKey)

	address, err := manager.GenerateAddress(pubKey.SerializeCompressed(), config)

	if err != nil {
		return nil, err
	}

	return &core.Wallet{
		Address:    address,
		PrivateKey: hex.EncodeToString(privateKey),
		Derivation: "",
	}, nil
}

func (manager *BitcoinWalletManager) RestoreWalletFromString(privateKey string, config core.BIPConfig) (*core.Wallet, error) {
	privKey, err := hex.DecodeString(privateKey)
	if err != nil {
		return nil, err
	}
	return manager.RestoreWallet(privKey, config)
}

func (manager *BitcoinWalletManager) RestoreWalletFromMnemonic(mnemonic string, password string, config core.BIPConfig) (*core.Wallet, error) {
	seed := core.SeedFromMnemonic(mnemonic, password)
	masterKey, err := core.MasterKeyFromSeed(seed, core.Mainnet)
	if err != nil {
		return nil, err
	}

	var path core.DerivationPathItem

	switch config {
	case core.BIP44:
		path = core.PathBip(44, 0, &core.WalletZeroPath)
	case core.BIP49:
		path = core.PathBip(49, 0, &core.WalletZeroPath)
	case core.BIP84:
		path = core.PathBip(84, 0, &core.WalletZeroPath)
	}

	privateKey, err := core.HDWallet(masterKey, path)
	if err != nil {
		return nil, err
	}

	publicKey, err := privateKey.ECPubKey()
	if err != nil {
		return nil, err
	}

	privateKeyEcdsa, err := privateKey.ECPrivKey()
	if err != nil {
		return nil, err
	}

	address, err := manager.GenerateAddress(publicKey.SerializeCompressed(), config)
	if err != nil {
		return nil, err
	}
	privKey := hex.EncodeToString(privateKeyEcdsa.D.Bytes())

	return &core.Wallet{
		Address:    address,
		PrivateKey: privKey,
		Derivation: "",
	}, nil
}

func (manager *BitcoinWalletManager) GenerateAddress(pubKey []byte, _ core.BIPConfig) (string, error) {
	var networkBytes []byte
	switch manager.NetworkType {
	case core.Mainnet:
		networkBytes = []byte("80")
	case core.Testnet:
		networkBytes = []byte("ef")
	}
	publicKey := append(networkBytes, pubKey...)
	hash256 := sha256.New()
	hash256.Write(publicKey)
	hashedPubKey := hash256.Sum(nil)
	hash256 = sha256.New()
	hash256.Write(hashedPubKey)
	hashedPubKey = hash256.Sum(nil)[:4]

	ripemd160Hasher := ripemd160.New()
	ripemd160Hasher.Write(hashedPubKey)
	hashedPubKeyRipemd160 := ripemd160Hasher.Sum(nil)

	fullPayload := append(hashedPubKeyRipemd160, hashedPubKey...)

	address := base58.Encode(fullPayload)
	return address, nil
}

func (manager *BitcoinWalletManager) GenerateAddressFromString(pubKey string, config core.BIPConfig) (string, error) {
	publicKey := base58.Decode(pubKey)
	return manager.GenerateAddress(publicKey, config)
}
