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

// Create new Product
func (s *shrikeServiceServer) CreateProduct(ctx context.Context, req *v1.CreateProductRequest) (*v1.CreateProductResponse, error) {
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
	// insert Product entity data
	err = c.QueryRowContext(ctx, "INSERT INTO product (title, product_type) VALUES($1, $2)  RETURNING id;",
		req.Item.Title, req.Item.ProductType).Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into Product-> "+err.Error())
	}

	// get ID of creates Product
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created Product-> "+err.Error())
	}

	return &v1.CreateProductResponse{
		Api: apiVersion,
		ID:  id,
	}, nil
}

// Get product by id.
func (s *shrikeServiceServer) GetProduct(ctx context.Context, req *v1.GetProductRequest) (*v1.GetProductResponse, error) {
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

	// query Product by ID
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, title, product_type FROM product WHERE id=$1",
		req.ID)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from Product-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from Product-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Product with ID='%s' is not found",
			req.ID))
	}

	// scan Product data into protobuf model
	var product v1.Product
	var createdAt time.Time
	var updatedAt time.Time

	if err := rows.Scan(&product.ID, &createdAt, &updatedAt, &product.Title, &product.ProductType); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from Product row-> "+err.Error())
	}

	// Convert time.Time from database into proto timestamp.
	product.CreatedAt, err = ptypes.TimestampProto(createdAt)
	if err != nil {
		return nil, status.Error(codes.Unknown, "createdAt field has invalid format-> "+err.Error())
	}
	product.UpdatedAt, err = ptypes.TimestampProto(updatedAt)
	if err != nil {
		return nil, status.Error(codes.Unknown, "createdAt field has invalid format-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple Product rows with ID='%s'",
			req.ID))
	}

	return &v1.GetProductResponse{
		Api:  apiVersion,
		Item: &product,
	}, nil

}

// Read all Product
func (s *shrikeServiceServer) ListProduct(ctx context.Context, req *v1.ListProductRequest) (*v1.ListProductResponse, error) {
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

	// Generate SQL to select all columns in Product Table
	// Then generate filtering and ordering sql and finally run query.

	baseSQL := "SELECT id, created_at, updated_at, title, product_type FROM product"
	querySQL := queries.BuildProductFilters(req.Filters, req.Ordering, req.Limit)
	SQL := fmt.Sprintf("%s %s", baseSQL, querySQL)
	rows, err := c.QueryContext(ctx, SQL)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from Product-> "+err.Error())
	}
	defer rows.Close()

	// Variables to store results returned by database.
	list := []*v1.Product{}
	var createdAt time.Time
	var updatedAt time.Time

	for rows.Next() {
		product := new(v1.Product)
		if err := rows.Scan(&product.ID, &createdAt, &updatedAt, &product.Title, &product.ProductType); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from Product row-> "+err.Error())
		}
		// Convert time.Time from database into proto timestamp.
		product.CreatedAt, err = ptypes.TimestampProto(createdAt)
		if err != nil {
			return nil, status.Error(codes.Unknown, "createdAt field has invalid format-> "+err.Error())
		}
		product.UpdatedAt, err = ptypes.TimestampProto(updatedAt)
		if err != nil {
			return nil, status.Error(codes.Unknown, "createdAt field has invalid format-> "+err.Error())
		}

		list = append(list, product)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from Product-> "+err.Error())
	}

	return &v1.ListProductResponse{
		Api:   apiVersion,
		Items: list,
	}, nil
}

// Update Product
func (s *shrikeServiceServer) UpdateProduct(ctx context.Context, req *v1.UpdateProductRequest) (*v1.UpdateProductResponse, error) {
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

	// update product
	res, err := c.ExecContext(ctx, "UPDATE product SET title=$2, product_type=$3 WHERE id=$1",
		req.Item.ID, req.Item.Title, req.Item.ProductType)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update Product-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Product with ID='%s' is not found",
			req.Item.ID))
	}

	return &v1.UpdateProductResponse{
		Api:     apiVersion,
		Updated: rows,
	}, nil
}

// Delete product
func (s *shrikeServiceServer) DeleteProduct(ctx context.Context, req *v1.DeleteProductRequest) (*v1.DeleteProductResponse, error) {
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

	// delete product
	res, err := c.ExecContext(ctx, "DELETE FROM product WHERE id=$1", req.ID)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete Product-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Product with ID='%s' is not found",
			req.ID))
	}

	return &v1.DeleteProductResponse{
		Api:     apiVersion,
		Deleted: rows,
	}, nil
}
