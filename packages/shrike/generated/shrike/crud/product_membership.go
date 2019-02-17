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

// NewShrikeServiceServer creates ProductMembership service
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

// Create new ProductMembership
func (s *shrikeServiceServer) CreateProductMembership(ctx context.Context, req *v1.CreateProductMembershipRequest) (*v1.CreateProductMembershipResponse, error) {
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
	// insert ProductMembership entity data
	err = c.QueryRowContext(ctx, "INSERT INTO product_membership (id, created_at, updated_at, cause, product, ) VALUES($1, $2, $3, $4, $5, )  RETURNING id;",
		 req.ItemID  req.ItemCreatedAt  req.ItemUpdatedAt  req.ItemCause  req.ItemProduct ).Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into ProductMembership-> "+err.Error())
	}

	// get ID of creates ProductMembership
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created ProductMembership-> "+err.Error())
	}

	return &v1.CreateProductMembershipResponse{
		Api: apiVersion,
		Id:  id,
	}, nil
}

// Get product_membership by id.
func (s *shrikeServiceServer) GetProductMembership(ctx context.Context, req *v1.GetProductMembershipRequest) (*v1.GetProductMembershipResponse, error) {
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

	// query ProductMembership by ID
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, cause, product,  FROM product_membership WHERE id=$1",
		req.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from ProductMembership-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from ProductMembership-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("ProductMembership with ID='%d' is not found",
			req.Id))
	}

	// get ProductMembership data
	var productmembership v1.ProductMembership
	if err := rows.Scan( &productmembership.ID,  &productmembership.CreatedAt,  &productmembership.UpdatedAt,  &productmembership.Cause,  &productmembership.Product, ); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from ProductMembership row-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple ProductMembership rows with ID='%d'",
			req.Id))
	}

	return &v1.GetProductMembershipResponse{
		Api:  apiVersion,
		Item: &productmembership,
	}, nil

}

// Read all ProductMembership
func (s *shrikeServiceServer) ListProductMembership(ctx context.Context, req *v1.ListProductMembershipRequest) (*v1.ListProductMembershipResponse, error) {
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

	// get ProductMembership list
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, cause, product,  FROM product_membership")
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from ProductMembership-> "+err.Error())
	}
	defer rows.Close()

	list := []*v1.ProductMembership{}
	for rows.Next() {
		productmembership := new(v1.ProductMembership)
		if err := rows.Scan( &productmembership.ID,  &productmembership.CreatedAt,  &productmembership.UpdatedAt,  &productmembership.Cause,  &productmembership.Product, ); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from ProductMembership row-> "+err.Error())
		}
		list = append(list, productmembership)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from ProductMembership-> "+err.Error())
	}

	return &v1.ListProductMembershipResponse{
		Api:   apiVersion,
		Items: list,
	}, nil
}

// Update ProductMembership
func (s *shrikeServiceServer) UpdateProductMembership(ctx context.Context, req *v1.UpdateProductMembershipRequest) (*v1.UpdateProductMembershipResponse, error) {
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

	// update product_membership
	res, err := c.ExecContext(ctx, "UPDATE product_membership SET title=$1 WHERE id=$2",
		req.Item.Title, req.Item.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update ProductMembership-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("ProductMembership with ID='%d' is not found",
			req.Item.Id))
	}

	return &v1.UpdateProductMembershipResponse{
		Api:     apiVersion,
		Updated: rows,
	}, nil
}

// Delete product_membership
func (s *shrikeServiceServer) DeleteProductMembership(ctx context.Context, req *v1.DeleteProductMembershipRequest) (*v1.DeleteProductMembershipResponse, error) {
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

	// delete product_membership
	res, err := c.ExecContext(ctx, "DELETE FROM product_membership WHERE id=$1", req.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete ProductMembership-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("ProductMembership with ID='%d' is not found",
			req.Id))
	}

	return &v1.DeleteProductMembershipResponse{
		Api:     apiVersion,
		Deleted: rows,
	}, nil
}
