package server

import (
	"context"
	"github.com/leonwind/cli2cloud/server/serverpb"
	"sync"
	"sync/atomic"
)

type Server struct {
	serverpb.UnimplementedCli2CloudServer
	mu sync.RWMutex

	ctr    uint32
	subCtr uint32

	subChs map[uint32]map[uint32]chan string
}

func NewServer() *Server {
	s := new(Server)
	return s
}

func (s *Server) Publish(stream serverpb.Cli2Cloud_PublishServer) error {
	return nil
}

func (s *Server) Subscribe(regMsg *serverpb.RegisterMessage, stream serverpb.Cli2Cloud_SubscribeServer) error {

	return nil
}

func (s *Server) Register(_ context.Context, _ *serverpb.Empty) (*serverpb.RegisterMessage, error) {
	return &serverpb.RegisterMessage{ClientID: atomic.AddUint32(&s.ctr, 1)}, nil
}
