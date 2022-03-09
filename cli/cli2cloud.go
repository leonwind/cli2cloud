package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/leonwind/cli2cloud/cli/crypto"
	"github.com/leonwind/cli2cloud/cli/streams"
	"github.com/leonwind/cli2cloud/service/api/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

const (
	serverIP             = "167.99.140.19:50051"
	randomPasswordLength = 16
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

func sendPipedMessages(c proto.Cli2CloudClient, ctx context.Context, password *string) error {
	stream, err := c.Publish(ctx)
	if err != nil {
		return err
	}

	var s *crypto.StreamEncrypter
	if password != nil {
		s, err = crypto.NewStreamEncrypter(*password)
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
	fmt.Printf("Share and monitor it live from https://cli2cloud.com/%s\n\n", clientId.Id)
	// Wait 2 seconds for user to copy the client ID
	time.Sleep(2 * time.Second)

	// Create a messages stream which is reading from both Stdout and Stdin
	streamMessages := make(chan string)
	go streams.CreateStreams(streamMessages)

	for row := range streamMessages {
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

func parseFlags() *string {
	var passwordFlag stringFlag
	flag.Var(&passwordFlag, "encrypt", "Password to encrypt your data with.")
	generatePassword := flag.Bool("encrypt-random", false, "Generate a random password to encrypt your data.")
	flag.Parse()

	if passwordFlag.set && passwordFlag.value == "" {
		log.Fatal("Password can not be empty.")
	}

	if passwordFlag.set && *generatePassword {
		log.Fatal("Can't set a password and generate one.")
	}

	var password *string = nil
	var err error = nil

	if passwordFlag.set {
		password = &passwordFlag.value
	} else if *generatePassword {
		password, err = crypto.GeneratePassword(randomPasswordLength)
		if err != nil {
			log.Fatal("Error while generating the random password", err)
		}

		fmt.Printf("Your password: %s\n", *password)
	}

	return password
}

func main() {
	password := parseFlags()

	conn, err := grpc.Dial(serverIP, grpc.WithTransportCredentials(insecure.NewCredentials()))
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
