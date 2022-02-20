package api

import (
	"fmt"
	"io"
	"log"
	"service/api/pb"
)

const (
	idLength = 6
)

func (s *Service) Publish(stream pb.Cli2Cloud_PublishServer) error {
	line := 0

	for {
		var content *pb.Content
		content, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&pb.Empty{})
		}

		if err != nil {
			return err
		}

		message := fmt.Sprintf("Client %s, line %d: %s", content.Client.Id, line, content.Payload)
		log.Println(message)
		line++
	}
}
