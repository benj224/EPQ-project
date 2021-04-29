from BlockChain import BlockChain, Block, Transaction
import json

chain = BlockChain()
chain.addGenesisBlock()

with open("blockchain.json", "w") as outFile:
    json.dump(chain.encodeChainJSON, outFile)
