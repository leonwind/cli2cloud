package api

import (
	"crypto/md5"
	"fmt"
	"google.golang.org/grpc/peer"
	"math/big"
	"service/servicepb"
	"strconv"
	"time"
)

const (
	idLength = 6
)

func (s *Service) Publish(stream servicepb.Cli2Cloud_PublishServer) error {
	p, ok := peer.FromContext(stream.Context())
	if !ok {
		return fmt.Errorf("failed to extract peer-info")
	}
	clientID := createNewID(p.Addr.String())

	return nil
}

// Create valid and unique ID for a client based on ones ip address and the current timestamp.
func createNewID(ipAddr string) string {
	// Create a unique ID for the client by hashing the ip address and the current
	// timestamp and encode the hash using base62 ([0-9][A-Z][a-z]).
	// Use MD5 since since it is significantly faster than SHA2 and not security relevant.
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	hash := md5.Sum([]byte(ipAddr + timestamp))

	// Encode hash into base62 and use the first 6 characters as the ID:
	// 62^6 ~ 56E9 different unique IDs
	uniqueID := encodeBase62(hash)[:idLength]
	return uniqueID
}

func encodeBase62(toEncode [16]byte) string {
	encoded := big.NewInt(0)
	encoded.SetBytes(toEncode[:])
	return encoded.Text(62)
}
