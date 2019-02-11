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

// NewShrikeServiceServer creates Photo service
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
func (s *shrikeServiceServer) CreatePhoto(ctx context.Context, req *v1.CreatePhotoRequest) (*v1.CreatePhotoResponse, error) {
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
	// insert Photo entity data
	err = c.QueryRowContext(ctx, "INSERT INTO photo (title) VALUES($1)  RETURNING id;",
		req.Item.Title).Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into Photo-> "+err.Error())
	}

	// get ID of creates Photo
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created Photo-> "+err.Error())
	}

	return &v1.CreatePhotoResponse{
		Api: apiVersion,
		Id:  id,
	}, nil
}

// Get photo by id.
func (s *shrikeServiceServer) GetPhoto(ctx context.Context, req *v1.GetPhotoRequest) (*v1.GetPhotoResponse, error) {
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

	// query Photo by ID
	rows, err := c.QueryContext(ctx, "SELECT id, title FROM photo WHERE id=$1",
		req.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from Photo-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from Photo-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Photo with ID='%d' is not found",
			req.Id))
	}

	// get Photo data
	var td v1.Photo
	if err := rows.Scan(&td.Id, &td.Title); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from Photo row-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple Photo rows with ID='%d'",
			req.Id))
	}

	return &v1.GetPhotoResponse{
		Api:  apiVersion,
		Item: &td,
	}, nil

}

// Read all todo tasks
func (s *shrikeServiceServer) ListPhoto(ctx context.Context, req *v1.ListPhotoRequest) (*v1.ListPhotoResponse, error) {
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

	// get Photo list
	rows, err := c.QueryContext(ctx, "SELECT id,title FROM Photo")
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from Photo-> "+err.Error())
	}
	defer rows.Close()

	list := []*v1.Photo{}
	for rows.Next() {
		td := new(v1.Photo)
		if err := rows.Scan(&td.Id, &td.Title); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from Photo row-> "+err.Error())
		}
		list = append(list, td)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from Photo-> "+err.Error())
	}

	return &v1.ListPhotoResponse{
		Api:   apiVersion,
		Items: list,
	}, nil
}

// Update todo task
func (s *shrikeServiceServer) UpdatePhoto(ctx context.Context, req *v1.UpdatePhotoRequest) (*v1.UpdatePhotoResponse, error) {
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

	// update photo
	res, err := c.ExecContext(ctx, "UPDATE photo SET title=$1 WHERE id=$2",
		req.Item.Title, req.Item.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update photo-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("photo with ID='%d' is not found",
			req.Item.Id))
	}

	return &v1.UpdatePhotoResponse{
		Api:     apiVersion,
		Updated: rows,
	}, nil
}

// Delete photo
func (s *shrikeServiceServer) DeletePhoto(ctx context.Context, req *v1.DeletePhotoRequest) (*v1.DeletePhotoResponse, error) {
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

	// delete photo
	res, err := c.ExecContext(ctx, "DELETE FROM photo WHERE id=$1", req.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete photo-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("photo with ID='%d' is not found",
			req.Id))
	}

	return &v1.DeletePhotoResponse{
		Api:     apiVersion,
		Deleted: rows,
	}, nil
}
