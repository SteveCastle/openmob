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

// Create new ContactMembership
func (s *shrikeServiceServer) CreateContactMembership(ctx context.Context, req *v1.CreateContactMembershipRequest) (*v1.CreateContactMembershipResponse, error) {
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
	// insert ContactMembership entity data
	err = c.QueryRowContext(ctx, "INSERT INTO contact_membership (cause, contact) VALUES($1, $2)  RETURNING id;",
		req.Item.Cause, req.Item.Contact).Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into ContactMembership-> "+err.Error())
	}

	// get ID of creates ContactMembership
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created ContactMembership-> "+err.Error())
	}

	return &v1.CreateContactMembershipResponse{
		Api: apiVersion,
		ID:  id,
	}, nil
}

// Get contact_membership by id.
func (s *shrikeServiceServer) GetContactMembership(ctx context.Context, req *v1.GetContactMembershipRequest) (*v1.GetContactMembershipResponse, error) {
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

	// query ContactMembership by ID
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, cause, contact FROM contact_membership WHERE id=$1",
		req.ID)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from ContactMembership-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from ContactMembership-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("ContactMembership with ID='%s' is not found",
			req.ID))
	}

	// scan ContactMembership data into protobuf model
	var contactmembership v1.ContactMembership
	var createdAt time.Time
	var updatedAt time.Time

	if err := rows.Scan(&contactmembership.ID, &createdAt, &updatedAt, &contactmembership.Cause, &contactmembership.Contact); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from ContactMembership row-> "+err.Error())
	}

	// Convert time.Time from database into proto timestamp.
	contactmembership.CreatedAt, err = ptypes.TimestampProto(createdAt)
	if err != nil {
		return nil, status.Error(codes.Unknown, "createdAt field has invalid format-> "+err.Error())
	}
	contactmembership.UpdatedAt, err = ptypes.TimestampProto(updatedAt)
	if err != nil {
		return nil, status.Error(codes.Unknown, "updatedAt field has invalid format-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple ContactMembership rows with ID='%s'",
			req.ID))
	}

	return &v1.GetContactMembershipResponse{
		Api:  apiVersion,
		Item: &contactmembership,
	}, nil

}

// Read all ContactMembership
func (s *shrikeServiceServer) ListContactMembership(ctx context.Context, req *v1.ListContactMembershipRequest) (*v1.ListContactMembershipResponse, error) {
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

	// get ContactMembership list
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, cause, contact FROM contact_membership")
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from ContactMembership-> "+err.Error())
	}
	defer rows.Close()

	// Variables to store results returned by database.
	list := []*v1.ContactMembership{}
	var createdAt time.Time
	var updatedAt time.Time

	for rows.Next() {
		contactmembership := new(v1.ContactMembership)
		if err := rows.Scan(&contactmembership.ID, &createdAt, &updatedAt, &contactmembership.Cause, &contactmembership.Contact); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from ContactMembership row-> "+err.Error())
		}
		// Convert time.Time from database into proto timestamp.
		contactmembership.CreatedAt, err = ptypes.TimestampProto(createdAt)
		if err != nil {
			return nil, status.Error(codes.Unknown, "createdAt field has invalid format-> "+err.Error())
		}
		contactmembership.UpdatedAt, err = ptypes.TimestampProto(updatedAt)
		if err != nil {
			return nil, status.Error(codes.Unknown, "updatedAt field has invalid format-> "+err.Error())
		}

		list = append(list, contactmembership)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from ContactMembership-> "+err.Error())
	}

	return &v1.ListContactMembershipResponse{
		Api:   apiVersion,
		Items: list,
	}, nil
}

// Update ContactMembership
func (s *shrikeServiceServer) UpdateContactMembership(ctx context.Context, req *v1.UpdateContactMembershipRequest) (*v1.UpdateContactMembershipResponse, error) {
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

	// update contact_membership
	res, err := c.ExecContext(ctx, "UPDATE contact_membership SET cause=$2, contact=$3 WHERE id=$1",
		req.Item.ID, req.Item.Cause, req.Item.Contact)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update ContactMembership-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("ContactMembership with ID='%s' is not found",
			req.Item.ID))
	}

	return &v1.UpdateContactMembershipResponse{
		Api:     apiVersion,
		Updated: rows,
	}, nil
}

// Delete contact_membership
func (s *shrikeServiceServer) DeleteContactMembership(ctx context.Context, req *v1.DeleteContactMembershipRequest) (*v1.DeleteContactMembershipResponse, error) {
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

	// delete contact_membership
	res, err := c.ExecContext(ctx, "DELETE FROM contact_membership WHERE id=$1", req.ID)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete ContactMembership-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("ContactMembership with ID='%s' is not found",
			req.ID))
	}

	return &v1.DeleteContactMembershipResponse{
		Api:     apiVersion,
		Deleted: rows,
	}, nil
}
