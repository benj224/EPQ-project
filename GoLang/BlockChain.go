package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"time"
)

//TODO:
//Add a hash calculation function to transaction struct - done
//Test block hashing and think about multiple nonce possibility
//Look back at node.go when at home
//think about when adding new block to json blockchain must have the latest version of the chain idk how that will work

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
		Hash:         NewBlockHash(currentTime, transactions, prevhash),
		Nonce:        0,
	}
}

func NewTransaction(sender string, reciever string, ammount float64, signature []byte) Transaction {
	currentTime := time.Now()
	return Transaction{
		Timestamp: currentTime,
		Sender:    sender,
		Recipient: reciever,
		Ammount:   ammount,
		Hash:      NewTransactionHash(currentTime, sender, reciever, ammount, signature),
		Signature: []byte("sample"),
	}
}

func NewBlockHash(time time.Time, transactions []Transaction, prevHash []byte) []byte {
	input := append(prevHash, time.String()...)
	for transaction := range transactions {
		input = append(input, string(rune(transaction))...)
	}
	hash := sha256.Sum256(input)
	return hash[:]
}

func NewTransactionHash(time time.Time, sender string, reciever string, ammount float64, signature []byte) []byte {
	input := append(signature, time.String()...)
	input = append(input, sender...)
	input = append(input, reciever...)
	input = append(input, []byte(fmt.Sprint(ammount))...)
	hash := sha256.Sum256(input)
	return hash[:]
}

// func (block Block) CalculateBlockHash() []byte {
// 	input := append(block.PrevHash, block.Timestamp.String()...)
// 	for transaction := range block.Transactions {
// 		input = append(input, string(rune(transaction))...)
// 	}
// 	hash := sha256.Sum256(input)
// 	return hash[:]
// }

func CalculateBlockHash(b Block) []byte {
	input := append(b.PrevHash, b.Timestamp.String()...)
	for transaction := range b.Transactions {
		input = append(input, string(rune(transaction))...)
	}
	hash := sha256.Sum256(input)
	return hash[:]
}

func (b Block) MineBlock(difficulty int) {
	puzzle := strings.Repeat("0", difficulty)
	for string(CalculateBlockHash(b)) != puzzle { //idk why CalBlkHash works but CalculateBlockHash doesnt
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

}
