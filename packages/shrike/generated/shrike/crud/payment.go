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

// NewShrikeServiceServer creates Payment service
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
func (s *shrikeServiceServer) CreatePayment(ctx context.Context, req *v1.CreatePaymentRequest) (*v1.CreatePaymentResponse, error) {
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
	// insert Payment entity data
	err = c.QueryRowContext(ctx, "INSERT INTO payment (title) VALUES($1)  RETURNING id;",
		req.Item.Title).Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into Payment-> "+err.Error())
	}

	// get ID of creates Payment
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created Payment-> "+err.Error())
	}

	return &v1.CreatePaymentResponse{
		Api: apiVersion,
		Id:  id,
	}, nil
}

// Get payment by id.
func (s *shrikeServiceServer) GetPayment(ctx context.Context, req *v1.GetPaymentRequest) (*v1.GetPaymentResponse, error) {
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

	// query Payment by ID
	rows, err := c.QueryContext(ctx, "SELECT id, title FROM payment WHERE id=$1",
		req.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from Payment-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from Payment-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Payment with ID='%d' is not found",
			req.Id))
	}

	// get Payment data
	var td v1.Payment
	if err := rows.Scan(&td.Id, &td.Title); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from Payment row-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple Payment rows with ID='%d'",
			req.Id))
	}

	return &v1.GetPaymentResponse{
		Api:  apiVersion,
		Item: &td,
	}, nil

}

// Read all todo tasks
func (s *shrikeServiceServer) ListPayment(ctx context.Context, req *v1.ListPaymentRequest) (*v1.ListPaymentResponse, error) {
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

	// get Payment list
	rows, err := c.QueryContext(ctx, "SELECT id,title FROM payment")
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from Payment-> "+err.Error())
	}
	defer rows.Close()

	list := []*v1.Payment{}
	for rows.Next() {
		td := new(v1.Payment)
		if err := rows.Scan(&td.Id, &td.Title); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from Payment row-> "+err.Error())
		}
		list = append(list, td)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from Payment-> "+err.Error())
	}

	return &v1.ListPaymentResponse{
		Api:   apiVersion,
		Items: list,
	}, nil
}

// Update todo task
func (s *shrikeServiceServer) UpdatePayment(ctx context.Context, req *v1.UpdatePaymentRequest) (*v1.UpdatePaymentResponse, error) {
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

	// update payment
	res, err := c.ExecContext(ctx, "UPDATE payment SET title=$1 WHERE id=$2",
		req.Item.Title, req.Item.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update Payment-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Payment with ID='%d' is not found",
			req.Item.Id))
	}

	return &v1.UpdatePaymentResponse{
		Api:     apiVersion,
		Updated: rows,
	}, nil
}

// Delete payment
func (s *shrikeServiceServer) DeletePayment(ctx context.Context, req *v1.DeletePaymentRequest) (*v1.DeletePaymentResponse, error) {
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

	// delete payment
	res, err := c.ExecContext(ctx, "DELETE FROM payment WHERE id=$1", req.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete Payment-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Payment with ID='%d' is not found",
			req.Id))
	}

	return &v1.DeletePaymentResponse{
		Api:     apiVersion,
		Deleted: rows,
	}, nil
}