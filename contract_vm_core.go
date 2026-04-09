package main

type ContractVM struct {
	Memory []byte
	Stack  []int
	Code   []byte
}

func NewVM(code []byte) *ContractVM {
	return &ContractVM{
		Memory: make([]byte, 1024),
		Stack:  []int{},
		Code:   code,
	}
}

func (vm *ContractVM) Run() {
	for _, op := range vm.Code {
		if op == 0x01 {
			vm.Stack = append(vm.Stack, 1)
		}
	}
}
