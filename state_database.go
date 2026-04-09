package main

type StateDB struct {
	AccountMap map[string]float64
	ContractMap map[string]string
}

func NewStateDB() *StateDB {
	return &StateDB{
		AccountMap: make(map[string]float64),
		ContractMap: make(map[string]string),
	}
}

func (db *StateDB) SetBalance(addr string, amount float64) {
	db.AccountMap[addr] = amount
}

func (db *StateDB) GetBalance(addr string) float64 {
	return db.AccountMap[addr]
}
