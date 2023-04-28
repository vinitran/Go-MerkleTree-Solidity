package merkleTree

import (
	"fmt"
	solmerkle "github.com/0xKiwi/sol-merkle-tree-go"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

type MerkleTree struct {
	//mk merkle.Tree
	mk    *solmerkle.MerkleTree
	leave [][]byte
}

type Proof struct {
	Leaf  string   `json:"leaf"`
	Proof []string `json:"proof"`
}

func NewTree(data []Data) (MerkleTree, error) {
	var leaves [][]byte
	for _, d := range data {
		leaves = append(leaves, d.Hash())
	}

	tree, err := solmerkle.GenerateTreeFromHashedItems(leaves)
	if err != nil {
		return MerkleTree{}, fmt.Errorf("could not generate trie: %v", err)
	}

	return MerkleTree{
		mk:    tree,
		leave: leaves,
	}, nil
}

func (tree MerkleTree) Root() string {
	return hexutil.Encode(tree.mk.Root())
}

func (tree MerkleTree) ProofByIndex(index int) ([]string, error) {
	proof, err := tree.Proof(tree.leave[index])
	if err != nil {
		return nil, err
	}

	return proof, nil
}

func (tree MerkleTree) Proof(data []byte) ([]string, error) {
	var proofStringArr []string
	proof, err := tree.mk.MerkleProof(data)
	if err != nil {
		return nil, fmt.Errorf("could not generate proof: %v", err)
	}

	for _, p := range proof {
		proofStringArr = append(proofStringArr, hexutil.Encode(p))
	}

	return proofStringArr, nil
}

func (tree MerkleTree) AllProof() ([]Proof, error) {
	var proof []Proof

	for _, leaf := range tree.leave {
		leafString := hexutil.Encode(leaf)
		pr, err := tree.Proof(leaf)
		if err != nil {
			return nil, err
		}

		proof = append(proof, Proof{
			Leaf:  leafString,
			Proof: pr,
		})
	}

	return proof, nil
}

func (tree MerkleTree) Verify(index int) (bool, error) {
	root := tree.mk.Root()
	proof, err := tree.mk.MerkleProof(tree.leave[index])
	if err != nil {
		return false, fmt.Errorf("could not generate proof: %v", err)
	}
	leaf := tree.leave[index]
	return solmerkle.VerifyMerkleBranch(root, leaf, proof), nil
}
