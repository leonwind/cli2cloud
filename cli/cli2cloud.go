package main

import (
	"bufio"
	"context"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"flag"
	"fmt"
	"io"
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

	var block cipher.Block
	if password != nil {
		keyphrase := sha256.Sum256([]byte(*password))
		block, err = aes.NewCipher(keyphrase[:])
		if err != nil {
			log.Fatal("Couldn't create a cipher. ", err)
		}

		gcm, err := cipher.NewGCM(block)
		if err != nil {
			log.Fatal("Couldn't init new GCM. ", err)
		}

		nonce := make([]byte, gcm.NonceSize())
		if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
			fmt.Println("Couldn't populate nounce. ", err)
		}
	}

	client, err := c.RegisterClient(ctx, &proto.Empty{})
	fmt.Printf("Your client ID: %s\n", client.Id)
	fmt.Printf("Share and monitor it live from cli2cloud.com/%s\n\n\n", client.Id)
	// Wait 3 seconds for user to copy the client ID
	time.Sleep(3 * time.Second)

	// TODO: Scan Stderr as well
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		row := scanner.Text()

		// Print original input to client as well
		fmt.Println(row)

		if block != nil {
			var encrypted []byte
			block.Encrypt(encrypted, []byte(row))
			row = string(encrypted)
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
	keyphrase := flag.String("encrypt", "", "Keyphrase to encrypt your data with.")
	flag.Parse()

	conn, err := grpc.Dial(":50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("Unable to connect to grpc", err)
	}

	client := proto.NewCli2CloudClient(conn)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if err := sendPipedMessages(client, ctx, keyphrase); err != nil {
		log.Fatal("Error while sending to server", err)
	}

	err = conn.Close()
	if err != nil {
		log.Fatal("Unable to close connection", err)
	}
}
