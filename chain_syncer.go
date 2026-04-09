package main

func SyncLocalChain(local, remote []Block) []Block {
	if len(remote) > len(local) {
		return remote
	}
	return local
}
