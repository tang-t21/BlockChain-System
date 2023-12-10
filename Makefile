miner: miner.go block.go blockchain.go Transaction.go TxIn.go TxOut.go utils.go UTXOset.go wallet.go
	go build -o miner miner.go block.go blockchain.go Transaction.go TxIn.go TxOut.go utils.go UTXOset.go wallet.go

client: client.go Transaction.go TxIn.go TxOut.go utils.go UTXOset.go wallet.go
	go build -o client client.go Transaction.go TxIn.go TxOut.go utils.go UTXOset.go wallet.go

deploy: miner
	./deploy.sh