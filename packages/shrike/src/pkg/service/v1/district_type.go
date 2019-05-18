package v1

import (
	"context"

	v1 "github.com/SteveCastle/openmob/packages/shrike/src/pkg/api/v1"
	"github.com/SteveCastle/openmob/packages/shrike/src/pkg/models"
)

// Create new DistrictType
func (s *shrikeServiceServer) CreateDistrictType(ctx context.Context, req *v1.CreateDistrictTypeRequest) (*v1.CreateDistrictTypeResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a DistrictType Manager
	m := models.NewDistrictTypeManager(s.db)

	// Get a list of districtTypes given filters, ordering, and limit rules.
	id, err := m.Create(ctx, req.Item)
	if err != nil {
		return nil, err
	}
	return &v1.CreateDistrictTypeResponse{
		Api: apiVersion,
		ID:  *id,
	}, nil
}

// Get districtType by id.
func (s *shrikeServiceServer) GetDistrictType(ctx context.Context, req *v1.GetDistrictTypeRequest) (*v1.GetDistrictTypeResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a DistrictType Manager
	m := models.NewDistrictTypeManager(s.db)

	// Get a list of districtTypes given filters, ordering, and limit rules.
	districtType, err := m.Get(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	return &v1.GetDistrictTypeResponse{
		Api:  apiVersion,
		Item: m.GetProto(districtType),
	}, nil

}

// Read all DistrictType
func (s *shrikeServiceServer) ListDistrictType(ctx context.Context, req *v1.ListDistrictTypeRequest) (*v1.ListDistrictTypeResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}

	// Create a DistrictType Manager
	m := models.NewDistrictTypeManager(s.db)

	// Get a list of districtTypes given filters, ordering, and limit rules.
	list, err := m.List(ctx, req.Filters, req.Ordering, req.Limit)
	if err != nil {
		return nil, err
	}

	return &v1.ListDistrictTypeResponse{
		Api:   apiVersion,
		Items: m.GetProtoList(list),
	}, nil
}

// Update DistrictType
func (s *shrikeServiceServer) UpdateDistrictType(ctx context.Context, req *v1.UpdateDistrictTypeRequest) (*v1.UpdateDistrictTypeResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a DistrictType Manager
	m := models.NewDistrictTypeManager(s.db)

	// Get a list of districtTypes given filters, ordering, and limit rules.
	rows, err := m.Update(ctx, req.Item)
	if err != nil {
		return nil, err
	}

	return &v1.UpdateDistrictTypeResponse{
		Api:     apiVersion,
		Updated: *rows,
	}, nil
}

// Delete districtType
func (s *shrikeServiceServer) DeleteDistrictType(ctx context.Context, req *v1.DeleteDistrictTypeRequest) (*v1.DeleteDistrictTypeResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a DistrictType Manager
	m := models.NewDistrictTypeManager(s.db)

	// Get a list of districtTypes given filters, ordering, and limit rules.
	rows, err := m.Delete(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	return &v1.DeleteDistrictTypeResponse{
		Api:     apiVersion,
		Deleted: *rows,
	}, nil
}
