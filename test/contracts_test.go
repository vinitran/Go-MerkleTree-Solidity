package test

import (
	"MerkleProof/cmd/contracts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"math/big"
	"testing"
)

func TestVerify(t *testing.T) {
	rootTest := "0x6fde6b67f907f5c5f936177161c2be0ea6f56a6bfacc82de27ed1896a45b3dca"
	proofTest := []string{
		"0xeb02c421cfa48976e66dfb29120745909ea3a0f843456c263cf8f1253483e283",
		"0x878a6177db1c361556ac046bbc6b4a34663c582653f718c4077482b82a5ce824",
	}
	addressTest := common.HexToAddress("0x1111111111111111111111111111111111111112")
	amountTest, _ := new(big.Int).SetString("6000000000000000000", 10)

	contract, err := contracts.InitContract()
	if err != nil {
		t.Error(err)
	}

	valid, err := contract.Verify(rootTest, proofTest, addressTest, amountTest)
	if err != nil {
		t.Error(err)
	}

	if valid == false {
		t.Error("Dont match")
	}
}

func TestCheckHash_0(t *testing.T) {
	addressTest := common.HexToAddress("0x1111111111111111111111111111111111111111")
	amountTest, _ := new(big.Int).SetString("5000000000000000000", 10)
	//leafData := append(
	//	addressTest.Bytes(),
	//	common.LeftPadBytes(amountTest.Bytes(), 32)...,
	//)
	//leaf := crypto.Keccak256(leafData)

	contract, err := contracts.InitContract()
	if err != nil {
		t.Error(err)
	}

	valid, err := contract.CheckHash(addressTest, amountTest, "0xeb02c421cfa48976e66dfb29120745909ea3a0f843456c263cf8f1253483e283")
	if err != nil {
		t.Error(err)
	}

	if valid == false {
		t.Error("Dont match")
	}

}

func TestCheckHash_1(t *testing.T) {
	addressTest := common.HexToAddress("0x1111111111111111111111111111111111111111")
	amountTest, _ := new(big.Int).SetString("5000000000000000000", 10)
	leafData := append(
		addressTest.Bytes(),
		common.LeftPadBytes(amountTest.Bytes(), 32)...,
	)
	leaf := hexutil.Encode(crypto.Keccak256(leafData))

	contract, err := contracts.InitContract()
	if err != nil {
		t.Error(err)
	}

	valid, err := contract.CheckHash(addressTest, amountTest, leaf)
	if err != nil {
		t.Error(err)
	}

	if valid == false {
		t.Error("Dont match")
	}

}
