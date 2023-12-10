package main

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"flag"
	"fmt"
	"log"
	"math"
	"math/big"
	"net"
	pb "p1/proto-blockchain"
	"sync"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var maxNonce = math.MaxInt64

const leading_zeros = 20
const fake_leading_zeros = leading_zeros - 1
const total_bits = 256

type BlockAddr struct {
	Selfblock *Block
	Addr      string
}

// Define the struct Miner
type Miner struct {
	Blockchain *BlockChain
	UtxoSet    *UTXOset
	Honest     bool
	TxChannel  chan Transaction
	BlockChan  chan *BlockAddr
	pb.UnimplementedBlockChainServer
}

// Constructor
func NewMiner(honest bool) *Miner {
	genesisblock := NewGenesisBlock()
	miner := &Miner{
		Blockchain: NewBlockChain([]*Block{genesisblock}),
		UtxoSet:    &UTXOset{make(map[string][]TxOutIndex)},
		Honest:     honest,
		TxChannel:  make(chan Transaction, 1),
		BlockChan:  make(chan *BlockAddr, 20),
	}
	return miner
}

func CheckInputInList(input TxIn, spendable_output_list []TxOutIndex) bool {
	ret := false
	for _, txoutindex := range spendable_output_list {
		if bytes.Equal(txoutindex.TxId, input.TxId) && txoutindex.OutIndex == input.OutIndex {
			ret = true
			break
		}
	}
	return ret
}

func (miner *Miner) VerifyTx(tx Transaction) bool {
	if tx.IsCoinBaseTransaction() {
		owner_string := hex.EncodeToString(tx.TxOutput[0].PubKeyHash)
		for _, txoutindex := range miner.UtxoSet.UtxoSet[owner_string] {
			if bytes.Equal(txoutindex.TxId, tx.TxId) {
				return false
			}
		}
		return true
	}
	if !tx.Verify() {
		return false
	}
	pubkeyhash := sha256.Sum256(tx.TxInput[0].PubKey)
	query_string := hex.EncodeToString(pubkeyhash[:])
	spendable_output_list := miner.UtxoSet.UtxoSet[query_string]
	for _, input := range tx.TxInput {
		if CheckInputInList(input, spendable_output_list) == false {
			return false
		}
	}
	return true
}

// grpc interface to receive transaction
func (miner *Miner) GetTransaction(ctx context.Context, in *pb.Transaction) (*pb.Response, error) {
	tx := DeserializeTransaction(in.GetStx())
	miner.TxChannel <- tx
	return &pb.Response{Result: "Got Tx!"}, nil
}

func (miner *Miner) GetUtxoSet(ctx context.Context, in *pb.UtxoRequest) (*pb.Utxo, error) {
	return &pb.Utxo{SUtxo: miner.UtxoSet.Serialize()}, nil
}

func (miner *Miner) GetBlock(ctx context.Context, sblock *pb.Block) (*pb.Response, error) {
	block := DeserializeBlock(sblock.GetSblock())
	if !block.Verify() {
		fmt.Println("\033[0m\033[1;31mInvalid Block!\033[0m")
		return &pb.Response{Result: "Invalid Block!"}, nil
	}
	miner.BlockChan <- &BlockAddr{Selfblock: block, Addr: sblock.GetAddr()}
	fmt.Printf("\033[0m\033[1;32mGot block:%d from %s\033[0m\n", block.ID, sblock.GetAddr())
	return &pb.Response{Result: "Got Block!"}, nil
}

// broadcast to other miners when puzzle solved
func (miner *Miner) BroadcastBlock(block *Block) {
	var wg sync.WaitGroup
	wg.Add(len(address_list) - 1)
	for _, addr := range address_list {
		if addr == address_list[*addr_idx] {
			continue
		}
		go func(addr string) {
			// fmt.Println("\033[0m\033[1;31mSending block to other miners!\033[0m")
			conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
			if err != nil {
				log.Fatalf("did not connect: %v", err)
			}
			// fmt.Printf("\033[0m\033[1;32mConnection to %s established!\033[0m\n", addr)
			c := pb.NewBlockChainClient(conn)
			ctx, _ := context.WithCancel(context.Background())
			_, _ = c.GetBlock(ctx, &pb.Block{Sblock: block.Serialize(), Addr: address_list[*addr_idx]})
			conn.Close()
			wg.Done()
		}(addr)
	}
	wg.Wait()
}

func (miner *Miner) compute_nonce(block *Block) int64 {
	fmt.Printf("\033[0m\033[1;34mMining new block %d\033[0m\n", block.ID)
	var hash [32]byte
	var hash_int big.Int
	nonce := int64(0)
	threshold := big.NewInt(1)
	if miner.Honest {
		threshold = threshold.Lsh(threshold, total_bits-leading_zeros)
	} else {
		threshold = threshold.Lsh(threshold, total_bits-fake_leading_zeros)
	}

	for true {
		select {
		case other_blockaddr := <-miner.BlockChan:
			// fmt.Println("4:\n" + miner.UtxoSet.String())
			// for _, tx := range block.Transactions {
			// 	fmt.Println(tx)
			// }
			miner.UtxoSet.UndoTransactions(block.Transactions)
			miner.HandleOtherBlock(other_blockaddr.Selfblock, other_blockaddr.Addr)
			var new_txs []Transaction
			block.PrevHash = miner.Blockchain.ComputeTailHash()
			block.ID = len(miner.Blockchain.Blocks)
			// fmt.Printf("Block ID \033[0m\033[1;31mAfter handle other block: %d\033[0m\n", block.ID)
			for _, tx := range block.Transactions {
				if miner.VerifyTx(tx) {
					new_txs = append(new_txs, tx)
					miner.UtxoSet.UpdateUTXOOneTransaction(&tx)
				}
			}
			block.Transactions = new_txs
			block.SetHeader()
			// nonce = 0

		default:
			select {
			case tx := <-miner.TxChannel:
				if miner.VerifyTx(tx) {
					fmt.Printf("\033[0m\033[1;32mGot Valid Tx:%x\033[0m\n", tx.TxId)
					block.Transactions = append(block.Transactions, tx)
					block.SetHeader()
					miner.UtxoSet.UpdateUTXOOneTransaction(&tx)

					// nonce = 0
				} else {
					fmt.Printf("\033[0m\033[1;31mInvalid Tx:%x\033[0m\n", tx.TxId)
				}
			default:
			}
		}
		block.Nonce = nonce
		hash_input := bytes.Join([][]byte{block.HashBlock(), IntToHex(int64(nonce))}, []byte{})

		hash = sha256.Sum256(hash_input)
		hash_int.SetBytes(hash[:])
		// fmt.Printf("hash value:%x\n", hash_int)
		if hash_int.Cmp(threshold) < 0 {
			fmt.Println("\033[0m\033[1;32mNew block mined out !\033[0m")
			break
		}

		if nonce >= int64(maxNonce) {
			nonce = 0
		}
		nonce++
	}
	return nonce
}

// mining process
func (miner *Miner) MineBlock() {
	prevHash := miner.Blockchain.ComputeTailHash()
	id := len(miner.Blockchain.Blocks)
	block := NewBlock(prevHash, []Transaction{}, id, 0)
	nonce := miner.compute_nonce(block)
	block.Nonce = nonce

	is_valid := miner.Blockchain.VerifyPrevHash(block)
	if is_valid {
		miner.Blockchain.AddBlock(block)
		// fmt.Println("1:\n" + miner.UtxoSet.String())
	} else {
		fmt.Println("\033[0m\033[1;31mCritical Error: Supposed to successfully add!\033[0m")
	}
	miner.BroadcastBlock(block)
}

// This function is executed only after the miner stop the on-going mining computation and recover the UTXO
func (miner *Miner) HandleOtherBlock(block *Block, addr string) {
	// fmt.Println("\033[0m\033[1;31mHandling other block!\033[0m")
	is_valid := miner.Blockchain.VerifyPrevHash(block)
	if is_valid {
		miner.Blockchain.AddBlock(block)
		miner.UtxoSet.UpdateUTXOOneBlockAppended(block)
		// fmt.Println("2:\n" + miner.UtxoSet.String())
	} else {
		if block.ID <= len(miner.Blockchain.Blocks)-1 {
			fmt.Printf("\033[0m\033[1;35mShorter Than Me!\033[0m\n")
			return
		} else {
			conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
			if err != nil {
				log.Fatalf("did not connect: %v", err)
			}
			c := pb.NewBlockChainClient(conn)
			ctx, _ := context.WithCancel(context.Background())
			Sblockchain, _ := c.GetBlockChain(ctx, &pb.ChainRequest{Request: "Give me whole Chain!"})
			fmt.Printf("\033[0m\033[1;32mGot blockchain from %s\033[0m\n", addr)
			long_blockchain := DeserializeBlockChain(Sblockchain.GetSchain())
			fmt.Printf("other_blockchain:%+v\n", long_blockchain)
			fmt.Printf("my_blockchain:%+v\n", miner.Blockchain)
			conn.Close()
			first_diff_id := miner.Blockchain.CompareBlockchain(long_blockchain)
			if first_diff_id == -1 {
				log.Panic("ERROR: Same Chain but cannot be added!")
			}
			diff_chain := &BlockChain{long_blockchain.Blocks[first_diff_id:]}
			if !diff_chain.Verify() {
				log.Printf("\033[0m\033[1;31mERROR:Got Invalid blockchain!\033[0m\n")
				return
			}

			miner.UndoBlocks(first_diff_id)
			for i := first_diff_id; i < len(long_blockchain.Blocks); i++ {
				miner.Blockchain.AddBlock(long_blockchain.Blocks[i])
				// fmt.Println("3:\n" + miner.UtxoSet.String())
				miner.UtxoSet.UpdateUTXOOneBlockAppended(long_blockchain.Blocks[i])
			}
		}
	}
}

func (miner *Miner) GetBlockChain(ctx context.Context, chainreq *pb.ChainRequest) (*pb.Chain, error) {
	return &pb.Chain{Schain: miner.Blockchain.Serialize()}, nil
}

func (miner *Miner) UndoBlocks(start_id int) {
	for i := len(miner.Blockchain.Blocks) - 1; i >= start_id; i-- {
		miner.UtxoSet.UndoTransactions(miner.Blockchain.Blocks[i].Transactions)
	}
	miner.Blockchain.Blocks = miner.Blockchain.Blocks[:start_id]
}

var (
	addr_idx  = flag.Int("addr_idx", 0, "The miner address index in addr_list")
	self_addr = flag.String("self_addr", "localhost:8090", "The miner address")
	is_honest = flag.Bool("is_honest", true, "whether the miner is honest")
)

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", address_list[*addr_idx])
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	server := NewMiner(*is_honest)
	go func() {
		for true {
			server.MineBlock()
		}
	}()
	pb.RegisterBlockChainServer(s, server)

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}

//longest chain to handle fork
