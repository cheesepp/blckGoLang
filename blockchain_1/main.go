package main

import (
	"fmt"
	"strconv"

	"github.com/cheesepp/golang-blockchain/blockchain"
)

func main() {
	chain := blockchain.InitBlockChain()

	chain.AddBlock("ngu")
	chain.AddBlock("ngoc")
	chain.AddBlock("ghe")

	for _, block := range chain.Blocks {
		fmt.Printf("Previous hash: %s\n", block.PrevHash)
		fmt.Printf("Data in Block: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)

		pow := blockchain.NewProof(block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()
	}
}
