package api

import (
	"fmt"
	"log"
	"service/api/proto"
	"time"
)

func (s *Service) Subscribe(client *proto.Client, stream proto.Cli2Cloud_SubscribeServer) error {
	ctx := stream.Context()
	var row int64 = 0
	log.Printf("Subscribe for client %s\n", client.Id)

	for {
		select {

		case <-ctx.Done():
			return nil

		default:
			contents, err := s.db.ReadContent(client, row)
			if err != nil {
				return err
			}

			for _, content := range contents {
				fmt.Printf("Sending %s for client %s\n", content.Payload, client.Id)
				if err := stream.Send(content); err != nil {
					return err
				}
			}

			time.Sleep(500 * time.Millisecond)

			row += int64(len(contents))
		}
	}
}
