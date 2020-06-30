package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
)

func main() {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		fmt.Println(err)
	}
	secretMessage := "The ants in france stay mainly on the plants"

	publicKey := privateKey.PublicKey

	encryptedMessage := encrypt(secretMessage, publicKey)

	fmt.Println("Cipher Text:", encryptedMessage)

	decrypt(encryptedMessage, *privateKey)

}

func encrypt(message string, key rsa.PublicKey) string {
	label := []byte("OAEP Encrypted")
	random := rand.Reader
	// if label is not needed, nil 
	cipherText, err := rsa.EncryptOAEP(sha256.New(), random, &key, []byte(message), label)
	
	checkError(err)
	
	return base64.StdEncoding.EncodeToString(cipherText)
}

func decrypt(cipherText string, privateKey rsa.PrivateKey) string {
	cipher, err := base64.StdEncoding.DecodeString(cipherText)
	
	checkError(err)

	label := []byte("OAEP Encrypted")
	rng := rand.Reader
	plainText, err := rsa.DecryptOAEP(sha256.New(), rng, &privateKey, cipher, label)

	checkError(err)

	fmt.Println("Original string: ", string(plainText))
	return string(plainText)
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}