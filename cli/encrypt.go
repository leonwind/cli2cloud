// Encrypt the data stream using the AES128-CBC Mode.

package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"golang.org/x/crypto/pbkdf2"
	"hash"
)

type StreamEncrypter struct {
	Block cipher.Block
	Mode  cipher.BlockMode
	IV    []byte
	Salt  []byte
	Mac   hash.Hash
}

const (
	numPBKDF2Iterations = 1024
	keyLength           = 32 // bytes = 256 bits
	saltLength          = 32 // bytes = 256 bits
)

// Kdf derives a new key with a length of 32 bytes based on the user password and on a newly created salt.
func kdf(password []byte) ([]byte, []byte, error) {
	salt := make([]byte, saltLength)
	if _, err := rand.Read(salt); err != nil {
		return nil, nil, err
	}

	key := pbkdf2.Key(password, salt, numPBKDF2Iterations, keyLength, sha256.New)
	return key, salt, nil
}

// NewStreamEncrypter provides a struct with all required information to encrypt a data stream.
func NewStreamEncrypter(password string) (*StreamEncrypter, error) {
	key, salt, err := kdf([]byte(password))
	if err != nil {
		return nil, err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	iv := make([]byte, block.BlockSize())

	_, err = rand.Read(iv)
	fmt.Println("Key: ", hex.EncodeToString(key))
	fmt.Println("IV: ", hex.EncodeToString(iv))
	fmt.Println("Salt: ", hex.EncodeToString(salt))
	if err != nil {
		return nil, err
	}

	blockMode := cipher.NewCBCEncrypter(block, iv)
	mac := hmac.New(sha256.New, key)
	fmt.Println("BlockSize: ", block.BlockSize())

	return &StreamEncrypter{
		Block: block,
		Mode:  blockMode,
		IV:    iv,
		Salt:  salt,
		Mac:   mac,
	}, nil
}

// Encrypt encrypts the given row and returns the byte array encoded as a Hex string.
func (s *StreamEncrypter) Encrypt(row string) (*string, error) {
	plaintext := pkcs7Padding([]byte(row), s.Block.BlockSize())
	fmt.Println("Received: ", hex.EncodeToString(plaintext))
	encrypted := make([]byte, len(plaintext))
	s.Mode.CryptBlocks(encrypted, plaintext)

	hexString := hex.EncodeToString(encrypted)
	return &hexString, nil
}

func pkcs7Padding(src []byte, blockSize int) []byte {
	paddingLength := blockSize - (len(src) % blockSize)
	fmt.Println(blockSize, len(src), paddingLength)
	padding := bytes.Repeat([]byte{byte(paddingLength)}, paddingLength)
	return append(src, padding...)
}

func writeHash(encrypted []byte, mac hash.Hash) error {
	m, err := mac.Write(encrypted)
	if err != nil {
		return err
	}

	if m != len(encrypted) {
		return fmt.Errorf("can't write all bytes to HMAC")
	}

	return nil
}
