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

// NewShrikeServiceServer creates CustomerCart service
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
func (s *shrikeServiceServer) CreateCustomerCart(ctx context.Context, req *v1.CreateCustomerCartRequest) (*v1.CreateCustomerCartResponse, error) {
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
	// insert CustomerCart entity data
	err = c.QueryRowContext(ctx, "INSERT INTO customercart (title) VALUES($1)  RETURNING id;",
		req.Item.Title).Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into CustomerCart-> "+err.Error())
	}

	// get ID of creates CustomerCart
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created CustomerCart-> "+err.Error())
	}

	return &v1.CreateCustomerCartResponse{
		Api: apiVersion,
		Id:  id,
	}, nil
}

// Get customercart by id.
func (s *shrikeServiceServer) GetCustomerCart(ctx context.Context, req *v1.GetCustomerCartRequest) (*v1.GetCustomerCartResponse, error) {
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

	// query CustomerCart by ID
	rows, err := c.QueryContext(ctx, "SELECT id, title FROM customercart WHERE id=$1",
		req.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from CustomerCart-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from CustomerCart-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("CustomerCart with ID='%d' is not found",
			req.Id))
	}

	// get CustomerCart data
	var td v1.CustomerCart
	if err := rows.Scan(&td.Id, &td.Title); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from CustomerCart row-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple CustomerCart rows with ID='%d'",
			req.Id))
	}

	return &v1.GetCustomerCartResponse{
		Api:  apiVersion,
		Item: &td,
	}, nil

}

// Read all todo tasks
func (s *shrikeServiceServer) ListCustomerCart(ctx context.Context, req *v1.ListCustomerCartRequest) (*v1.ListCustomerCartResponse, error) {
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

	// get CustomerCart list
	rows, err := c.QueryContext(ctx, "SELECT id,title FROM CustomerCart")
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from CustomerCart-> "+err.Error())
	}
	defer rows.Close()

	list := []*v1.CustomerCart{}
	for rows.Next() {
		td := new(v1.CustomerCart)
		if err := rows.Scan(&td.Id, &td.Title); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from CustomerCart row-> "+err.Error())
		}
		list = append(list, td)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from CustomerCart-> "+err.Error())
	}

	return &v1.ListCustomerCartResponse{
		Api:   apiVersion,
		Items: list,
	}, nil
}

// Update todo task
func (s *shrikeServiceServer) UpdateCustomerCart(ctx context.Context, req *v1.UpdateCustomerCartRequest) (*v1.UpdateCustomerCartResponse, error) {
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

	// update customercart
	res, err := c.ExecContext(ctx, "UPDATE customercart SET title=$1 WHERE id=$2",
		req.Item.Title, req.Item.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update customercart-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("customercart with ID='%d' is not found",
			req.Item.Id))
	}

	return &v1.UpdateCustomerCartResponse{
		Api:     apiVersion,
		Updated: rows,
	}, nil
}

// Delete customercart
func (s *shrikeServiceServer) DeleteCustomerCart(ctx context.Context, req *v1.DeleteCustomerCartRequest) (*v1.DeleteCustomerCartResponse, error) {
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

	// delete customercart
	res, err := c.ExecContext(ctx, "DELETE FROM customercart WHERE id=$1", req.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete customercart-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("customercart with ID='%d' is not found",
			req.Id))
	}

	return &v1.DeleteCustomerCartResponse{
		Api:     apiVersion,
		Deleted: rows,
	}, nil
}
