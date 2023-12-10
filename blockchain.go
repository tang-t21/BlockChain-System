package main

import (
	"bytes"
	"encoding/gob"
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
			log.Printf("\033[0m\033[1;31mERROR: Verify blockchain failed!\033[0m\n")
			return false
		}
	}
	return true
}
func (blockchain *BlockChain) CompareBlockchain(other_chain *BlockChain) int {
	for i := 0; i < len(other_chain.Blocks); i++ {
		if i == len(blockchain.Blocks) {
			return i
		}
		if !bytes.Equal(blockchain.Blocks[i].PrevHash, other_chain.Blocks[i].PrevHash) {
			return i - 1
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
