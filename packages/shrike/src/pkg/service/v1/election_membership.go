package v1

import (
	"context"
	"fmt"

	v1 "github.com/SteveCastle/openmob/packages/shrike/src/pkg/api/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

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
	err = c.QueryRowContext(ctx, "INSERT INTO election_membership (id, created_at, updated_at, cause, election) VALUES($1, $2, $3, $4, $5)  RETURNING id;",
		 req.Item.ID,  req.Item.CreatedAt,  req.Item.UpdatedAt,  req.Item.Cause,  req.Item.Election, ).Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into ElectionMembership-> "+err.Error())
	}

	// get ID of creates ElectionMembership
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created ElectionMembership-> "+err.Error())
	}

	return &v1.CreateElectionMembershipResponse{
		Api: apiVersion,
		ID:  id,
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
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, cause, election FROM election_membership WHERE id=$1",
		req.ID)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from ElectionMembership-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from ElectionMembership-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("ElectionMembership with ID='%d' is not found",
			req.ID))
	}

	// get ElectionMembership data
	var electionmembership v1.ElectionMembership
	if err := rows.Scan( &electionmembership.ID,  &electionmembership.CreatedAt,  &electionmembership.UpdatedAt,  &electionmembership.Cause,  &electionmembership.Election, ); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from ElectionMembership row-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple ElectionMembership rows with ID='%d'",
			req.ID))
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
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, cause, election FROM election_membership")
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
	res, err := c.ExecContext(ctx, "UPDATE election_membership SET id=$1, created_at=$2, updated_at=$3, cause=$4, election=$5 WHERE id=$1",
		req.Item.ID,req.Item.CreatedAt,req.Item.UpdatedAt,req.Item.Cause,req.Item.Election, )
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update ElectionMembership-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("ElectionMembership with ID='%d' is not found",
			req.Item.ID))
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
	res, err := c.ExecContext(ctx, "DELETE FROM election_membership WHERE id=$1", req.ID)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete ElectionMembership-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("ElectionMembership with ID='%d' is not found",
			req.ID))
	}

	return &v1.DeleteElectionMembershipResponse{
		Api:     apiVersion,
		Deleted: rows,
	}, nil
}
