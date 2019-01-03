package blockchain

// Represent the single block
type Block struct {
	Hash     []byte // Hash of this block
	Data     []byte // data contained in this block
	PrevHash []byte // Last block hash, needed to link the new block has to the chain
	Nonce    int
}

// Represent the chain of blocks
type BlockChain struct {
	Blocks []*Block
}

// Takes data and the previous hash from the last block and returns a pointer to a block
func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), prevHash, 0}
	pow := Proof(block)
	nonce, hash := pow.RunProof()
	block.Hash = hash[:]
	block.Nonce = nonce
	return block
}

// Gets pointer for blockchain
func (chain *BlockChain) AddNewBlock(data string) {
	prevBlock := chain.Blocks[len(chain.Blocks)-1]
	newBlock := CreateBlock(data, prevBlock.Hash)
	chain.Blocks = append(chain.Blocks, newBlock)
}

// The starting block
func Alpha() *Block {
	return CreateBlock("Alpha", []byte{})
}

func InitBlockChain() *BlockChain {
	return &BlockChain{[]*Block{Alpha()}}
}
