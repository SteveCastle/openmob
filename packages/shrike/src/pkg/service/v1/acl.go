package v1

import (
	"context"

	v1 "github.com/SteveCastle/openmob/packages/shrike/src/pkg/api/v1"
	"github.com/SteveCastle/openmob/packages/shrike/src/pkg/models"
)

// Create new ACL
func (s *shrikeServiceServer) CreateACL(ctx context.Context, req *v1.CreateACLRequest) (*v1.CreateACLResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a ACL Manager
	m := models.NewACLManager(s.db)

	// Get a list of acls given filters, ordering, and limit rules.
	id, err := m.Create(ctx, req.Item)
	if err != nil {
		return nil, err
	}
	return &v1.CreateACLResponse{
		Api: apiVersion,
		ID:  *id,
	}, nil
}

// Get acl by id.
func (s *shrikeServiceServer) GetACL(ctx context.Context, req *v1.GetACLRequest) (*v1.GetACLResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a ACL Manager
	m := models.NewACLManager(s.db)

	// Get a list of acls given filters, ordering, and limit rules.
	acl, err := m.Get(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	return &v1.GetACLResponse{
		Api:  apiVersion,
		Item: m.GetProto(acl),
	}, nil

}

// Read all ACL
func (s *shrikeServiceServer) ListACL(ctx context.Context, req *v1.ListACLRequest) (*v1.ListACLResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}

	// Create a ACL Manager
	m := models.NewACLManager(s.db)

	// Get a list of acls given filters, ordering, and limit rules.
	list, err := m.List(ctx, req.Filters, req.Ordering, req.Limit)
	if err != nil {
		return nil, err
	}

	return &v1.ListACLResponse{
		Api:   apiVersion,
		Items: m.GetProtoList(list),
	}, nil
}

// Update ACL
func (s *shrikeServiceServer) UpdateACL(ctx context.Context, req *v1.UpdateACLRequest) (*v1.UpdateACLResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a ACL Manager
	m := models.NewACLManager(s.db)

	// Get a list of acls given filters, ordering, and limit rules.
	rows, err := m.Update(ctx, req.Item)
	if err != nil {
		return nil, err
	}

	return &v1.UpdateACLResponse{
		Api:     apiVersion,
		Updated: *rows,
	}, nil
}

// Delete acl
func (s *shrikeServiceServer) DeleteACL(ctx context.Context, req *v1.DeleteACLRequest) (*v1.DeleteACLResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a ACL Manager
	m := models.NewACLManager(s.db)

	// Get a list of acls given filters, ordering, and limit rules.
	rows, err := m.Delete(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	return &v1.DeleteACLResponse{
		Api:     apiVersion,
		Deleted: *rows,
	}, nil
}
