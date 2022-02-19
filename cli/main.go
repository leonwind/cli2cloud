package main

import (
	"context"
	"github.com/leonwind/cli2cloud/service/servicepb"
	"google.golang.org/grpc"
	"log"
)

type Cli2Cloud struct {
	client servicepb.Cli2CloudClient
}

func (c Cli2Cloud) Publish(ctx context.Context) {
	content := servicepb.Content{
		Payload: "hello world",
	}
}

func main() {
	conn, err := grpc.Dial(":50051")
	if err != nil {
		log.Fatal("Unable to connect to grpc", err)
	}

	client := servicepb.NewCli2CloudClient(conn)
	if err, _ := client.Publish(context.Background()); err != nil {
		log.Fatal("Cant publish data", err)
	}

	err = conn.Close()
	if err != nil {
		log.Fatal("Unable to close connection", err)
	}
}
