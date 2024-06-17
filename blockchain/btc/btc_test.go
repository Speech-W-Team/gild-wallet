package btc

import (
	"encoding/hex"
	"gild-wallet/core"
	"testing"
)

const MasterKeyBIP44Expected = "xprv9s21ZrQH143K4JeCpHjRXGrTHY2rqLH6NCSmMTsyGUJrQ4jFy1cGrp1iNccyXQ9eM1Y8cLtK5bjSX9x92BcKJJZagehuaWfgPUmrw8SaFGf"
const SeedExpected = "c0f5f121ad57f4ee9499c94ea873f92768edcd19a466acb516884ee4902987511d3e0c1e3b0497d901df69a91971113968da21f2f09a2b9033daac1bbceb9455"
const Mnemonic = "other budget write video mimic captain cargo anger emerge chalk neck series"
const PrivateBTCBIP44KeyExpected = "xprvA2HswdqmdjVDBFnXQ84XU8n9T2rnkNuSLNodKH4pK78i3UGARtTnGaeGd1EqJjSvyZ5cjHWiVJv8sh4HdiCfQ8PNEuxD8rHPGM9dKh4yiFz"
const PrivateBTCBIP49KeyExpected = "xprvA3mS85KzYDdF2fUNci7TFaaj9ufDsz47pNM8PojrKQUTruef7SoLcS6oX74dNEk6kzz75FbBC5JHxSNoXVYjrkXgAyNLvQ8viPTTD3xHpS1"
const PrivateBTCBIP84KeyExpected = "xprvA3ZXfC49e7BzwatzgaF1P7iBtCd7iqaiWjHLCxWgmD39VSS6X4HtD9uUStvU4jgmgRie7tughU8i7kXYzgqW3ssMRg7SAjahPabB8P4ori3"

func TestBitcoinWalletManager_GenerateAddress44(t *testing.T) {
	seed := core.SeedFromMnemonic(Mnemonic, "")
	if SeedExpected != hex.EncodeToString(seed) {
		t.Fatalf("seed != %s, got %s", SeedExpected, seed)
	}
	masterKey44, err := core.MasterKeyFromSeed(seed, core.Mainnet)
	if err != nil {
		t.Fatal(err)
	}
	if masterKey44.String() != MasterKeyBIP44Expected {
		t.Fatalf("masterKey44 != %s, got %s", MasterKeyBIP44Expected, masterKey44)
	}
	derivatedPrivateKey44, err := core.HDWallet(
		masterKey44,
		core.PathBip(44, 0, nil),
	)
	if err != nil {
		t.Fatal(err)
	}

	if derivatedPrivateKey44.String() != PrivateBTCBIP44KeyExpected {
		t.Fatalf("walletAddress != %s, got %s", PrivateBTCBIP44KeyExpected, derivatedPrivateKey44.String())
	}
}

func TestBitcoinWalletManager_GenerateAddress49(t *testing.T) {
	seed := core.SeedFromMnemonic(Mnemonic, "")
	if SeedExpected != hex.EncodeToString(seed) {
		t.Fatalf("seed != %s, got %s", SeedExpected, seed)
	}
	masterKey, err := core.MasterKeyFromSeed(seed, core.Mainnet)
	if err != nil {
		t.Fatal(err)
	}
	if masterKey.String() != MasterKeyBIP44Expected {
		t.Fatalf("masterKey != %s, got %s", MasterKeyBIP44Expected, masterKey)
	}
	derivatedPrivateKey, err := core.HDWallet(
		masterKey,
		core.PathBip(49, 0, &core.WalletZeroPath),
	)
	if err != nil {
		t.Fatal(err)
	}

	if derivatedPrivateKey.String() != PrivateBTCBIP49KeyExpected {
		t.Fatalf("walletAddress != %s, got %s", PrivateBTCBIP49KeyExpected, derivatedPrivateKey.String())
	}
}

func TestBitcoinWalletManager_GenerateAddress84(t *testing.T) {
	seed := core.SeedFromMnemonic(Mnemonic, "")
	if SeedExpected != hex.EncodeToString(seed) {
		t.Fatalf("seed != %s, got %s", SeedExpected, seed)
	}
	masterKey, err := core.MasterKeyFromSeed(seed, core.Mainnet)
	if err != nil {
		t.Fatal(err)
	}
	if masterKey.String() != MasterKeyBIP44Expected {
		t.Fatalf("masterKey != %s, got %s", MasterKeyBIP44Expected, masterKey)
	}
	derivatedPrivateKey, err := core.HDWallet(
		masterKey,
		core.PathBip(84, 0, &core.WalletZeroPath),
	)
	if err != nil {
		t.Fatal(err)
	}

	if derivatedPrivateKey.String() != PrivateBTCBIP84KeyExpected {
		t.Fatalf("walletAddress != %s, got %s", PrivateBTCBIP84KeyExpected, derivatedPrivateKey.String())
	}

}
