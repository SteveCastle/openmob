package v1

import (
	"context"

	v1 "github.com/SteveCastle/openmob/packages/shrike/src/pkg/api/v1"
	"github.com/SteveCastle/openmob/packages/shrike/src/pkg/models"
)

// Create new Account
func (s *shrikeServiceServer) CreateAccount(ctx context.Context, req *v1.CreateAccountRequest) (*v1.CreateAccountResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a Account Manager
	m := models.NewAccountManager(s.db)

	// Get a list of accounts given filters, ordering, and limit rules.
	id, err := m.Create(ctx, req.Item)
	if err != nil {
		return nil, err
	}
	return &v1.CreateAccountResponse{
		Api: apiVersion,
		ID:  *id,
	}, nil
}

// Get account by id.
func (s *shrikeServiceServer) GetAccount(ctx context.Context, req *v1.GetAccountRequest) (*v1.GetAccountResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a Account Manager
	m := models.NewAccountManager(s.db)

	// Get a list of accounts given filters, ordering, and limit rules.
	account, err := m.Get(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	return &v1.GetAccountResponse{
		Api:  apiVersion,
		Item: m.GetProto(account),
	}, nil

}

// Read all Account
func (s *shrikeServiceServer) ListAccount(ctx context.Context, req *v1.ListAccountRequest) (*v1.ListAccountResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}

	// Create a Account Manager
	m := models.NewAccountManager(s.db)

	// Get a list of accounts given filters, ordering, and limit rules.
	list, err := m.List(ctx, req.Filters, req.Ordering, req.Limit)
	if err != nil {
		return nil, err
	}

	return &v1.ListAccountResponse{
		Api:   apiVersion,
		Items: m.GetProtoList(list),
	}, nil
}

// Update Account
func (s *shrikeServiceServer) UpdateAccount(ctx context.Context, req *v1.UpdateAccountRequest) (*v1.UpdateAccountResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a Account Manager
	m := models.NewAccountManager(s.db)

	// Get a list of accounts given filters, ordering, and limit rules.
	rows, err := m.Update(ctx, req.Item)
	if err != nil {
		return nil, err
	}

	return &v1.UpdateAccountResponse{
		Api:     apiVersion,
		Updated: *rows,
	}, nil
}

// Delete account
func (s *shrikeServiceServer) DeleteAccount(ctx context.Context, req *v1.DeleteAccountRequest) (*v1.DeleteAccountResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a Account Manager
	m := models.NewAccountManager(s.db)

	// Get a list of accounts given filters, ordering, and limit rules.
	rows, err := m.Delete(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	return &v1.DeleteAccountResponse{
		Api:     apiVersion,
		Deleted: *rows,
	}, nil
}
