// This code is just for testing and verification of the encryption functions.
// They are not required to the actual workflow since the decryption will happen in the browser.

package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"golang.org/x/crypto/pbkdf2"
	"hash"
)

type StreamDecrypter struct {
	Block  cipher.Block
	Stream cipher.Stream
	Mac    hash.Hash
}

// KdfWithSalt derives a new key with a length of 32 bytes based on the user password and an existing salt.
func kdfWithSalt(password string, salt []byte) []byte {
	return pbkdf2.Key([]byte(password), salt, numPBKDF2Iterations, keyLength, sha256.New)
}

// NewStreamDecrypter provides a struct with all required information to decrypt the encrypted data stream.
func NewStreamDecrypter(password string, salt []byte, iv []byte) (*StreamDecrypter, error) {
	key := kdfWithSalt(password, salt)

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	stream := cipher.NewCTR(block, iv)
	mac := hmac.New(sha256.New, key)

	return &StreamDecrypter{
		Block:  block,
		Stream: stream,
		Mac:    mac,
	}, nil
}

// Decrypt decrypts the encrypted rows which are encoded as a hex string.
func (s *StreamDecrypter) Decrypt(row string) (*string, error) {
	encrypted, err := hex.DecodeString(row)
	if err != nil {
		return nil, err
	}

	if err := writeHash(encrypted, s.Mac); err != nil {
		return nil, err
	}

	bytePlaintext := make([]byte, len(encrypted))
	s.Stream.XORKeyStream(bytePlaintext, encrypted)

	plaintext := string(bytePlaintext)
	return &plaintext, nil
}
