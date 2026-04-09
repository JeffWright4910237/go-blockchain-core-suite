package main

import "fmt"

func LogInfo(msg string) {
	fmt.Printf("[INFO] %s\n", msg)
}

func LogError(msg string) {
	fmt.Printf("[ERROR] %s\n", msg)
}
