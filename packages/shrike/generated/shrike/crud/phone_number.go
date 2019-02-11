package v1

import (
	"context"
	"database/sql"
	"fmt"

	v1 "github.com/SteveCastle/openmob/packages/shrike/src/pkg/api/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	// apiVersion is version of API is provided by server
	apiVersion = "v1"
)

// shrikeServiceServer is implementation of v1.ShrikeServiceServer proto interface
type shrikeServiceServer struct {
	db *sql.DB
}

// NewShrikeServiceServer creates PhoneNumber service
func NewShrikeServiceServer(db *sql.DB) v1.ShrikeServiceServer {
	return &shrikeServiceServer{db: db}
}

// checkAPI checks if the API version requested by client is supported by server
func (s *shrikeServiceServer) checkAPI(api string) error {
	// API version is "" means use current version of the service
	if len(api) > 0 {
		if apiVersion != api {
			return status.Errorf(codes.Unimplemented,
				"unsupported API version: service implements API version '%s', but asked for '%s'", apiVersion, api)
		}
	}
	return nil
}

// connect returns SQL database connection from the pool
func (s *shrikeServiceServer) connect(ctx context.Context) (*sql.Conn, error) {
	c, err := s.db.Conn(ctx)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to connect to database-> "+err.Error())
	}
	return c, nil
}

// Create new todo task
func (s *shrikeServiceServer) CreatePhoneNumber(ctx context.Context, req *v1.CreatePhoneNumberRequest) (*v1.CreatePhoneNumberResponse, error) {
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
	// insert PhoneNumber entity data
	err = c.QueryRowContext(ctx, "INSERT INTO phone_number (title) VALUES($1)  RETURNING id;",
		req.Item.Title).Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into PhoneNumber-> "+err.Error())
	}

	// get ID of creates PhoneNumber
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created PhoneNumber-> "+err.Error())
	}

	return &v1.CreatePhoneNumberResponse{
		Api: apiVersion,
		Id:  id,
	}, nil
}

// Get phone_number by id.
func (s *shrikeServiceServer) GetPhoneNumber(ctx context.Context, req *v1.GetPhoneNumberRequest) (*v1.GetPhoneNumberResponse, error) {
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

	// query PhoneNumber by ID
	rows, err := c.QueryContext(ctx, "SELECT id, title FROM phone_number WHERE id=$1",
		req.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from PhoneNumber-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from PhoneNumber-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("PhoneNumber with ID='%d' is not found",
			req.Id))
	}

	// get PhoneNumber data
	var td v1.PhoneNumber
	if err := rows.Scan(&td.Id, &td.Title); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from PhoneNumber row-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple PhoneNumber rows with ID='%d'",
			req.Id))
	}

	return &v1.GetPhoneNumberResponse{
		Api:  apiVersion,
		Item: &td,
	}, nil

}

// Read all todo tasks
func (s *shrikeServiceServer) ListPhoneNumber(ctx context.Context, req *v1.ListPhoneNumberRequest) (*v1.ListPhoneNumberResponse, error) {
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

	// get PhoneNumber list
	rows, err := c.QueryContext(ctx, "SELECT id,title FROM phone_number")
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from PhoneNumber-> "+err.Error())
	}
	defer rows.Close()

	list := []*v1.PhoneNumber{}
	for rows.Next() {
		td := new(v1.PhoneNumber)
		if err := rows.Scan(&td.Id, &td.Title); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from PhoneNumber row-> "+err.Error())
		}
		list = append(list, td)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from PhoneNumber-> "+err.Error())
	}

	return &v1.ListPhoneNumberResponse{
		Api:   apiVersion,
		Items: list,
	}, nil
}

// Update todo task
func (s *shrikeServiceServer) UpdatePhoneNumber(ctx context.Context, req *v1.UpdatePhoneNumberRequest) (*v1.UpdatePhoneNumberResponse, error) {
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

	// update phone_number
	res, err := c.ExecContext(ctx, "UPDATE phone_number SET title=$1 WHERE id=$2",
		req.Item.Title, req.Item.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update PhoneNumber-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("PhoneNumber with ID='%d' is not found",
			req.Item.Id))
	}

	return &v1.UpdatePhoneNumberResponse{
		Api:     apiVersion,
		Updated: rows,
	}, nil
}

// Delete phone_number
func (s *shrikeServiceServer) DeletePhoneNumber(ctx context.Context, req *v1.DeletePhoneNumberRequest) (*v1.DeletePhoneNumberResponse, error) {
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

	// delete phone_number
	res, err := c.ExecContext(ctx, "DELETE FROM phone_number WHERE id=$1", req.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete PhoneNumber-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("PhoneNumber with ID='%d' is not found",
			req.Id))
	}

	return &v1.DeletePhoneNumberResponse{
		Api:     apiVersion,
		Deleted: rows,
	}, nil
}
