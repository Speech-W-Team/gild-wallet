package eth

import (
	"crypto/ecdsa"
	"encoding/hex"
	"gild-wallet/wallets/blockchains"
	"github.com/ethereum/go-ethereum/crypto"
	"golang.org/x/crypto/sha3"
)

func (eth *ETH) GenerateAddress(pubKey []byte, networkType blockchains.NetworkType) (string, error) {
	pubKeyECDSA, err := crypto.DecompressPubkey(pubKey)
	if err != nil {
		return "", nil
	}

	address := crypto.PubkeyToAddress(*pubKeyECDSA)
	return address.Hex(), nil
}

func (eth *ETH) GenerateAddressFromString(pubKey string, networkType blockchains.NetworkType) (string, error) {

}

// Temporarily disabled
func generateAddress(pubKey *ecdsa.PublicKey) string {
	pubKeyBytes := append(pubKey.X.Bytes(), pubKey.Y.Bytes()...)
	hash := sha3.NewLegacyKeccak256()
	hash.Write(pubKeyBytes[1:])
	address := hash.Sum(nil)[12:]
	return "0x" + hex.EncodeToString(address)
}
