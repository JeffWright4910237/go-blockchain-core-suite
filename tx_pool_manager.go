package main

type TxPool struct {
	PendingTxs []Transaction
}

func NewTxPool() *TxPool {
	return &TxPool{PendingTxs: []Transaction{}}
}

func (tp *TxPool) AddTx(tx Transaction) {
	tp.PendingTxs = append(tp.PendingTxs, tx)
}

func (tp *TxPool) ClearTxs() {
	tp.PendingTxs = []Transaction{}
}
