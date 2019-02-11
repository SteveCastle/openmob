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

// NewShrikeServiceServer creates LiveEvent service
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
func (s *shrikeServiceServer) CreateLiveEvent(ctx context.Context, req *v1.CreateLiveEventRequest) (*v1.CreateLiveEventResponse, error) {
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
	// insert LiveEvent entity data
	err = c.QueryRowContext(ctx, "INSERT INTO live_event (title) VALUES($1)  RETURNING id;",
		req.Item.Title).Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into LiveEvent-> "+err.Error())
	}

	// get ID of creates LiveEvent
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created LiveEvent-> "+err.Error())
	}

	return &v1.CreateLiveEventResponse{
		Api: apiVersion,
		Id:  id,
	}, nil
}

// Get live_event by id.
func (s *shrikeServiceServer) GetLiveEvent(ctx context.Context, req *v1.GetLiveEventRequest) (*v1.GetLiveEventResponse, error) {
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

	// query LiveEvent by ID
	rows, err := c.QueryContext(ctx, "SELECT id, title FROM live_event WHERE id=$1",
		req.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from LiveEvent-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from LiveEvent-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("LiveEvent with ID='%d' is not found",
			req.Id))
	}

	// get LiveEvent data
	var td v1.LiveEvent
	if err := rows.Scan(&td.Id, &td.Title); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from LiveEvent row-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple LiveEvent rows with ID='%d'",
			req.Id))
	}

	return &v1.GetLiveEventResponse{
		Api:  apiVersion,
		Item: &td,
	}, nil

}

// Read all todo tasks
func (s *shrikeServiceServer) ListLiveEvent(ctx context.Context, req *v1.ListLiveEventRequest) (*v1.ListLiveEventResponse, error) {
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

	// get LiveEvent list
	rows, err := c.QueryContext(ctx, "SELECT id,title FROM live_event")
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from LiveEvent-> "+err.Error())
	}
	defer rows.Close()

	list := []*v1.LiveEvent{}
	for rows.Next() {
		td := new(v1.LiveEvent)
		if err := rows.Scan(&td.Id, &td.Title); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from LiveEvent row-> "+err.Error())
		}
		list = append(list, td)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from LiveEvent-> "+err.Error())
	}

	return &v1.ListLiveEventResponse{
		Api:   apiVersion,
		Items: list,
	}, nil
}

// Update todo task
func (s *shrikeServiceServer) UpdateLiveEvent(ctx context.Context, req *v1.UpdateLiveEventRequest) (*v1.UpdateLiveEventResponse, error) {
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

	// update live_event
	res, err := c.ExecContext(ctx, "UPDATE live_event SET title=$1 WHERE id=$2",
		req.Item.Title, req.Item.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update LiveEvent-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("LiveEvent with ID='%d' is not found",
			req.Item.Id))
	}

	return &v1.UpdateLiveEventResponse{
		Api:     apiVersion,
		Updated: rows,
	}, nil
}

// Delete live_event
func (s *shrikeServiceServer) DeleteLiveEvent(ctx context.Context, req *v1.DeleteLiveEventRequest) (*v1.DeleteLiveEventResponse, error) {
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

	// delete live_event
	res, err := c.ExecContext(ctx, "DELETE FROM live_event WHERE id=$1", req.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete LiveEvent-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("LiveEvent with ID='%d' is not found",
			req.Id))
	}

	return &v1.DeleteLiveEventResponse{
		Api:     apiVersion,
		Deleted: rows,
	}, nil
}
