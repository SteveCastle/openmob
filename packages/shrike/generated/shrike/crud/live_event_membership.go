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

// NewShrikeServiceServer creates LiveEventMembership service
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
func (s *shrikeServiceServer) CreateLiveEventMembership(ctx context.Context, req *v1.CreateLiveEventMembershipRequest) (*v1.CreateLiveEventMembershipResponse, error) {
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
	// insert LiveEventMembership entity data
	err = c.QueryRowContext(ctx, "INSERT INTO live_event_membership (title) VALUES($1)  RETURNING id;",
		req.Item.Title).Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into LiveEventMembership-> "+err.Error())
	}

	// get ID of creates LiveEventMembership
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created LiveEventMembership-> "+err.Error())
	}

	return &v1.CreateLiveEventMembershipResponse{
		Api: apiVersion,
		Id:  id,
	}, nil
}

// Get live_event_membership by id.
func (s *shrikeServiceServer) GetLiveEventMembership(ctx context.Context, req *v1.GetLiveEventMembershipRequest) (*v1.GetLiveEventMembershipResponse, error) {
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

	// query LiveEventMembership by ID
	rows, err := c.QueryContext(ctx, "SELECT id, title FROM live_event_membership WHERE id=$1",
		req.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from LiveEventMembership-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from LiveEventMembership-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("LiveEventMembership with ID='%d' is not found",
			req.Id))
	}

	// get LiveEventMembership data
	var td v1.LiveEventMembership
	if err := rows.Scan(&td.Id, &td.Title); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from LiveEventMembership row-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple LiveEventMembership rows with ID='%d'",
			req.Id))
	}

	return &v1.GetLiveEventMembershipResponse{
		Api:  apiVersion,
		Item: &td,
	}, nil

}

// Read all todo tasks
func (s *shrikeServiceServer) ListLiveEventMembership(ctx context.Context, req *v1.ListLiveEventMembershipRequest) (*v1.ListLiveEventMembershipResponse, error) {
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

	// get LiveEventMembership list
	rows, err := c.QueryContext(ctx, "SELECT id,title FROM live_event_membership")
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from LiveEventMembership-> "+err.Error())
	}
	defer rows.Close()

	list := []*v1.LiveEventMembership{}
	for rows.Next() {
		td := new(v1.LiveEventMembership)
		if err := rows.Scan(&td.Id, &td.Title); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from LiveEventMembership row-> "+err.Error())
		}
		list = append(list, td)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from LiveEventMembership-> "+err.Error())
	}

	return &v1.ListLiveEventMembershipResponse{
		Api:   apiVersion,
		Items: list,
	}, nil
}

// Update todo task
func (s *shrikeServiceServer) UpdateLiveEventMembership(ctx context.Context, req *v1.UpdateLiveEventMembershipRequest) (*v1.UpdateLiveEventMembershipResponse, error) {
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

	// update live_event_membership
	res, err := c.ExecContext(ctx, "UPDATE live_event_membership SET title=$1 WHERE id=$2",
		req.Item.Title, req.Item.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update LiveEventMembership-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("LiveEventMembership with ID='%d' is not found",
			req.Item.Id))
	}

	return &v1.UpdateLiveEventMembershipResponse{
		Api:     apiVersion,
		Updated: rows,
	}, nil
}

// Delete live_event_membership
func (s *shrikeServiceServer) DeleteLiveEventMembership(ctx context.Context, req *v1.DeleteLiveEventMembershipRequest) (*v1.DeleteLiveEventMembershipResponse, error) {
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

	// delete live_event_membership
	res, err := c.ExecContext(ctx, "DELETE FROM live_event_membership WHERE id=$1", req.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete LiveEventMembership-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("LiveEventMembership with ID='%d' is not found",
			req.Id))
	}

	return &v1.DeleteLiveEventMembershipResponse{
		Api:     apiVersion,
		Deleted: rows,
	}, nil
}
