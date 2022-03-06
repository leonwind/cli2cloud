package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"github.com/leonwind/cli2cloud/service/api/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"os"
	"time"
)

type stringFlag struct {
	set   bool
	value string
}

func (sf *stringFlag) Set(x string) error {
	sf.value = x
	sf.set = true
	return nil
}

func (sf *stringFlag) String() string {
	return sf.value
}

func sendPipedMessages(c proto.Cli2CloudClient, ctx context.Context, password stringFlag) error {
	stream, err := c.Publish(ctx)
	if err != nil {
		return err
	}

	var s *StreamEncrypter
	if password.set {
		if password.value == "" {
			log.Fatal("Password cannot be empty.")
		}

		s, err = NewStreamEncrypter(password.value)
		if err != nil {
			log.Fatal("Can't create a Stream Encrypter.", err)
		}
	}

	client := proto.Client{
		Encrypted: s != nil,
		Salt:      s.GetSaltAsHex(),
		Iv:        s.GetIVAsHex(),
	}

	clientId, err := c.RegisterClient(ctx, &client)
	fmt.Printf("Your client ID: %s\n", clientId.Id)
	fmt.Printf("Share and monitor it live from cli2cloud.com/%s\n\n", clientId.Id)
	// Wait 2 seconds for user to copy the client ID
	time.Sleep(2 * time.Second)

	// TODO: Scan Stderr as well
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		row := scanner.Text()
		// Print original input to client as well
		fmt.Println(row)

		if s != nil {
			encryptedRow, err := s.Encrypt(row)
			if err != nil {
				log.Println("Can't encrypt the data.", err)
				return err
			}
			row = *encryptedRow
		}

		content := proto.PublishRequest{
			Payload:  &proto.Payload{Body: row},
			ClientId: clientId,
		}

		if err := stream.Send(&content); err != nil {
			return err
		}
	}

	_, err = stream.CloseAndRecv()
	return err
}

func main() {
	var password stringFlag
	flag.Var(&password, "encrypt", "Password to encrypto your data with.")
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

	if err := conn.Close(); err != nil {
		log.Fatal("Unable to close connection.", err)
	}
}
