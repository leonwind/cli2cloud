package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/leonwind/cli2cloud/service/api/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func sendPipedMessages(c proto.Cli2CloudClient, ctx context.Context, password *string) error {
	stream, err := c.Publish(ctx)
	if err != nil {
		return err
	}

	var s *StreamEncrypter
	if password != nil {
		s, err = NewStreamEncrypter(*password)
		if err != nil {
			log.Fatal("Can't create a Stream Encrypter.", err)
		}
	}

	client, err := c.RegisterClient(ctx, &proto.Empty{})
	fmt.Printf("Your client ID: %s\n", client.Id)
	fmt.Printf("Share and monitor it live from cli2cloud.com/%s\n\n", client.Id)
	// Wait 3 seconds for user to copy the client ID
	time.Sleep(3 * time.Second)

	// TODO: Scan Stderr as well
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		row := scanner.Text()
		// Print original input to client as well
		fmt.Println(row)

		if s != nil {
			encryptedRow, err := s.Encrypt(row)
			if err != nil {
				log.Fatal("Can't encrypt the data.", err)
			}
			row = *encryptedRow
			fmt.Printf("Encrypted row: %s\n", row)
		}

		content := proto.Content{
			Payload: row,
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
	password := flag.String("encrypt", "", "Password to encrypt your data with.")
	flag.Parse()

	conn, err := grpc.Dial(":50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("Unable to connect to gRPC server.", err)
	}

	client := proto.NewCli2CloudClient(conn)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if err := sendPipedMessages(client, ctx, password); err != nil {
		log.Fatal("Error while sending to server.", err)
	}

	err = conn.Close()
	if err != nil {
		log.Fatal("Unable to close connection.", err)
	}
}
