# Basic framework for Blockchain tech in Go
Every block inside the blockchain references the previous block that was created inside the BC. We derive the hash inside of our block from the data inside of the block and the previous hash that has been passed to the block.

Then we try to power the actual network by running proof of work algorithm to secure the blockchain by forcing the network to do work to add a block to the chain. By doing that the blocks and data inside the blocks are more secure.

Validation of the proof, when a user does the work to sign a block, a proof of this work needs to be provided => work must be hard to do but proving this work must be relatively easy.

Requirements to meet the hash => first few bytes must contain 0s.

Blocks are stored with metadata which descirbes all of the blocks in the chain.
Chain state object stores the state of a chain and all of the current transaction outputs as well as few pieces of metadata.

To increase performances each block is split in its separate file so we don't have to open up more blocks in case we want to read only one hash or block.

### Run the app to see the usage
> ```$ go run main.go``` 

### Run the app to print the chain
> ```$ go run main.go print``` 

### Run the app to add a block
> ```$ go run main.go add -block "Beta block to be added"``` 

### Build an executable file
> ```$ go build main.go``` 
