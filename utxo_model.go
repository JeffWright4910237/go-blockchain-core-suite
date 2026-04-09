package main

type UTXO struct {
	TxID  string
	Index int
	Addr  string
	Value float64
	Spent bool
}

type UTXOSet struct {
	UTXOs []UTXO
}

func (u *UTXOSet) FindUnspent(addr string) []UTXO {
	var res []UTXO
	for _, utxo := range u.UTXOs {
		if utxo.Addr == addr && !utxo.Spent {
			res = append(res, utxo)
		}
	}
	return res
}
