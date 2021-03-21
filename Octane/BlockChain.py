import datetime
from Crypto.PublicKey import RSA
import hashlib
import json


class BlockChain():
    def __init__(self):
        self.chain = [self.addGenesisBlock()]
        self.pendingTransactions = []
        self.dificulty = 2
        self.miningReward = 10
        self.blockSize = 5
        self.nodes = set()

    ##function for registering nodes

    ##function for updating chain 
    def resolveChainConflicts(self):
        neighborNodes = self.nodes
        updatedChain = None
        chainUpdated = False

        currentLength = len(self.chain)
        for node in neighborNodes:
            ##return the nodes current chain

            if True: #change to request return status
                length = 0 ##length of neighbor nodes chain
                chain = None ##neighbors chain

                if True:  ##Length of neighbors chain bigger that current chain and thier chain is valid
                    currentLength = length
                    updatedChain = chain
                    chainUpdated = True
        
        if chainUpdated:
            self.chain = None ##decode the returned jason of the newest chain
            return True
        return False

    ##function to create blocks from pending transactions
    def minePendingTransactions(self, miner):
        pendingTransactionLength = len(self.pendingTransactions)
        if(pendingTransactionLength <= 1):
            print("Not enough Transactions to mine, Must be greater than 1")
            return False
        else:
            for i in range(0, pendingTransactionLength, self.blockSize): ##create blocks from pending transactions by itterating through at steps of blockSize
                endTransaction = i + self.blockSize ##sets the bounds for the next transactions to be put in a block
                if i >= pendingTransactionLength:  ##if the next iteration gives an upper bound higher than the length of list set the upper bound to length of list
                    end = pendingTransactionLength

                transactionSlice = self.pendingTransactions[i:end]

                newBlock = Block(transactionSlice, datetime.now().strftime("%m/%d/%Y, %H:%M:%S"), len(self.chain))
                prevHashVal = self.getLastBlock().hash ##need to add this function
                newBlock.prev = prevHashVal
                newBlock.mineBlock(self.dificulty) ##add this function
                self.chain.append(newBlock)
            print("MIning Transaction success")

            payMiner = Transaction("Miner Rewards", miner, self.miningReward)
            self.pendingTransactions = [payMiner]
        return True

    ##function for adding a transaction

    ##function to get last block
    def getLastBlock(self):
        return self.chain[-1]

    ##function to add genesisBlock
    def addGenesisBlock(self):
        nullTransaction = []
        nullTransaction.append(Transaction("none", "none", 0))
        genesis = Block(nullTransaction, datetime.now().strftime("%m/%d/%Y, %H:%M:%S"), 0)
        genesis.prev = "none"
        return genesis

    ##function to validate the chain
    def isValidChain(self):
        for i in range(1, len(self.chain)):
            block1 = self.chain[i-1]
            block2 = self.chain[i]

            if not block2.hasValidTransactions():##create this funciton
                print("error3")
                return False

            if block2.hash != block2.calculateHash():
                print("error4")
                return False

            if block2.prev != block1.hash:
                print("error5")
                return False
        return True

    ##function to generate RSA keys
    def generateKeys(self):
        key = RSA.generate(2048)
        privateKey = key.export_key()
        fileOut = open("Private.pem", "wb")
        fileOut.write(privateKey)

        publicKey = key.publickey().export_key()
        fileOut = open("Reciever.pem", "wb")
        fileOut.write(publicKey)

        print(publicKey.decode("ASCII"))
        return key.publickey().export_key().decode("ASCII")

    ##function to turn blocks to json
    def encodeChainJSON(self):
        JSONBlockArray = []
        for block in self.chain:
            blockJSON = {}
            blockJSON["hash"] = block.hash
            blockJSON['index'] = block.index
            blockJSON['prev'] = block.prev
            blockJSON['time'] = block.time
            blockJSON['nonse'] = block.nonse

            transactionsJSON = []
            transactionJSON = {}
            for transaction in block.transactions:
                transactionJSON['time'] = transaction.time
                transactionJSON['sender'] = transaction.sender
                transactionJSON['reciever'] = transaction.reciever
                transactionJSON['ammount'] = transaction.ammount
                transactionJSON['hash'] = transaction.hash
                transactionsJSON.append(transactionJSON)

            blockJSON["transactions"] = transactionsJSON
            JSONBlockArray.append(blockJSON)

            return JSONBlockArray


    ##function to get blocks from json
    def chainJSONdecode(self, chainJSON):
        chain=[]
        for blockJSON in chainJSON:

            tArr = []
            for tJSON in blockJSON['transactions']:
                transaction = Transaction(tJSON['sender'], tJSON['reciever'], tJSON['ammount'])
                transaction.time = tJSON['time']
                transaction.hash = tJSON['hash']
                tArr.append(transaction)


            block = Block(tArr, blockJSON['time'], blockJSON['index'])
            block.hash = blockJSON['hash']
            block.prev =blockJSON['prev']
            block.nonse = blockJSON['nonse']

            chain.append(block)
            return chain

    ##function to get balance
    def getBalance(self, person):
        balance = 0 
        for i in range(1, len(self.chain)):
            block = self.chain[i]
            try:
                for j in range(0, len(block.transactions)):
                    transaction = block.transactions[j]
                    if(transaction.sender == person):
                        balance -= transaction.amt
                    if(transaction.reciever == person):
                        balance += transaction.amt
            except AttributeError:
                print("no transaction")
        return balance + 100 ## might chance this



class Block (object):
    def __init__(self, transactions, time, index):
        self.index = index
        self.transactions = transactions
        self.time = time
        self.prev = ""
        self.nonse = 0
        self.hash = self.calculateHash() ##self.calculate hash

        ##function to calculate the hash#
    def calculateHash(self):
        transactionsToHash = ""
        for transaction in self.transactions:
            transactionsToHash += transaction.hash
        hashString = str(self.time) + transactionsToHash + self.prev + str(self.nonse)
        hashEncoded = json.dumps(hashString, sort_keys = True).endcode()
        return hashlib.sha256(hashEncoded).hexdigest() ##swap out here for better hash function

        ##function to mine block
    def mineBlock(self, difficulty):
        ary = []
        for i in range(difficulty):
            ary.append(i)
        aryStr = map(str, ary)
        hashPuzzle = "".join(aryStr)
        while self.hash[0:difficulty] != hashPuzzle:
            self.nonse += 1
            self.hash = self.calculateHash()
        print("block ", self.hash," mined")
        return True


        ##function for validating transactions
    def hasValidTransactions(self):
        for i in range(0, len(self.transactions)):
            transaction = self.transactions[i]
            if not transaction.isValidTransaction():
                return False
            return True

        ##function for json encoding self
    def JSONencode(self):
        return jsonpickle.encode(self)


class Transaction (object):
    def __init__(self, sender, reciever, ammount):
        self.sender = sender
        self.reciever = reciever
        self.ammount = ammount
        self.time = datetime.now().strftime("%m/%d/%Y, %H:%M:%S")
        self.hash = self.calculateHash
        self.signature = ""


    ##function to calculate hash 
    def calculateHash(self):
        hashString = self.sender + self.reciever + str(self.ammount) + str(self.time)
        hashEncoded = json.dumps(hashString, sort_keys = True).encode()
        return hashlib.sha256(hashEncoded).hexdigest()

    ##function to check validity
    def isValidTransaction(self):
        if self.hash != self.calculateHash():
            return False
        if self.sender == self.reciever:
            return False
        if self.sender == "Miner Rewards":
            return True
        if not self.signature or len(self.signature) == 0:
            print("No Signature")
            return False
        return True

    ##function to sign transaction
    def signTransaction(self, key, sender):
        if(self.hash != self.calculateHash):
            print("transaction error tampered")
            return False

        if(str(key.publickey().export_key()) != str(senderKey.publickey().export_key())):
            print("transaction to be signed form a nother wallet")
            return False

        self.signature = key.sign(self.hash, "") ##might need to base 64 encode
        return true


