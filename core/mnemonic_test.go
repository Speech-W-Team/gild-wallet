package core

import (
	"encoding/hex"
	"testing"
)

const MasterKeyBIP44Expected = "xprv9s21ZrQH143K4JeCpHjRXGrTHY2rqLH6NCSmMTsyGUJrQ4jFy1cGrp1iNccyXQ9eM1Y8cLtK5bjSX9x92BcKJJZagehuaWfgPUmrw8SaFGf"
const SeedExpected = "c0f5f121ad57f4ee9499c94ea873f92768edcd19a466acb516884ee4902987511d3e0c1e3b0497d901df69a91971113968da21f2f09a2b9033daac1bbceb9455"
const Mnemonic = "other budget write video mimic captain cargo anger emerge chalk neck series"
const PrivateTRXKeyExpected = "xprvA2748ftwe6ivbVP1EYgMT2m1zYZ1Dvb2fZJUWusSxB1jgrPjTrHddx35cXs4HtGL47Mi3k1mWLKCmYAE5jPZhacbhEzMruvSb97Arw9UNgg"
const PrivateETHKeyExpected = "xprvA24oktLuUyiMiw3GRrA8NTxLToF3ZAUh3ozGgkHsG31jVaMoZRvdDYqvn9S7qDm3nGasySSfCpU9CBu2TeCpaWLSv94QYNcQPVAjPE6SSDg"
const PrivateBTCBIP44KeyExpected = "xprvA2HswdqmdjVDBFnXQ84XU8n9T2rnkNuSLNodKH4pK78i3UGARtTnGaeGd1EqJjSvyZ5cjHWiVJv8sh4HdiCfQ8PNEuxD8rHPGM9dKh4yiFz"

func TestTRXHDWalletFromMnemonic(t *testing.T) {
	seed := SeedFromMnemonic(Mnemonic, "")
	if SeedExpected != hex.EncodeToString(seed) {
		t.Fatalf("seed != %s, got %s", SeedExpected, seed)
	}
	masterKey, err := MasterKeyFromSeed(seed, Mainnet)
	if err != nil {
		t.Fatal(err)
	}
	if masterKey.String() != MasterKeyBIP44Expected {
		t.Fatalf("masterKey != %s, got %s", MasterKeyBIP44Expected, masterKey)
	}
	derivatedPrivateKey, err := HDWallet(
		masterKey,
		DerivationPathItem{
			Path:     44,
			Hardened: true,
			Child: &DerivationPathItem{
				Path:     195,
				Hardened: true,
				Child: &DerivationPathItem{
					0,
					true,
					&DerivationPathItem{
						0,
						false,
						nil,
					},
				},
			},
		},
	)
	if err != nil {
		t.Fatal(err)
	}

	if derivatedPrivateKey.String() != PrivateTRXKeyExpected {
		t.Fatalf("walletAddress != %s, got %s", PrivateTRXKeyExpected, derivatedPrivateKey.String())
	}
}

func TestETHHDWalletFromMnemonic(t *testing.T) {
	seed := SeedFromMnemonic(Mnemonic, "")
	if SeedExpected != hex.EncodeToString(seed) {
		t.Fatalf("seed != %s, got %s", SeedExpected, seed)
	}
	masterKey, err := MasterKeyFromSeed(seed, Mainnet)
	if err != nil {
		t.Fatal(err)
	}
	if masterKey.String() != MasterKeyBIP44Expected {
		t.Fatalf("masterKey != %s, got %s", MasterKeyBIP44Expected, masterKey)
	}
	derivatedPrivateKey, err := HDWallet(
		masterKey,
		DerivationPathItem{
			Path:     44,
			Hardened: true,
			Child: &DerivationPathItem{
				Path:     60,
				Hardened: true,
				Child: &DerivationPathItem{
					0,
					true,
					&DerivationPathItem{
						0,
						false,
						nil,
					},
				},
			},
		},
	)
	if err != nil {
		t.Fatal(err)
	}

	if derivatedPrivateKey.String() != PrivateETHKeyExpected {
		t.Fatalf("walletAddress != %s, got %s", PrivateETHKeyExpected, derivatedPrivateKey.String())
	}
}

func TestBTCBIP44HDWalletFromMnemonic(t *testing.T) {
	seed := SeedFromMnemonic(Mnemonic, "")
	if SeedExpected != hex.EncodeToString(seed) {
		t.Fatalf("seed != %s, got %s", SeedExpected, seed)
	}
	masterKey, err := MasterKeyFromSeed(seed, Mainnet)
	if err != nil {
		t.Fatal(err)
	}
	if masterKey.String() != MasterKeyBIP44Expected {
		t.Fatalf("masterKey != %s, got %s", MasterKeyBIP44Expected, masterKey)
	}
	derivatedPrivateKey, err := HDWallet(
		masterKey,
		DerivationPathItem{
			Path:     44,
			Hardened: true,
			Child: &DerivationPathItem{
				Path:     0,
				Hardened: true,
				Child: &DerivationPathItem{
					0,
					true,
					&DerivationPathItem{
						0,
						false,
						nil,
					},
				},
			},
		},
	)
	if err != nil {
		t.Fatal(err)
	}

	if derivatedPrivateKey.String() != PrivateBTCBIP44KeyExpected {
		t.Fatalf("walletAddress != %s, got %s", PrivateBTCBIP44KeyExpected, derivatedPrivateKey.String())
	}
}
