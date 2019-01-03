package blockchain

// Represent the chain of blocks
type BlockChain struct {
	Blocks []*Block
}

// Gets pointer for blockchain
func (chain *BlockChain) AddNewBlock(data string) {
	prevBlock := chain.Blocks[len(chain.Blocks)-1]
	newBlock := CreateBlock(data, prevBlock.Hash)
	chain.Blocks = append(chain.Blocks, newBlock)
}

func InitBlockChain() *BlockChain {
	return &BlockChain{[]*Block{Alpha()}}
}
