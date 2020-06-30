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
	cipherText, err := rsa.EncryptOAEP(sha256.New(), random, &key, []byte(message), label)
	if err != nil {
		fmt.Println(err)
	}
	return base64.StdEncoding.EncodeToString(cipherText)
}

func decrypt(cipherText string, privateKey rsa.PrivateKey) string {
	cipher, err := base64.RawStdEncoding.DecodeString(cipherText)
	if err != nil {
		fmt.Println(err)
	}
	label := []byte("OAEP Encrypted")
	rng := rand.Reader
	plainText, err := rsa.DecryptOAEP(sha256.New(), rng, &privateKey, cipher, label)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Original string: ", string(plainText))
	return string(plainText)
}