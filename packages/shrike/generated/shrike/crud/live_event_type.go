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

// NewShrikeServiceServer creates LiveEventType service
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
func (s *shrikeServiceServer) CreateLiveEventType(ctx context.Context, req *v1.CreateLiveEventTypeRequest) (*v1.CreateLiveEventTypeResponse, error) {
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
	// insert LiveEventType entity data
	err = c.QueryRowContext(ctx, "INSERT INTO liveeventtype (title) VALUES($1)  RETURNING id;",
		req.Item.Title).Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into LiveEventType-> "+err.Error())
	}

	// get ID of creates LiveEventType
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created LiveEventType-> "+err.Error())
	}

	return &v1.CreateLiveEventTypeResponse{
		Api: apiVersion,
		Id:  id,
	}, nil
}

// Get liveeventtype by id.
func (s *shrikeServiceServer) GetLiveEventType(ctx context.Context, req *v1.GetLiveEventTypeRequest) (*v1.GetLiveEventTypeResponse, error) {
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

	// query LiveEventType by ID
	rows, err := c.QueryContext(ctx, "SELECT id, title FROM liveeventtype WHERE id=$1",
		req.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from LiveEventType-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from LiveEventType-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("LiveEventType with ID='%d' is not found",
			req.Id))
	}

	// get LiveEventType data
	var td v1.LiveEventType
	if err := rows.Scan(&td.Id, &td.Title); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from LiveEventType row-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple LiveEventType rows with ID='%d'",
			req.Id))
	}

	return &v1.GetLiveEventTypeResponse{
		Api:  apiVersion,
		Item: &td,
	}, nil

}

// Read all todo tasks
func (s *shrikeServiceServer) ListLiveEventType(ctx context.Context, req *v1.ListLiveEventTypeRequest) (*v1.ListLiveEventTypeResponse, error) {
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

	// get LiveEventType list
	rows, err := c.QueryContext(ctx, "SELECT id,title FROM LiveEventType")
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from LiveEventType-> "+err.Error())
	}
	defer rows.Close()

	list := []*v1.LiveEventType{}
	for rows.Next() {
		td := new(v1.LiveEventType)
		if err := rows.Scan(&td.Id, &td.Title); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from LiveEventType row-> "+err.Error())
		}
		list = append(list, td)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from LiveEventType-> "+err.Error())
	}

	return &v1.ListLiveEventTypeResponse{
		Api:   apiVersion,
		Items: list,
	}, nil
}

// Update todo task
func (s *shrikeServiceServer) UpdateLiveEventType(ctx context.Context, req *v1.UpdateLiveEventTypeRequest) (*v1.UpdateLiveEventTypeResponse, error) {
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

	// update liveeventtype
	res, err := c.ExecContext(ctx, "UPDATE liveeventtype SET title=$1 WHERE id=$2",
		req.Item.Title, req.Item.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update liveeventtype-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("liveeventtype with ID='%d' is not found",
			req.Item.Id))
	}

	return &v1.UpdateLiveEventTypeResponse{
		Api:     apiVersion,
		Updated: rows,
	}, nil
}

// Delete liveeventtype
func (s *shrikeServiceServer) DeleteLiveEventType(ctx context.Context, req *v1.DeleteLiveEventTypeRequest) (*v1.DeleteLiveEventTypeResponse, error) {
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

	// delete liveeventtype
	res, err := c.ExecContext(ctx, "DELETE FROM liveeventtype WHERE id=$1", req.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete liveeventtype-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("liveeventtype with ID='%d' is not found",
			req.Id))
	}

	return &v1.DeleteLiveEventTypeResponse{
		Api:     apiVersion,
		Deleted: rows,
	}, nil
}
