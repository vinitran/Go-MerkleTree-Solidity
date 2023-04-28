package merkleTree

import (
	"fmt"
	"github.com/ComposableFi/go-merkle-trees/hasher"
	"github.com/ComposableFi/go-merkle-trees/merkle"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

func InitAndVerify(data []Data) {
	var leaves [][]byte
	for _, d := range data {
		leaves = append(leaves, d.Hash())
	}

	//init markle tree with type hash
	merkleTree := merkle.NewTree(hasher.Keccak256Hasher{})
	//add leaves to merkle tree
	merkleTree, err := merkleTree.FromLeaves(leaves)
	if err != nil {
		fmt.Printf("err: %e\n", err)
		return
	}

	root := merkleTree.Root()
	fmt.Printf("Merkle root is %s \n", hexutil.Encode(merkleTree.Root()))

	leafIndexInArray := []uint64{1}
	proof := merkleTree.Proof(leafIndexInArray)

	fmt.Printf("Merkle proof hashes are:\n")
	for _, v := range proof.ProofHashesHex() {
		fmt.Printf(" - 0x%v\n", v)
	}

	// verify merkle proof
	verifyResult, err := proof.Verify(root)
	if err != nil {
		fmt.Printf("err: %e\n", err)
		return
	}

	if !verifyResult {
		fmt.Printf("err: %s\n", "Merkle proof verify result is false")
		return
	}

	fmt.Printf("Merkle proof verify result is %v\n", verifyResult)
}
