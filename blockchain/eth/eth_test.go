package eth

import (
	"encoding/hex"
	"gild-wallet/core"
	"testing"
)

const addressBIP44Expected = "0x107B57cDF9F2309d0e62c402AaDc863Fc97B0cB3"
const publicBIP44Key = "02ee0693d9221d7baff9977698658b2417404da7e0576cb6653af8b9c5fa8b7cc6"

const masterKeyBIP44Expected = "xprv9s21ZrQH143K4JeCpHjRXGrTHY2rqLH6NCSmMTsyGUJrQ4jFy1cGrp1iNccyXQ9eM1Y8cLtK5bjSX9x92BcKJJZagehuaWfgPUmrw8SaFGf"
const seedExpected = "c0f5f121ad57f4ee9499c94ea873f92768edcd19a466acb516884ee4902987511d3e0c1e3b0497d901df69a91971113968da21f2f09a2b9033daac1bbceb9455"
const mnemonicTest = "other budget write video mimic captain cargo anger emerge chalk neck series"
const PrivateETHKeyExpected = "xprvA3tRgN6NM8znU4uinPYJcXD2SwQafRjGDxrhtCS6bzxyjGegCWMJvCFBPdSXAfTMeY3YLZMySPFFgYLwUHs1GGhEhQXibcxf8JJzrdNaBrW"
const privateWalletKeyExpected = "9e5c0cac2b7cfc1d2e91f40df68db051fb4a9aeed8d903c5b890e01346c6793e"
const addressExpected = "0x107B57cDF9F2309d0e62c402AaDc863Fc97B0cB3"

func TestEthereumWalletManager_GenerateWallet(t *testing.T) {
	manager := EthereumWalletManager{core.Mainnet}
	_, _, err := manager.GenerateWallet(core.BIP44)
	if err != nil {
		t.Fatal(err)
	}
}

func TestEthereumWalletManager_GenerateAddress(t *testing.T) {
	manager := EthereumWalletManager{core.Mainnet}

	address, err := manager.GenerateAddressFromString(publicBIP44Key, core.BIP44)
	if err != nil {
		t.Fatal(err)
	}

	if address != addressBIP44Expected {
		t.Fatalf("invalid address: %s, required: %s", address, addressBIP44Expected)
	}
}

func TestEthereumWalletManager_RestoreWallet(t *testing.T) {
	manager := EthereumWalletManager{core.Mainnet}
	seed := core.SeedFromMnemonic(mnemonicTest, "")
	if seedExpected != hex.EncodeToString(seed) {
		t.Fatalf("seed != %s, got %s", seedExpected, seed)
	}
	masterKey, err := core.MasterKeyFromSeed(seed, core.Mainnet)
	if err != nil {
		t.Fatal(err)
	}
	if masterKey.String() != masterKeyBIP44Expected {
		t.Fatalf("masterKey != %s, got %s", masterKeyBIP44Expected, masterKey)
	}
	derivatedPrivateKey, err := core.HDWallet(masterKey, core.PathBip(44, 60, &core.WalletZeroPath))
	if err != nil {
		t.Fatal(err)
	}

	if derivatedPrivateKey.String() != PrivateETHKeyExpected {
		t.Fatalf("privateKey != %s, got %s", PrivateETHKeyExpected, derivatedPrivateKey.String())
	}

	privateKey, err := derivatedPrivateKey.ECPrivKey()
	if err != nil {
		t.Fatal(err)
	}

	wallet, err := manager.RestoreWalletFromString(hex.EncodeToString(privateKey.ToECDSA().D.Bytes()), core.BIP44)
	if err != nil {
		t.Fatal(err)
	}

	if wallet.Address != addressExpected {
		t.Fatal("invalid address", wallet.Address)
	}

	if wallet.PrivateKey != privateWalletKeyExpected {
		t.Fatal("invalid private key got:", wallet.PrivateKey, "expected:", privateWalletKeyExpected)
	}
}

func TestEthereumWalletManager_RestoreWalletFromMnemonic(t *testing.T) {
	manager := EthereumWalletManager{core.Mainnet}
	wallet, err := manager.RestoreWalletFromMnemonic(mnemonicTest, "", core.BIP44)
	if err != nil {
		t.Fatal(err)
	}

	if wallet.PrivateKey != privateWalletKeyExpected {
		t.Fatal("invalid private key got:", wallet.PrivateKey, "expected:", privateWalletKeyExpected)
	}

	if wallet.Address != addressExpected {
		t.Fatal("invalid address", wallet.Address)
	}
}
