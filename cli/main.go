package main

import (
	"context"
	"fmt"
	"github.com/leonwind/cli2cloud/service/api/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func sendMessages(c pb.Cli2CloudClient, ctx context.Context) error {
	stream, err := c.Publish(ctx)
	if err != nil {
		return err
	}

	client, err := c.RegisterClient(ctx, &pb.Empty{})
	fmt.Println(client.Id)

	for i := 0; i < 10; i++ {
		content := pb.Content{
			Payload: fmt.Sprintf("Hello World %d", i),
			Client:  client,
		}

		if err := stream.Send(&content); err != nil {
			return err
		}
	}

	_, err = stream.CloseAndRecv()
	return err
}

func main() {
	conn, err := grpc.Dial(":50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("Unable to connect to grpc", err)
	}

	client := pb.NewCli2CloudClient(conn)
	if err := sendMessages(client, context.Background()); err != nil {
		log.Fatal("Error while sending to server", err)
	}

	err = conn.Close()
	if err != nil {
		log.Fatal("Unable to close connection", err)
	}
}
