package main

import (
	"bytes"
	"crypto/sha512"
)

// Create the hash based on the previous hash and the data
func (b *Block) DeriveHash() {
	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
	hash := sha512.Sum512(info)
	// Push the created hash into the block field
	b.Hash = hash[:]
}
