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

// Create new PhoneNumber
func (s *shrikeServiceServer) CreatePhoneNumber(ctx context.Context, req *v1.CreatePhoneNumberRequest) (*v1.CreatePhoneNumberResponse, error) {
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
	// insert PhoneNumber entity data
	err = c.QueryRowContext(ctx, "INSERT INTO phone_number (phone_number) VALUES($1)  RETURNING id;",
		req.Item.PhoneNumber).Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into PhoneNumber-> "+err.Error())
	}

	// get ID of creates PhoneNumber
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created PhoneNumber-> "+err.Error())
	}

	return &v1.CreatePhoneNumberResponse{
		Api: apiVersion,
		ID:  id,
	}, nil
}

// Get phone_number by id.
func (s *shrikeServiceServer) GetPhoneNumber(ctx context.Context, req *v1.GetPhoneNumberRequest) (*v1.GetPhoneNumberResponse, error) {
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

	// query PhoneNumber by ID
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, phone_number FROM phone_number WHERE id=$1",
		req.ID)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from PhoneNumber-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from PhoneNumber-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("PhoneNumber with ID='%s' is not found",
			req.ID))
	}

	// scan PhoneNumber data into protobuf model
	var phonenumber v1.PhoneNumber
	var createdAt time.Time
	var updatedAt time.Time

	if err := rows.Scan(&phonenumber.ID, &createdAt, &updatedAt, &phonenumber.PhoneNumber); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from PhoneNumber row-> "+err.Error())
	}

	// Convert time.Time from database into proto timestamp.
	phonenumber.CreatedAt, err = ptypes.TimestampProto(createdAt)
	if err != nil {
		return nil, status.Error(codes.Unknown, "createdAt field has invalid format-> "+err.Error())
	}
	phonenumber.UpdatedAt, err = ptypes.TimestampProto(updatedAt)
	if err != nil {
		return nil, status.Error(codes.Unknown, "updatedAt field has invalid format-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple PhoneNumber rows with ID='%s'",
			req.ID))
	}

	return &v1.GetPhoneNumberResponse{
		Api:  apiVersion,
		Item: &phonenumber,
	}, nil

}

// Read all PhoneNumber
func (s *shrikeServiceServer) ListPhoneNumber(ctx context.Context, req *v1.ListPhoneNumberRequest) (*v1.ListPhoneNumberResponse, error) {
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

	// get PhoneNumber list
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, phone_number FROM phone_number")
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from PhoneNumber-> "+err.Error())
	}
	defer rows.Close()

	// Variables to store results returned by database.
	list := []*v1.PhoneNumber{}
	var createdAt time.Time
	var updatedAt time.Time

	for rows.Next() {
		phonenumber := new(v1.PhoneNumber)
		if err := rows.Scan(&phonenumber.ID, &createdAt, &updatedAt, &phonenumber.PhoneNumber); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from PhoneNumber row-> "+err.Error())
		}
		// Convert time.Time from database into proto timestamp.
		phonenumber.CreatedAt, err = ptypes.TimestampProto(createdAt)
		if err != nil {
			return nil, status.Error(codes.Unknown, "createdAt field has invalid format-> "+err.Error())
		}
		phonenumber.UpdatedAt, err = ptypes.TimestampProto(updatedAt)
		if err != nil {
			return nil, status.Error(codes.Unknown, "updatedAt field has invalid format-> "+err.Error())
		}

		list = append(list, phonenumber)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from PhoneNumber-> "+err.Error())
	}

	return &v1.ListPhoneNumberResponse{
		Api:   apiVersion,
		Items: list,
	}, nil
}

// Update PhoneNumber
func (s *shrikeServiceServer) UpdatePhoneNumber(ctx context.Context, req *v1.UpdatePhoneNumberRequest) (*v1.UpdatePhoneNumberResponse, error) {
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

	// update phone_number
	res, err := c.ExecContext(ctx, "UPDATE phone_number SET phone_number=$2 WHERE id=$1",
		req.Item.ID, req.Item.PhoneNumber)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update PhoneNumber-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("PhoneNumber with ID='%s' is not found",
			req.Item.ID))
	}

	return &v1.UpdatePhoneNumberResponse{
		Api:     apiVersion,
		Updated: rows,
	}, nil
}

// Delete phone_number
func (s *shrikeServiceServer) DeletePhoneNumber(ctx context.Context, req *v1.DeletePhoneNumberRequest) (*v1.DeletePhoneNumberResponse, error) {
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

	// delete phone_number
	res, err := c.ExecContext(ctx, "DELETE FROM phone_number WHERE id=$1", req.ID)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete PhoneNumber-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("PhoneNumber with ID='%s' is not found",
			req.ID))
	}

	return &v1.DeletePhoneNumberResponse{
		Api:     apiVersion,
		Deleted: rows,
	}, nil
}
