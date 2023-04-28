package contracts

import (
	"MerkleProof/app/contracts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
)

type Contract struct {
	c *contracts.MerkleTree
}

func InitContract() (Contract, error) {
	client, err := ethclient.Dial(RPC)
	if err != nil {
		return Contract{}, err
	}

	address := common.HexToAddress(AddressContract)
	contract, err := contracts.NewMerkleTree(address, client)
	if err != nil {
		return Contract{}, err
	}

	return Contract{c: contract}, nil
}

func (contract Contract) Verify(root string, proof []string, addr common.Address, amount *big.Int) (bool, error) {
	vrf, err := contract.c.Verify(&bind.CallOpts{}, byte32(root), byte32Array(proof), addr, amount)
	if err != nil {
		return false, err
	}

	return vrf, nil
}

func (contract Contract) CheckHash(addr common.Address, amount *big.Int, leaf string) (bool, error) {
	vrf, err := contract.c.CheckHash(&bind.CallOpts{}, addr, amount, byte32(leaf))
	if err != nil {
		return false, nil
	}

	return vrf, nil
}
