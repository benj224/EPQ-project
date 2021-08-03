package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

//basic block struct
type Block struct {
	Timestamp    time.Time
	Transactions []Transaction
	PrevHash     []byte
	Hash         []byte
	Nonce        int64
}

//BlockChian struct
type BlockChain struct {
	BlockChain []Block
}

//Transaction struct
type Transaction struct {
	Timestamp time.Time
	Sender    string
	Recipient string
	Ammount   float64
	Hash      []byte
	Signature []byte
}

//creates a new block
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

//creates a new transaction
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

//creates block hash
func NewBlockHash(time time.Time, transactions []Transaction, prevHash []byte) []byte {
	input := append(prevHash, time.String()...)
	for transaction := range transactions {
		input = append(input, string(rune(transaction))...)
	}
	hash := sha256.Sum256(input)
	return hash[:]
}

//calculates transaction hash
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

//calculates a block hash
func CalculateBlockHash(b Block) []byte {
	input := append(b.PrevHash, b.Timestamp.String()...)
	for transaction := range b.Transactions {
		input = append(input, string(rune(transaction))...)
	}
	hash := sha256.Sum256(input)
	return hash[:]
}

//mines a block
func (b Block) MineBlock(difficulty int) {
	puzzle := strings.Repeat("0", difficulty)
	for string(CalculateBlockHash(b)) != puzzle { //idk why CalBlkHash works but CalculateBlockHash doesnt
		b.Nonce += 1
	}
	fmt.Println("Block ", b.Hash, " mined with nonce", b.Nonce)
}

//prints block info to terminal
func printBlockInformation(block *Block) {
	fmt.Printf("\ttime: %s\n", block.Timestamp.String())
	fmt.Printf("\tprevHash: %x\n", block.PrevHash)
	fmt.Printf("\thash: %x\n", block.Hash)
	printTransactions(block)
}

//prints transaction to terminal
func printTransactions(block *Block) {
	fmt.Println("\tTransactions:")
	for i, transaction := range block.Transactions {
		fmt.Printf("\t\t%v: %s sent %s %v\n", i, transaction.Sender, transaction.Recipient, transaction.Ammount)
	}
}

//prints blockchain info to terminal
func printBlockChain(chain *BlockChain) {
	for _, block := range chain.BlockChain {
		printBlockInformation(&block)
	}
}

//converts blockchain to JSON
func BlockChainToJSON(chain BlockChain) {
	file, err := json.MarshalIndent(chain, "", " ")
	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(file))
	_ = ioutil.WriteFile("test.json", file, 0644)
}

//gets blockchain from given url -- probably change this to single block in future
func GetJSON(url string) BlockChain {
	spaceClient := http.Client{
		Timeout: time.Second * 2,
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", "test")

	res, getErr := spaceClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	nuBlock := BlockChain{}
	jsonErr := json.Unmarshal(body, &nuBlock)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	return nuBlock

}
