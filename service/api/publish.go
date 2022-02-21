package api

import (
	"fmt"
	"io"
	"log"
	"service/api/proto"
)

func (s *Service) Publish(stream proto.Cli2Cloud_PublishServer) error {
	var row int64 = 0

	for {
		var content *proto.Content
		content, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&proto.Empty{})
		}

		if err != nil {
			return err
		}

		message := fmt.Sprintf("Client %s, line %d: %s", content.Client.Id, row, content.Payload)
		log.Println(message)

		if err := s.db.WriteContent(content, row); err != nil {
			log.Println("Couldn't write content to psql", err)
			return err
		}

		row++
	}
}
