package v1

import (
	"context"

	v1 "github.com/SteveCastle/openmob/packages/shrike/src/pkg/api/v1"
	"github.com/SteveCastle/openmob/packages/shrike/src/pkg/models/v1"
)

// Create new Activity
func (s *shrikeServiceServer) CreateActivity(ctx context.Context, req *v1.CreateActivityRequest) (*v1.CreateActivityResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a Activity Manager
	m := models.NewActivityManager(s.db)

	// Get a list of activitys given filters, ordering, and limit rules.
	id, err := m.CreateActivity(ctx, req.Item)
	if err != nil {
		return nil, err
	}
	return &v1.CreateActivityResponse{
		Api: apiVersion,
		ID:  *id,
	}, nil
}

// Get activity by id.
func (s *shrikeServiceServer) GetActivity(ctx context.Context, req *v1.GetActivityRequest) (*v1.GetActivityResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a Activity Manager
	m := models.NewActivityManager(s.db)

	// Get a list of activitys given filters, ordering, and limit rules.
	activity, err := m.GetActivity(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	return &v1.GetActivityResponse{
		Api:  apiVersion,
		Item: m.GetProto(activity),
	}, nil

}

// Read all Activity
func (s *shrikeServiceServer) ListActivity(ctx context.Context, req *v1.ListActivityRequest) (*v1.ListActivityResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}

	// Create a Activity Manager
	m := models.NewActivityManager(s.db)

	// Get a list of activitys given filters, ordering, and limit rules.
	list, err := m.ListActivity(ctx, req.Filters, req.Ordering, req.Limit)
	if err != nil {
		return nil, err
	}

	return &v1.ListActivityResponse{
		Api:   apiVersion,
		Items: m.GetProtoList(list),
	}, nil
}

// Update Activity
func (s *shrikeServiceServer) UpdateActivity(ctx context.Context, req *v1.UpdateActivityRequest) (*v1.UpdateActivityResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a Activity Manager
	m := models.NewActivityManager(s.db)

	// Get a list of activitys given filters, ordering, and limit rules.
	rows, err := m.UpdateActivity(ctx, req.Item)
	if err != nil {
		return nil, err
	}

	return &v1.UpdateActivityResponse{
		Api:     apiVersion,
		Updated: *rows,
	}, nil
}

// Delete activity
func (s *shrikeServiceServer) DeleteActivity(ctx context.Context, req *v1.DeleteActivityRequest) (*v1.DeleteActivityResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a Activity Manager
	m := models.NewActivityManager(s.db)

	// Get a list of activitys given filters, ordering, and limit rules.
	rows, err := m.DeleteActivity(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	return &v1.DeleteActivityResponse{
		Api:     apiVersion,
		Deleted: *rows,
	}, nil
}
