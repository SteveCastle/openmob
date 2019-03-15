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

// Create new CustomerCart
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
	var id string
	// insert CustomerCart entity data
	err = c.QueryRowContext(ctx, "INSERT INTO customer_cart () VALUES()  RETURNING id;").Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into CustomerCart-> "+err.Error())
	}

	// get ID of creates CustomerCart
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created CustomerCart-> "+err.Error())
	}

	return &v1.CreateCustomerCartResponse{
		Api: apiVersion,
		ID:  id,
	}, nil
}

// Get customer_cart by id.
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
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at FROM customer_cart WHERE id=$1",
		req.ID)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from CustomerCart-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from CustomerCart-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("CustomerCart with ID='%s' is not found",
			req.ID))
	}

	// scan CustomerCart data into protobuf model
	var customercart v1.CustomerCart
	var createdAt time.Time
	var updatedAt time.Time

	if err := rows.Scan(&customercart.ID, &createdAt, &updatedAt); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from CustomerCart row-> "+err.Error())
	}

	// Convert time.Time from database into proto timestamp.
	customercart.CreatedAt, err = ptypes.TimestampProto(createdAt)
	if err != nil {
		return nil, status.Error(codes.Unknown, "createdAt field has invalid format-> "+err.Error())
	}
	customercart.UpdatedAt, err = ptypes.TimestampProto(updatedAt)
	if err != nil {
		return nil, status.Error(codes.Unknown, "updatedAt field has invalid format-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple CustomerCart rows with ID='%s'",
			req.ID))
	}

	return &v1.GetCustomerCartResponse{
		Api:  apiVersion,
		Item: &customercart,
	}, nil

}

// Read all CustomerCart
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

	// Generate SQL to select all columns in CustomerCart Table
	// Then generate filtering and ordering sql and finally run query.

	baseSQL := "SELECT id, created_at, updated_at FROM customer_cart"
	querySQL := queries.BuildCustomerCartFilters(req.Filters, req.Ordering, req.Limit)
	SQL := fmt.Sprintf("%s %s", baseSQL, querySQL)
	rows, err := c.QueryContext(ctx, SQL)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from CustomerCart-> "+err.Error())
	}
	defer rows.Close()

	// Variables to store results returned by database.
	list := []*v1.CustomerCart{}
	var createdAt time.Time
	var updatedAt time.Time

	for rows.Next() {
		customercart := new(v1.CustomerCart)
		if err := rows.Scan(&customercart.ID, &createdAt, &updatedAt); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from CustomerCart row-> "+err.Error())
		}
		// Convert time.Time from database into proto timestamp.
		customercart.CreatedAt, err = ptypes.TimestampProto(createdAt)
		if err != nil {
			return nil, status.Error(codes.Unknown, "createdAt field has invalid format-> "+err.Error())
		}
		customercart.UpdatedAt, err = ptypes.TimestampProto(updatedAt)
		if err != nil {
			return nil, status.Error(codes.Unknown, "updatedAt field has invalid format-> "+err.Error())
		}

		list = append(list, customercart)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from CustomerCart-> "+err.Error())
	}

	return &v1.ListCustomerCartResponse{
		Api:   apiVersion,
		Items: list,
	}, nil
}

// Update CustomerCart
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

	// update customer_cart
	res, err := c.ExecContext(ctx, "UPDATE customer_cart SET  WHERE id=$1",
		req.Item.ID)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update CustomerCart-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("CustomerCart with ID='%s' is not found",
			req.Item.ID))
	}

	return &v1.UpdateCustomerCartResponse{
		Api:     apiVersion,
		Updated: rows,
	}, nil
}

// Delete customer_cart
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

	// delete customer_cart
	res, err := c.ExecContext(ctx, "DELETE FROM customer_cart WHERE id=$1", req.ID)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete CustomerCart-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("CustomerCart with ID='%s' is not found",
			req.ID))
	}

	return &v1.DeleteCustomerCartResponse{
		Api:     apiVersion,
		Deleted: rows,
	}, nil
}
