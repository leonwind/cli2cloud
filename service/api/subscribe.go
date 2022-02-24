package api

import (
	"fmt"
	"service/api/proto"
	"time"
)

func (s *Service) Subscribe(client *proto.Client, stream proto.Cli2Cloud_SubscribeServer) error {
	ctx := stream.Context()
	var row int64 = 0

	for {
		select {

		case <-ctx.Done():
			return nil

		default:
			contents, err := s.db.ReadContent(client, row)
			if err != nil {
				return err
			}

			if row == 0 && (contents == nil || len(contents) == 0) {
				return fmt.Errorf("no output for client %s found", client.Id)
			}

			for _, content := range contents {
				fmt.Printf("Sending %s for client %s\n", content.Payload, client.Id)
				if err := stream.Send(content); err != nil {
					return err
				}
			}

			// Prevent database system calls spamming
			time.Sleep(500 * time.Millisecond)

			row += int64(len(contents))
		}
	}
}
