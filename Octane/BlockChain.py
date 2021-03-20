class BlockChain():
    def __init__():
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

            payMiner = Transaction("Miner Rewards", miner, self.minerRewards)
            self.pendingTransactions = [payMiner]
        return True

    ##function for adding a transaction

    ##function to get last block
    def getLastBlock():
        return self.chain[-1]

    ##function to add genesisBlock
    def addGenesisBlock(self):
        nullTransaction = []
        nullTransaction.append(Transaction("none", "none", 0))
        genesis = Block()
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

            if block2.prev != b1.hash:
                print("error5")
                return False
        return True

    ##function to generate RSA keys

    ##function to turn blocks to json

    ##function to get blocks from json

    ##function to get balance



class Block (object):
    def __init__(self, transactions, time, index)
    self.index = index
    self.transactions = transactions
    self.time = time
    self.prev = ""
    self.nonse = 0
    self.hash = "" ##self.calculate hash

    ##function to calculate the hash

    ##function to mine block

    ##function for validating transactions

    ##function for json encoding self


class Transaction (object):
    def __init__(self, sender, reciever, ammount):
        self.sender = sender
        self.reciever = reciever
        self.ammount = ammount
        self.time = datetime.now().strftime("%m/%d/%Y, %H:%M:%S")


    ##function to calculate hash

    ##function to check validity

    ##function to sign transaction

