package btc

import (
	"crypto/sha256"
	"gild-wallet/cryptography"
	"gild-wallet/wallets/blockchains"
	"golang.org/x/crypto/ripemd160"
)

func (btc *BTC) GenerateAddress(pubKey []byte, networkType blockchains.NetworkType) (string, error) {
	var networkBytes []byte
	switch networkType {
	case blockchains.Mainnet:
		networkBytes = []byte("80")
	case blockchains.Testnet:
		networkBytes = []byte("ef")
	}
	publicKey := append(networkBytes, pubKey...)
	//publicKey = append(publicKey, []byte("01")...)
	hash256 := sha256.New()
	hash256.Write(publicKey)
	hashedPubKey := hash256.Sum(nil)
	hash256 = sha256.New()
	hash256.Write(hashedPubKey)
	hashedPubKey = hash256.Sum(nil)[:4]

	ripemd160Hasher := ripemd160.New()
	ripemd160Hasher.Write(hashedPubKey)
	hashedPubKeyRipemd160 := ripemd160Hasher.Sum(nil)

	//versionedPayload := append([]byte{0x80}, hashedPubKeyRipemd160...)

	//hash256 = sha256.New()
	//hash256.Write(versionedPayload)
	//checksum := hash256.Sum(nil)
	//
	//hash256 = sha256.New()
	//hash256.Write(checksum)
	//checksum = hash256.Sum(nil)[:4]

	fullPayload := append(hashedPubKeyRipemd160, hashedPubKey...)

	address := cryptography.Base58Encode(fullPayload)

	return address, nil
}

func generateAddress(pubKey []byte) (string, error) {

	// Step 1: Perform SHA-256 hashing on the public key
	//hashedPubKey := crypto.Keccak256Hash(pubKey).Bytes()

	// Step 2: Perform RIPEMD-160 hashing on the result of SHA-256
	hash160 := ripemd160.New()
	hash160.Write(pubKey)
	publicRIPEMD160 := hash160.Sum(nil)

	// Step 3: Add version byte in front of RIPEMD-160 hash (0x00 for Main Network)
	versionedPayload := append([]byte{0x00}, publicRIPEMD160...)

	// Step 4: Perform SHA-256 hash on the extended RIPEMD-160 result
	firstSHA := sha256.New()
	firstSHA.Write(versionedPayload)
	firstHash := firstSHA.Sum(nil)

	// Step 5: Perform SHA-256 hash on the result of the previous SHA-256 hash
	secondSHA := sha256.New()
	secondSHA.Write(firstHash)
	secondHash := secondSHA.Sum(nil)

	// Step 6: Take the first 4 bytes of the second SHA-256 hash. This is the address checksum
	addressChecksum := secondHash[:4]

	// Step 7: Add the 4 checksum bytes from stage 7 at the end of extended RIPEMD-160 hash from stage 4. This is the 25-byte binary Bitcoin Address.
	binaryAddress := append(versionedPayload, addressChecksum...)

	// Convert binary to a Base58Check encoded string
	walletAddress := cryptography.Base58Encode(binaryAddress)
	return walletAddress, nil
}
