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

// NewShrikeServiceServer creates VolunteerOpportunityMembership service
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
func (s *shrikeServiceServer) CreateVolunteerOpportunityMembership(ctx context.Context, req *v1.CreateVolunteerOpportunityMembershipRequest) (*v1.CreateVolunteerOpportunityMembershipResponse, error) {
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
	// insert VolunteerOpportunityMembership entity data
	err = c.QueryRowContext(ctx, "INSERT INTO volunteeropportunitymembership (title) VALUES($1)  RETURNING id;",
		req.Item.Title).Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into VolunteerOpportunityMembership-> "+err.Error())
	}

	// get ID of creates VolunteerOpportunityMembership
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created VolunteerOpportunityMembership-> "+err.Error())
	}

	return &v1.CreateVolunteerOpportunityMembershipResponse{
		Api: apiVersion,
		Id:  id,
	}, nil
}

// Get volunteeropportunitymembership by id.
func (s *shrikeServiceServer) GetVolunteerOpportunityMembership(ctx context.Context, req *v1.GetVolunteerOpportunityMembershipRequest) (*v1.GetVolunteerOpportunityMembershipResponse, error) {
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

	// query VolunteerOpportunityMembership by ID
	rows, err := c.QueryContext(ctx, "SELECT id, title FROM volunteeropportunitymembership WHERE id=$1",
		req.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from VolunteerOpportunityMembership-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from VolunteerOpportunityMembership-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("VolunteerOpportunityMembership with ID='%d' is not found",
			req.Id))
	}

	// get VolunteerOpportunityMembership data
	var td v1.VolunteerOpportunityMembership
	if err := rows.Scan(&td.Id, &td.Title); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from VolunteerOpportunityMembership row-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple VolunteerOpportunityMembership rows with ID='%d'",
			req.Id))
	}

	return &v1.GetVolunteerOpportunityMembershipResponse{
		Api:  apiVersion,
		Item: &td,
	}, nil

}

// Read all todo tasks
func (s *shrikeServiceServer) ListVolunteerOpportunityMembership(ctx context.Context, req *v1.ListVolunteerOpportunityMembershipRequest) (*v1.ListVolunteerOpportunityMembershipResponse, error) {
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

	// get VolunteerOpportunityMembership list
	rows, err := c.QueryContext(ctx, "SELECT id,title FROM VolunteerOpportunityMembership")
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from VolunteerOpportunityMembership-> "+err.Error())
	}
	defer rows.Close()

	list := []*v1.VolunteerOpportunityMembership{}
	for rows.Next() {
		td := new(v1.VolunteerOpportunityMembership)
		if err := rows.Scan(&td.Id, &td.Title); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from VolunteerOpportunityMembership row-> "+err.Error())
		}
		list = append(list, td)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from VolunteerOpportunityMembership-> "+err.Error())
	}

	return &v1.ListVolunteerOpportunityMembershipResponse{
		Api:   apiVersion,
		Items: list,
	}, nil
}

// Update todo task
func (s *shrikeServiceServer) UpdateVolunteerOpportunityMembership(ctx context.Context, req *v1.UpdateVolunteerOpportunityMembershipRequest) (*v1.UpdateVolunteerOpportunityMembershipResponse, error) {
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

	// update volunteeropportunitymembership
	res, err := c.ExecContext(ctx, "UPDATE volunteeropportunitymembership SET title=$1 WHERE id=$2",
		req.Item.Title, req.Item.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update volunteeropportunitymembership-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("volunteeropportunitymembership with ID='%d' is not found",
			req.Item.Id))
	}

	return &v1.UpdateVolunteerOpportunityMembershipResponse{
		Api:     apiVersion,
		Updated: rows,
	}, nil
}

// Delete volunteeropportunitymembership
func (s *shrikeServiceServer) DeleteVolunteerOpportunityMembership(ctx context.Context, req *v1.DeleteVolunteerOpportunityMembershipRequest) (*v1.DeleteVolunteerOpportunityMembershipResponse, error) {
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

	// delete volunteeropportunitymembership
	res, err := c.ExecContext(ctx, "DELETE FROM volunteeropportunitymembership WHERE id=$1", req.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete volunteeropportunitymembership-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("volunteeropportunitymembership with ID='%d' is not found",
			req.Id))
	}

	return &v1.DeleteVolunteerOpportunityMembershipResponse{
		Api:     apiVersion,
		Deleted: rows,
	}, nil
}
