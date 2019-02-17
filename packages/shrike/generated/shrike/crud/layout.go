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

// NewShrikeServiceServer creates Layout service
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

// Create new Layout
func (s *shrikeServiceServer) CreateLayout(ctx context.Context, req *v1.CreateLayoutRequest) (*v1.CreateLayoutResponse, error) {
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
	// insert Layout entity data
	err = c.QueryRowContext(ctx, "INSERT INTO layout ( id  created_at  updated_at  layout_type ) VALUES( $1 $2 $3 $4)  RETURNING id;",
		 req.ItemID  req.ItemCreatedAt  req.ItemUpdatedAt  req.ItemLayoutType ).Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into Layout-> "+err.Error())
	}

	// get ID of creates Layout
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created Layout-> "+err.Error())
	}

	return &v1.CreateLayoutResponse{
		Api: apiVersion,
		Id:  id,
	}, nil
}

// Get layout by id.
func (s *shrikeServiceServer) GetLayout(ctx context.Context, req *v1.GetLayoutRequest) (*v1.GetLayoutResponse, error) {
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

	// query Layout by ID
	rows, err := c.QueryContext(ctx, "SELECT id, title FROM layout WHERE id=$1",
		req.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from Layout-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from Layout-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Layout with ID='%d' is not found",
			req.Id))
	}

	// get Layout data
	var layout v1.Layout
	if err := rows.Scan(&layout.Id, &layout.Title); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from Layout row-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple Layout rows with ID='%d'",
			req.Id))
	}

	return &v1.GetLayoutResponse{
		Api:  apiVersion,
		Item: &layout,
	}, nil

}

// Read all Layout
func (s *shrikeServiceServer) ListLayout(ctx context.Context, req *v1.ListLayoutRequest) (*v1.ListLayoutResponse, error) {
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

	// get Layout list
	rows, err := c.QueryContext(ctx, "SELECT id,title FROM layout")
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from Layout-> "+err.Error())
	}
	defer rows.Close()

	list := []*v1.Layout{}
	for rows.Next() {
		layout := new(v1.Layout)
		if err := rows.Scan(&layout.Id, &layout.Title); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from Layout row-> "+err.Error())
		}
		list = append(list, layout)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from Layout-> "+err.Error())
	}

	return &v1.ListLayoutResponse{
		Api:   apiVersion,
		Items: list,
	}, nil
}

// Update Layout
func (s *shrikeServiceServer) UpdateLayout(ctx context.Context, req *v1.UpdateLayoutRequest) (*v1.UpdateLayoutResponse, error) {
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

	// update layout
	res, err := c.ExecContext(ctx, "UPDATE layout SET title=$1 WHERE id=$2",
		req.Item.Title, req.Item.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update Layout-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Layout with ID='%d' is not found",
			req.Item.Id))
	}

	return &v1.UpdateLayoutResponse{
		Api:     apiVersion,
		Updated: rows,
	}, nil
}

// Delete layout
func (s *shrikeServiceServer) DeleteLayout(ctx context.Context, req *v1.DeleteLayoutRequest) (*v1.DeleteLayoutResponse, error) {
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

	// delete layout
	res, err := c.ExecContext(ctx, "DELETE FROM layout WHERE id=$1", req.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete Layout-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Layout with ID='%d' is not found",
			req.Id))
	}

	return &v1.DeleteLayoutResponse{
		Api:     apiVersion,
		Deleted: rows,
	}, nil
}
