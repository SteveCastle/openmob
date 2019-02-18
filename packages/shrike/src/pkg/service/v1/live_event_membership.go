package v1

import (
	"context"
	"fmt"

	v1 "github.com/SteveCastle/openmob/packages/shrike/src/pkg/api/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Create new LiveEventMembership
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
	err = c.QueryRowContext(ctx, "INSERT INTO live_event_membership (id, created_at, updated_at, cause, live_event) VALUES($1, $2, $3, $4, $5)  RETURNING id;",
		 req.Item.ID,  req.Item.CreatedAt,  req.Item.UpdatedAt,  req.Item.Cause,  req.Item.LiveEvent, ).Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into LiveEventMembership-> "+err.Error())
	}

	// get ID of creates LiveEventMembership
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created LiveEventMembership-> "+err.Error())
	}

	return &v1.CreateLiveEventMembershipResponse{
		Api: apiVersion,
		ID:  id,
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
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, cause, live_event FROM live_event_membership WHERE id=$1",
		req.ID)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from LiveEventMembership-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from LiveEventMembership-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("LiveEventMembership with ID='%d' is not found",
			req.ID))
	}

	// get LiveEventMembership data
	var liveeventmembership v1.LiveEventMembership
	if err := rows.Scan( &liveeventmembership.ID,  &liveeventmembership.CreatedAt,  &liveeventmembership.UpdatedAt,  &liveeventmembership.Cause,  &liveeventmembership.LiveEvent, ); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from LiveEventMembership row-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple LiveEventMembership rows with ID='%d'",
			req.ID))
	}

	return &v1.GetLiveEventMembershipResponse{
		Api:  apiVersion,
		Item: &liveeventmembership,
	}, nil

}

// Read all LiveEventMembership
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
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, cause, live_event FROM live_event_membership")
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from LiveEventMembership-> "+err.Error())
	}
	defer rows.Close()

	list := []*v1.LiveEventMembership{}
	for rows.Next() {
		liveeventmembership := new(v1.LiveEventMembership)
		if err := rows.Scan( &liveeventmembership.ID,  &liveeventmembership.CreatedAt,  &liveeventmembership.UpdatedAt,  &liveeventmembership.Cause,  &liveeventmembership.LiveEvent, ); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from LiveEventMembership row-> "+err.Error())
		}
		list = append(list, liveeventmembership)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from LiveEventMembership-> "+err.Error())
	}

	return &v1.ListLiveEventMembershipResponse{
		Api:   apiVersion,
		Items: list,
	}, nil
}

// Update LiveEventMembership
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
	res, err := c.ExecContext(ctx, "UPDATE live_event_membership SET $1 ,$2 ,$3 ,$4 ,$5  WHERE id=$1",
		req.Item.ID,req.Item.CreatedAt,req.Item.UpdatedAt,req.Item.Cause,req.Item.LiveEvent, )
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update LiveEventMembership-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("LiveEventMembership with ID='%d' is not found",
			req.Item.ID))
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
	res, err := c.ExecContext(ctx, "DELETE FROM live_event_membership WHERE id=$1", req.ID)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete LiveEventMembership-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("LiveEventMembership with ID='%d' is not found",
			req.ID))
	}

	return &v1.DeleteLiveEventMembershipResponse{
		Api:     apiVersion,
		Deleted: rows,
	}, nil
}
