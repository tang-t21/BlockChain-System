package main

type TxIn struct {
	TxId      []byte
	OutIndex  int
	Txout     TxOut
	Signature []byte
	PubKey    []byte
}
