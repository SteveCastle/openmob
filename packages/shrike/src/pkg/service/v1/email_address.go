package v1

import (
	"context"
	"fmt"

	v1 "github.com/SteveCastle/openmob/packages/shrike/src/pkg/api/v1"
	"github.com/SteveCastle/openmob/packages/shrike/src/pkg/queries"
	"github.com/golang/protobuf/ptypes"
	"github.com/lib/pq"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Create new EmailAddress
func (s *shrikeServiceServer) CreateEmailAddress(ctx context.Context, req *v1.CreateEmailAddressRequest) (*v1.CreateEmailAddressResponse, error) {
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
	// insert EmailAddress entity data
	err = c.QueryRowContext(ctx, "INSERT INTO email_address (address) VALUES($1)  RETURNING id;",
		req.Item.Address).Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into EmailAddress-> "+err.Error())
	}

	// get ID of creates EmailAddress
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created EmailAddress-> "+err.Error())
	}

	return &v1.CreateEmailAddressResponse{
		Api: apiVersion,
		ID:  id,
	}, nil
}

// Get email_address by id.
func (s *shrikeServiceServer) GetEmailAddress(ctx context.Context, req *v1.GetEmailAddressRequest) (*v1.GetEmailAddressResponse, error) {
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

	// query EmailAddress by ID
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, address FROM email_address WHERE id=$1",
		req.ID)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from EmailAddress-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from EmailAddress-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("EmailAddress with ID='%s' is not found",
			req.ID))
	}

	// scan EmailAddress data into protobuf model
	var emailaddress v1.EmailAddress
	var createdAt pq.NullTime
	var updatedAt pq.NullTime

	if err := rows.Scan(&emailaddress.ID, &createdAt, &updatedAt, &emailaddress.Address); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from EmailAddress row-> "+err.Error())
	}

	// Convert pq.NullTime from database into proto timestamp.
	if createdAt.Valid {
		emailaddress.CreatedAt, err = ptypes.TimestampProto(createdAt.Time)
		if err != nil {
			return nil, status.Error(codes.Unknown, "createdAt field has invalid format-> "+err.Error())
		}
	}
	if updatedAt.Valid {
		emailaddress.UpdatedAt, err = ptypes.TimestampProto(updatedAt.Time)
		if err != nil {
			return nil, status.Error(codes.Unknown, "createdAt field has invalid format-> "+err.Error())
		}
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple EmailAddress rows with ID='%s'",
			req.ID))
	}

	return &v1.GetEmailAddressResponse{
		Api:  apiVersion,
		Item: &emailaddress,
	}, nil

}

// Read all EmailAddress
func (s *shrikeServiceServer) ListEmailAddress(ctx context.Context, req *v1.ListEmailAddressRequest) (*v1.ListEmailAddressResponse, error) {
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

	// Generate SQL to select all columns in EmailAddress Table
	// Then generate filtering and ordering sql and finally run query.
	querySQL := queries.BuildEmailAddressListQuery(req.Filters, req.Ordering, req.Limit)
	// Execute query and scan into return type.
	rows, err := c.QueryContext(ctx, querySQL)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from EmailAddress-> "+err.Error())
	}
	defer rows.Close()

	// Variables to store results returned by database.
	list := []*v1.EmailAddress{}
	var createdAt pq.NullTime
	var updatedAt pq.NullTime

	for rows.Next() {
		emailaddress := new(v1.EmailAddress)
		if err := rows.Scan(&emailaddress.ID, &createdAt, &updatedAt, &emailaddress.Address); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from EmailAddress row-> "+err.Error())
		}
		// Convert pq.NullTime from database into proto timestamp.
		if createdAt.Valid {
			emailaddress.CreatedAt, err = ptypes.TimestampProto(createdAt.Time)
			if err != nil {
				return nil, status.Error(codes.Unknown, "createdAt field has invalid format-> "+err.Error())
			}
		}
		if updatedAt.Valid {
			emailaddress.UpdatedAt, err = ptypes.TimestampProto(updatedAt.Time)
			if err != nil {
				return nil, status.Error(codes.Unknown, "createdAt field has invalid format-> "+err.Error())
			}
		}

		list = append(list, emailaddress)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from EmailAddress-> "+err.Error())
	}

	return &v1.ListEmailAddressResponse{
		Api:   apiVersion,
		Items: list,
	}, nil
}

// Update EmailAddress
func (s *shrikeServiceServer) UpdateEmailAddress(ctx context.Context, req *v1.UpdateEmailAddressRequest) (*v1.UpdateEmailAddressResponse, error) {
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

	// update email_address
	res, err := c.ExecContext(ctx, "UPDATE email_address SET address=$2 WHERE id=$1",
		req.Item.ID, req.Item.Address)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update EmailAddress-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("EmailAddress with ID='%s' is not found",
			req.Item.ID))
	}

	return &v1.UpdateEmailAddressResponse{
		Api:     apiVersion,
		Updated: rows,
	}, nil
}

// Delete email_address
func (s *shrikeServiceServer) DeleteEmailAddress(ctx context.Context, req *v1.DeleteEmailAddressRequest) (*v1.DeleteEmailAddressResponse, error) {
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

	// delete email_address
	res, err := c.ExecContext(ctx, "DELETE FROM email_address WHERE id=$1", req.ID)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete EmailAddress-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("EmailAddress with ID='%s' is not found",
			req.ID))
	}

	return &v1.DeleteEmailAddressResponse{
		Api:     apiVersion,
		Deleted: rows,
	}, nil
}
