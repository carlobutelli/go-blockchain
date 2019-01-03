package main

import (
	"bytes"
	"crypto/sha512"
	"fmt"
)

/*
Every block inside the blockchain references the previous block that was created inside the BC.
We derive the hash inside of our block from the data inside of the block and the previous hash
that has been passed to the block.
*/
type Block struct {
	Hash     []byte // Hash of this block
	Data     []byte // data contained in this block
	PrevHash []byte // Last block hash, needed to link the new block has to the chain
}

type BlockChain struct {
	blocks []*Block
}

// Create the hash based on the previous hash and the data
func (b *Block) DeriveHash() {
	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
	hash := sha512.Sum512(info)
	// Push the created hash into the block field
	b.Hash = hash[:]
}

// Takes data and the previous hash from the last block and returns a pointer to a block
func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), prevHash}
	block.DeriveHash()
	return block
}

// Gets pointer for blockchain
func (chain *BlockChain) AddBlock(data string) {
	prevBlock := chain.blocks[len(chain.blocks)-1]
	newBlock := CreateBlock(data, prevBlock.Hash)
	chain.blocks = append(chain.blocks, newBlock)
}

func Alpha() *Block {
	return CreateBlock("Alpha", []byte{})
}

func IntiBlockChain() *BlockChain {
	return &BlockChain{[]*Block{Alpha()}}
}

func main() {
	chain := IntiBlockChain()

	chain.AddBlock("Beta Block")
	chain.AddBlock("Gamma Block")
	chain.AddBlock("Omega Block")

	for _, block := range chain.blocks {
		fmt.Printf("Data in block: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
	}
}
