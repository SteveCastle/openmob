package v1

import (
	"context"
	"fmt"
	"time"

	v1 "github.com/SteveCastle/openmob/packages/shrike/src/pkg/api/v1"
	"github.com/SteveCastle/openmob/packages/shrike/src/pkg/queries"
	"github.com/golang/protobuf/ptypes"
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
	var id string
	// insert Account entity data
	err = c.QueryRowContext(ctx, "INSERT INTO account (username) VALUES($1)  RETURNING id;",
		req.Item.Username).Scan(&id)
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
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, username FROM account WHERE id=$1",
		req.ID)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from Account-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from Account-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Account with ID='%s' is not found",
			req.ID))
	}

	// scan Account data into protobuf model
	var account v1.Account
	var createdAt time.Time
	var updatedAt time.Time

	if err := rows.Scan(&account.ID, &createdAt, &updatedAt, &account.Username); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from Account row-> "+err.Error())
	}

	// Convert time.Time from database into proto timestamp.
	account.CreatedAt, err = ptypes.TimestampProto(createdAt)
	if err != nil {
		return nil, status.Error(codes.Unknown, "createdAt field has invalid format-> "+err.Error())
	}
	account.UpdatedAt, err = ptypes.TimestampProto(updatedAt)
	if err != nil {
		return nil, status.Error(codes.Unknown, "updatedAt field has invalid format-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple Account rows with ID='%s'",
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
	queries.BuildAccountFilters(req.Filters, req.Ordering, req.Limit)
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, username FROM account")
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from Account-> "+err.Error())
	}
	defer rows.Close()

	// Variables to store results returned by database.
	list := []*v1.Account{}
	var createdAt time.Time
	var updatedAt time.Time

	for rows.Next() {
		account := new(v1.Account)
		if err := rows.Scan(&account.ID, &createdAt, &updatedAt, &account.Username); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from Account row-> "+err.Error())
		}
		// Convert time.Time from database into proto timestamp.
		account.CreatedAt, err = ptypes.TimestampProto(createdAt)
		if err != nil {
			return nil, status.Error(codes.Unknown, "createdAt field has invalid format-> "+err.Error())
		}
		account.UpdatedAt, err = ptypes.TimestampProto(updatedAt)
		if err != nil {
			return nil, status.Error(codes.Unknown, "updatedAt field has invalid format-> "+err.Error())
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
	res, err := c.ExecContext(ctx, "UPDATE account SET username=$2 WHERE id=$1",
		req.Item.ID, req.Item.Username)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update Account-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Account with ID='%s' is not found",
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
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Account with ID='%s' is not found",
			req.ID))
	}

	return &v1.DeleteAccountResponse{
		Api:     apiVersion,
		Deleted: rows,
	}, nil
}
