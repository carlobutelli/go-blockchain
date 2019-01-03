package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"
	"strconv"

	"github.com/carlitos26/go-blockchain/blockchain"
	"github.com/carlitos26/go-blockchain/errors"
)

type CommandLine struct {
	blockchain *blockchain.BlockChain
}

func (cl *CommandLine) printCommandUsage() {
	fmt.Println("Usage:")
	fmt.Println("add -block BLOCK_DATA - to add a block to the chain")
	fmt.Println("print - Prints all the blocks in the chain")
}

func (cl *CommandLine) validateArguments() {
	if len(os.Args) < 2 {
		cl.printCommandUsage()
		// exiting the app by shutting down the go routine
		runtime.Goexit()
	}
}

func (cl *CommandLine) addBlock(data string) {
	cl.blockchain.AddNewBlock(data)
	fmt.Println("Block Added!")
}

func (cl *CommandLine) printWholeChain() {
	iterator := cl.blockchain.Iterator()
	for {
		block := iterator.Next()
		fmt.Printf("Previous hash: %x\n", block.PrevHash)
		fmt.Printf("Block's Data:  %s\n", block.Data)
		fmt.Printf("Current Hash:  %x\n", block.Hash)

		pow := blockchain.Proof(block)
		fmt.Printf("Proof of Work: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()
		if len(block.PrevHash) == 0 {
			break
		}
	}
}

func (cl *CommandLine) run() {
	cl.validateArguments()
	// Create flags for user to type in
	addBlockCmd := flag.NewFlagSet("add", flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("print", flag.ExitOnError)
	addBlockData := addBlockCmd.String("block", "", "Block Data")

	switch os.Args[1] {
	case "add":
		err := addBlockCmd.Parse(os.Args[2:])
		errors.Handle(err)

	case "print":
		err := printChainCmd.Parse(os.Args[2:])
		errors.Handle(err)

	default:
		cl.printCommandUsage()
		runtime.Goexit()
	}

	// Checks the returning values
	if addBlockCmd.Parsed() {
		if *addBlockData == "" {
			addBlockCmd.Usage()
			runtime.Goexit()
		}
		cl.addBlock(*addBlockData)
	}

	if printChainCmd.Parsed() {
		cl.printWholeChain()
	}
}

func main() {
	defer os.Exit(0)
	chain := blockchain.InitBlockChain()
	defer chain.Database.Close()

	cl := CommandLine{chain}
	cl.run()
}
