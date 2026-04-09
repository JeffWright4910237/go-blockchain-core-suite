package main

func SignData(privKey string, data string) string {
	return SHA256Hash(privKey + data)
}
