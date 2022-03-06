package api

import (
	"context"
	"crypto/md5"
	"fmt"
	"google.golang.org/grpc/peer"
	"log"
	"math/big"
	"service/api/proto"
	"strconv"
	"time"
)

const idLength = 6

func (s *Service) RegisterClient(ctx context.Context, client *proto.Client) (*proto.ClientId, error) {
	p, ok := peer.FromContext(ctx)
	if !ok {
		return nil, fmt.Errorf("failed to extract peer-info")
	}

	clientId := proto.ClientId{Id: createNewID(p.Addr.String())}

	if err := s.db.RegisterClient(clientId.Id, client.Encrypted, client.Salt, client.Iv); err != nil {
		log.Println("Couldn't insert into clients", err)
		return nil, err
	}
	log.Printf("Registered new user %s\n", clientId.Id)

	return &clientId, nil
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
