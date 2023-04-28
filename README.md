# Go-Sol-MerkleTree
Go-Sol-MerkleTree is a Golang library that provides a simple and efficient way to hash data and push it into a Merkle tree within an smart contract for airdrop. The library takes in a [data.csv](https://github.com/vinitran/Go-MerkleTree-Solidity/blob/main/data.csv) file, which contains addresses and corresponding amounts, and outputs a [data.json](https://github.com/vinitran/Go-MerkleTree-Solidity/blob/main/data.json) file.

This library is well-suited for use with the standard tree of [Openzzeppelin](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/master/contracts/utils/cryptography/MerkleProof.sol), ensuring compatibility and ease of integration.

Furthermore, I have developed [an instance contract](https://github.com/vinitran/Sol-Merkle-Contract) that utilizes this library, providing an example implementation for other developers to reference


# About Merkle Tree
Merkle tree is a tree in which every "leaf" (node) is labelled with the cryptographic hash of a data block, and every node that is not a leaf (called a branch, inner node, or inode) is labelled with the cryptographic hash of the labels of its child nodes. A hash tree allows efficient and secure verification of the contents of a large data structure. A hash tree is a generalization of a hash list and a hash chain.
## Installation

Use the package, run the below command in cmd

```bash
git clone https://github.com/vinitran/Go-MerkleTree-Solidity.git

cd Go-MerkleTree-Solidity

go run cmd/main.go
```

## Usage
Read file cmd/main.go
## Contributing

Pull requests are welcome. For major changes, please open an issue first
to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License

[MIT](https://choosealicense.com/licenses/mit/)
