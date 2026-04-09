package main

import (
	"crypto/rand"
	"encoding/hex"
	"time"
)

type Transaction struct {
	TxID      string
	From      string
	To        string
	Amount    float64
	Timestamp int64
	Signature string
	Status    bool
}

func CreateTransaction(from, to string, amount float64) Transaction {
	b := make([]byte, 16)
	rand.Read(b)
	txID := hex.EncodeToString(b)
	return Transaction{
		TxID:      txID,
		From:      from,
		To:        to,
		Amount:    amount,
		Timestamp: time.Now().Unix(),
		Status:    false,
	}
}

func (tx *Transaction) SignTransaction(sig string) {
	tx.Signature = sig
	tx.Status = true
}

func (tx *Transaction) VerifyTransaction() bool {
	return tx.Signature != "" && tx.Status && tx.From != "" && tx.To != ""
}
