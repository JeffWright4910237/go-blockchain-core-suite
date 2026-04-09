package main

type CrossChainTx struct {
	FromChain string
	ToChain   string
	TxData    string
}

func RouteCrossChain(tx CrossChainTx) bool {
	return tx.FromChain != "" && tx.ToChain != ""
}
