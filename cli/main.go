package main

import (
	"context"
	"fmt"
	"github.com/leonwind/cli2cloud/service/api/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

func sendMessages(c pb.Cli2CloudClient, ctx context.Context) error {
	stream, err := c.Publish(ctx)
	if err != nil {
		return err
	}

	client, err := c.RegisterClient(ctx, &pb.Empty{})
	fmt.Println(client.Id)

	for i := 0; i < 1000; i++ {
		content := pb.Content{
			Payload: fmt.Sprintf("Hello World %d", i),
			Client:  client,
		}

		if err := stream.Send(&content); err != nil {
			return err
		}
		time.Sleep(1 * time.Second)
	}

	_, err = stream.CloseAndRecv()
	return err
}

func receiveMessages(c pb.Cli2CloudClient, ctx context.Context) error {
	client := &pb.Client{Id: "1WRTFE"}
	stream, err := c.Subscribe(ctx, client)
	if err != nil {
		return err
	}

	for {
		content, err := stream.Recv()
		if err != nil {
			return err
		}

		// We need to reference row since it is optional and thus a pointer.
		log.Printf("Received %s as row %d\n", content.Payload, *content.Row)
	}
}

func main() {
	conn, err := grpc.Dial(":50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("Unable to connect to grpc", err)
	}

	client := pb.NewCli2CloudClient(conn)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	//if err := sendMessages(client, ctx); err != nil {
	//log.Fatal("Error while sending to server", err)
	//}

	if err := receiveMessages(client, ctx); err != nil {
		log.Fatal("Error while receiving from server", err)
	}

	err = conn.Close()
	if err != nil {
		log.Fatal("Unable to close connection", err)
	}
}
