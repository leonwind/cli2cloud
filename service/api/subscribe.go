package api

import (
	"log"
	"service/api/proto"
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
				if err := stream.Send(content); err != nil {
					return err
				}
			}

			row += int64(len(contents))
		}
	}
}
