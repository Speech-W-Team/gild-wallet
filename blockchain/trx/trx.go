package trx

import (
	"crypto/sha256"
	"encoding/hex"
	"gild-wallet/core"
	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcutil/base58"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/tyler-smith/go-bip39"
)

type TronWalletManager struct {
	core.NetworkType
}

func (manager *TronWalletManager) GenerateWallet(config core.BIPConfig) (*core.Wallet, string, error) {
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

func (manager *TronWalletManager) RestoreWallet(privateKey []byte, config core.BIPConfig) (*core.Wallet, error) {
	_, pubKey := btcec.PrivKeyFromBytes(privateKey)

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

func (manager *TronWalletManager) RestoreWalletFromString(privateKey string, config core.BIPConfig) (*core.Wallet, error) {
	privKey, err := hex.DecodeString(privateKey)
	if err != nil {
		return nil, err
	}
	return manager.RestoreWallet(privKey, config)
}

func (manager *TronWalletManager) RestoreWalletFromMnemonic(mnemonic string, password string, config core.BIPConfig) (*core.Wallet, error) {
	seed := core.SeedFromMnemonic(mnemonic, password)
	masterKey, err := core.MasterKeyFromSeed(seed, manager.NetworkType)
	if err != nil {
		return nil, err
	}

	privateKey, err := core.HDWallet(masterKey, core.PathBip(core.BIP44, 195, &core.WalletZeroPath))
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
	privKey := hex.EncodeToString(privateKeyEcdsa.ToECDSA().D.Bytes())

	return &core.Wallet{
		Address:    address,
		PrivateKey: privKey,
		Derivation: "",
	}, nil
}

func (manager *TronWalletManager) GenerateAddressFromString(pubKey string, config core.BIPConfig) (string, error) {
	publicKey, err := hex.DecodeString(pubKey)
	if err != nil {
		return "", err
	}
	return manager.GenerateAddress(publicKey, config)
}

func (manager *TronWalletManager) GenerateAddress(pubKey []byte, _ core.BIPConfig) (string, error) {
	publicKeyEcdsa, err := crypto.DecompressPubkey(pubKey)
	if err != nil {
		return "", err
	}

	publicKey := crypto.PubkeyToAddress(*publicKeyEcdsa).Bytes()

	initialAddress := append([]byte{0x41}, publicKey...)

	hash2 := sha256.New()
	hash2.Write(initialAddress)
	hash := hash2.Sum(nil)

	hash2 = sha256.New()
	hash2.Write(hash)
	hash = hash2.Sum(nil)

	verificationCode := hash[:4]
	initialAddress = append(initialAddress, verificationCode...)

	return base58.Encode(initialAddress), nil
}
