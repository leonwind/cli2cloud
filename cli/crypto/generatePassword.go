package crypto

import (
	"crypto/rand"
	"math/big"
)

const base = 62

func GeneratePassword(passwordLength int) (*string, error) {
	randomBytes := make([]byte, 32)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return nil, err
	}

	encoded := big.NewInt(0)
	encoded.SetBytes(randomBytes[:])
	password := encoded.Text(base)[:passwordLength]
	return &password, nil
}
