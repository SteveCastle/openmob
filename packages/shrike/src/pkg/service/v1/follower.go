package v1

import (
	"context"

	v1 "github.com/SteveCastle/openmob/packages/shrike/src/pkg/api/v1"
	"github.com/SteveCastle/openmob/packages/shrike/src/pkg/models/v1"
)

// Create new Follower
func (s *shrikeServiceServer) CreateFollower(ctx context.Context, req *v1.CreateFollowerRequest) (*v1.CreateFollowerResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a Follower Manager
	m := models.NewFollowerManager(s.db)

	// Get a list of followers given filters, ordering, and limit rules.
	id, err := m.Create(ctx, req.Item)
	if err != nil {
		return nil, err
	}
	return &v1.CreateFollowerResponse{
		Api: apiVersion,
		ID:  *id,
	}, nil
}

// Get follower by id.
func (s *shrikeServiceServer) GetFollower(ctx context.Context, req *v1.GetFollowerRequest) (*v1.GetFollowerResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a Follower Manager
	m := models.NewFollowerManager(s.db)

	// Get a list of followers given filters, ordering, and limit rules.
	follower, err := m.Get(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	return &v1.GetFollowerResponse{
		Api:  apiVersion,
		Item: m.GetProto(follower),
	}, nil

}

// Read all Follower
func (s *shrikeServiceServer) ListFollower(ctx context.Context, req *v1.ListFollowerRequest) (*v1.ListFollowerResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}

	// Create a Follower Manager
	m := models.NewFollowerManager(s.db)

	// Get a list of followers given filters, ordering, and limit rules.
	list, err := m.List(ctx, req.Filters, req.Ordering, req.Limit)
	if err != nil {
		return nil, err
	}

	return &v1.ListFollowerResponse{
		Api:   apiVersion,
		Items: m.GetProtoList(list),
	}, nil
}

// Update Follower
func (s *shrikeServiceServer) UpdateFollower(ctx context.Context, req *v1.UpdateFollowerRequest) (*v1.UpdateFollowerResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a Follower Manager
	m := models.NewFollowerManager(s.db)

	// Get a list of followers given filters, ordering, and limit rules.
	rows, err := m.Update(ctx, req.Item)
	if err != nil {
		return nil, err
	}

	return &v1.UpdateFollowerResponse{
		Api:     apiVersion,
		Updated: *rows,
	}, nil
}

// Delete follower
func (s *shrikeServiceServer) DeleteFollower(ctx context.Context, req *v1.DeleteFollowerRequest) (*v1.DeleteFollowerResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a Follower Manager
	m := models.NewFollowerManager(s.db)

	// Get a list of followers given filters, ordering, and limit rules.
	rows, err := m.Delete(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	return &v1.DeleteFollowerResponse{
		Api:     apiVersion,
		Deleted: *rows,
	}, nil
}
