package contracts

import "github.com/ethereum/go-ethereum/common/hexutil"

const (
	RPC             = "https://bsc-testnet.nodereal.io/v1/c206afd242af4b9a8eb1fe07aab5c60b"
	AddressContract = "0xBEe5a81210728B7f5AbEae40f161A447ab6c852e"
)

func byte32(data string) [32]byte {
	value := [32]byte{}
	dcm, _ := hexutil.Decode(data)
	copy(value[:], dcm)
	return value
}

func byte32Array(data []string) [][32]byte {
	var array [][32]byte
	for _, dt := range data {
		array = append(array, byte32(dt))
	}

	return array
}
