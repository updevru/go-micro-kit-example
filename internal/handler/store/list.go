package store

import (
	pb "github.com/updevru/go-micro-kit-example/gen/store"
)

func (s *Handler) List(in *pb.ListRequest, stream pb.Store_ListServer) error {

	items, err := s.store.List()
	if err != nil {
		return err
	}

	for item := range items {
		err := stream.Send(mapResponse(item))
		if err != nil {
			return err
		}
	}

	return nil
}
