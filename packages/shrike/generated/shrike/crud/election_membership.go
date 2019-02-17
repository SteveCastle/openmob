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

// NewShrikeServiceServer creates ElectionMembership service
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

// Create new ElectionMembership
func (s *shrikeServiceServer) CreateElectionMembership(ctx context.Context, req *v1.CreateElectionMembershipRequest) (*v1.CreateElectionMembershipResponse, error) {
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
	// insert ElectionMembership entity data
	err = c.QueryRowContext(ctx, "INSERT INTO election_membership (id, created_at, updated_at, cause, election, ) VALUES($1, $2, $3, $4, $5, )  RETURNING id;",
		 req.ItemID  req.ItemCreatedAt  req.ItemUpdatedAt  req.ItemCause  req.ItemElection ).Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into ElectionMembership-> "+err.Error())
	}

	// get ID of creates ElectionMembership
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created ElectionMembership-> "+err.Error())
	}

	return &v1.CreateElectionMembershipResponse{
		Api: apiVersion,
		Id:  id,
	}, nil
}

// Get election_membership by id.
func (s *shrikeServiceServer) GetElectionMembership(ctx context.Context, req *v1.GetElectionMembershipRequest) (*v1.GetElectionMembershipResponse, error) {
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

	// query ElectionMembership by ID
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, cause, election,  FROM election_membership WHERE id=$1",
		req.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from ElectionMembership-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from ElectionMembership-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("ElectionMembership with ID='%d' is not found",
			req.Id))
	}

	// get ElectionMembership data
	var electionmembership v1.ElectionMembership
	if err := rows.Scan( &electionmembership.ID,  &electionmembership.CreatedAt,  &electionmembership.UpdatedAt,  &electionmembership.Cause,  &electionmembership.Election, ); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from ElectionMembership row-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple ElectionMembership rows with ID='%d'",
			req.Id))
	}

	return &v1.GetElectionMembershipResponse{
		Api:  apiVersion,
		Item: &electionmembership,
	}, nil

}

// Read all ElectionMembership
func (s *shrikeServiceServer) ListElectionMembership(ctx context.Context, req *v1.ListElectionMembershipRequest) (*v1.ListElectionMembershipResponse, error) {
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

	// get ElectionMembership list
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, cause, election,  FROM election_membership")
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from ElectionMembership-> "+err.Error())
	}
	defer rows.Close()

	list := []*v1.ElectionMembership{}
	for rows.Next() {
		electionmembership := new(v1.ElectionMembership)
		if err := rows.Scan( &electionmembership.ID,  &electionmembership.CreatedAt,  &electionmembership.UpdatedAt,  &electionmembership.Cause,  &electionmembership.Election, ); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from ElectionMembership row-> "+err.Error())
		}
		list = append(list, electionmembership)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from ElectionMembership-> "+err.Error())
	}

	return &v1.ListElectionMembershipResponse{
		Api:   apiVersion,
		Items: list,
	}, nil
}

// Update ElectionMembership
func (s *shrikeServiceServer) UpdateElectionMembership(ctx context.Context, req *v1.UpdateElectionMembershipRequest) (*v1.UpdateElectionMembershipResponse, error) {
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

	// update election_membership
	res, err := c.ExecContext(ctx, "UPDATE election_membership SET title=$1 WHERE id=$2",
		req.Item.Title, req.Item.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update ElectionMembership-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("ElectionMembership with ID='%d' is not found",
			req.Item.Id))
	}

	return &v1.UpdateElectionMembershipResponse{
		Api:     apiVersion,
		Updated: rows,
	}, nil
}

// Delete election_membership
func (s *shrikeServiceServer) DeleteElectionMembership(ctx context.Context, req *v1.DeleteElectionMembershipRequest) (*v1.DeleteElectionMembershipResponse, error) {
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

	// delete election_membership
	res, err := c.ExecContext(ctx, "DELETE FROM election_membership WHERE id=$1", req.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete ElectionMembership-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("ElectionMembership with ID='%d' is not found",
			req.Id))
	}

	return &v1.DeleteElectionMembershipResponse{
		Api:     apiVersion,
		Deleted: rows,
	}, nil
}
