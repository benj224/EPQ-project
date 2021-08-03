package main

import (
	"fmt"
)

//TODO:
//Add a hash calculation function to transaction struct - done
//Test block hashing and think about multiple nonce possibility
//Look back at node.go when at home
//think about when adding new block to json blockchain must have the latest version of the chain idk how that will work

func main() {
	//place holder sample signature
	SAMPLE_SIGNATURE := []byte("benj224")

	//add some place holder transactions for testing
	genesisTransactions := []Transaction{
		NewTransaction("Genesis", "Reciever", 1000, SAMPLE_SIGNATURE),
		NewTransaction("Genesis", "Reciever2", 100, SAMPLE_SIGNATURE),
	}

	//create test block genesis
	genesisBlock := NewBlock(genesisTransactions, []byte("genesis"))
	blockChain := BlockChain{
		BlockChain: []Block{genesisBlock},
	}

	//write block to JSON and console
	BlockChainToJSON(blockChain)

	fmt.Println("added genesis")
	fmt.Scanln() //wait for input

	//create a second place holder block after input
	block1Transactions := []Transaction{
		NewTransaction("block1", "Reciever", 1000, SAMPLE_SIGNATURE),
		NewTransaction("block1", "Reciever2", 100, SAMPLE_SIGNATURE),
	}

	//add block to blockchain
	block1Block := NewBlock(block1Transactions, genesisBlock.Hash)
	blockChain.BlockChain = append(blockChain.BlockChain, block1Block)

	//write to JSON and console
	BlockChainToJSON(blockChain)

	newChain := GetJSON("localhost:8080")
	printBlockChain(&newChain)

}
