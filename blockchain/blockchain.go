package blockchain

import (
	"fmt"

	"github.com/carlitos26/go-blockchain/errors"
	"github.com/dgraph-io/badger"
)

const (
	dbPath = "./data/blocks"
)

// Represent the chain of blocks
type BlockChain struct {
	LastHash []byte
	Database *badger.DB
}

// Create iterator to iterate through blocks into the blockchain
type BlockChainIterator struct {
	CurrentHash []byte
	Database    *badger.DB
}

// Gets pointer for blockchain
func (chain *BlockChain) AddNewBlock(data string) {
	var lastHash []byte
	err := chain.Database.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte("lh"))
		errors.Handle(err)
		lastHash, err = item.Value()
		return err
	})
	errors.Handle(err)
	newBlock := CreateBlock(data, lastHash)

	// put the new block in the db and assign it's has to last hash
	err = chain.Database.Update(func(txn *badger.Txn) error {
		err := txn.Set(newBlock.Hash, newBlock.Serialize())
		errors.Handle(err)
		err = txn.Set([]byte("lh"), newBlock.Hash)
		lastHash = newBlock.Hash
		return err
	})
	errors.Handle(err)
}

func InitBlockChain() *BlockChain {
	var lastHash []byte
	dbOpts := badger.DefaultOptions
	dbOpts.Dir = dbPath
	dbOpts.ValueDir = dbPath

	// get pointer to the database based on the set options
	db, err := badger.Open(dbOpts)
	errors.Handle(err)

	// Access the db to read/write transactions. It takes a closure with a pointer to badger transaction
	// First check if the blockchain has already been stored in the db, if yes, create a new blockchain
	// instance in memory and get the last hash, then push this block to the chain
	err = db.Update(func(txn *badger.Txn) error {
		// pass the key "lh" which stands for last hash
		if _, err := txn.Get([]byte("lh")); err == badger.ErrKeyNotFound {
			// if err, there is no blockchain in the db
			fmt.Println("No existing blockchain found")

			alpha := Alpha()
			fmt.Println("Alpha created and proved")

			// Use the alpha hash as key for the alpha block, serialize the alpha block itself and put it int the db
			err := txn.Set(alpha.Hash, alpha.Serialize())
			errors.Handle(err)

			// 'cause alpha is the first block we set its hash as last hash
			err = txn.Set([]byte("lh"), alpha.Hash)
			lastHash = alpha.Hash
			return err
		} else {
			// Blockchain is present in the db so get the last hash
			item, err := txn.Get([]byte("lh"))
			errors.Handle(err)
			lastHash, err = item.Value()
			return err
		}
	})
	errors.Handle(err)
	blockchain := BlockChain{lastHash, db}
	return &blockchain
}

func (chain *BlockChain) Iterator() *BlockChainIterator {
	iterator := &BlockChainIterator{chain.LastHash, chain.Database}
	return iterator
}

func (iterator *BlockChainIterator) Next() *Block {
	var block *Block
	err := iterator.Database.View(func(txn *badger.Txn) error {
		item, err := txn.Get(iterator.CurrentHash)
		encodedBlock, err := item.Value()
		// Deserialize the encoded block
		block = Deserialize(encodedBlock)
		return err
	})
	errors.Handle(err)

	// Change the iterator current hash to the block previous hash to fetch the next out of the db
	iterator.CurrentHash = block.PrevHash
	return block
}
