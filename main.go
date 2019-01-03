package main

import (
	"fmt"

	"github.com/carlitos26/go-blockchain/blockchain"
)

func main() {
	chain := blockchain.InitBlockChain()

	chain.AddBlock("Beta Block")
	chain.AddBlock("Gamma Block")
	chain.AddBlock("Omega Block")

	for _, block := range chain.Blocks {
		fmt.Printf("Data in block: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
	}
}
