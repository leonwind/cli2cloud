package api

import (
	"google.golang.org/grpc"
	"log"
	"net"
	"service/api/pb"
	"service/internal/storage"
)

type Service struct {
	pb.UnimplementedCli2CloudServer
	db storage.Database
}

func NewServer() (*Service, error) {
	psql, err := storage.InitPostgres()
	if err != nil {
		return nil, err
	}
	log.Println("Connected to database")

	service := Service{
		db: psql,
	}

	return &service, nil
}

func (s *Service) Start(ip string) error {
	lis, err := net.Listen("tcp", ip)
	if err != nil {
		return err
	}

	server := grpc.NewServer()
	pb.RegisterCli2CloudServer(server, s)
	log.Println("Registered server...")

	if err := server.Serve(lis); err != nil {
		return err
	}

	return nil
}
