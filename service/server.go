package server

import (
	"crypto/md5"
	"fmt"
	"github.com/leonwind/cli2cloud/server/serverpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
	"math/big"
	"net"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
)

type Server struct {
	serverpb.UnimplementedCli2CloudServer
	mu          sync.RWMutex
	subServices map[string]*subService
}

func NewServer() *Server {
	s := new(Server)
	s.subServices = make(map[string]*subService)
	return s
}

func (s *Server) Start(ip string) error {
	lis, err := net.Listen("tcp", ip)
	if err != nil {
		return err
	}
	server := grpc.NewServer()
	serverpb.RegisterCli2CloudServer(server, s)
	if err := server.Serve(lis); err != nil {
		return err
	}
	return nil
}

func (s *Server) Publish(stream serverpb.Cli2Cloud_PublishServer) error {
	p, ok := peer.FromContext(stream.Context())
	if !ok {
		return fmt.Errorf("failed to extract peer-info")
	}
	clientID := createUniqueID(p.Addr.String())
	md := metadata.New(map[string]string{"clientid": clientID})
	if err := stream.SendHeader(md); err != nil {
		return err
	}

	subSrv := newSubService()
	s.mu.Lock()
	s.subServices[clientID] = subSrv
	s.mu.Unlock()

	for {
		output, err := stream.Recv()
		if err != nil {
			return fmt.Errorf("failed to receive output: %v", err)
		}
		subSrv.addNewOutput(output)
		// TODO: psql stuff
	}
}

func (s *Server) Subscribe(_ *serverpb.Empty, stream serverpb.Cli2Cloud_SubscribeServer) error {
	md, ok := metadata.FromIncomingContext(stream.Context())
	if !ok {
		return fmt.Errorf("failed to get metadata")
	}
	if len(md.Get("clientid")) <= 0 {
		return fmt.Errorf("failed to get clientID from metadata")
	}

	clientID := md.Get("clientid")[0]
	// TODO also lookup psql
	s.mu.RLock()
	subSrv, ok := s.subServices[clientID]
	s.mu.RUnlock()
	if !ok {
		return fmt.Errorf("nothing published for this client")
	}

	ch, id := subSrv.newSubscription()
	defer subSrv.cancelSubscription(id)
	for output := range ch {
		err := stream.Send(output)
		if err != nil {
			return fmt.Errorf("failed to send output: %v", err)
		}
	}
	return nil
}

type subService struct {
	mu   sync.Mutex
	ctr  uint32
	subs map[uint32]chan *serverpb.Output
}

func newSubService() *subService {
	sub := new(subService)
	sub.subs = make(map[uint32]chan *serverpb.Output)
	return sub
}

func (sub *subService) addNewOutput(output *serverpb.Output) {
	sub.mu.Lock()
	for _, ch := range sub.subs {
		ch <- output
	}
	sub.mu.Unlock()
}

func (sub *subService) newSubscription() (<-chan *serverpb.Output, uint32) {
	subID := atomic.AddUint32(&sub.ctr, 1)
	ch := make(chan *serverpb.Output, 1024)
	sub.mu.Lock()
	sub.subs[subID] = ch
	sub.mu.Unlock()
	return ch, subID
}

func (sub *subService) cancelSubscription(subID uint32) {
	sub.mu.Lock()
	delete(sub.subs, subID)
	sub.mu.Unlock()
}

func createUniqueID(ipAddr string) string {
	timestamp := strconv.FormatInt(time.Now().UnixNano(), 10)
	hash := md5.Sum([]byte(ipAddr + timestamp))
	return encodeBase62(hash)[:5]
}

func encodeBase62(toEncode [16]byte) string {
	encoded := big.NewInt(0)
	encoded.SetBytes(toEncode[:])
	return encoded.Text(62)
}
