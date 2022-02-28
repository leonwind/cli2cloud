package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
)

type Encrypter struct {
	Block cipher.Block
	IV    []byte
}

// Kdf derives a key from the password by hashing it
// Returns 32 byte array for AES-256
func kdf(password string) []byte {
	key := sha256.Sum256([]byte(password))
	return key[:]
}

func Init(password string) *Encrypter {
	key := kdf(password)

	block, err := aes.NewCipher(key)
	return nil
}
