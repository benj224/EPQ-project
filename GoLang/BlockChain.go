package main

import (
	"fmt"
	"net/http"
)

//TODO:
//Add a hash calculation function to transaction struct - done
//Test block hashing and think about multiple nonce possibility
//Look back at node.go when at home
//think about when adding new block to json blockchain must have the latest version of the chain idk how that will work

func main() {

	SAMPLE_SIGNATURE := []byte("benj224")

	genesisTransactions := []Transaction{
		NewTransaction("Genesis", "Reciever", 1000, SAMPLE_SIGNATURE),
		NewTransaction("Genesis", "Reciever2", 100, SAMPLE_SIGNATURE),
	}

	genesisBlock := NewBlock(genesisTransactions, []byte("genesis"))

	blockChain := BlockChain{
		BlockChain: []Block{genesisBlock},
	}

	BlockChainToJSON(blockChain)
	printBlockChain(&blockChain)

	http.HandleFunc("/", ReturnJSON)
	http.ListenAndServe(":8080", nil)

	fmt.Println("dfjaskdfjask")
}
