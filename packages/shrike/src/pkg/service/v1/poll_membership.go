package v1

import (
	"context"
	"fmt"
	"time"

	v1 "github.com/SteveCastle/openmob/packages/shrike/src/pkg/api/v1"
	"github.com/SteveCastle/openmob/packages/shrike/src/pkg/queries"
	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

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
	var id string
	// insert PollMembership entity data
	err = c.QueryRowContext(ctx, "INSERT INTO poll_membership (cause, petition) VALUES($1, $2)  RETURNING id;",
		req.Item.Cause, req.Item.Petition).Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into PollMembership-> "+err.Error())
	}

	// get ID of creates PollMembership
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created PollMembership-> "+err.Error())
	}

	return &v1.CreatePollMembershipResponse{
		Api: apiVersion,
		ID:  id,
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
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, cause, petition FROM poll_membership WHERE id=$1",
		req.ID)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from PollMembership-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from PollMembership-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("PollMembership with ID='%s' is not found",
			req.ID))
	}

	// scan PollMembership data into protobuf model
	var pollmembership v1.PollMembership
	var createdAt time.Time
	var updatedAt time.Time

	if err := rows.Scan(&pollmembership.ID, &createdAt, &updatedAt, &pollmembership.Cause, &pollmembership.Petition); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from PollMembership row-> "+err.Error())
	}

	// Convert time.Time from database into proto timestamp.
	pollmembership.CreatedAt, err = ptypes.TimestampProto(createdAt)
	if err != nil {
		return nil, status.Error(codes.Unknown, "createdAt field has invalid format-> "+err.Error())
	}
	pollmembership.UpdatedAt, err = ptypes.TimestampProto(updatedAt)
	if err != nil {
		return nil, status.Error(codes.Unknown, "updatedAt field has invalid format-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple PollMembership rows with ID='%s'",
			req.ID))
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

	// Generate SQL to select all columns in PollMembership Table
	// Then generate filtering and ordering sql and finally run query.

	baseSQL := "SELECT id, created_at, updated_at, cause, petition FROM poll_membership"
	querySQL := queries.BuildPollMembershipFilters(req.Filters, req.Ordering, req.Limit)
	SQL := fmt.Sprintf("%s %s", baseSQL, querySQL)
	rows, err := c.QueryContext(ctx, SQL)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from PollMembership-> "+err.Error())
	}
	defer rows.Close()

	// Variables to store results returned by database.
	list := []*v1.PollMembership{}
	var createdAt time.Time
	var updatedAt time.Time

	for rows.Next() {
		pollmembership := new(v1.PollMembership)
		if err := rows.Scan(&pollmembership.ID, &createdAt, &updatedAt, &pollmembership.Cause, &pollmembership.Petition); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from PollMembership row-> "+err.Error())
		}
		// Convert time.Time from database into proto timestamp.
		pollmembership.CreatedAt, err = ptypes.TimestampProto(createdAt)
		if err != nil {
			return nil, status.Error(codes.Unknown, "createdAt field has invalid format-> "+err.Error())
		}
		pollmembership.UpdatedAt, err = ptypes.TimestampProto(updatedAt)
		if err != nil {
			return nil, status.Error(codes.Unknown, "updatedAt field has invalid format-> "+err.Error())
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
	res, err := c.ExecContext(ctx, "UPDATE poll_membership SET cause=$2, petition=$3 WHERE id=$1",
		req.Item.ID, req.Item.Cause, req.Item.Petition)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update PollMembership-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("PollMembership with ID='%s' is not found",
			req.Item.ID))
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
	res, err := c.ExecContext(ctx, "DELETE FROM poll_membership WHERE id=$1", req.ID)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete PollMembership-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("PollMembership with ID='%s' is not found",
			req.ID))
	}

	return &v1.DeletePollMembershipResponse{
		Api:     apiVersion,
		Deleted: rows,
	}, nil
}
