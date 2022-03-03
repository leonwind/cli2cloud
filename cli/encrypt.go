package main

import (
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
	Block  cipher.Block
	Stream cipher.Stream
	IV     []byte
	Salt   []byte
	Mac    hash.Hash
}

const (
	numPBKDF2Iterations = 1024
	keyLength           = 32 // bytes
	saltLength          = 32 // bytes
)

// kdf derives a new key with a length of 32 bytes based on the user password and on a newly created salt.
func kdf(password string) ([]byte, []byte, error) {
	salt := make([]byte, saltLength)
	_, err := rand.Read(salt)
	if err != nil {
		return nil, nil, err
	}

	key := pbkdf2.Key([]byte(password), salt, numPBKDF2Iterations, keyLength, sha256.New)
	return key, salt, nil
}

// NewStreamEncrypter provides a struct with all required information to encrypt a data stream.
func NewStreamEncrypter(password string) (*StreamEncrypter, error) {
	key, salt, err := kdf(password)
	if err != nil {
		return nil, err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	iv := make([]byte, block.BlockSize())
	_, err = rand.Read(iv)
	if err != nil {
		return nil, err
	}

	stream := cipher.NewCTR(block, iv)
	mac := hmac.New(sha256.New, key)

	return &StreamEncrypter{
		Block:  block,
		Stream: stream,
		IV:     iv,
		Salt:   salt,
		Mac:    mac,
	}, nil
}

// Encrypt encrypts the given row and returns the byte array encoded as a Hex string.
func (s *StreamEncrypter) Encrypt(row string) (*string, error) {
	plaintext := []byte(row)
	encrypted := make([]byte, len(plaintext))
	s.Stream.XORKeyStream(encrypted, plaintext)

	err := writeHash(encrypted, s.Mac)
	if err != nil {
		return nil, err
	}

	hexString := hex.EncodeToString(encrypted)
	return &hexString, nil
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
