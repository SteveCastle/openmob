package v1

import (
	"context"
	"fmt"

	v1 "github.com/SteveCastle/openmob/packages/shrike/src/pkg/api/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

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
	err = c.QueryRowContext(ctx, "INSERT INTO product_membership (id, created_at, updated_at, cause, product) VALUES($1, $2, $3, $4, $5)  RETURNING id;",
		 req.Item.ID,  req.Item.CreatedAt,  req.Item.UpdatedAt,  req.Item.Cause,  req.Item.Product, ).Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into ProductMembership-> "+err.Error())
	}

	// get ID of creates ProductMembership
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created ProductMembership-> "+err.Error())
	}

	return &v1.CreateProductMembershipResponse{
		Api: apiVersion,
		ID:  id,
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
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, cause, product FROM product_membership WHERE id=$1",
		req.ID)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from ProductMembership-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from ProductMembership-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("ProductMembership with ID='%d' is not found",
			req.ID))
	}

	// get ProductMembership data
	var productmembership v1.ProductMembership
	if err := rows.Scan( &productmembership.ID,  &productmembership.CreatedAt,  &productmembership.UpdatedAt,  &productmembership.Cause,  &productmembership.Product, ); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from ProductMembership row-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple ProductMembership rows with ID='%d'",
			req.ID))
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
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, cause, product FROM product_membership")
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
	res, err := c.ExecContext(ctx, "UPDATE product_membership SET id=$1, created_at=$2, updated_at=$3, cause=$4, product=$5 WHERE id=$1",
		req.Item.ID,req.Item.CreatedAt,req.Item.UpdatedAt,req.Item.Cause,req.Item.Product, )
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update ProductMembership-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("ProductMembership with ID='%d' is not found",
			req.Item.ID))
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
	res, err := c.ExecContext(ctx, "DELETE FROM product_membership WHERE id=$1", req.ID)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete ProductMembership-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("ProductMembership with ID='%d' is not found",
			req.ID))
	}

	return &v1.DeleteProductMembershipResponse{
		Api:     apiVersion,
		Deleted: rows,
	}, nil
}