package v1

import (
	"context"

	v1 "github.com/SteveCastle/openmob/packages/shrike/src/pkg/api/v1"
	"github.com/SteveCastle/openmob/packages/shrike/src/pkg/models"
)

// Create new Photo
func (s *shrikeServiceServer) CreatePhoto(ctx context.Context, req *v1.CreatePhotoRequest) (*v1.CreatePhotoResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a Photo Manager
	m := models.NewPhotoManager(s.db)

	// Get a list of photos given filters, ordering, and limit rules.
	id, err := m.Create(ctx, req.Item)
	if err != nil {
		return nil, err
	}
	return &v1.CreatePhotoResponse{
		Api: apiVersion,
		ID:  *id,
	}, nil
}

// Get photo by id.
func (s *shrikeServiceServer) GetPhoto(ctx context.Context, req *v1.GetPhotoRequest) (*v1.GetPhotoResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a Photo Manager
	m := models.NewPhotoManager(s.db)

	// Get a list of photos given filters, ordering, and limit rules.
	photo, err := m.Get(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	return &v1.GetPhotoResponse{
		Api:  apiVersion,
		Item: m.GetProto(photo),
	}, nil

}

// Read all Photo
func (s *shrikeServiceServer) ListPhoto(ctx context.Context, req *v1.ListPhotoRequest) (*v1.ListPhotoResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}

	// Create a Photo Manager
	m := models.NewPhotoManager(s.db)

	// Get a list of photos given filters, ordering, and limit rules.
	list, err := m.List(ctx, req.Filters, req.Ordering, req.Limit)
	if err != nil {
		return nil, err
	}

	return &v1.ListPhotoResponse{
		Api:   apiVersion,
		Items: m.GetProtoList(list),
	}, nil
}

// Update Photo
func (s *shrikeServiceServer) UpdatePhoto(ctx context.Context, req *v1.UpdatePhotoRequest) (*v1.UpdatePhotoResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a Photo Manager
	m := models.NewPhotoManager(s.db)

	// Get a list of photos given filters, ordering, and limit rules.
	rows, err := m.Update(ctx, req.Item)
	if err != nil {
		return nil, err
	}

	return &v1.UpdatePhotoResponse{
		Api:     apiVersion,
		Updated: *rows,
	}, nil
}

// Delete photo
func (s *shrikeServiceServer) DeletePhoto(ctx context.Context, req *v1.DeletePhotoRequest) (*v1.DeletePhotoResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a Photo Manager
	m := models.NewPhotoManager(s.db)

	// Get a list of photos given filters, ordering, and limit rules.
	rows, err := m.Delete(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	return &v1.DeletePhotoResponse{
		Api:     apiVersion,
		Deleted: *rows,
	}, nil
}
