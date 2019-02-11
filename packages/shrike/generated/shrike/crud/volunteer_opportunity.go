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

// NewShrikeServiceServer creates VolunteerOpportunity service
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
func (s *shrikeServiceServer) CreateVolunteerOpportunity(ctx context.Context, req *v1.CreateVolunteerOpportunityRequest) (*v1.CreateVolunteerOpportunityResponse, error) {
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
	// insert VolunteerOpportunity entity data
	err = c.QueryRowContext(ctx, "INSERT INTO volunteeropportunity (title) VALUES($1)  RETURNING id;",
		req.Item.Title).Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into VolunteerOpportunity-> "+err.Error())
	}

	// get ID of creates VolunteerOpportunity
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created VolunteerOpportunity-> "+err.Error())
	}

	return &v1.CreateVolunteerOpportunityResponse{
		Api: apiVersion,
		Id:  id,
	}, nil
}

// Get volunteeropportunity by id.
func (s *shrikeServiceServer) GetVolunteerOpportunity(ctx context.Context, req *v1.GetVolunteerOpportunityRequest) (*v1.GetVolunteerOpportunityResponse, error) {
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

	// query VolunteerOpportunity by ID
	rows, err := c.QueryContext(ctx, "SELECT id, title FROM volunteeropportunity WHERE id=$1",
		req.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from VolunteerOpportunity-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from VolunteerOpportunity-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("VolunteerOpportunity with ID='%d' is not found",
			req.Id))
	}

	// get VolunteerOpportunity data
	var td v1.VolunteerOpportunity
	if err := rows.Scan(&td.Id, &td.Title); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from VolunteerOpportunity row-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple VolunteerOpportunity rows with ID='%d'",
			req.Id))
	}

	return &v1.GetVolunteerOpportunityResponse{
		Api:  apiVersion,
		Item: &td,
	}, nil

}

// Read all todo tasks
func (s *shrikeServiceServer) ListVolunteerOpportunity(ctx context.Context, req *v1.ListVolunteerOpportunityRequest) (*v1.ListVolunteerOpportunityResponse, error) {
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

	// get VolunteerOpportunity list
	rows, err := c.QueryContext(ctx, "SELECT id,title FROM VolunteerOpportunity")
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from VolunteerOpportunity-> "+err.Error())
	}
	defer rows.Close()

	list := []*v1.VolunteerOpportunity{}
	for rows.Next() {
		td := new(v1.VolunteerOpportunity)
		if err := rows.Scan(&td.Id, &td.Title); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from VolunteerOpportunity row-> "+err.Error())
		}
		list = append(list, td)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from VolunteerOpportunity-> "+err.Error())
	}

	return &v1.ListVolunteerOpportunityResponse{
		Api:   apiVersion,
		Items: list,
	}, nil
}

// Update todo task
func (s *shrikeServiceServer) UpdateVolunteerOpportunity(ctx context.Context, req *v1.UpdateVolunteerOpportunityRequest) (*v1.UpdateVolunteerOpportunityResponse, error) {
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

	// update volunteeropportunity
	res, err := c.ExecContext(ctx, "UPDATE volunteeropportunity SET title=$1 WHERE id=$2",
		req.Item.Title, req.Item.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update volunteeropportunity-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("volunteeropportunity with ID='%d' is not found",
			req.Item.Id))
	}

	return &v1.UpdateVolunteerOpportunityResponse{
		Api:     apiVersion,
		Updated: rows,
	}, nil
}

// Delete volunteeropportunity
func (s *shrikeServiceServer) DeleteVolunteerOpportunity(ctx context.Context, req *v1.DeleteVolunteerOpportunityRequest) (*v1.DeleteVolunteerOpportunityResponse, error) {
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

	// delete volunteeropportunity
	res, err := c.ExecContext(ctx, "DELETE FROM volunteeropportunity WHERE id=$1", req.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete volunteeropportunity-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("volunteeropportunity with ID='%d' is not found",
			req.Id))
	}

	return &v1.DeleteVolunteerOpportunityResponse{
		Api:     apiVersion,
		Deleted: rows,
	}, nil
}
