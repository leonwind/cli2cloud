package api

import (
	"fmt"
	"log"
	"service/api/proto"
	"time"
)

func (s *Service) Subscribe(clientId *proto.ClientId, stream proto.Cli2Cloud_SubscribeServer) error {
	ctx := stream.Context()
	var line int64 = 0

	for {
		select {

		case <-ctx.Done():
			return nil

		default:
			rows, err := s.db.ReadContent(clientId.Id, line)
			if err != nil {
				return err
			}

			if line == 0 && (rows == nil || len(rows) == 0) {
				return fmt.Errorf("no output for client %s found", clientId.Id)
			}

			for _, row := range rows {
				if err := stream.Send(&proto.Payload{Body: row.Content}); err != nil {
					return err
				}

				log.Printf("Sending %s for client %s\n", row.Content, clientId.Id)
			}

			// Prevent database system calls spamming
			time.Sleep(500 * time.Millisecond)

			line += int64(len(rows))
		}
	}
}
