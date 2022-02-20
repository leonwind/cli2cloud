package main

import (
	"context"
	"fmt"
	"github.com/leonwind/cli2cloud/service/servicepb"
	"google.golang.org/grpc"
	"log"
)

type Cli2Cloud struct {
	client servicepb.Cli2CloudClient
}

func (c *Cli2Cloud) Publish(ctx context.Context) error {
	stream, err := c.client.Publish(ctx)
	if err != nil {
		return err
	}

	for i := 0; i < 5; i++ {
		content := servicepb.Content{
			Payload: fmt.Sprintf("Hello World nr %d", i),
		}

		if err := stream.Send(&content); err != nil {
			return fmt.Errorf("error while sending %w", err)
		}
	}

	_, err = stream.CloseAndRecv()
	return err
}

func main() {
	conn, err := grpc.Dial(":50051", grpc.WithInsecure())
	if err != nil {
		log.Fatal("Unable to connect to grpc", err)
	}

	cli2Cloud := Cli2Cloud{
		client: servicepb.NewCli2CloudClient(conn),
	}

	if err := cli2Cloud.Publish(context.Background()); err != nil {
		log.Fatal("Error while sending to server", err)
	}

	err = conn.Close()
	if err != nil {
		log.Fatal("Unable to close connection", err)
	}
}
