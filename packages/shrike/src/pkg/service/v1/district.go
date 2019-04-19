package v1

import (
	"context"

	v1 "github.com/SteveCastle/openmob/packages/shrike/src/pkg/api/v1"
	"github.com/SteveCastle/openmob/packages/shrike/src/pkg/models/v1"
)

// Create new District
func (s *shrikeServiceServer) CreateDistrict(ctx context.Context, req *v1.CreateDistrictRequest) (*v1.CreateDistrictResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a District Manager
	m := models.NewDistrictManager(s.db)

	// Get a list of districts given filters, ordering, and limit rules.
	id, err := m.Create(ctx, req.Item)
	if err != nil {
		return nil, err
	}
	return &v1.CreateDistrictResponse{
		Api: apiVersion,
		ID:  *id,
	}, nil
}

// Get district by id.
func (s *shrikeServiceServer) GetDistrict(ctx context.Context, req *v1.GetDistrictRequest) (*v1.GetDistrictResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a District Manager
	m := models.NewDistrictManager(s.db)

	// Get a list of districts given filters, ordering, and limit rules.
	district, err := m.Get(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	return &v1.GetDistrictResponse{
		Api:  apiVersion,
		Item: m.GetProto(district),
	}, nil

}

// Read all District
func (s *shrikeServiceServer) ListDistrict(ctx context.Context, req *v1.ListDistrictRequest) (*v1.ListDistrictResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}

	// Create a District Manager
	m := models.NewDistrictManager(s.db)

	// Get a list of districts given filters, ordering, and limit rules.
	list, err := m.List(ctx, req.Filters, req.Ordering, req.Limit)
	if err != nil {
		return nil, err
	}

	return &v1.ListDistrictResponse{
		Api:   apiVersion,
		Items: m.GetProtoList(list),
	}, nil
}

// Update District
func (s *shrikeServiceServer) UpdateDistrict(ctx context.Context, req *v1.UpdateDistrictRequest) (*v1.UpdateDistrictResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a District Manager
	m := models.NewDistrictManager(s.db)

	// Get a list of districts given filters, ordering, and limit rules.
	rows, err := m.Update(ctx, req.Item)
	if err != nil {
		return nil, err
	}

	return &v1.UpdateDistrictResponse{
		Api:     apiVersion,
		Updated: *rows,
	}, nil
}

// Delete district
func (s *shrikeServiceServer) DeleteDistrict(ctx context.Context, req *v1.DeleteDistrictRequest) (*v1.DeleteDistrictResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a District Manager
	m := models.NewDistrictManager(s.db)

	// Get a list of districts given filters, ordering, and limit rules.
	rows, err := m.Delete(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	return &v1.DeleteDistrictResponse{
		Api:     apiVersion,
		Deleted: *rows,
	}, nil
}
