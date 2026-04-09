package main

type MultiSigWallet struct {
	Owners   []string
	Required int
}

func (ms *MultiSigWallet) VerifySigns(signs []string) bool {
	return len(signs) >= ms.Required
}
