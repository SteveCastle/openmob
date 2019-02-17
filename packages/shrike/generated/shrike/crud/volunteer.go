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

// NewShrikeServiceServer creates Volunteer service
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

// Create new Volunteer
func (s *shrikeServiceServer) CreateVolunteer(ctx context.Context, req *v1.CreateVolunteerRequest) (*v1.CreateVolunteerResponse, error) {
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
	// insert Volunteer entity data
	err = c.QueryRowContext(ctx, "INSERT INTO volunteer (id, created_at, updated_at, volunteer_opportunity, contact, cause, ) VALUES($1, $2, $3, $4, $5, $6, )  RETURNING id;",
		 req.ItemID  req.ItemCreatedAt  req.ItemUpdatedAt  req.ItemVolunteerOpportunity  req.ItemContact  req.ItemCause ).Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into Volunteer-> "+err.Error())
	}

	// get ID of creates Volunteer
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created Volunteer-> "+err.Error())
	}

	return &v1.CreateVolunteerResponse{
		Api: apiVersion,
		Id:  id,
	}, nil
}

// Get volunteer by id.
func (s *shrikeServiceServer) GetVolunteer(ctx context.Context, req *v1.GetVolunteerRequest) (*v1.GetVolunteerResponse, error) {
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

	// query Volunteer by ID
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, volunteer_opportunity, contact, cause,  FROM volunteer WHERE id=$1",
		req.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from Volunteer-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from Volunteer-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Volunteer with ID='%d' is not found",
			req.Id))
	}

	// get Volunteer data
	var volunteer v1.Volunteer
	if err := rows.Scan( &volunteer.ID,  &volunteer.CreatedAt,  &volunteer.UpdatedAt,  &volunteer.VolunteerOpportunity,  &volunteer.Contact,  &volunteer.Cause, ); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from Volunteer row-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple Volunteer rows with ID='%d'",
			req.Id))
	}

	return &v1.GetVolunteerResponse{
		Api:  apiVersion,
		Item: &volunteer,
	}, nil

}

// Read all Volunteer
func (s *shrikeServiceServer) ListVolunteer(ctx context.Context, req *v1.ListVolunteerRequest) (*v1.ListVolunteerResponse, error) {
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

	// get Volunteer list
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, volunteer_opportunity, contact, cause,  FROM volunteer")
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from Volunteer-> "+err.Error())
	}
	defer rows.Close()

	list := []*v1.Volunteer{}
	for rows.Next() {
		volunteer := new(v1.Volunteer)
		if err := rows.Scan( &volunteer.ID,  &volunteer.CreatedAt,  &volunteer.UpdatedAt,  &volunteer.VolunteerOpportunity,  &volunteer.Contact,  &volunteer.Cause, ); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from Volunteer row-> "+err.Error())
		}
		list = append(list, volunteer)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from Volunteer-> "+err.Error())
	}

	return &v1.ListVolunteerResponse{
		Api:   apiVersion,
		Items: list,
	}, nil
}

// Update Volunteer
func (s *shrikeServiceServer) UpdateVolunteer(ctx context.Context, req *v1.UpdateVolunteerRequest) (*v1.UpdateVolunteerResponse, error) {
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

	// update volunteer
	res, err := c.ExecContext(ctx, "UPDATE volunteer SET title=$1 WHERE id=$2",
		req.Item.Title, req.Item.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update Volunteer-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Volunteer with ID='%d' is not found",
			req.Item.Id))
	}

	return &v1.UpdateVolunteerResponse{
		Api:     apiVersion,
		Updated: rows,
	}, nil
}

// Delete volunteer
func (s *shrikeServiceServer) DeleteVolunteer(ctx context.Context, req *v1.DeleteVolunteerRequest) (*v1.DeleteVolunteerResponse, error) {
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

	// delete volunteer
	res, err := c.ExecContext(ctx, "DELETE FROM volunteer WHERE id=$1", req.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete Volunteer-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Volunteer with ID='%d' is not found",
			req.Id))
	}

	return &v1.DeleteVolunteerResponse{
		Api:     apiVersion,
		Deleted: rows,
	}, nil
}
