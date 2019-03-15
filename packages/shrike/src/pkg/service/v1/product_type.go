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

// Create new ProductType
func (s *shrikeServiceServer) CreateProductType(ctx context.Context, req *v1.CreateProductTypeRequest) (*v1.CreateProductTypeResponse, error) {
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
	// insert ProductType entity data
	err = c.QueryRowContext(ctx, "INSERT INTO product_type (title) VALUES($1)  RETURNING id;",
		req.Item.Title).Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into ProductType-> "+err.Error())
	}

	// get ID of creates ProductType
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created ProductType-> "+err.Error())
	}

	return &v1.CreateProductTypeResponse{
		Api: apiVersion,
		ID:  id,
	}, nil
}

// Get product_type by id.
func (s *shrikeServiceServer) GetProductType(ctx context.Context, req *v1.GetProductTypeRequest) (*v1.GetProductTypeResponse, error) {
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

	// query ProductType by ID
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, title FROM product_type WHERE id=$1",
		req.ID)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from ProductType-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from ProductType-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("ProductType with ID='%s' is not found",
			req.ID))
	}

	// scan ProductType data into protobuf model
	var producttype v1.ProductType
	var createdAt time.Time
	var updatedAt time.Time

	if err := rows.Scan(&producttype.ID, &createdAt, &updatedAt, &producttype.Title); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from ProductType row-> "+err.Error())
	}

	// Convert time.Time from database into proto timestamp.
	producttype.CreatedAt, err = ptypes.TimestampProto(createdAt)
	if err != nil {
		return nil, status.Error(codes.Unknown, "createdAt field has invalid format-> "+err.Error())
	}
	producttype.UpdatedAt, err = ptypes.TimestampProto(updatedAt)
	if err != nil {
		return nil, status.Error(codes.Unknown, "updatedAt field has invalid format-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple ProductType rows with ID='%s'",
			req.ID))
	}

	return &v1.GetProductTypeResponse{
		Api:  apiVersion,
		Item: &producttype,
	}, nil

}

// Read all ProductType
func (s *shrikeServiceServer) ListProductType(ctx context.Context, req *v1.ListProductTypeRequest) (*v1.ListProductTypeResponse, error) {
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

	// get ProductType list
	queries.BuildProductTypeFilters(req.Filters, req.Ordering, req.Limit)
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, title FROM product_type")
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from ProductType-> "+err.Error())
	}
	defer rows.Close()

	// Variables to store results returned by database.
	list := []*v1.ProductType{}
	var createdAt time.Time
	var updatedAt time.Time

	for rows.Next() {
		producttype := new(v1.ProductType)
		if err := rows.Scan(&producttype.ID, &createdAt, &updatedAt, &producttype.Title); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from ProductType row-> "+err.Error())
		}
		// Convert time.Time from database into proto timestamp.
		producttype.CreatedAt, err = ptypes.TimestampProto(createdAt)
		if err != nil {
			return nil, status.Error(codes.Unknown, "createdAt field has invalid format-> "+err.Error())
		}
		producttype.UpdatedAt, err = ptypes.TimestampProto(updatedAt)
		if err != nil {
			return nil, status.Error(codes.Unknown, "updatedAt field has invalid format-> "+err.Error())
		}

		list = append(list, producttype)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from ProductType-> "+err.Error())
	}

	return &v1.ListProductTypeResponse{
		Api:   apiVersion,
		Items: list,
	}, nil
}

// Update ProductType
func (s *shrikeServiceServer) UpdateProductType(ctx context.Context, req *v1.UpdateProductTypeRequest) (*v1.UpdateProductTypeResponse, error) {
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

	// update product_type
	res, err := c.ExecContext(ctx, "UPDATE product_type SET title=$2 WHERE id=$1",
		req.Item.ID, req.Item.Title)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update ProductType-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("ProductType with ID='%s' is not found",
			req.Item.ID))
	}

	return &v1.UpdateProductTypeResponse{
		Api:     apiVersion,
		Updated: rows,
	}, nil
}

// Delete product_type
func (s *shrikeServiceServer) DeleteProductType(ctx context.Context, req *v1.DeleteProductTypeRequest) (*v1.DeleteProductTypeResponse, error) {
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

	// delete product_type
	res, err := c.ExecContext(ctx, "DELETE FROM product_type WHERE id=$1", req.ID)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete ProductType-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("ProductType with ID='%s' is not found",
			req.ID))
	}

	return &v1.DeleteProductTypeResponse{
		Api:     apiVersion,
		Deleted: rows,
	}, nil
}
