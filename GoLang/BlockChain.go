package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"time"
)

//TODO:
//Add a hash calculation function to transaction struct
//Test block hashing and think about multiple nonce possibility
//Look back at node.go when at home

type Block struct {
	Timestamp    time.Time
	Transactions []Transaction
	PrevHash     []byte
	Hash         []byte
	Nonce        int64
}

type BlockChain struct {
	BlockChain []Block
}

type Transaction struct {
	Timestamp time.Time
	Sender    string
	Recipient string
	Ammount   float64
	Hash      []byte
	Signature []byte
}

func NewBlock(transactions []Transaction, prevhash []byte) Block {
	currentTime := time.Now()
	return Block{
		Timestamp:    currentTime,
		Transactions: transactions,
		PrevHash:     prevhash,
		Hash:         NewHash(currentTime, transactions, prevhash),
		Nonce:        0,
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

func (block Block) CalculateHash() []byte {
	input := append(block.PrevHash, block.Time.String()...)
	for transaction := range block.Transactions {
		input = append(input, string(rune(transaction))...)
	}
	hash := sha256.Sum256(input)
	return hash[:]
}

func (b Block) MineBlock(difficulty int) {
	puzzle := Repeat("0", difficulty)
	for string(b.CalculateHash) != puzzle {
		b.Nonce += 1
	}
	fmt.Println("Block ", b.Hash, " mined with nonce", b.Nonce)
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

func BlockChainToJSON(chain BlockChain) {
	file, err := json.MarshalIndent(chain.BlockChain, "", " ")
	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(file))
	_ = ioutil.WriteFile("test.json", file, 0644)
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
