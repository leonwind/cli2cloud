package session_client

import (
	"context"
	"github.com/leonwind/cli2cloud/service/serverpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type SubClient struct {
	serverClient serverpb.Cli2CloudClient
}

func NewSubClient(serviceIP string) (*SubClient, error) {
	sc := new(SubClient)
	conn, err := grpc.Dial(serviceIP, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	sc.serverClient = serverpb.NewCli2CloudClient(conn)
	return sc, nil
}

func (sc *SubClient) Subscribe(sessionToken string, ch chan<- *serverpb.Output, stop chan bool) error {
	md := metadata.New(map[string]string{"sessionid": sessionToken})
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	stream, err := sc.serverClient.Subscribe(ctx, &serverpb.Empty{})
	if err != nil {
		return err
	}
	outputCh := make(chan *serverpb.Output, 1)
	errCh := make(chan error, 1)
	go func() {
		for {
			output, err := stream.Recv()
			outputCh <- output
			errCh <- err
		}
	}()
	for {
		select {
		case output := <-outputCh:
			ch <- output
		case <-stop:
			return nil
		}
	}
}
