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

// NewShrikeServiceServer creates ProductType service
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
	var id int64
	// insert ProductType entity data
	err = c.QueryRowContext(ctx, "INSERT INTO product_type ( id  created_at  updated_at  title ) VALUES( $1 $2 $3 $4)  RETURNING id;",
		 req.ItemID  req.ItemCreatedAt  req.ItemUpdatedAt  req.ItemTitle ).Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into ProductType-> "+err.Error())
	}

	// get ID of creates ProductType
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created ProductType-> "+err.Error())
	}

	return &v1.CreateProductTypeResponse{
		Api: apiVersion,
		Id:  id,
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
	rows, err := c.QueryContext(ctx, "SELECT id, title FROM product_type WHERE id=$1",
		req.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from ProductType-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from ProductType-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("ProductType with ID='%d' is not found",
			req.Id))
	}

	// get ProductType data
	var producttype v1.ProductType
	if err := rows.Scan(&producttype.Id, &producttype.Title); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from ProductType row-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple ProductType rows with ID='%d'",
			req.Id))
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
	rows, err := c.QueryContext(ctx, "SELECT id,title FROM product_type")
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from ProductType-> "+err.Error())
	}
	defer rows.Close()

	list := []*v1.ProductType{}
	for rows.Next() {
		producttype := new(v1.ProductType)
		if err := rows.Scan(&producttype.Id, &producttype.Title); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from ProductType row-> "+err.Error())
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
	res, err := c.ExecContext(ctx, "UPDATE product_type SET title=$1 WHERE id=$2",
		req.Item.Title, req.Item.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update ProductType-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("ProductType with ID='%d' is not found",
			req.Item.Id))
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
	res, err := c.ExecContext(ctx, "DELETE FROM product_type WHERE id=$1", req.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete ProductType-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("ProductType with ID='%d' is not found",
			req.Id))
	}

	return &v1.DeleteProductTypeResponse{
		Api:     apiVersion,
		Deleted: rows,
	}, nil
}
