package api

import (
	"fmt"
	"io"
	"log"
	"service/api/proto"
	"service/internal/storage"
)

func (s *Service) Publish(stream proto.Cli2Cloud_PublishServer) error {
	var line int64 = 0

	for {
		var request *proto.PublishRequest
		request, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&proto.Empty{})
		}

		if err != nil {
			return err
		}

		row := storage.Row{
			Content: request.Payload.Body,
			Line:    line,
		}

		message := fmt.Sprintf("Received from client %s, line %d: %s", request.ClientId.Id, line, row.Content)
		log.Println(message)

		if err := s.db.WriteContent(request.ClientId.Id, row); err != nil {
			log.Println("Couldn't write content to psql", err)
			return err
		}

		line++
	}
}
