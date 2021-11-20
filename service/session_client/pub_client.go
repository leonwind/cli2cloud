package session_client

import (
	"context"
	"fmt"
	"github.com/leonwind/cli2cloud/service/serverpb"
	"google.golang.org/grpc"
)

type PubClient struct {
	sessionID string
	pubStream serverpb.Cli2Cloud_PublishClient
}

func NewPubClient(serviceIP string) (*PubClient, error) {
	pc := new(PubClient)
	conn, err := grpc.Dial(serviceIP, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	serverClient := serverpb.NewCli2CloudClient(conn)
	publishStream, err := serverClient.Publish(context.Background())
	if err != nil {
		return nil, err
	}
	md, err := publishStream.Header()
	if err != nil {
		return nil, fmt.Errorf("failed to extract metadata: %v", err)
	}
	if len(md.Get("sessionid")) <= 0 {
		return nil, fmt.Errorf("failed to get id from metadata")
	}
	pc.sessionID = md.Get("sessionid")[0]
	return pc, nil
}

func (pc *PubClient) Publish(output *serverpb.Output) error {
	return pc.pubStream.Send(output)
}

func (pc *PubClient) GetSessionID() string {
	return pc.sessionID
}
