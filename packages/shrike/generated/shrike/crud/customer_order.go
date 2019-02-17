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

// NewShrikeServiceServer creates CustomerOrder service
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

// Create new CustomerOrder
func (s *shrikeServiceServer) CreateCustomerOrder(ctx context.Context, req *v1.CreateCustomerOrderRequest) (*v1.CreateCustomerOrderResponse, error) {
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
	// insert CustomerOrder entity data
	err = c.QueryRowContext(ctx, "INSERT INTO customer_order (id, created_at, updated_at, customer_cart, ) VALUES($1, $2, $3, $4, )  RETURNING id;",
		 req.ItemID  req.ItemCreatedAt  req.ItemUpdatedAt  req.ItemCustomerCart ).Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into CustomerOrder-> "+err.Error())
	}

	// get ID of creates CustomerOrder
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created CustomerOrder-> "+err.Error())
	}

	return &v1.CreateCustomerOrderResponse{
		Api: apiVersion,
		Id:  id,
	}, nil
}

// Get customer_order by id.
func (s *shrikeServiceServer) GetCustomerOrder(ctx context.Context, req *v1.GetCustomerOrderRequest) (*v1.GetCustomerOrderResponse, error) {
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

	// query CustomerOrder by ID
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, customer_cart,  FROM customer_order WHERE id=$1",
		req.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from CustomerOrder-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from CustomerOrder-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("CustomerOrder with ID='%d' is not found",
			req.Id))
	}

	// get CustomerOrder data
	var customerorder v1.CustomerOrder
	if err := rows.Scan( &customerorder.ID,  &customerorder.CreatedAt,  &customerorder.UpdatedAt,  &customerorder.CustomerCart, ); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from CustomerOrder row-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple CustomerOrder rows with ID='%d'",
			req.Id))
	}

	return &v1.GetCustomerOrderResponse{
		Api:  apiVersion,
		Item: &customerorder,
	}, nil

}

// Read all CustomerOrder
func (s *shrikeServiceServer) ListCustomerOrder(ctx context.Context, req *v1.ListCustomerOrderRequest) (*v1.ListCustomerOrderResponse, error) {
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

	// get CustomerOrder list
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, customer_cart,  FROM customer_order")
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from CustomerOrder-> "+err.Error())
	}
	defer rows.Close()

	list := []*v1.CustomerOrder{}
	for rows.Next() {
		customerorder := new(v1.CustomerOrder)
		if err := rows.Scan( &customerorder.ID,  &customerorder.CreatedAt,  &customerorder.UpdatedAt,  &customerorder.CustomerCart, ); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from CustomerOrder row-> "+err.Error())
		}
		list = append(list, customerorder)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from CustomerOrder-> "+err.Error())
	}

	return &v1.ListCustomerOrderResponse{
		Api:   apiVersion,
		Items: list,
	}, nil
}

// Update CustomerOrder
func (s *shrikeServiceServer) UpdateCustomerOrder(ctx context.Context, req *v1.UpdateCustomerOrderRequest) (*v1.UpdateCustomerOrderResponse, error) {
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

	// update customer_order
	res, err := c.ExecContext(ctx, "UPDATE customer_order SET title=$1 WHERE id=$2",
		req.Item.Title, req.Item.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update CustomerOrder-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("CustomerOrder with ID='%d' is not found",
			req.Item.Id))
	}

	return &v1.UpdateCustomerOrderResponse{
		Api:     apiVersion,
		Updated: rows,
	}, nil
}

// Delete customer_order
func (s *shrikeServiceServer) DeleteCustomerOrder(ctx context.Context, req *v1.DeleteCustomerOrderRequest) (*v1.DeleteCustomerOrderResponse, error) {
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

	// delete customer_order
	res, err := c.ExecContext(ctx, "DELETE FROM customer_order WHERE id=$1", req.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete CustomerOrder-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("CustomerOrder with ID='%d' is not found",
			req.Id))
	}

	return &v1.DeleteCustomerOrderResponse{
		Api:     apiVersion,
		Deleted: rows,
	}, nil
}
