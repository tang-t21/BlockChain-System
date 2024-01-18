package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
)

type BlockChain struct {
	Blocks []*Block
}

func NewBlockChain(Blocks []*Block) *BlockChain {
	return &BlockChain{Blocks}
}

func (blockchain *BlockChain) ComputeTailHash() []byte {
	tail_block := blockchain.Blocks[len(blockchain.Blocks)-1]
	tail_hash := tail_block.HashBlock()
	return tail_hash
}

func (blockchain *BlockChain) VerifyPrevHash(block *Block) bool {
	tail_block := blockchain.Blocks[len(blockchain.Blocks)-1]
	tail_hash := tail_block.HashBlock()
	if !bytes.Equal(tail_hash, block.PrevHash) {
		// log.Printf("ERROR: Invalid block! Wrong previous hash!")
		return false
	}
	return true
}

func (blockchain *BlockChain) AddBlock(block *Block) {
	var data [][]byte
	for _, tx := range block.Transactions {
		data = append(data, tx.TxId)
	}
	block.MerkleTree = NewMerkleTree(data)
	blockchain.Blocks = append(blockchain.Blocks, block)
	// fmt.Printf("Block %d added!\n", block.ID)
	block.Print()
}

func (blockchain *BlockChain) Verify() bool {
	for i, block := range blockchain.Blocks[1:] {
		if !block.Verify() {
			log.Printf("\033[0m\033[1;31mERROR: Invalid block found in blockchain!\033[0m\n")
			return false
		}
		prev_hash := blockchain.Blocks[i].HashBlock()
		if !bytes.Equal(block.PrevHash, prev_hash) {
			fmt.Printf("failure id:%d\n", i)
			log.Printf("\033[0m\033[1;31mERROR: Verify blockchain failed!\033[0m\n")
			return false
		}
	}
	return true
}
func (blockchain *BlockChain) CompareBlockchain(other_chain *BlockChain, probe_id int) int {
	for i := 0; i < len(other_chain.Blocks); i++ {
		if i == len(blockchain.Blocks[probe_id:]) {
			return i
		}
		if !bytes.Equal(blockchain.Blocks[i+probe_id].HashBlock(), other_chain.Blocks[i].HashBlock()) {
			return i
		}
	}
	return -1
}

func (blockchain *BlockChain) Serialize() []byte {
	var serialized_blockchain bytes.Buffer
	encoder := gob.NewEncoder(&serialized_blockchain)

	err := encoder.Encode(blockchain)
	if err != nil {
		log.Panic(err)
	}

	return serialized_blockchain.Bytes()
}

func DeserializeBlockChain(data []byte) *BlockChain {
	var blockchain BlockChain
	decoder := gob.NewDecoder(bytes.NewReader(data))
	err := decoder.Decode(&blockchain)
	if err != nil {
		log.Panic(err)
	}
	return &blockchain
}
