package main

import (
	"encoding/hex"
	"encrypt-string/utils"
	"fmt"
)

func main() {
	/*
					  bytes   hex
		   AES-128    16      32
		   AES-192    24      48
		   AES-256    32      64
	*/
	AESKey, _ := hex.DecodeString("6368616e676520746869732070617373776f726420746f206120736563726574")

	// Encrypt
	base64Ciphertext, _ := utils.EncryptBytes(AESKey, []byte("I am A Message"))
	fmt.Printf("%x\n", base64Ciphertext)

	// Decrypt
	plaintext, _ := utils.DecryptBytes(AESKey, base64Ciphertext)
	fmt.Printf("%s\n", plaintext)
}
