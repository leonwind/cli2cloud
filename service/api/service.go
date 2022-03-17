package api

import (
	"github.com/leonwind/cli2cloud/service/api/proto"
	"github.com/leonwind/cli2cloud/service/internal/storage"
	"google.golang.org/grpc"
	"log"
	"net"
)

type Service struct {
	proto.UnimplementedCli2CloudServer
	db storage.Database
}

func NewServer(dbUrl string) (*Service, error) {
	psql, err := storage.InitPostgres(dbUrl)
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
	proto.RegisterCli2CloudServer(server, s)
	log.Println("Registered server...")

	if err := server.Serve(lis); err != nil {
		return err
	}

	return nil
}
