package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/leonwind/cli2cloud/service/api/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func sendPipedMessages(c proto.Cli2CloudClient, ctx context.Context) error {
	stream, err := c.Publish(ctx)
	if err != nil {
		return err
	}

	client, err := c.RegisterClient(ctx, &proto.Empty{})
	fmt.Printf("Your client ID: %s\n", client.Id)
	fmt.Printf("Share it at cli2cloud.com/%s\n\n\n", client.Id)
	// Wait 3 seconds for user to copy the client ID
	time.Sleep(3 * time.Second)

	// TODO: Scan Stderr as well
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		content := proto.Content{
			Payload: fmt.Sprintf(scanner.Text()),
			Client:  client,
		}

		if err := stream.Send(&content); err != nil {
			return err
		}

		// Print original input to client as well
		fmt.Println(scanner.Text())
	}

	_, err = stream.CloseAndRecv()
	return err
}

func main() {
	conn, err := grpc.Dial(":50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("Unable to connect to grpc", err)
	}

	client := proto.NewCli2CloudClient(conn)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if err := sendPipedMessages(client, ctx); err != nil {
		log.Fatal("Error while sending to server", err)
	}

	err = conn.Close()
	if err != nil {
		log.Fatal("Unable to close connection", err)
	}
}
