package cryptography

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/rand"
	"github.com/ethereum/go-ethereum/crypto"
	"math/big"
)

func GenerateKeyPair() (*ecdsa.PrivateKey, *ecdsa.PublicKey, error) {
	privKey, err := ecdsa.GenerateKey(crypto.S256(), rand.Reader)
	if err != nil {
		return nil, nil, err
	}

	pubKey := &privKey.PublicKey
	return privKey, pubKey, nil
}

func Base58Encode(input []byte) string {
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

func Base58Decode(input []byte) []byte {
	alphabet := "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"

	result := big.NewInt(0)
	for _, b := range input {
		charIndex := bytes.IndexByte([]byte(alphabet), b)
		if charIndex == -1 {
			panic("Invalid character for Base58 encoding")
		}
		result.Mul(result, big.NewInt(58))
		result.Add(result, big.NewInt(int64(charIndex)))
	}

	decoded := result.Bytes()
	// Add leading zeros back in.
	for i := range input {
		if i > 0 && input[i] != '1' {
			break
		}
		decoded = append([]byte{0x00}, decoded...)
	}

	return decoded
}
