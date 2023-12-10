package main

import (
	"crypto/sha256"
	"math/rand"
	"time"
)

func test_pow() {
	wallet_A := CreateWallet()
	wallet_B := CreateWallet()

	A_PubKeyHash := sha256.Sum256(wallet_A.PublicKey)
	B_PubKeyHash := sha256.Sum256(wallet_B.PublicKey)

	A_id := A_PubKeyHash[:]
	B_id := B_PubKeyHash[:]

	BroadcastTx(NewCoinBaseTransaction(A_id, 5000))
	BroadcastTx(NewCoinBaseTransaction(B_id, 5000))
	time.Sleep(30 * time.Second)
	rand.Seed(time.Now().UnixNano())
	var wallet *Wallet
	to := make([]byte, 32)
	for i := 0; i < 20; i++ {
		if rand.Float64() < 0.5 {
			wallet = wallet_A
			to = B_id
		} else {
			wallet = wallet_B
			to = A_id
		}
		utxo := FetchUTXOSetFromMiner(rand.Intn(len(address_list)))
		// fmt.Printf("got utxo: %+v \n", utxo)
		amount := 1 + rand.Intn(10)
		BroadcastTx(NewTransaction(wallet, to, amount, &utxo))
		time.Sleep(15 * time.Second)
	}
}
