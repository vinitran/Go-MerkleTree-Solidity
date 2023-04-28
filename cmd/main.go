package main

import "MerkleProof/cmd/merkleTree"

func main() {
	data := []merkleTree.Data{
		{Address: "0x1111111111111111111111111111111111111111", Amount: "5000000000000000000"},
		{Address: "0x1111111111111111111111111111111111111112", Amount: "6000000000000000000"},
		{Address: "0x1111111111111111111111111111111111111113", Amount: "7000000000000000000"},
		{Address: "0x1111111111111111111111111111111111111114", Amount: "8000000000000000000"},
	}

	merkleTree.InitAndVerify(data)

	//err = CallContract(merkleTree.Root(), proof.ProofHashes(), data[1].GetAddress(), data[1].GetAmount())
	//if err != nil {
	//	log.Fatal(err)
	//}

}
