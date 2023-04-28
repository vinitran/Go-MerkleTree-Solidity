package main

import (
	"MerkleProof/cmd/merkleTree"
	"fmt"
)

func main() {
	//get data from csv file
	csvData, err := ReadCsv(csvPath)
	if err != nil {
		fmt.Println(err)
		return
	}

	data := merkleTree.ParseToDataMerkleTree(csvData)

	mkTree, err := merkleTree.NewTree(data)
	if err != nil {
		fmt.Println(err)
		return
	}

	proof, err := mkTree.AllProof()
	if err != nil {
		fmt.Println(err)
		return
	}

	type TreeData struct {
		Root  string             `json:"root"`
		Proof []merkleTree.Proof `json:"proof"`
	}

	treeData := TreeData{
		Root:  mkTree.Root(),
		Proof: proof,
	}

	err = WriteDataToFileAsJSON(treeData, jsonPath)
	if err != nil {
		fmt.Println(err)
		return
	}
}
