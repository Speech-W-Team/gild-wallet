package trx

import (
	"encoding/hex"
	"gild-wallet/core"
	"testing"
)

const addressBIP44Expected = "TWxWU4brH7GkL1hiyXGjLbbs6YDvah7zmd"
const publicBIP44Key = "02f936b7de77f74d4b93d86f1ea41eb0b7a9926bef4cdce496a097b1d5b263d066"

const masterKeyBIP44Expected = "xprv9s21ZrQH143K4JeCpHjRXGrTHY2rqLH6NCSmMTsyGUJrQ4jFy1cGrp1iNccyXQ9eM1Y8cLtK5bjSX9x92BcKJJZagehuaWfgPUmrw8SaFGf"
const seedExpected = "c0f5f121ad57f4ee9499c94ea873f92768edcd19a466acb516884ee4902987511d3e0c1e3b0497d901df69a91971113968da21f2f09a2b9033daac1bbceb9455"
const mnemonicTest = "other budget write video mimic captain cargo anger emerge chalk neck series"
const privateTRXKeyExpected = "xprvA3uyKz1bbVY3VPYthFUN6UC6H4F96eaN9nuQ8VuTWU5wP77sqpiAbCBB8hob9ptXZDye5d23ymcVQVCXnnZqPeTPoFT4bJSudTe3bJNNNqj"
const privateWalletKeyExpected = "bce327739ae6a657e47ff7352755d788db3b997603e8064207dd695f3dea915b"
const addressExpected = "TWxWU4brH7GkL1hiyXGjLbbs6YDvah7zmd"

func TestTronWalletManager_GenerateWallet(t *testing.T) {
	manager := TronWalletManager{core.Mainnet}
	_, _, err := manager.GenerateWallet(core.BIP44)
	if err != nil {
		t.Fatal(err)
	}
}

func TestTronWalletManager_GenerateAddress(t *testing.T) {
	manager := TronWalletManager{core.Mainnet}

	address, err := manager.GenerateAddressFromString(publicBIP44Key, core.BIP44)
	if err != nil {
		t.Fatal(err)
	}

	if address != addressBIP44Expected {
		t.Fatalf("invalid address: %s, required: %s", address, addressBIP44Expected)
	}
}

func TestTronWalletManager_RestoreWallet(t *testing.T) {
	manager := TronWalletManager{core.Mainnet}
	seed := core.SeedFromMnemonic(mnemonicTest, "")
	if seedExpected != hex.EncodeToString(seed) {
		t.Fatalf("seed != %s, got %s", seedExpected, seed)
	}
	masterKey, err := core.MasterKeyFromSeed(seed, manager.NetworkType)
	if err != nil {
		t.Fatal(err)
	}
	if masterKey.String() != masterKeyBIP44Expected {
		t.Fatalf("masterKey != %s, got %s", masterKeyBIP44Expected, masterKey)
	}
	derivatedPrivateKey, err := core.HDWallet(masterKey, core.PathBip(core.BIP44, 195, &core.WalletZeroPath))
	if err != nil {
		t.Fatal(err)
	}

	if derivatedPrivateKey.String() != privateTRXKeyExpected {
		t.Fatalf("privateKey != %s, got %s", privateTRXKeyExpected, derivatedPrivateKey.String())
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
		t.Fatal("invalid address")
	}

	if wallet.PrivateKey != privateWalletKeyExpected {
		t.Fatal("invalid private key got:", wallet.PrivateKey, "expected:", privateWalletKeyExpected)
	}
}

func TestTronWalletManager_RestoreWalletFromMnemonic(t *testing.T) {
	manager := TronWalletManager{core.Mainnet}
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
