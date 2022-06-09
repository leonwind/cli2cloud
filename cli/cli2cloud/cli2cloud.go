package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/leonwind/cli2cloud/cli/cli2cloud/crypto"
	"github.com/leonwind/cli2cloud/cli/cli2cloud/proto"
	"github.com/leonwind/cli2cloud/cli/cli2cloud/streams"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"strconv"
	"time"
)

const (
	//serverIP = "localhost:50051" // local dev
	serverIP             = "cli2cloud.com:50051" // production
	randomPasswordLength = 16
	defaultDelayLength   = 3
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

func sendPipedMessages(c proto.Cli2CloudClient, ctx context.Context, password *string, delay time.Duration) error {
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

	keyURLSuffix := ""
	if password != nil {
		keyURLSuffix = fmt.Sprintf("#key=%s", *password)
	}

	fmt.Printf("Share and monitor it live from https://cli2cloud.com/%s%s\n\n", clientId.Id, keyURLSuffix)
	// Wait delay seconds for user to copy the client ID
	time.Sleep(delay * time.Second)

	// Create a messages stream which is reading from both Stdout and Stdin
	streamMessages := make(chan interface{})
	//defer close(streamMessages)
	go streams.CreateStreams(streamMessages)

	for res := range streamMessages {
		switch res.(type) {
		case bool:
			fmt.Println("Close channel")
			//_, err = stream.CloseAndRecv()
			//return err
			break
		default:
			row := res.(string)
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
	}

	_, err = stream.CloseAndRecv()
	return err
}

func parseFlags() (*string, time.Duration) {
	var passwordFlag stringFlag
	flag.Var(&passwordFlag, "encrypt", "Password to encrypt your data with.")
	generatePassword := flag.Bool("encrypt-random", false, "Generate a random password to encrypt your data.")

	var delayFlag stringFlag
	flag.Var(&delayFlag, "delay", "Delay before printing the command output to copy the client ID.")

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

	var delay int
	if delayFlag.set {
		delay, err = strconv.Atoi(delayFlag.value)
		if err != nil || delay < 0 {
			log.Fatal("Delay parameter argument is non parseable ", err)
		}
	} else {
		delay = defaultDelayLength
	}

	return password, time.Duration(delay)
}

func main() {
	password, delay := parseFlags()

	conn, err := grpc.Dial(serverIP, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("Unable to connect to gRPC server.", err)
	}

	client := proto.NewCli2CloudClient(conn)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if err := sendPipedMessages(client, ctx, password, delay); err != nil {
		log.Fatal("Error while sending to server.", err)
	}

	if err := conn.Close(); err != nil {
		log.Fatal("Unable to close connection.", err)
	}
}
