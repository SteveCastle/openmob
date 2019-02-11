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

// Create new todo task
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
	err = c.QueryRowContext(ctx, "INSERT INTO productmembership (title) VALUES($1)  RETURNING id;",
		req.Item.Title).Scan(&id)
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

// Get productmembership by id.
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
	rows, err := c.QueryContext(ctx, "SELECT id, title FROM productmembership WHERE id=$1",
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
	var td v1.ProductMembership
	if err := rows.Scan(&td.Id, &td.Title); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from ProductMembership row-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple ProductMembership rows with ID='%d'",
			req.Id))
	}

	return &v1.GetProductMembershipResponse{
		Api:  apiVersion,
		Item: &td,
	}, nil

}

// Read all todo tasks
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
	rows, err := c.QueryContext(ctx, "SELECT id,title FROM ProductMembership")
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from ProductMembership-> "+err.Error())
	}
	defer rows.Close()

	list := []*v1.ProductMembership{}
	for rows.Next() {
		td := new(v1.ProductMembership)
		if err := rows.Scan(&td.Id, &td.Title); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from ProductMembership row-> "+err.Error())
		}
		list = append(list, td)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from ProductMembership-> "+err.Error())
	}

	return &v1.ListProductMembershipResponse{
		Api:   apiVersion,
		Items: list,
	}, nil
}

// Update todo task
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

	// update productmembership
	res, err := c.ExecContext(ctx, "UPDATE productmembership SET title=$1 WHERE id=$2",
		req.Item.Title, req.Item.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update productmembership-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("productmembership with ID='%d' is not found",
			req.Item.Id))
	}

	return &v1.UpdateProductMembershipResponse{
		Api:     apiVersion,
		Updated: rows,
	}, nil
}

// Delete productmembership
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

	// delete productmembership
	res, err := c.ExecContext(ctx, "DELETE FROM productmembership WHERE id=$1", req.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete productmembership-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("productmembership with ID='%d' is not found",
			req.Id))
	}

	return &v1.DeleteProductMembershipResponse{
		Api:     apiVersion,
		Deleted: rows,
	}, nil
}
