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

// Create new Cause
func (s *shrikeServiceServer) CreateCause(ctx context.Context, req *v1.CreateCauseRequest) (*v1.CreateCauseResponse, error) {
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
	// insert Cause entity data
	err = c.QueryRowContext(ctx, "INSERT INTO cause (title, slug, summary) VALUES($1, $2, $3)  RETURNING id;",
		req.Item.Title, req.Item.Slug, req.Item.Summary).Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into Cause-> "+err.Error())
	}

	// get ID of creates Cause
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created Cause-> "+err.Error())
	}

	return &v1.CreateCauseResponse{
		Api: apiVersion,
		ID:  id,
	}, nil
}

// Get cause by id.
func (s *shrikeServiceServer) GetCause(ctx context.Context, req *v1.GetCauseRequest) (*v1.GetCauseResponse, error) {
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

	// query Cause by ID
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, title, slug, summary FROM cause WHERE id=$1",
		req.ID)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from Cause-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from Cause-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Cause with ID='%d' is not found",
			req.ID))
	}

	// scan Cause data into protobuf model
	var cause v1.Cause
	var createdAt time.Time
	var updatedAt time.Time

	if err := rows.Scan(&cause.ID, &createdAt, &updatedAt, &cause.Title, &cause.Slug, &cause.Summary); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from Cause row-> "+err.Error())
	}

	// Convert time.Time from database into proto timestamp.
	cause.CreatedAt, err = ptypes.TimestampProto(createdAt)
	if err != nil {
		return nil, status.Error(codes.Unknown, "createdAt field has invalid format-> "+err.Error())
	}
	cause.UpdatedAt, err = ptypes.TimestampProto(updatedAt)
	if err != nil {
		return nil, status.Error(codes.Unknown, "updatedAt field has invalid format-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple Cause rows with ID='%d'",
			req.ID))
	}

	return &v1.GetCauseResponse{
		Api:  apiVersion,
		Item: &cause,
	}, nil

}

// Read all Cause
func (s *shrikeServiceServer) ListCause(ctx context.Context, req *v1.ListCauseRequest) (*v1.ListCauseResponse, error) {
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

	// get Cause list
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, title, slug, summary FROM cause")
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from Cause-> "+err.Error())
	}
	defer rows.Close()

	// Variables to store results returned by database.
	list := []*v1.Cause{}
	var createdAt time.Time
	var updatedAt time.Time

	for rows.Next() {
		cause := new(v1.Cause)
		if err := rows.Scan(&cause.ID, &createdAt, &updatedAt, &cause.Title, &cause.Slug, &cause.Summary); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from Cause row-> "+err.Error())
		}
		// Convert time.Time from database into proto timestamp.
		cause.CreatedAt, err = ptypes.TimestampProto(createdAt)
		if err != nil {
			return nil, status.Error(codes.Unknown, "createdAt field has invalid format-> "+err.Error())
		}
		cause.UpdatedAt, err = ptypes.TimestampProto(updatedAt)
		if err != nil {
			return nil, status.Error(codes.Unknown, "updatedAt field has invalid format-> "+err.Error())
		}

		list = append(list, cause)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from Cause-> "+err.Error())
	}

	return &v1.ListCauseResponse{
		Api:   apiVersion,
		Items: list,
	}, nil
}

// Update Cause
func (s *shrikeServiceServer) UpdateCause(ctx context.Context, req *v1.UpdateCauseRequest) (*v1.UpdateCauseResponse, error) {
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

	// update cause
	res, err := c.ExecContext(ctx, "UPDATE cause SET title=$2, slug=$3, summary=$4 WHERE id=$1",
		req.Item.ID, req.Item.Title, req.Item.Slug, req.Item.Summary)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update Cause-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Cause with ID='%d' is not found",
			req.Item.ID))
	}

	return &v1.UpdateCauseResponse{
		Api:     apiVersion,
		Updated: rows,
	}, nil
}

// Delete cause
func (s *shrikeServiceServer) DeleteCause(ctx context.Context, req *v1.DeleteCauseRequest) (*v1.DeleteCauseResponse, error) {
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

	// delete cause
	res, err := c.ExecContext(ctx, "DELETE FROM cause WHERE id=$1", req.ID)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete Cause-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Cause with ID='%d' is not found",
			req.ID))
	}

	return &v1.DeleteCauseResponse{
		Api:     apiVersion,
		Deleted: rows,
	}, nil
}
