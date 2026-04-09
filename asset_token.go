package main

type Token struct {
	Name     string
	Symbol   string
	Total    float64
	Decimals int
}

func TransferToken(am *AccountManager, from, to string, amount float64) bool {
	if am.Accounts[from].Balance < amount {
		return false
	}
	am.Accounts[from].Balance -= amount
	am.Accounts[to].Balance += amount
	return true
}
