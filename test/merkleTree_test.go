package test

import (
	"MerkleProof/cmd/contracts"
	"MerkleProof/cmd/merkleTree"
	"fmt"
	"testing"
)

func TestMerkleTree(t *testing.T) {
	data := []merkleTree.Data{
		{Address: "0x1111111111111111111111111111111111111111", Amount: "5000000000000000000"},
		{Address: "0x1111111111111111111111111111111111111112", Amount: "6000000000000000000"},
		{Address: "0x1111111111111111111111111111111111111113", Amount: "7000000000000000000"},
		{Address: "0x1111111111111111111111111111111111111114", Amount: "8000000000000000000"},
	}

	mkTree, err := merkleTree.NewTree(data)
	if err != nil {
		t.Error(err)
	}

	indexTest := 0

	proof, err := mkTree.ProofByIndex(indexTest)
	if err != nil {
		t.Error(err)
	}

	fmt.Println("root: ", mkTree.Root())
	fmt.Println("proof: ", proof)

	vrf, err := mkTree.Verify(indexTest)
	if err != nil {
		t.Error(err)
	}

	fmt.Println("verify: ", vrf)

	contract, err := contracts.InitContract()
	if err != nil {
		t.Error(err)
	}

	valid, err := contract.Verify(mkTree.Root(), proof, data[indexTest].GetAddress(), data[indexTest].GetAmount())
	if err != nil {
		t.Error(err)
	}

	fmt.Println("check in contract ", valid)
}
