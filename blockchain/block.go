package blockchain

import (
	"bytes"
	"encoding/gob"

	"github.com/carlitos26/go-blockchain/errors"
)

// Represent the single block
type Block struct {
	Hash     []byte // Hash of this block
	Data     []byte // data contained in this block
	PrevHash []byte // Last block hash, needed to link the new block has to the chain
	Nonce    int
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

// The starting block
func Alpha() *Block {
	return CreateBlock("Alpha", []byte{})
}

// Serializer
func (b *Block) Serialize() []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)

	err := encoder.Encode(b)
	errors.Handle(err)
	// return bytes portion of our result
	return result.Bytes()
}

// Deserializer
func Deserialize(data []byte) *Block {
	var block Block
	// Create a bytes reader for new decoder
	decoder := gob.NewDecoder(bytes.NewReader(data))
	err := decoder.Decode(&block)
	errors.Handle(err)
	return &block
}
