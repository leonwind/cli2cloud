package api

import (
	"google.golang.org/grpc"
	"log"
	"net"
	"service/servicepb"
)

type Service struct {
	servicepb.UnimplementedCli2CloudServer
}

func NewServer() *Service {
	return &Service{}
}

func (s *Service) Start(ip string) error {
	lis, err := net.Listen("tcp", ip)
	if err != nil {
		return err
	}

	server := grpc.NewServer()
	servicepb.RegisterCli2CloudServer(server, s)
	log.Println("Registered server...")
	if err := server.Serve(lis); err != nil {
		return err
	}
	return nil
}
