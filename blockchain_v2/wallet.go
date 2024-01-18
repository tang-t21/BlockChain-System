package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"log"
)

type Wallet struct {
	PrivateKey ecdsa.PrivateKey
	PublicKey  []byte // used as id, no address for a wallet
}

func CreateWallet() *Wallet {
	curve := elliptic.P256()
	key, err := ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {
		log.Panic(err)
	}
	pubKey := append(key.PublicKey.X.Bytes(), key.PublicKey.Y.Bytes()...)
	wallet := Wallet{
		PrivateKey: *key,
		PublicKey:  pubKey,
	}
	return &wallet
}

func (w *Wallet) SignTransaction(tx *Transaction) {
	// TODO: IF COINBASE, NO SIGN
	if tx.IsCoinBaseTransaction() {
		return
	}
	// sign the combination of TxId in each TxInput and OutIndex
	for i, in := range tx.TxInput {
		outindexbyte := IntToBytes(in.OutIndex)
		data_to_sign := append(in.TxId, outindexbyte...)
		sig, err := ecdsa.SignASN1(rand.Reader, &w.PrivateKey, data_to_sign)
		if err != nil {
			log.Panic(err)
		}
		tx.TxInput[i].Signature = sig
	}
}
