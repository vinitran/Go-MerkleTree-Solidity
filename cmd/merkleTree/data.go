package merkleTree

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"math/big"
)

type Data struct {
	Address string `json:"address"`
	Amount  string `json:"amount"`
}

func (d Data) Hash() []byte {
	address := common.HexToAddress(d.Address)
	myAmount, _ := new(big.Int).SetString(d.Amount, 10)

	packed := append(
		address.Bytes(),
		common.LeftPadBytes(myAmount.Bytes(), 32)...,
	)

	return crypto.Keccak256(packed)
}

func (d Data) GetAmount() *big.Int {
	myAmount, _ := new(big.Int).SetString(d.Amount, 10)
	return myAmount
}

func (d Data) GetAddress() common.Address {
	return common.HexToAddress(d.Address)
}

func ParseToDataMerkleTree(tmp [][]string) []Data {
	var data []Data
	for _, t := range tmp {
		dt := Data{
			Address: t[0],
			Amount:  t[1][1:],
		}
		data = append(data, dt)
	}

	return data
}
