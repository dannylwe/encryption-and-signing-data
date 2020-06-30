package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
	"log"
)

func main() {
	msg := []byte("verify this message")

	msgHash := sha256.New()
	// The hash is what we actually sign
	_, err := msgHash.Write(msg)
	if err != nil {
		fmt.Println(err)
	}

	msgHashSum := msgHash.Sum(nil)
	privateKey, _ := rsa.GenerateKey(rand.Reader, 2048)
	publicKey := privateKey.PublicKey
	signiture, err := rsa.SignPSS(rand.Reader, privateKey, crypto.SHA256, msgHashSum, nil)
	if err != nil {
		fmt.Println(err)
	}

	err = rsa.VerifyPSS(&publicKey, crypto.SHA256, msgHashSum, signiture, nil)
	if err != nil {
		log.Fatal("could not verify signiture: ", err)
	}
	fmt.Println("signiture verified")
}