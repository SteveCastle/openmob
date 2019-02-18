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

// Create new Contact
func (s *shrikeServiceServer) CreateContact(ctx context.Context, req *v1.CreateContactRequest) (*v1.CreateContactResponse, error) {
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
	// insert Contact entity data
	err = c.QueryRowContext(ctx, "INSERT INTO contact () VALUES()  RETURNING id;").Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into Contact-> "+err.Error())
	}

	// get ID of creates Contact
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created Contact-> "+err.Error())
	}

	return &v1.CreateContactResponse{
		Api: apiVersion,
		ID:  id,
	}, nil
}

// Get contact by id.
func (s *shrikeServiceServer) GetContact(ctx context.Context, req *v1.GetContactRequest) (*v1.GetContactResponse, error) {
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

	// query Contact by ID
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at FROM contact WHERE id=$1",
		req.ID)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from Contact-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from Contact-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Contact with ID='%d' is not found",
			req.ID))
	}

	// scan Contact data into protobuf model
	var contact v1.Contact
	var createdAt time.Time
	var updatedAt time.Time

	if err := rows.Scan(&contact.ID, &createdAt, &updatedAt); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from Contact row-> "+err.Error())
	}

	// Convert time.Time from database into proto timestamp.
	contact.CreatedAt, err = ptypes.TimestampProto(createdAt)
	if err != nil {
		return nil, status.Error(codes.Unknown, "createdAt field has invalid format-> "+err.Error())
	}
	contact.UpdatedAt, err = ptypes.TimestampProto(updatedAt)
	if err != nil {
		return nil, status.Error(codes.Unknown, "updatedAt field has invalid format-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple Contact rows with ID='%d'",
			req.ID))
	}

	return &v1.GetContactResponse{
		Api:  apiVersion,
		Item: &contact,
	}, nil

}

// Read all Contact
func (s *shrikeServiceServer) ListContact(ctx context.Context, req *v1.ListContactRequest) (*v1.ListContactResponse, error) {
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

	// get Contact list
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at FROM contact")
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from Contact-> "+err.Error())
	}
	defer rows.Close()

	// Variables to store results returned by database.
	list := []*v1.Contact{}
	var createdAt time.Time
	var updatedAt time.Time

	for rows.Next() {
		contact := new(v1.Contact)
		if err := rows.Scan(&contact.ID, &createdAt, &updatedAt); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from Contact row-> "+err.Error())
		}
		// Convert time.Time from database into proto timestamp.
		contact.CreatedAt, err = ptypes.TimestampProto(createdAt)
		if err != nil {
			return nil, status.Error(codes.Unknown, "createdAt field has invalid format-> "+err.Error())
		}
		contact.UpdatedAt, err = ptypes.TimestampProto(updatedAt)
		if err != nil {
			return nil, status.Error(codes.Unknown, "updatedAt field has invalid format-> "+err.Error())
		}

		list = append(list, contact)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from Contact-> "+err.Error())
	}

	return &v1.ListContactResponse{
		Api:   apiVersion,
		Items: list,
	}, nil
}

// Update Contact
func (s *shrikeServiceServer) UpdateContact(ctx context.Context, req *v1.UpdateContactRequest) (*v1.UpdateContactResponse, error) {
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

	// update contact
	res, err := c.ExecContext(ctx, "UPDATE contact SET  WHERE id=$1",
		req.Item.ID)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update Contact-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Contact with ID='%d' is not found",
			req.Item.ID))
	}

	return &v1.UpdateContactResponse{
		Api:     apiVersion,
		Updated: rows,
	}, nil
}

// Delete contact
func (s *shrikeServiceServer) DeleteContact(ctx context.Context, req *v1.DeleteContactRequest) (*v1.DeleteContactResponse, error) {
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

	// delete contact
	res, err := c.ExecContext(ctx, "DELETE FROM contact WHERE id=$1", req.ID)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete Contact-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Contact with ID='%d' is not found",
			req.ID))
	}

	return &v1.DeleteContactResponse{
		Api:     apiVersion,
		Deleted: rows,
	}, nil
}
