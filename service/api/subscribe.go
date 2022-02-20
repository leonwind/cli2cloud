package api

import "service/api/pb"

func (s *Service) Subscribe(client *pb.Client, stream pb.Cli2Cloud_SubscribeServer) error {
	var row int64 = 0
	ctx := stream.Context()

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
