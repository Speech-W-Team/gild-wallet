package trx

import (
	"crypto/sha256"
	"encoding/hex"
	"gild-wallet/core"
	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcutil/base58"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/tyler-smith/go-bip39"
)

var pathBip44 = core.DerivationPathItem{
	Path:     44,
	Hardened: true,
	Child: &core.DerivationPathItem{
		Path:     195,
		Hardened: true,
		Child: &core.DerivationPathItem{
			Path:     0,
			Hardened: true,
			Child: &core.DerivationPathItem{
				Path:     0,
				Hardened: false,
				Child: &core.DerivationPathItem{
					Path:     0,
					Hardened: false,
					Child:    nil,
				},
			},
		},
	},
}

type TronWalletManager struct{}

func (manager *TronWalletManager) GenerateWallet() (*core.Wallet, string, error) {
	entropy, err := bip39.NewEntropy(128)
	if err != nil {
		return nil, "", err
	}
	mnemonic, err := bip39.NewMnemonic(entropy)
	if err != nil {
		return nil, "", err
	}

	wallet, err := manager.RestoreWalletFromMnemonic(mnemonic, "")
	return wallet, mnemonic, err
}

func (manager *TronWalletManager) RestoreWallet(privateKey []byte) (*core.Wallet, error) {
	_, pubKey := btcec.PrivKeyFromBytes(btcec.S256(), privateKey)

	address, err := manager.GenerateAddress(pubKey.SerializeCompressed())

	if err != nil {
		return nil, err
	}

	return &core.Wallet{
		Address:    address,
		PrivateKey: hex.EncodeToString(privateKey),
		Derivation: "",
	}, nil
}

func (manager *TronWalletManager) RestoreWalletFromString(privateKey string) (*core.Wallet, error) {
	privKey, err := hex.DecodeString(privateKey)
	if err != nil {
		return nil, err
	}
	return manager.RestoreWallet(privKey)
}

func (manager *TronWalletManager) RestoreWalletFromMnemonic(mnemonic string, password string) (*core.Wallet, error) {
	seed := core.SeedFromMnemonic(mnemonic, password)
	masterKey, err := core.MasterKeyFromSeed(seed)
	if err != nil {
		return nil, err
	}

	privateKey, err := core.HDWallet(masterKey, pathBip44)
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

	address, err := manager.GenerateAddress(publicKey.SerializeCompressed())
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

func (manager *TronWalletManager) GenerateAddressFromString(pubKey string) (string, error) {
	publicKey, err := hex.DecodeString(pubKey)
	if err != nil {
		return "", err
	}
	return manager.GenerateAddress(publicKey)
}

func (manager *TronWalletManager) GenerateAddress(pubKey []byte) (string, error) {
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
