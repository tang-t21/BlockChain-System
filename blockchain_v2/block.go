package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"fmt"
	"log"
	"math/big"
	"time"
)

type Block struct {
	// modified by v2
	Nonce        int64
	PrevHash     []byte
	Header       []byte
	Transactions []Transaction
	ID           int
	Timestamp    int64
	MerkleTree   *Tree
}

func NewBlock(prevHash []byte, transactions []Transaction, id int, nonce int64) *Block {
	block := &Block{nonce, prevHash, []byte{}, transactions, id, int64(time.Now().Unix()), nil}
	block.SetHeader()
	return block
}

func (block *Block) SetHeader() {
	// modified by v2
	var data [][]byte
	for _, tx := range block.Transactions {
		data = append(data, tx.TxId)
	}
	block.MerkleTree = NewMerkleTree(data)
	if len(block.Transactions) == 0 {
		var empty []byte
		hash := sha256.Sum256(empty)
		block.Header = hash[:]
	} else {
		block.Header = block.MerkleTree.Root.Data
	}
}

func NewGenesisBlock() *Block {
	block := &Block{0, []byte{}, []byte{}, []Transaction{}, 0, 0, nil}
	block.SetHeader()
	return block
}

func (block *Block) HashBlock() []byte {
	data := bytes.Join([][]byte{IntToHex(int64(block.Nonce)), block.PrevHash, block.Header, IntToHex(int64(block.ID)), IntToHex(block.Timestamp)}, []byte{})
	hash := sha256.Sum256(data)
	return hash[:]
}

func (block *Block) Verify() bool {
	nonce := block.Nonce
	threshold := big.NewInt(1)
	threshold = threshold.Lsh(threshold, total_bits-leading_zeros)
	hash_input := bytes.Join([][]byte{block.HashBlock(), IntToHex(int64(nonce))}, []byte{})
	var hash_int big.Int
	hash := sha256.Sum256(hash_input)
	hash_int.SetBytes(hash[:])
	if hash_int.Cmp(threshold) >= 0 {
		return false
	}
	for _, tx := range block.Transactions {
		if !tx.Verify() {
			log.Printf("ERROR: Invalid block! Invalid Transaction!")
			return false
		}
	}
	return true
}

func (block *Block) Serialize() []byte {
	var serialized_block bytes.Buffer
	encoder := gob.NewEncoder(&serialized_block)

	err := encoder.Encode(block)
	if err != nil {
		log.Panic(err)
	}

	return serialized_block.Bytes()
}

func DeserializeBlock(data []byte) *Block {
	var block Block

	decoder := gob.NewDecoder(bytes.NewReader(data))
	err := decoder.Decode(&block)
	if err != nil {
		log.Panic(err)
	}

	return &block
}

func (b *Block) Print() {
	fmt.Printf(" \033[0m\033[1;33m================================ Block: %d ================================\033[0m \n", b.ID)
	formatTimeStr := time.Unix(b.Timestamp, 0).Format("2006-01-02 15:04:05")
	fmt.Printf("Timestamp:%s\n", formatTimeStr)
	fmt.Printf("Header: %x \n", b.Header)
	fmt.Printf("PrevBlockHash: %x \n", b.PrevHash)
	fmt.Printf("Nounce: %x \n", b.Nonce)
	fmt.Printf("\033[0m\033[1;33m%s\033[0m \n", b.MerkleTree)
	for _, tx := range b.Transactions {
		fmt.Println(tx)
	}
	fmt.Printf("ThisBlockHash: %x \n", b.HashBlock())
	fmt.Printf(" ******************************** Block: %d ********************************* \n", b.ID)
	fmt.Printf("\n")
}
