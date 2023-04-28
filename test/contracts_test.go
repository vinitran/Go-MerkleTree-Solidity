package test

import (
	"MerkleProof/cmd/contracts"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"math/big"
	"testing"
)

const (
	address = "0x1111111111111111111111111111111111111111"
	amount  = "5000000000000000000"
)

func TestVerify(t *testing.T) {
	rootTest := "0x8ccde925183babcbedbd7cd9f6bba4f66905419df0136726df2a3151b30aaf66"
	proofTest := []string{
		"0x2945234a39fd175d084a3bc3349c65f9680c1b83c7f6b3c5da64d16a38a9f683",
		"0x75c5c39494133e59e3a85ffd2c820a113c9ecf29f6e7e030a4ae6f0c02f3e087",
	}
	addressTest := common.HexToAddress("0x1111111111111111111111111111111111111111")
	amountTest, _ := new(big.Int).SetString("5000000000000000000", 10)

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
	addressTest := common.HexToAddress(address)
	amountTest, _ := new(big.Int).SetString(amount, 10)
	leaf := "0xeb02c421cfa48976e66dfb29120745909ea3a0f843456c263cf8f1253483e283"

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

func TestCheckHash_1(t *testing.T) {
	addressTest := common.HexToAddress(address)
	amountTest, _ := new(big.Int).SetString(amount, 10)
	leafData := append(
		addressTest.Bytes(),
		common.LeftPadBytes(amountTest.Bytes(), 32)...,
	)
	leaf := hexutil.Encode(crypto.Keccak256(leafData))
	fmt.Println(leaf)
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
