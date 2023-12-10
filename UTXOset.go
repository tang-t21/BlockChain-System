package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"encoding/hex"
	"fmt"
	"log"
	"strings"
)

type TxOutRet struct {
	OutIndex int
	Txout    TxOut
}

type TxOutIndex struct {
	TxId     []byte
	OutIndex int
	Out      TxOut
}

type UTXOset struct {
	UtxoSet map[string][]TxOutIndex // A map from client id (publickey) (coonvert from []byte to string) to all his spendable output
}

func (u *UTXOset) FindSpendableOutputsForAmount(pub_key_hash []byte, amount_needed int) (int, map[string][]TxOutRet) {
	// return the sum of funds in validoutputs, a map of [transaction_id] [](output index)
	query_string := hex.EncodeToString(pub_key_hash)
	txoutindexs := u.UtxoSet[query_string]
	ret_map := make(map[string][]TxOutRet)
	ret_aval := 0

	for _, txoutindex := range txoutindexs {
		txid_string := hex.EncodeToString(txoutindex.TxId)
		ret_map[txid_string] = append(ret_map[txid_string], TxOutRet{OutIndex: txoutindex.OutIndex, Txout: txoutindex.Out})
		ret_aval = ret_aval + txoutindex.Out.Value
		if ret_aval >= amount_needed {
			return ret_aval, ret_map
		}
	}
	return ret_aval, ret_map
}

func (u *UTXOset) UpdateUTXOOneBlockAppended(b *Block) {
	for _, tx := range b.Transactions {
		u.UpdateUTXOOneTransaction(&tx)
	}
}

func (u *UTXOset) UpdateUTXOOneTransaction(tx *Transaction) {
	if !tx.IsCoinBaseTransaction() {
		for _, input := range tx.TxInput {
			PubKeyHash := sha256.Sum256(input.PubKey)
			query_string := hex.EncodeToString(PubKeyHash[:])
			index := -1
			for i, txoutindex := range u.UtxoSet[query_string] {
				if bytes.Equal(txoutindex.TxId, input.TxId) && txoutindex.OutIndex == input.OutIndex {
					index = i
					break
				}
			}
			if index == -1 {
				log.Panic("tnnd\n")
			}
			if index == len(u.UtxoSet[query_string])-1 {
				u.UtxoSet[query_string] = u.UtxoSet[query_string][:index]
			} else {
				u.UtxoSet[query_string] = append(u.UtxoSet[query_string][:index], u.UtxoSet[query_string][index+1:]...)
			}
		}
	}
	for i, output := range tx.TxOutput {
		query_string := hex.EncodeToString(output.PubKeyHash[:])
		txoutindex := TxOutIndex{
			TxId:     tx.TxId,
			OutIndex: i,
			Out:      output,
		}
		u.UtxoSet[query_string] = append(u.UtxoSet[query_string], txoutindex)
	}
}

func CheckTxIdInList(txid string, id_list []string) bool {
	for _, i := range id_list {
		if i == txid {
			return true
		}
	}
	return false
}

func (u *UTXOset) UndoTransactions(txs []Transaction) {
	id_list := make([]string, 0)
	for _, i := range txs {
		txid := hex.EncodeToString(i.TxId)
		id_list = append(id_list, txid)
	}

	for wallet_id, spendable_list := range u.UtxoSet {
		deleted := 0
		for idx, out := range spendable_list {
			txid_string := hex.EncodeToString(out.TxId)
			if CheckTxIdInList(txid_string, id_list) {
				real_idx := idx - deleted
				if real_idx == len(u.UtxoSet[wallet_id])-1 {
					u.UtxoSet[wallet_id] = u.UtxoSet[wallet_id][:real_idx]
				} else {
					u.UtxoSet[wallet_id] = append(u.UtxoSet[wallet_id][:real_idx], u.UtxoSet[wallet_id][real_idx+1:]...)
				}
				deleted = deleted + 1
			}
		}
	}

	for _, tx := range txs {
		if tx.IsCoinBaseTransaction() {
			continue
		}
		for _, input := range tx.TxInput {
			txid_string := hex.EncodeToString(input.TxId)
			if CheckTxIdInList(txid_string, id_list) == false {
				wallet_id := sha256.Sum256(input.PubKey)
				wallet_id_ := wallet_id[:]
				wallet_id_string := hex.EncodeToString(wallet_id_)
				u.UtxoSet[wallet_id_string] = append(u.UtxoSet[wallet_id_string], TxOutIndex{TxId: input.TxId, OutIndex: input.OutIndex, Out: input.Txout})
			}
		}
	}

}

func (u *UTXOset) Serialize() []byte {
	var BytesBuffer bytes.Buffer
	var encoder = gob.NewEncoder(&BytesBuffer)
	err := encoder.Encode(u)
	if err != nil {
		log.Panic(err)
	}
	return BytesBuffer.Bytes()
}

func DeserializeUTXOset(data []byte) UTXOset {
	var utxoset UTXOset
	var decoder = gob.NewDecoder(bytes.NewReader(data))
	err := decoder.Decode(&utxoset)
	if err != nil {
		log.Panic(err)
	}
	return utxoset
}

func (u *UTXOset) String() string {
	var lines []string
	lines = append(lines, fmt.Sprintf("----utxo length"))
	for i, _ := range u.UtxoSet {
		lines = append(lines, fmt.Sprintf("    i: %s   length: %d", i, len(u.UtxoSet[i])))
	}
	return strings.Join(lines, "\n")
}
