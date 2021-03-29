package main

import (
	"fmt"
	"strconv"

	"github.com/jongregis/goBlockchain/blockchain"
)

// func (b *Block) DeriveHash() {
// 	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})

// 	hash := sha256.Sum256(info)

// 	b.Hash = hash[:]
// }

func main() {

	// init the blockcahin
	chain := blockchain.InitBlockChain()

	// adding blocks to the blockcahin
	chain.AddBlock("first block after genesis")
	chain.AddBlock("second block after genesis")
	chain.AddBlock("third block after genesis")

	// prints out the data/run pow from each block in the blockchain
	for _, block := range chain.Blocks {
		fmt.Printf("Previous hash: %x\n", block.PrevHash)
		fmt.Printf("data: %s\n", block.Data)
		fmt.Printf("hash: %x\n", block.Hash)

		//runs the pow algo on each block
		pow := blockchain.NewProofOfWork(block)
		fmt.Printf("Pow: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()
	}
}
