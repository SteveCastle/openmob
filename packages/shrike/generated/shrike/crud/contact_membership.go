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

// NewShrikeServiceServer creates ContactMembership service
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
	var id int64
	// insert ContactMembership entity data
	err = c.QueryRowContext(ctx, "INSERT INTO contact_membership (id, created_at, updated_at, cause, contact, ) VALUES($1, $2, $3, $4, $5, )  RETURNING id;",
		 req.ItemID  req.ItemCreatedAt  req.ItemUpdatedAt  req.ItemCause  req.ItemContact ).Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into ContactMembership-> "+err.Error())
	}

	// get ID of creates ContactMembership
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created ContactMembership-> "+err.Error())
	}

	return &v1.CreateContactMembershipResponse{
		Api: apiVersion,
		Id:  id,
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
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, cause, contact,  FROM contact_membership WHERE id=$1",
		req.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from ContactMembership-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from ContactMembership-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("ContactMembership with ID='%d' is not found",
			req.Id))
	}

	// get ContactMembership data
	var contactmembership v1.ContactMembership
	if err := rows.Scan( &contactmembership.ID,  &contactmembership.CreatedAt,  &contactmembership.UpdatedAt,  &contactmembership.Cause,  &contactmembership.Contact, ); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from ContactMembership row-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple ContactMembership rows with ID='%d'",
			req.Id))
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
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, cause, contact,  FROM contact_membership")
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from ContactMembership-> "+err.Error())
	}
	defer rows.Close()

	list := []*v1.ContactMembership{}
	for rows.Next() {
		contactmembership := new(v1.ContactMembership)
		if err := rows.Scan( &contactmembership.ID,  &contactmembership.CreatedAt,  &contactmembership.UpdatedAt,  &contactmembership.Cause,  &contactmembership.Contact, ); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from ContactMembership row-> "+err.Error())
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
	res, err := c.ExecContext(ctx, "UPDATE contact_membership SET title=$1 WHERE id=$2",
		req.Item.Title, req.Item.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update ContactMembership-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("ContactMembership with ID='%d' is not found",
			req.Item.Id))
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
	res, err := c.ExecContext(ctx, "DELETE FROM contact_membership WHERE id=$1", req.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete ContactMembership-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("ContactMembership with ID='%d' is not found",
			req.Id))
	}

	return &v1.DeleteContactMembershipResponse{
		Api:     apiVersion,
		Deleted: rows,
	}, nil
}
