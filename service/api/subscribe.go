package api

import "service/servicepb"

func (s *Service) Subscribe(_ *servicepb.Empty, stream servicepb.Cli2Cloud_SubscribeServer) error {
	return nil
}
