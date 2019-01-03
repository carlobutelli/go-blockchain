package blockchain

import (
	"bytes"
	"crypto/sha512"
	"encoding/binary"
	"fmt"
	"log"
	"math"
	"math/big"
)

/* Proof of work algorithm. Secure the blockchain by forcing the network to do work to add a block to
   the chain. Computational power.
*/

const Difficulty = 12

type ProofOfWork struct {
	Block  *Block
	Target *big.Int
}

// Grab data from the block
func NewProof(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(512-Difficulty)) // shift left
	pow := &ProofOfWork{b, target}
	return pow
}

// Create counter at 0. Nonce is an arbitrary number that can be used just once
func (pow *ProofOfWork) InitData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			pow.Block.Data,
			pow.Block.PrevHash,
			ToHex(int64(nonce)),
			ToHex(int64(Difficulty)),
		},
		[]byte{},
	)
	return data
}

func ToHex(num int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}
	return buff.Bytes()
}

// Create hash for data and counter
func (pow *ProofOfWork) RunProof() (int, []byte) {
	var intHash big.Int
	var hash [64]byte

	nonce := 0
	for nonce < math.MaxInt64 {
		data := pow.InitData(nonce)
		hash = sha512.Sum512(data)

		fmt.Printf("\r%x", hash)

		intHash.SetBytes(hash[:])
		if intHash.Cmp(pow.Target) == -1 {
			break // hash is less than the target
		} else {
			nonce++
		}
	}
	fmt.Println()
	return nonce, hash[:]
}

// Check if hash meets requirements:
// firsts bytes must contains 0s

func (pow *ProofOfWork) Validate() bool {
	var intHash big.Int
	data := pow.InitData(pow.Block.Nonce)

	hash := sha512.Sum512(data)
	intHash.SetBytes(hash[:])

	return intHash.Cmp(pow.Target) == -1
}
