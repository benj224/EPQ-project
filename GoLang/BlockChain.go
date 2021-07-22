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
	timestamp    time.Time
	transactions []Transaction
	prevHash     []byte
	Hash         []byte
}

type BlockChain struct {
	blockChain []Block
}

type Transaction struct {
	timestamp time.Time
	sender    string
	recipient string
	ammount   float64
}

func NewBlock(transactions []Transaction, prevhash []byte) Block {
	currentTime := time.Now()
	return Block{
		timestamp:    currentTime,
		transactions: transactions,
		prevHash:     prevhash,
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
	fmt.Printf("\ttime: %s\n", block.timestamp.String())
	fmt.Printf("\tprevHash: %x\n", block.prevHash)
	fmt.Printf("\thash: %x\n", block.Hash)
	printTransactions(block)
}

func printTransactions(block *Block) {
	fmt.Println("\tTransactions:")
	for i, transaction := range block.transactions {
		fmt.Printf("\t\t%v: %s sent %s %v\n", i, transaction.sender, transaction.recipient, transaction.ammount)
	}
}

func printBlockChain(chain *BlockChain) {
	for _, block := range chain.blockChain {
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
		blockChain: []Block{genesisBlock},
	}

	BlockChainToJSON(blockChain)
	printBlockChain(&blockChain)

}

func BlockChainToJSON(chain BlockChain) {
	file, err := json.MarshalIndent(chain.blockChain, "", " ")
	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(file))
	_ = ioutil.WriteFile("test.json", file, 0644)
}
