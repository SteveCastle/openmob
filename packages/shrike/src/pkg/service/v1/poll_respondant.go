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

// Create new PollRespondant
func (s *shrikeServiceServer) CreatePollRespondant(ctx context.Context, req *v1.CreatePollRespondantRequest) (*v1.CreatePollRespondantResponse, error) {
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
	// insert PollRespondant entity data
	err = c.QueryRowContext(ctx, "INSERT INTO poll_respondant (poll, contact, cause) VALUES($1, $2, $3)  RETURNING id;",
		req.Item.Poll, req.Item.Contact, req.Item.Cause).Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into PollRespondant-> "+err.Error())
	}

	// get ID of creates PollRespondant
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created PollRespondant-> "+err.Error())
	}

	return &v1.CreatePollRespondantResponse{
		Api: apiVersion,
		ID:  id,
	}, nil
}

// Get poll_respondant by id.
func (s *shrikeServiceServer) GetPollRespondant(ctx context.Context, req *v1.GetPollRespondantRequest) (*v1.GetPollRespondantResponse, error) {
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

	// query PollRespondant by ID
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, poll, contact, cause FROM poll_respondant WHERE id=$1",
		req.ID)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from PollRespondant-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from PollRespondant-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("PollRespondant with ID='%s' is not found",
			req.ID))
	}

	// scan PollRespondant data into protobuf model
	var pollrespondant v1.PollRespondant
	var createdAt time.Time
	var updatedAt time.Time

	if err := rows.Scan(&pollrespondant.ID, &createdAt, &updatedAt, &pollrespondant.Poll, &pollrespondant.Contact, &pollrespondant.Cause); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from PollRespondant row-> "+err.Error())
	}

	// Convert time.Time from database into proto timestamp.
	pollrespondant.CreatedAt, err = ptypes.TimestampProto(createdAt)
	if err != nil {
		return nil, status.Error(codes.Unknown, "createdAt field has invalid format-> "+err.Error())
	}
	pollrespondant.UpdatedAt, err = ptypes.TimestampProto(updatedAt)
	if err != nil {
		return nil, status.Error(codes.Unknown, "updatedAt field has invalid format-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple PollRespondant rows with ID='%s'",
			req.ID))
	}

	return &v1.GetPollRespondantResponse{
		Api:  apiVersion,
		Item: &pollrespondant,
	}, nil

}

// Read all PollRespondant
func (s *shrikeServiceServer) ListPollRespondant(ctx context.Context, req *v1.ListPollRespondantRequest) (*v1.ListPollRespondantResponse, error) {
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

	// get PollRespondant list
	queries.BuildPollRespondantFilters(req.Filters, req.Ordering, req.Limit)
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, poll, contact, cause FROM poll_respondant")
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from PollRespondant-> "+err.Error())
	}
	defer rows.Close()

	// Variables to store results returned by database.
	list := []*v1.PollRespondant{}
	var createdAt time.Time
	var updatedAt time.Time

	for rows.Next() {
		pollrespondant := new(v1.PollRespondant)
		if err := rows.Scan(&pollrespondant.ID, &createdAt, &updatedAt, &pollrespondant.Poll, &pollrespondant.Contact, &pollrespondant.Cause); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from PollRespondant row-> "+err.Error())
		}
		// Convert time.Time from database into proto timestamp.
		pollrespondant.CreatedAt, err = ptypes.TimestampProto(createdAt)
		if err != nil {
			return nil, status.Error(codes.Unknown, "createdAt field has invalid format-> "+err.Error())
		}
		pollrespondant.UpdatedAt, err = ptypes.TimestampProto(updatedAt)
		if err != nil {
			return nil, status.Error(codes.Unknown, "updatedAt field has invalid format-> "+err.Error())
		}

		list = append(list, pollrespondant)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from PollRespondant-> "+err.Error())
	}

	return &v1.ListPollRespondantResponse{
		Api:   apiVersion,
		Items: list,
	}, nil
}

// Update PollRespondant
func (s *shrikeServiceServer) UpdatePollRespondant(ctx context.Context, req *v1.UpdatePollRespondantRequest) (*v1.UpdatePollRespondantResponse, error) {
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

	// update poll_respondant
	res, err := c.ExecContext(ctx, "UPDATE poll_respondant SET poll=$2, contact=$3, cause=$4 WHERE id=$1",
		req.Item.ID, req.Item.Poll, req.Item.Contact, req.Item.Cause)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update PollRespondant-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("PollRespondant with ID='%s' is not found",
			req.Item.ID))
	}

	return &v1.UpdatePollRespondantResponse{
		Api:     apiVersion,
		Updated: rows,
	}, nil
}

// Delete poll_respondant
func (s *shrikeServiceServer) DeletePollRespondant(ctx context.Context, req *v1.DeletePollRespondantRequest) (*v1.DeletePollRespondantResponse, error) {
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

	// delete poll_respondant
	res, err := c.ExecContext(ctx, "DELETE FROM poll_respondant WHERE id=$1", req.ID)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete PollRespondant-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("PollRespondant with ID='%s' is not found",
			req.ID))
	}

	return &v1.DeletePollRespondantResponse{
		Api:     apiVersion,
		Deleted: rows,
	}, nil
}
