package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
)

func main() {
	fmt.Println("Hello world")
}

func encrypt(message string, key rsa.PublicKey) string {
	label := []byte("OAEP Encrypted")
	random := rand.Reader
	// if label is not needed, nil 
	cipherText, err := rsa.EncryptOAEP(sha256.New(), random, &key, message, label)
	if err != nil {
		fmt.Println(err)
	}
	return base64.StdEncoding.EncodeToString(cipherText)
}