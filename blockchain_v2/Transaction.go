package main

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/sha256"
	"encoding/gob"
	"encoding/hex"
	"fmt"
	"log"
	"math/big"
	"strings"
)

type Transaction struct {
	TxId     []byte
	TxInput  []TxIn
	TxOutput []TxOut
}

func (tx Transaction) Hash() []byte {
	var hash [32]byte
	tx.TxId = []byte{}
	hash = sha256.Sum256(tx.Serialize())
	return hash[:]
}

func NewCoinBaseTransaction(to []byte, amount int) *Transaction {
	input := TxIn{
		TxId:      []byte{},
		OutIndex:  -1,
		Txout:     TxOut{},
		Signature: nil,
		PubKey:    []byte{},
	}
	output := TxOut{
		Value:      amount,
		PubKeyHash: to,
	}
	tx := Transaction{
		TxId:     nil,
		TxInput:  []TxIn{input},
		TxOutput: []TxOut{output},
	}
	tx.TxId = tx.Hash()
	return &tx
}

func NewTransaction(wallet *Wallet, to []byte, amount int, u *UTXOset) *Transaction {
	var inputs []TxIn
	var outputs []TxOut

	from := sha256.Sum256(wallet.PublicKey) //get sender id
	avalibale, validOutputs := u.FindSpendableOutputsForAmount(from[:], amount)
	if avalibale < amount {
		// log.Panic("ERROR: NOT ENOUGH FUNDS !!! \n")
		log.Printf("\033[0m\033[1;31mERROR:\033[0m INVALID TRANSACTION !!! NOT ENOUGH FUNDS \n")
		return nil
	}
	for txid, out_list := range validOutputs {
		txID, _ := hex.DecodeString(txid)
		for _, out := range out_list {
			input := TxIn{
				TxId:      txID,
				OutIndex:  out.OutIndex,
				Txout:     out.Txout,
				Signature: nil,
				PubKey:    wallet.PublicKey,
			}
			inputs = append(inputs, input)
		}
	}
	// give the money to 'to'
	out_to := TxOut{
		Value:      amount,
		PubKeyHash: to,
	}
	outputs = append(outputs, out_to)
	// give the exceedding money back to myself
	if avalibale > amount {
		out_from := TxOut{
			Value:      avalibale - amount,
			PubKeyHash: from[:],
		}
		outputs = append(outputs, out_from)
	}
	tx := Transaction{
		TxId:     nil,
		TxInput:  inputs,
		TxOutput: outputs,
	}
	txid := tx.Hash()
	tx.TxId = txid

	wallet.SignTransaction(&tx)
	return &tx
}

func (tx *Transaction) IsCoinBaseTransaction() bool {
	if len(tx.TxInput) == 1 && len(tx.TxInput[0].TxId) == 0 && tx.TxInput[0].OutIndex == -1 {
		return true
	}
	return false
}

func (tx *Transaction) Verify() bool {
	// TODO if transaction is CoinBase transaction, return true
	if tx.IsCoinBaseTransaction() {
		return true
	}
	inamount := 0
	curve := elliptic.P256()
	for _, in := range tx.TxInput {
		x := big.Int{}
		y := big.Int{}
		keyLen := len(in.PubKey)
		x.SetBytes(in.PubKey[:(keyLen / 2)])
		y.SetBytes(in.PubKey[(keyLen / 2):])
		outindexbyte := IntToBytes(in.OutIndex)
		data_to_verify := append(in.TxId, outindexbyte...)
		rawPubKey := ecdsa.PublicKey{Curve: curve, X: &x, Y: &y}
		if ecdsa.VerifyASN1(&rawPubKey, data_to_verify, in.Signature) == false {
			fmt.Println("\033[0m\033[1;31m Wrong Signature! Fake Client !!!\033[0m")
			return false
		}
		inamount = inamount + in.Txout.Value
	}
	outamount := 0
	for _, out := range tx.TxOutput {
		outamount = outamount + out.Value
	}

	if inamount < outamount {
		fmt.Println("\033[0m\033[1;31m Not Enough Inputs! Fake Client !!!\033[0m")
		return false
	}

	return true
}

func (tx Transaction) Serialize() []byte {
	var BytesBuffer bytes.Buffer
	var encoder = gob.NewEncoder(&BytesBuffer)
	err := encoder.Encode(tx)
	if err != nil {
		log.Panic(err)
	}
	return BytesBuffer.Bytes()
}

func DeserializeTransaction(data []byte) Transaction {
	var tx Transaction
	var decoder = gob.NewDecoder(bytes.NewReader(data))
	err := decoder.Decode(&tx)
	if err != nil {
		log.Panic(err)
	}
	return tx
}

func (tx Transaction) String() string {
	var lines []string
	lines = append(lines, fmt.Sprintf("---Transaction %x:", tx.TxId))
	for i, input := range tx.TxInput {
		lines = append(lines, fmt.Sprintf("   INPUT %d:", i))
		lines = append(lines, fmt.Sprintf("           TxId:    %x", input.TxId))
		lines = append(lines, fmt.Sprintf("           OutIndex:    %d", input.OutIndex))
	}
	for i, output := range tx.TxOutput {
		lines = append(lines, fmt.Sprintf("   OUTPUT %d:", i))
		lines = append(lines, fmt.Sprintf("           Value:   %d", output.Value))
	}
	return strings.Join(lines, "\n")
}
