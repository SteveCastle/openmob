package v1

import (
	"context"
	"fmt"

	v1 "github.com/SteveCastle/openmob/packages/shrike/src/pkg/api/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Create new Account
func (s *shrikeServiceServer) CreateAccount(ctx context.Context, req *v1.CreateAccountRequest) (*v1.CreateAccountResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// get SQL connection from pool
	c, err := s.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()
	var id int64
	// insert Account entity data
	err = c.QueryRowContext(ctx, "INSERT INTO account (id, created_at, updated_at, username, ) VALUES($1, $2, $3, $4, )  RETURNING id;",
		req.Item.ID, req.Item.CreatedAt, req.Item.UpdatedAt, req.Item.Username).Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into Account-> "+err.Error())
	}

	// get ID of creates Account
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created Account-> "+err.Error())
	}

	return &v1.CreateAccountResponse{
		Api: apiVersion,
		ID:  id,
	}, nil
}

// Get account by id.
func (s *shrikeServiceServer) GetAccount(ctx context.Context, req *v1.GetAccountRequest) (*v1.GetAccountResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// get SQL connection from pool
	c, err := s.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	// query Account by ID
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, username,  FROM account WHERE id=$1",
		req.ID)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from Account-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from Account-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Account with ID='%d' is not found",
			req.ID))
	}

	// get Account data
	var account v1.Account
	if err := rows.Scan(&account.ID, &account.CreatedAt, &account.UpdatedAt, &account.Username); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from Account row-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple Account rows with ID='%d'",
			req.ID))
	}

	return &v1.GetAccountResponse{
		Api:  apiVersion,
		Item: &account,
	}, nil

}

// Read all Account
func (s *shrikeServiceServer) ListAccount(ctx context.Context, req *v1.ListAccountRequest) (*v1.ListAccountResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}

	// get SQL connection from pool
	c, err := s.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	// get Account list
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, username,  FROM account")
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from Account-> "+err.Error())
	}
	defer rows.Close()

	list := []*v1.Account{}
	for rows.Next() {
		account := new(v1.Account)
		if err := rows.Scan(&account.ID, &account.CreatedAt, &account.UpdatedAt, &account.Username); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from Account row-> "+err.Error())
		}
		list = append(list, account)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from Account-> "+err.Error())
	}

	return &v1.ListAccountResponse{
		Api:   apiVersion,
		Items: list,
	}, nil
}

// Update Account
func (s *shrikeServiceServer) UpdateAccount(ctx context.Context, req *v1.UpdateAccountRequest) (*v1.UpdateAccountResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}

	// get SQL connection from pool
	c, err := s.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	// update account
	res, err := c.ExecContext(ctx, "UPDATE account SET $1, $2, $3, $4,  WHERE id=$1",
		req.Item.ID, req.Item.CreatedAt, req.Item.UpdatedAt, req.Item.Username)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update Account-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Account with ID='%d' is not found",
			req.Item.ID))
	}

	return &v1.UpdateAccountResponse{
		Api:     apiVersion,
		Updated: rows,
	}, nil
}

// Delete account
func (s *shrikeServiceServer) DeleteAccount(ctx context.Context, req *v1.DeleteAccountRequest) (*v1.DeleteAccountResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}

	// get SQL connection from pool
	c, err := s.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	// delete account
	res, err := c.ExecContext(ctx, "DELETE FROM account WHERE id=$1", req.ID)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete Account-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Account with ID='%d' is not found",
			req.ID))
	}

	return &v1.DeleteAccountResponse{
		Api:     apiVersion,
		Deleted: rows,
	}, nil
}
