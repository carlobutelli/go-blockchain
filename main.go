package main

import (
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
		// fmt.Printf("Previous hash: %x\n", block.PrevHash)
		fmt.Printf("Data in block: %s\n", block.Data)
		fmt.Printf("Hash: %x\n\n", block.Hash)
	}
}
