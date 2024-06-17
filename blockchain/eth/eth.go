package eth

import (
	"crypto/ecdsa"
	"encoding/hex"
	"gild-wallet/core"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/tyler-smith/go-bip39"
	"log"
)

type EthereumWalletManager struct {
	core.NetworkType
}

func (manager *EthereumWalletManager) GenerateWallet(config core.BIPConfig) (*core.Wallet, string, error) {
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

func (manager *EthereumWalletManager) RestoreWallet(privateKey []byte, _ core.BIPConfig) (*core.Wallet, error) {
	privKey, err := crypto.ToECDSA(privateKey)
	if err != nil {
		return nil, err
	}

	publicKey := privKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()

	return &core.Wallet{
		Address:    address,
		PrivateKey: hex.EncodeToString(privateKey),
	}, err
}

func (manager *EthereumWalletManager) RestoreWalletFromString(privateKey string, config core.BIPConfig) (*core.Wallet, error) {
	privKey, err := hex.DecodeString(privateKey)
	if err != nil {
		return nil, err
	}

	return manager.RestoreWallet(privKey, config)
}

func (manager *EthereumWalletManager) RestoreWalletFromMnemonic(mnemonic string, password string, config core.BIPConfig) (*core.Wallet, error) {
	seed := core.SeedFromMnemonic(mnemonic, password)
	masterKey, err := core.MasterKeyFromSeed(seed, manager.NetworkType)
	if err != nil {
		return nil, err
	}
	privateKey, err := core.HDWallet(masterKey, core.PathBip(44, 60, &core.WalletZeroPath))
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
	privKey := hex.EncodeToString(crypto.FromECDSA(privateKeyEcdsa.ToECDSA()))

	return &core.Wallet{
		Address:    address,
		PrivateKey: privKey,
		Derivation: "",
	}, nil
}

func (manager *EthereumWalletManager) GenerateAddress(pubKey []byte, _ core.BIPConfig) (string, error) {
	pubKeyECDSA, err := crypto.DecompressPubkey(pubKey)
	if err != nil {
		return "", nil
	}

	address := crypto.PubkeyToAddress(*pubKeyECDSA)
	return address.Hex(), nil
}

func (manager *EthereumWalletManager) GenerateAddressFromString(pubKey string, config core.BIPConfig) (string, error) {
	publicKey, err := hex.DecodeString(pubKey)
	if err != nil {
		return "", err
	}

	return manager.GenerateAddress(publicKey, config)
}
