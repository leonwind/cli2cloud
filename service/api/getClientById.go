package api

import (
	"context"
	"log"
	"service/api/proto"
)

func (s *Service) GetClientById(_ context.Context, clientId *proto.ClientId) (*proto.Client, error) {
	encrypted, salt, iv, err := s.db.GetClientById(clientId.Id)
	if err != nil {
		log.Printf("Couldn't fetch client details for id %s.\n", clientId.Id)
		log.Println(err)
		return nil, err
	}

	client := proto.Client{
		Encrypted: encrypted,
		Salt:      salt,
		Iv:        iv,
		Timestamp: nil,
	}
	return &client, nil
}
