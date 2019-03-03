package v1

import (
	"context"
	"fmt"
	"time"

	v1 "github.com/SteveCastle/openmob/packages/shrike/src/pkg/api/v1"
	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Create new OwnerMembership
func (s *shrikeServiceServer) CreateOwnerMembership(ctx context.Context, req *v1.CreateOwnerMembershipRequest) (*v1.CreateOwnerMembershipResponse, error) {
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
	// insert OwnerMembership entity data
	err = c.QueryRowContext(ctx, "INSERT INTO owner_membership (cause, account) VALUES($1, $2)  RETURNING id;",
		req.Item.Cause, req.Item.Account).Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into OwnerMembership-> "+err.Error())
	}

	// get ID of creates OwnerMembership
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created OwnerMembership-> "+err.Error())
	}

	return &v1.CreateOwnerMembershipResponse{
		Api: apiVersion,
		ID:  id,
	}, nil
}

// Get owner_membership by id.
func (s *shrikeServiceServer) GetOwnerMembership(ctx context.Context, req *v1.GetOwnerMembershipRequest) (*v1.GetOwnerMembershipResponse, error) {
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

	// query OwnerMembership by ID
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, cause, account FROM owner_membership WHERE id=$1",
		req.ID)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from OwnerMembership-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from OwnerMembership-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("OwnerMembership with ID='%d' is not found",
			req.ID))
	}

	// scan OwnerMembership data into protobuf model
	var ownermembership v1.OwnerMembership
	var createdAt time.Time
	var updatedAt time.Time

	if err := rows.Scan(&ownermembership.ID, &createdAt, &updatedAt, &ownermembership.Cause, &ownermembership.Account); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from OwnerMembership row-> "+err.Error())
	}

	// Convert time.Time from database into proto timestamp.
	ownermembership.CreatedAt, err = ptypes.TimestampProto(createdAt)
	if err != nil {
		return nil, status.Error(codes.Unknown, "createdAt field has invalid format-> "+err.Error())
	}
	ownermembership.UpdatedAt, err = ptypes.TimestampProto(updatedAt)
	if err != nil {
		return nil, status.Error(codes.Unknown, "updatedAt field has invalid format-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple OwnerMembership rows with ID='%d'",
			req.ID))
	}

	return &v1.GetOwnerMembershipResponse{
		Api:  apiVersion,
		Item: &ownermembership,
	}, nil

}

// Read all OwnerMembership
func (s *shrikeServiceServer) ListOwnerMembership(ctx context.Context, req *v1.ListOwnerMembershipRequest) (*v1.ListOwnerMembershipResponse, error) {
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

	// get OwnerMembership list
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, cause, account FROM owner_membership")
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from OwnerMembership-> "+err.Error())
	}
	defer rows.Close()

	// Variables to store results returned by database.
	list := []*v1.OwnerMembership{}
	var createdAt time.Time
	var updatedAt time.Time

	for rows.Next() {
		ownermembership := new(v1.OwnerMembership)
		if err := rows.Scan(&ownermembership.ID, &createdAt, &updatedAt, &ownermembership.Cause, &ownermembership.Account); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from OwnerMembership row-> "+err.Error())
		}
		// Convert time.Time from database into proto timestamp.
		ownermembership.CreatedAt, err = ptypes.TimestampProto(createdAt)
		if err != nil {
			return nil, status.Error(codes.Unknown, "createdAt field has invalid format-> "+err.Error())
		}
		ownermembership.UpdatedAt, err = ptypes.TimestampProto(updatedAt)
		if err != nil {
			return nil, status.Error(codes.Unknown, "updatedAt field has invalid format-> "+err.Error())
		}

		list = append(list, ownermembership)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from OwnerMembership-> "+err.Error())
	}

	return &v1.ListOwnerMembershipResponse{
		Api:   apiVersion,
		Items: list,
	}, nil
}

// Update OwnerMembership
func (s *shrikeServiceServer) UpdateOwnerMembership(ctx context.Context, req *v1.UpdateOwnerMembershipRequest) (*v1.UpdateOwnerMembershipResponse, error) {
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

	// update owner_membership
	res, err := c.ExecContext(ctx, "UPDATE owner_membership SET cause=$2, account=$3 WHERE id=$1",
		req.Item.ID, req.Item.Cause, req.Item.Account)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update OwnerMembership-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("OwnerMembership with ID='%d' is not found",
			req.Item.ID))
	}

	return &v1.UpdateOwnerMembershipResponse{
		Api:     apiVersion,
		Updated: rows,
	}, nil
}

// Delete owner_membership
func (s *shrikeServiceServer) DeleteOwnerMembership(ctx context.Context, req *v1.DeleteOwnerMembershipRequest) (*v1.DeleteOwnerMembershipResponse, error) {
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

	// delete owner_membership
	res, err := c.ExecContext(ctx, "DELETE FROM owner_membership WHERE id=$1", req.ID)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete OwnerMembership-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("OwnerMembership with ID='%d' is not found",
			req.ID))
	}

	return &v1.DeleteOwnerMembershipResponse{
		Api:     apiVersion,
		Deleted: rows,
	}, nil
}
