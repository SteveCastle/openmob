package v1

import (
	"context"

	v1 "github.com/SteveCastle/openmob/packages/shrike/src/pkg/api/v1"
	"github.com/SteveCastle/openmob/packages/shrike/src/pkg/models"
)

// Create new Donor
func (s *shrikeServiceServer) CreateDonor(ctx context.Context, req *v1.CreateDonorRequest) (*v1.CreateDonorResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a Donor Manager
	m := models.NewDonorManager(s.db)

	// Get a list of donors given filters, ordering, and limit rules.
	id, err := m.Create(ctx, req.Item)
	if err != nil {
		return nil, err
	}
	return &v1.CreateDonorResponse{
		Api: apiVersion,
		ID:  *id,
	}, nil
}

// Get donor by id.
func (s *shrikeServiceServer) GetDonor(ctx context.Context, req *v1.GetDonorRequest) (*v1.GetDonorResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a Donor Manager
	m := models.NewDonorManager(s.db)

	// Get a list of donors given filters, ordering, and limit rules.
	donor, err := m.Get(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	return &v1.GetDonorResponse{
		Api:  apiVersion,
		Item: m.GetProto(donor),
	}, nil

}

// Read all Donor
func (s *shrikeServiceServer) ListDonor(ctx context.Context, req *v1.ListDonorRequest) (*v1.ListDonorResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}

	// Create a Donor Manager
	m := models.NewDonorManager(s.db)

	// Get a list of donors given filters, ordering, and limit rules.
	list, err := m.List(ctx, req.Filters, req.Ordering, req.Limit)
	if err != nil {
		return nil, err
	}

	return &v1.ListDonorResponse{
		Api:   apiVersion,
		Items: m.GetProtoList(list),
	}, nil
}

// Update Donor
func (s *shrikeServiceServer) UpdateDonor(ctx context.Context, req *v1.UpdateDonorRequest) (*v1.UpdateDonorResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a Donor Manager
	m := models.NewDonorManager(s.db)

	// Get a list of donors given filters, ordering, and limit rules.
	rows, err := m.Update(ctx, req.Item)
	if err != nil {
		return nil, err
	}

	return &v1.UpdateDonorResponse{
		Api:     apiVersion,
		Updated: *rows,
	}, nil
}

// Delete donor
func (s *shrikeServiceServer) DeleteDonor(ctx context.Context, req *v1.DeleteDonorRequest) (*v1.DeleteDonorResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a Donor Manager
	m := models.NewDonorManager(s.db)

	// Get a list of donors given filters, ordering, and limit rules.
	rows, err := m.Delete(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	return &v1.DeleteDonorResponse{
		Api:     apiVersion,
		Deleted: *rows,
	}, nil
}
