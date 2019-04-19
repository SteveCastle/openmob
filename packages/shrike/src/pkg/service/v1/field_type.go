package v1

import (
	"context"

	v1 "github.com/SteveCastle/openmob/packages/shrike/src/pkg/api/v1"
	"github.com/SteveCastle/openmob/packages/shrike/src/pkg/models/v1"
)

// Create new FieldType
func (s *shrikeServiceServer) CreateFieldType(ctx context.Context, req *v1.CreateFieldTypeRequest) (*v1.CreateFieldTypeResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a FieldType Manager
	m := models.NewFieldTypeManager(s.db)

	// Get a list of fieldTypes given filters, ordering, and limit rules.
	id, err := m.Create(ctx, req.Item)
	if err != nil {
		return nil, err
	}
	return &v1.CreateFieldTypeResponse{
		Api: apiVersion,
		ID:  *id,
	}, nil
}

// Get fieldType by id.
func (s *shrikeServiceServer) GetFieldType(ctx context.Context, req *v1.GetFieldTypeRequest) (*v1.GetFieldTypeResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a FieldType Manager
	m := models.NewFieldTypeManager(s.db)

	// Get a list of fieldTypes given filters, ordering, and limit rules.
	fieldType, err := m.Get(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	return &v1.GetFieldTypeResponse{
		Api:  apiVersion,
		Item: m.GetProto(fieldType),
	}, nil

}

// Read all FieldType
func (s *shrikeServiceServer) ListFieldType(ctx context.Context, req *v1.ListFieldTypeRequest) (*v1.ListFieldTypeResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}

	// Create a FieldType Manager
	m := models.NewFieldTypeManager(s.db)

	// Get a list of fieldTypes given filters, ordering, and limit rules.
	list, err := m.List(ctx, req.Filters, req.Ordering, req.Limit)
	if err != nil {
		return nil, err
	}

	return &v1.ListFieldTypeResponse{
		Api:   apiVersion,
		Items: m.GetProtoList(list),
	}, nil
}

// Update FieldType
func (s *shrikeServiceServer) UpdateFieldType(ctx context.Context, req *v1.UpdateFieldTypeRequest) (*v1.UpdateFieldTypeResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a FieldType Manager
	m := models.NewFieldTypeManager(s.db)

	// Get a list of fieldTypes given filters, ordering, and limit rules.
	rows, err := m.Update(ctx, req.Item)
	if err != nil {
		return nil, err
	}

	return &v1.UpdateFieldTypeResponse{
		Api:     apiVersion,
		Updated: *rows,
	}, nil
}

// Delete fieldType
func (s *shrikeServiceServer) DeleteFieldType(ctx context.Context, req *v1.DeleteFieldTypeRequest) (*v1.DeleteFieldTypeResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a FieldType Manager
	m := models.NewFieldTypeManager(s.db)

	// Get a list of fieldTypes given filters, ordering, and limit rules.
	rows, err := m.Delete(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	return &v1.DeleteFieldTypeResponse{
		Api:     apiVersion,
		Deleted: *rows,
	}, nil
}
