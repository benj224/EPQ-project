package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"time"
)

type Block struct {
	Timestamp    time.Time     `json:"timestamp"`
	Transactions []Transaction `json:"transactions"`
	PrevHash     []byte        `json:"prev_hash"`
	Hash         []byte        `json:"hash"`
}

type BlockChain struct {
	BlockChain []Block
}

type Transaction struct {
	Timestamp time.Time
	Sender    string
	Recipient string
	Ammount   float64
}

func NewBlock(transactions []Transaction, prevhash []byte) Block {
	currentTime := time.Now()
	return Block{
		Timestamp:    currentTime,
		Transactions: transactions,
		PrevHash:     prevhash,
		Hash:         NewHash(currentTime, transactions, prevhash),
	}
}

func NewHash(time time.Time, transactions []Transaction, prevHash []byte) []byte {
	input := append(prevHash, time.String()...)
	for transaction := range transactions {
		input = append(input, string(rune(transaction))...)
	}
	hash := sha256.Sum256(input)
	return hash[:]
}

func printBlockInformation(block *Block) {
	fmt.Printf("\ttime: %s\n", block.Timestamp.String())
	fmt.Printf("\tprevHash: %x\n", block.PrevHash)
	fmt.Printf("\thash: %x\n", block.Hash)
	printTransactions(block)
}

func printTransactions(block *Block) {
	fmt.Println("\tTransactions:")
	for i, transaction := range block.Transactions {
		fmt.Printf("\t\t%v: %s sent %s %v\n", i, transaction.Sender, transaction.Recipient, transaction.Ammount)
	}
}

func printBlockChain(chain *BlockChain) {
	for _, block := range chain.BlockChain {
		printBlockInformation(&block)
	}
}

func main() {

	genesisTransactions := []Transaction{
		{
			time.Now(),
			"Genesis",
			"Reciever",
			183698258,
		},
		{
			time.Now(),
			"Genesis",
			"Reciever2",
			100,
		},
	}

	genesisBlock := NewBlock(genesisTransactions, []byte("genesis"))

	blockChain := BlockChain{
		BlockChain: []Block{genesisBlock},
	}

	BlockChainToJSON(blockChain)
	printBlockChain(&blockChain)

}

func BlockChainToJSON(chain BlockChain) {
	file, err := json.MarshalIndent(chain.BlockChain, "", " ")
	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(file))
	err = ioutil.WriteFile("test.json", file, 0644)
	if err != nil {
		log.Println(err)
	}
}
