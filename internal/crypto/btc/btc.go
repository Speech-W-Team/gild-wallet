package btc

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	_ "errors"
	"math/big"

	"golang.org/x/crypto/ripemd160"
)

func GenerateWallet() (string, string, error) {
	privKey, pubKey, err := generateKeyPair()
	if err != nil {
		return "", "", err
	}

	address, err := generateAddress(pubKey)
	if err != nil {
		return "", "", err
	}

	privateKey := encodePrivateKey(privKey)
	return address, privateKey, nil
}

func generateKeyPair() (*ecdsa.PrivateKey, *ecdsa.PublicKey, error) {
	privKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return nil, nil, err
	}

	pubKey := &privKey.PublicKey
	return privKey, pubKey, nil
}

func generateAddress(pubKey *ecdsa.PublicKey) (string, error) {
	pubKeyBytes := append(pubKey.X.Bytes(), pubKey.Y.Bytes()...)
	hash256 := sha256.New()
	hash256.Write(pubKeyBytes)
	hashedPubKey := hash256.Sum(nil)

	ripemd160Hasher := ripemd160.New()
	ripemd160Hasher.Write(hashedPubKey)
	hashedPubKeyRipemd160 := ripemd160Hasher.Sum(nil)

	versionedPayload := append([]byte{0x00}, hashedPubKeyRipemd160...)

	hash256 = sha256.New()
	hash256.Write(versionedPayload)
	checksum := hash256.Sum(nil)

	hash256 = sha256.New()
	hash256.Write(checksum)
	checksum = hash256.Sum(nil)[:4]

	fullPayload := append(versionedPayload, checksum...)

	address := base58Encode(fullPayload)

	return address, nil
}

func encodePrivateKey(privKey *ecdsa.PrivateKey) string {
	return base58Encode(privKey.D.Bytes())
}

func base58Encode(input []byte) string {
	alphabet := "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"
	b58 := make([]byte, 0, len(input)*138/100+1)

	x := new(big.Int).SetBytes(input)
	mod := new(big.Int)
	zero := big.NewInt(0)
	base := big.NewInt(int64(len(alphabet)))

	for x.Cmp(zero) != 0 {
		x.DivMod(x, base, mod)
		b58 = append(b58, alphabet[mod.Int64()])
	}

	for _, b := range input {
		if b == 0x00 {
			b58 = append(b58, alphabet[0])
		} else {
			break
		}
	}

	for i := len(b58)/2 - 1; i >= 0; i-- {
		opp := len(b58) - 1 - i
		b58[i], b58[opp] = b58[opp], b58[i]
	}

	return string(b58)
}
