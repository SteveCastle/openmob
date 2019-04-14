package v1

import (
	"context"

	v1 "github.com/SteveCastle/openmob/packages/shrike/src/pkg/api/v1"
	"github.com/SteveCastle/openmob/packages/shrike/src/pkg/models/v1"
)

// Create new Company
func (s *shrikeServiceServer) CreateCompany(ctx context.Context, req *v1.CreateCompanyRequest) (*v1.CreateCompanyResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a Company Manager
	m := models.NewCompanyManager(s.db)

	// Get a list of companys given filters, ordering, and limit rules.
	id, err := m.CreateCompany(ctx, req.Item)
	if err != nil {
		return nil, err
	}
	return &v1.CreateCompanyResponse{
		Api: apiVersion,
		ID:  *id,
	}, nil
}

// Get company by id.
func (s *shrikeServiceServer) GetCompany(ctx context.Context, req *v1.GetCompanyRequest) (*v1.GetCompanyResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a Company Manager
	m := models.NewCompanyManager(s.db)

	// Get a list of companys given filters, ordering, and limit rules.
	company, err := m.GetCompany(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	return &v1.GetCompanyResponse{
		Api:  apiVersion,
		Item: m.GetProto(company),
	}, nil

}

// Read all Company
func (s *shrikeServiceServer) ListCompany(ctx context.Context, req *v1.ListCompanyRequest) (*v1.ListCompanyResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}

	// Create a Company Manager
	m := models.NewCompanyManager(s.db)

	// Get a list of companys given filters, ordering, and limit rules.
	list, err := m.ListCompany(ctx, req.Filters, req.Ordering, req.Limit)
	if err != nil {
		return nil, err
	}

	return &v1.ListCompanyResponse{
		Api:   apiVersion,
		Items: m.GetProtoList(list),
	}, nil
}

// Update Company
func (s *shrikeServiceServer) UpdateCompany(ctx context.Context, req *v1.UpdateCompanyRequest) (*v1.UpdateCompanyResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a Company Manager
	m := models.NewCompanyManager(s.db)

	// Get a list of companys given filters, ordering, and limit rules.
	rows, err := m.UpdateCompany(ctx, req.Item)
	if err != nil {
		return nil, err
	}

	return &v1.UpdateCompanyResponse{
		Api:     apiVersion,
		Updated: *rows,
	}, nil
}

// Delete company
func (s *shrikeServiceServer) DeleteCompany(ctx context.Context, req *v1.DeleteCompanyRequest) (*v1.DeleteCompanyResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a Company Manager
	m := models.NewCompanyManager(s.db)

	// Get a list of companys given filters, ordering, and limit rules.
	rows, err := m.DeleteCompany(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	return &v1.DeleteCompanyResponse{
		Api:     apiVersion,
		Deleted: *rows,
	}, nil
}
