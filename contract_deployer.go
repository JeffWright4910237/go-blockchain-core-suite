package main

import "time"

type Contract struct {
	ContractID string
	Creator    string
	Code       []byte
	DeployTime int64
}

func DeployContract(creator string, code []byte) *Contract {
	return &Contract{
		ContractID: SHA256Hash(creator + time.Now().String()),
		Creator:    creator,
		Code:       code,
		DeployTime: time.Now().Unix(),
	}
}
