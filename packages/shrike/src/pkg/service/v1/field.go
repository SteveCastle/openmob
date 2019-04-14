package v1

import (
	"context"

	v1 "github.com/SteveCastle/openmob/packages/shrike/src/pkg/api/v1"
	"github.com/SteveCastle/openmob/packages/shrike/src/pkg/models/v1"
)

// Create new Field
func (s *shrikeServiceServer) CreateField(ctx context.Context, req *v1.CreateFieldRequest) (*v1.CreateFieldResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a Field Manager
	m := models.NewFieldManager(s.db)

	// Get a list of fields given filters, ordering, and limit rules.
	id, err := m.CreateField(ctx, req.Item)
	if err != nil {
		return nil, err
	}
	return &v1.CreateFieldResponse{
		Api: apiVersion,
		ID:  *id,
	}, nil
}

// Get field by id.
func (s *shrikeServiceServer) GetField(ctx context.Context, req *v1.GetFieldRequest) (*v1.GetFieldResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a Field Manager
	m := models.NewFieldManager(s.db)

	// Get a list of fields given filters, ordering, and limit rules.
	field, err := m.GetField(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	return &v1.GetFieldResponse{
		Api:  apiVersion,
		Item: m.GetProto(field),
	}, nil

}

// Read all Field
func (s *shrikeServiceServer) ListField(ctx context.Context, req *v1.ListFieldRequest) (*v1.ListFieldResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}

	// Create a Field Manager
	m := models.NewFieldManager(s.db)

	// Get a list of fields given filters, ordering, and limit rules.
	list, err := m.ListField(ctx, req.Filters, req.Ordering, req.Limit)
	if err != nil {
		return nil, err
	}

	return &v1.ListFieldResponse{
		Api:   apiVersion,
		Items: m.GetProtoList(list),
	}, nil
}

// Update Field
func (s *shrikeServiceServer) UpdateField(ctx context.Context, req *v1.UpdateFieldRequest) (*v1.UpdateFieldResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a Field Manager
	m := models.NewFieldManager(s.db)

	// Get a list of fields given filters, ordering, and limit rules.
	rows, err := m.UpdateField(ctx, req.Item)
	if err != nil {
		return nil, err
	}

	return &v1.UpdateFieldResponse{
		Api:     apiVersion,
		Updated: *rows,
	}, nil
}

// Delete field
func (s *shrikeServiceServer) DeleteField(ctx context.Context, req *v1.DeleteFieldRequest) (*v1.DeleteFieldResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a Field Manager
	m := models.NewFieldManager(s.db)

	// Get a list of fields given filters, ordering, and limit rules.
	rows, err := m.DeleteField(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	return &v1.DeleteFieldResponse{
		Api:     apiVersion,
		Deleted: *rows,
	}, nil
}
