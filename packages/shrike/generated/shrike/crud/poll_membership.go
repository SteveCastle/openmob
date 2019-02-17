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

// NewShrikeServiceServer creates PollMembership service
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

// Create new PollMembership
func (s *shrikeServiceServer) CreatePollMembership(ctx context.Context, req *v1.CreatePollMembershipRequest) (*v1.CreatePollMembershipResponse, error) {
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
	// insert PollMembership entity data
	err = c.QueryRowContext(ctx, "INSERT INTO poll_membership (id, created_at, updated_at, cause, petition, ) VALUES($1, $2, $3, $4, $5, )  RETURNING id;",
		 req.ItemID  req.ItemCreatedAt  req.ItemUpdatedAt  req.ItemCause  req.ItemPetition ).Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into PollMembership-> "+err.Error())
	}

	// get ID of creates PollMembership
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created PollMembership-> "+err.Error())
	}

	return &v1.CreatePollMembershipResponse{
		Api: apiVersion,
		Id:  id,
	}, nil
}

// Get poll_membership by id.
func (s *shrikeServiceServer) GetPollMembership(ctx context.Context, req *v1.GetPollMembershipRequest) (*v1.GetPollMembershipResponse, error) {
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

	// query PollMembership by ID
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, cause, petition,  FROM poll_membership WHERE id=$1",
		req.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from PollMembership-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from PollMembership-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("PollMembership with ID='%d' is not found",
			req.Id))
	}

	// get PollMembership data
	var pollmembership v1.PollMembership
	if err := rows.Scan( &pollmembership.ID,  &pollmembership.CreatedAt,  &pollmembership.UpdatedAt,  &pollmembership.Cause,  &pollmembership.Petition, ); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from PollMembership row-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple PollMembership rows with ID='%d'",
			req.Id))
	}

	return &v1.GetPollMembershipResponse{
		Api:  apiVersion,
		Item: &pollmembership,
	}, nil

}

// Read all PollMembership
func (s *shrikeServiceServer) ListPollMembership(ctx context.Context, req *v1.ListPollMembershipRequest) (*v1.ListPollMembershipResponse, error) {
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

	// get PollMembership list
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, cause, petition,  FROM poll_membership")
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from PollMembership-> "+err.Error())
	}
	defer rows.Close()

	list := []*v1.PollMembership{}
	for rows.Next() {
		pollmembership := new(v1.PollMembership)
		if err := rows.Scan( &pollmembership.ID,  &pollmembership.CreatedAt,  &pollmembership.UpdatedAt,  &pollmembership.Cause,  &pollmembership.Petition, ); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from PollMembership row-> "+err.Error())
		}
		list = append(list, pollmembership)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from PollMembership-> "+err.Error())
	}

	return &v1.ListPollMembershipResponse{
		Api:   apiVersion,
		Items: list,
	}, nil
}

// Update PollMembership
func (s *shrikeServiceServer) UpdatePollMembership(ctx context.Context, req *v1.UpdatePollMembershipRequest) (*v1.UpdatePollMembershipResponse, error) {
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

	// update poll_membership
	res, err := c.ExecContext(ctx, "UPDATE poll_membership SET title=$1 WHERE id=$2",
		req.Item.Title, req.Item.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update PollMembership-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("PollMembership with ID='%d' is not found",
			req.Item.Id))
	}

	return &v1.UpdatePollMembershipResponse{
		Api:     apiVersion,
		Updated: rows,
	}, nil
}

// Delete poll_membership
func (s *shrikeServiceServer) DeletePollMembership(ctx context.Context, req *v1.DeletePollMembershipRequest) (*v1.DeletePollMembershipResponse, error) {
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

	// delete poll_membership
	res, err := c.ExecContext(ctx, "DELETE FROM poll_membership WHERE id=$1", req.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete PollMembership-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("PollMembership with ID='%d' is not found",
			req.Id))
	}

	return &v1.DeletePollMembershipResponse{
		Api:     apiVersion,
		Deleted: rows,
	}, nil
}
