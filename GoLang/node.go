package main

type node struct {
	address     string
	score       int64
	ping        int64
	chainLength int64
}

//make  a function to create a http server accessable by public_ip/chain and public_ip/chainlen

//public_ip/chainlen retrun current chain length

//public_ip/chain retrun missing blocks and check matching hashes

//function to get some_ip/chainlen and get the chain if given length is longer that current length
