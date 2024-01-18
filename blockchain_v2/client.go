package main

import (
	"context"
	"crypto/sha256"
	"flag"
	"fmt"
	"log"
	"math/rand"
	pb "p1/proto-blockchain"
	"sync"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// send transaction to miner

var miner_addr StringList

func BroadcastTx(tx *Transaction) {
	if tx == nil {
		return
	}
	var wg sync.WaitGroup
	wg.Add(len(miner_addr))
	for _, addr := range miner_addr {
		go func(addr string) {
			conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
			if err != nil {
				log.Fatalf("did not connect: %v", err)
			}
			c := pb.NewBlockChainClient(conn)
			ctx, _ := context.WithCancel(context.Background())
			_, _ = c.GetTransaction(ctx, &pb.Transaction{Stx: tx.Serialize()})
			conn.Close()
			wg.Done()
		}(addr)
	}
	wg.Wait()
}

func FetchUTXOSetFromMiner(index int) UTXOset {
	addr := miner_addr[index]
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	c := pb.NewBlockChainClient(conn)
	ctx, _ := context.WithCancel(context.Background())
	u, _ := c.GetUtxoSet(ctx, &pb.UtxoRequest{Request: "YourUTXOSetPlease"})
	conn.Close()
	return DeserializeUTXOset(u.SUtxo)
}

func test_normal() {
	wallet_A := CreateWallet()
	wallet_B := CreateWallet()

	A_PubKeyHash := sha256.Sum256(wallet_A.PublicKey)
	B_PubKeyHash := sha256.Sum256(wallet_B.PublicKey)

	A_id := A_PubKeyHash[:]
	B_id := B_PubKeyHash[:]

	BroadcastTx(NewCoinBaseTransaction(A_id, 4000))
	BroadcastTx(NewCoinBaseTransaction(B_id, 4000))
	time.Sleep(5 * time.Second)
	rand.Seed(time.Now().UnixNano())
	var wallet *Wallet
	to := make([]byte, 32)
	for i := 0; i < 200; i++ {
		if rand.Float64() < 0.5 {
			wallet = wallet_A
			to = B_id
		} else {
			wallet = wallet_B
			to = A_id
		}
		index := rand.Intn(len(miner_addr))
		utxo := FetchUTXOSetFromMiner(index)
		fmt.Printf("got utxo from: %s\n", miner_addr[index])
		amount := 1 + rand.Intn(10)
		BroadcastTx(NewTransaction(wallet, to, amount, &utxo))
		time.Sleep(500 * time.Millisecond)
	}
}

func test_poor() {
	wallet_A := CreateWallet()
	wallet_B := CreateWallet()
	wallet_C := CreateWallet()

	A_PubKeyHash := sha256.Sum256(wallet_A.PublicKey)
	B_PubKeyHash := sha256.Sum256(wallet_B.PublicKey)
	C_PubKeyHash := sha256.Sum256(wallet_C.PublicKey)

	A_id := A_PubKeyHash[:]
	B_id := B_PubKeyHash[:]
	C_id := C_PubKeyHash[:]

	BroadcastTx(NewCoinBaseTransaction(A_id, 1000))
	BroadcastTx(NewCoinBaseTransaction(B_id, 1000))
	BroadcastTx(NewCoinBaseTransaction(C_id, 1000))

	time.Sleep(10 * time.Second)

	utxo := FetchUTXOSetFromMiner(rand.Intn(len(miner_addr)))
	BroadcastTx(NewTransaction(wallet_A, B_id, 500, &utxo))
	time.Sleep(5 * time.Second)

	utxo = FetchUTXOSetFromMiner(rand.Intn(len(miner_addr)))
	BroadcastTx(NewTransaction(wallet_C, B_id, 500, &utxo))
	time.Sleep(5 * time.Second)
	utxo = FetchUTXOSetFromMiner(rand.Intn(len(miner_addr)))
	BroadcastTx(NewTransaction(wallet_B, A_id, 1500, &utxo))
	time.Sleep(5 * time.Second)
	utxo = FetchUTXOSetFromMiner(rand.Intn(len(miner_addr)))
	BroadcastTx(NewTransaction(wallet_A, C_id, 1000, &utxo))
	time.Sleep(5 * time.Second)
	utxo = FetchUTXOSetFromMiner(rand.Intn(len(miner_addr)))
	BroadcastTx(NewTransaction(wallet_A, C_id, 1500, &utxo))
	time.Sleep(5 * time.Second)
	utxo = FetchUTXOSetFromMiner(rand.Intn(len(miner_addr)))
	BroadcastTx(NewTransaction(wallet_C, B_id, 500, &utxo))
	time.Sleep(5 * time.Second)
	utxo = FetchUTXOSetFromMiner(rand.Intn(len(miner_addr)))
	BroadcastTx(NewTransaction(wallet_A, C_id, 1000, &utxo))
	time.Sleep(5 * time.Second)
	utxo = FetchUTXOSetFromMiner(rand.Intn(len(miner_addr)))
	BroadcastTx(NewTransaction(wallet_A, B_id, 1000, &utxo))
	time.Sleep(5 * time.Second)
	utxo = FetchUTXOSetFromMiner(rand.Intn(len(miner_addr)))
	BroadcastTx(NewTransaction(wallet_B, C_id, 500, &utxo))
}

func test_attack_invalid_txs() {
	wallet_A := CreateWallet()
	wallet_B := CreateWallet()
	wallet_C := CreateWallet()

	A_PubKeyHash := sha256.Sum256(wallet_A.PublicKey)
	B_PubKeyHash := sha256.Sum256(wallet_B.PublicKey)
	C_PubKeyHash := sha256.Sum256(wallet_C.PublicKey)

	A_id := A_PubKeyHash[:]
	B_id := B_PubKeyHash[:]
	C_id := C_PubKeyHash[:]

	BroadcastTx(NewCoinBaseTransaction(A_id, 5000))
	BroadcastTx(NewCoinBaseTransaction(B_id, 5000))
	BroadcastTx(NewCoinBaseTransaction(C_id, 5000))

	time.Sleep(10 * time.Second)

	var wallet *Wallet
	to := make([]byte, 32)
	var utxo UTXOset
	for i := 0; i < 200; i++ {
		p := rand.Float64()
		if p < 0.33 {
			if rand.Float64() < 0.5 {
				wallet = wallet_A
				to = B_id
			} else {
				wallet = wallet_B
				to = A_id
			}

		} else if p < 0.67 {
			if rand.Float64() < 0.5 {
				wallet = wallet_B
				to = C_id
			} else {
				wallet = wallet_C
				to = B_id
			}
		} else {
			if rand.Float64() < 0.5 {
				wallet = wallet_C
				to = A_id
			} else {
				wallet = wallet_A
				to = C_id
			}
		}
		index := rand.Intn(len(miner_addr))
		if i < 50 || rand.Float64() < 0.7 {
			utxo = FetchUTXOSetFromMiner(index)
		}
		fmt.Printf("got utxo from: %s\n", miner_addr[index])
		amount := 1 + rand.Intn(10)
		BroadcastTx(NewTransaction(wallet, to, amount, &utxo))
		time.Sleep(500 * time.Millisecond)
	}
}

func test_special_attack_invalid_txs() {
	wallet_A := CreateWallet()
	wallet_B := CreateWallet()
	wallet_C := CreateWallet()

	A_PubKeyHash := sha256.Sum256(wallet_A.PublicKey)
	B_PubKeyHash := sha256.Sum256(wallet_B.PublicKey)
	C_PubKeyHash := sha256.Sum256(wallet_C.PublicKey)

	A_id := A_PubKeyHash[:]
	B_id := B_PubKeyHash[:]
	C_id := C_PubKeyHash[:]

	BroadcastTx(NewCoinBaseTransaction(A_id, 1000))
	BroadcastTx(NewCoinBaseTransaction(B_id, 1000))
	BroadcastTx(NewCoinBaseTransaction(C_id, 1000))

	time.Sleep(10 * time.Second)

	utxo := FetchUTXOSetFromMiner(0)
	BroadcastTx(NewTransaction(wallet_A, B_id, 500, &utxo))
	time.Sleep(5 * time.Second)

	utxo = FetchUTXOSetFromMiner(0)
	BroadcastTx(NewTransaction(wallet_C, B_id, 500, &utxo))
	time.Sleep(5 * time.Second)
	utxo = FetchUTXOSetFromMiner(0)
	BroadcastTx(NewTransaction(wallet_B, A_id, 1500, &utxo))
	time.Sleep(5 * time.Second)
	utxo = FetchUTXOSetFromMiner(0)
	BroadcastTx(NewTransaction(wallet_A, C_id, 2000, &utxo))
	time.Sleep(5 * time.Second)
	BroadcastTx(NewTransaction(wallet_A, C_id, 1500, &utxo))
	time.Sleep(5 * time.Second)
	utxo = FetchUTXOSetFromMiner(rand.Intn(len(miner_addr)))
	BroadcastTx(NewTransaction(wallet_C, B_id, 500, &utxo))
	time.Sleep(5 * time.Second)
	utxo = FetchUTXOSetFromMiner(rand.Intn(len(miner_addr)))
	BroadcastTx(NewTransaction(wallet_A, C_id, 1000, &utxo))
	time.Sleep(5 * time.Second)
	utxo = FetchUTXOSetFromMiner(rand.Intn(len(miner_addr)))
	BroadcastTx(NewTransaction(wallet_A, B_id, 1000, &utxo))
	time.Sleep(5 * time.Second)
	utxo = FetchUTXOSetFromMiner(rand.Intn(len(miner_addr)))
	BroadcastTx(NewTransaction(wallet_B, C_id, 500, &utxo))
}

func test_fake_client() {
	wallet_A := CreateWallet()
	wallet_B := CreateWallet()

	A_PubKeyHash := sha256.Sum256(wallet_A.PublicKey)
	B_PubKeyHash := sha256.Sum256(wallet_B.PublicKey)

	A_id := A_PubKeyHash[:]
	B_id := B_PubKeyHash[:]

	BroadcastTx(NewCoinBaseTransaction(A_id, 5000))
	BroadcastTx(NewCoinBaseTransaction(B_id, 5000))
	time.Sleep(3 * time.Second)
	rand.Seed(time.Now().UnixNano())
	var wallet *Wallet
	to := make([]byte, 32)
	for i := 0; i < 40; i++ {
		if rand.Float64() < 0.5 {
			wallet = wallet_A
			to = B_id
		} else {
			wallet = wallet_B
			to = A_id
		}
		utxo := FetchUTXOSetFromMiner(rand.Intn(len(miner_addr)))
		// fmt.Printf("got utxo: %+v \n", utxo)
		amount := 1 + rand.Intn(10)
		tx := NewTransaction(wallet, to, amount, &utxo)
		if rand.Float64() < 0.1 {
			input_id := rand.Intn(len(tx.TxInput))
			tx.TxInput[input_id].Signature = []byte{'A', 'B', 'C'}
		}
		BroadcastTx(tx)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	flag.Var(&miner_addr, "miner_addr", "the list of address of miners")
	flag.Parse()
	test_normal()
	// test_poor()
	// test_fake_client()
	// test_attack_invalid_txs()
	// test_special_attack_invalid_txs()
}
