package v1

import (
	"context"

	v1 "github.com/SteveCastle/openmob/packages/shrike/src/pkg/api/v1"
	"github.com/SteveCastle/openmob/packages/shrike/src/pkg/models"
)

// Create new ActivityType
func (s *shrikeServiceServer) CreateActivityType(ctx context.Context, req *v1.CreateActivityTypeRequest) (*v1.CreateActivityTypeResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a ActivityType Manager
	m := models.NewActivityTypeManager(s.db)

	// Get a list of activityTypes given filters, ordering, and limit rules.
	id, err := m.Create(ctx, req.Item)
	if err != nil {
		return nil, err
	}
	return &v1.CreateActivityTypeResponse{
		Api: apiVersion,
		ID:  *id,
	}, nil
}

// Get activityType by id.
func (s *shrikeServiceServer) GetActivityType(ctx context.Context, req *v1.GetActivityTypeRequest) (*v1.GetActivityTypeResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a ActivityType Manager
	m := models.NewActivityTypeManager(s.db)

	// Get a list of activityTypes given filters, ordering, and limit rules.
	activityType, err := m.Get(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	return &v1.GetActivityTypeResponse{
		Api:  apiVersion,
		Item: m.GetProto(activityType),
	}, nil

}

// Read all ActivityType
func (s *shrikeServiceServer) ListActivityType(ctx context.Context, req *v1.ListActivityTypeRequest) (*v1.ListActivityTypeResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}

	// Create a ActivityType Manager
	m := models.NewActivityTypeManager(s.db)

	// Get a list of activityTypes given filters, ordering, and limit rules.
	list, err := m.List(ctx, req.Filters, req.Ordering, req.Limit)
	if err != nil {
		return nil, err
	}

	return &v1.ListActivityTypeResponse{
		Api:   apiVersion,
		Items: m.GetProtoList(list),
	}, nil
}

// Update ActivityType
func (s *shrikeServiceServer) UpdateActivityType(ctx context.Context, req *v1.UpdateActivityTypeRequest) (*v1.UpdateActivityTypeResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a ActivityType Manager
	m := models.NewActivityTypeManager(s.db)

	// Get a list of activityTypes given filters, ordering, and limit rules.
	rows, err := m.Update(ctx, req.Item)
	if err != nil {
		return nil, err
	}

	return &v1.UpdateActivityTypeResponse{
		Api:     apiVersion,
		Updated: *rows,
	}, nil
}

// Delete activityType
func (s *shrikeServiceServer) DeleteActivityType(ctx context.Context, req *v1.DeleteActivityTypeRequest) (*v1.DeleteActivityTypeResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a ActivityType Manager
	m := models.NewActivityTypeManager(s.db)

	// Get a list of activityTypes given filters, ordering, and limit rules.
	rows, err := m.Delete(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	return &v1.DeleteActivityTypeResponse{
		Api:     apiVersion,
		Deleted: *rows,
	}, nil
}
