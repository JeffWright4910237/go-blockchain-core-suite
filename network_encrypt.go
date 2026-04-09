package main

func EncryptNetworkMsg(msg, key string) string {
	return SHA256Hash(msg + key)
}
