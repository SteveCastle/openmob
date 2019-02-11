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

// NewShrikeServiceServer creates LayoutType service
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
func (s *shrikeServiceServer) CreateLayoutType(ctx context.Context, req *v1.CreateLayoutTypeRequest) (*v1.CreateLayoutTypeResponse, error) {
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
	// insert LayoutType entity data
	err = c.QueryRowContext(ctx, "INSERT INTO layout_type (title) VALUES($1)  RETURNING id;",
		req.Item.Title).Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into LayoutType-> "+err.Error())
	}

	// get ID of creates LayoutType
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created LayoutType-> "+err.Error())
	}

	return &v1.CreateLayoutTypeResponse{
		Api: apiVersion,
		Id:  id,
	}, nil
}

// Get layout_type by id.
func (s *shrikeServiceServer) GetLayoutType(ctx context.Context, req *v1.GetLayoutTypeRequest) (*v1.GetLayoutTypeResponse, error) {
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

	// query LayoutType by ID
	rows, err := c.QueryContext(ctx, "SELECT id, title FROM layout_type WHERE id=$1",
		req.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from LayoutType-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from LayoutType-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("LayoutType with ID='%d' is not found",
			req.Id))
	}

	// get LayoutType data
	var td v1.LayoutType
	if err := rows.Scan(&td.Id, &td.Title); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from LayoutType row-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple LayoutType rows with ID='%d'",
			req.Id))
	}

	return &v1.GetLayoutTypeResponse{
		Api:  apiVersion,
		Item: &td,
	}, nil

}

// Read all todo tasks
func (s *shrikeServiceServer) ListLayoutType(ctx context.Context, req *v1.ListLayoutTypeRequest) (*v1.ListLayoutTypeResponse, error) {
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

	// get LayoutType list
	rows, err := c.QueryContext(ctx, "SELECT id,title FROM layout_type")
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from LayoutType-> "+err.Error())
	}
	defer rows.Close()

	list := []*v1.LayoutType{}
	for rows.Next() {
		td := new(v1.LayoutType)
		if err := rows.Scan(&td.Id, &td.Title); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from LayoutType row-> "+err.Error())
		}
		list = append(list, td)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from LayoutType-> "+err.Error())
	}

	return &v1.ListLayoutTypeResponse{
		Api:   apiVersion,
		Items: list,
	}, nil
}

// Update todo task
func (s *shrikeServiceServer) UpdateLayoutType(ctx context.Context, req *v1.UpdateLayoutTypeRequest) (*v1.UpdateLayoutTypeResponse, error) {
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

	// update layout_type
	res, err := c.ExecContext(ctx, "UPDATE layout_type SET title=$1 WHERE id=$2",
		req.Item.Title, req.Item.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update LayoutType-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("LayoutType with ID='%d' is not found",
			req.Item.Id))
	}

	return &v1.UpdateLayoutTypeResponse{
		Api:     apiVersion,
		Updated: rows,
	}, nil
}

// Delete layout_type
func (s *shrikeServiceServer) DeleteLayoutType(ctx context.Context, req *v1.DeleteLayoutTypeRequest) (*v1.DeleteLayoutTypeResponse, error) {
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

	// delete layout_type
	res, err := c.ExecContext(ctx, "DELETE FROM layout_type WHERE id=$1", req.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete LayoutType-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("LayoutType with ID='%d' is not found",
			req.Id))
	}

	return &v1.DeleteLayoutTypeResponse{
		Api:     apiVersion,
		Deleted: rows,
	}, nil
}
