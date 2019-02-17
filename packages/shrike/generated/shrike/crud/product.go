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

// NewShrikeServiceServer creates Product service
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
	var id int64
	// insert Product entity data
	err = c.QueryRowContext(ctx, "INSERT INTO product (id, created_at, updated_at, title, product_type, ) VALUES($1, $2, $3, $4, $5, )  RETURNING id;",
		 req.ItemID  req.ItemCreatedAt  req.ItemUpdatedAt  req.ItemTitle  req.ItemProductType ).Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into Product-> "+err.Error())
	}

	// get ID of creates Product
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created Product-> "+err.Error())
	}

	return &v1.CreateProductResponse{
		Api: apiVersion,
		Id:  id,
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
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, title, product_type,  FROM product WHERE id=$1",
		req.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from Product-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from Product-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Product with ID='%d' is not found",
			req.Id))
	}

	// get Product data
	var product v1.Product
	if err := rows.Scan( &product.ID,  &product.CreatedAt,  &product.UpdatedAt,  &product.Title,  &product.ProductType, ); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from Product row-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple Product rows with ID='%d'",
			req.Id))
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

	// get Product list
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, title, product_type,  FROM product")
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from Product-> "+err.Error())
	}
	defer rows.Close()

	list := []*v1.Product{}
	for rows.Next() {
		product := new(v1.Product)
		if err := rows.Scan( &product.ID,  &product.CreatedAt,  &product.UpdatedAt,  &product.Title,  &product.ProductType, ); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from Product row-> "+err.Error())
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
	res, err := c.ExecContext(ctx, "UPDATE product SET title=$1 WHERE id=$2",
		req.Item.Title, req.Item.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update Product-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Product with ID='%d' is not found",
			req.Item.Id))
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
	res, err := c.ExecContext(ctx, "DELETE FROM product WHERE id=$1", req.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete Product-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Product with ID='%d' is not found",
			req.Id))
	}

	return &v1.DeleteProductResponse{
		Api:     apiVersion,
		Deleted: rows,
	}, nil
}
