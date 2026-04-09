package main

type Account struct {
	Address string
	Balance float64
	Nonce   uint64
}

type AccountManager struct {
	Accounts map[string]*Account
}

func NewAccountManager() *AccountManager {
	return &AccountManager{Accounts: make(map[string]*Account)}
}

func (am *AccountManager) CreateAccount() *Account {
	w := NewWallet()
	acc := &Account{
		Address: w.Address,
		Balance: 0,
		Nonce:   0,
	}
	am.Accounts[w.Address] = acc
	return acc
}
