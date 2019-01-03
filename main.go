package main

import (
	"fmt"
	"strconv"

	"github.com/carlitos26/go-blockchain/blockchain"
)

func main() {
	chain := blockchain.InitBlockChain()

	chain.AddNewBlock("Beta secuirty info")
	chain.AddNewBlock("Gamma transaction")
	chain.AddNewBlock("Omega detailed payment")

	for _, block := range chain.Blocks {
		fmt.Printf("Previous hash: %x\n", block.PrevHash)
		fmt.Printf("Block's Data:  %s\n", block.Data)
		fmt.Printf("Current Hash:  %x\n", block.Hash)

		pow := blockchain.Proof(block)
		fmt.Printf("Proof of Work: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()
	}
}
