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

// Create new Voter
func (s *shrikeServiceServer) CreateVoter(ctx context.Context, req *v1.CreateVoterRequest) (*v1.CreateVoterResponse, error) {
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
	// insert Voter entity data
	err = c.QueryRowContext(ctx, "INSERT INTO voter (contact, cause) VALUES($1, $2)  RETURNING id;",
		req.Item.Contact, req.Item.Cause).Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into Voter-> "+err.Error())
	}

	// get ID of creates Voter
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created Voter-> "+err.Error())
	}

	return &v1.CreateVoterResponse{
		Api: apiVersion,
		ID:  id,
	}, nil
}

// Get voter by id.
func (s *shrikeServiceServer) GetVoter(ctx context.Context, req *v1.GetVoterRequest) (*v1.GetVoterResponse, error) {
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

	// query Voter by ID
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, contact, cause FROM voter WHERE id=$1",
		req.ID)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from Voter-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from Voter-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Voter with ID='%s' is not found",
			req.ID))
	}

	// scan Voter data into protobuf model
	var voter v1.Voter
	var createdAt time.Time
	var updatedAt time.Time

	if err := rows.Scan(&voter.ID, &createdAt, &updatedAt, &voter.Contact, &voter.Cause); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from Voter row-> "+err.Error())
	}

	// Convert time.Time from database into proto timestamp.
	voter.CreatedAt, err = ptypes.TimestampProto(createdAt)
	if err != nil {
		return nil, status.Error(codes.Unknown, "createdAt field has invalid format-> "+err.Error())
	}
	voter.UpdatedAt, err = ptypes.TimestampProto(updatedAt)
	if err != nil {
		return nil, status.Error(codes.Unknown, "updatedAt field has invalid format-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple Voter rows with ID='%s'",
			req.ID))
	}

	return &v1.GetVoterResponse{
		Api:  apiVersion,
		Item: &voter,
	}, nil

}

// Read all Voter
func (s *shrikeServiceServer) ListVoter(ctx context.Context, req *v1.ListVoterRequest) (*v1.ListVoterResponse, error) {
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

	// get Voter list
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, contact, cause FROM voter")
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from Voter-> "+err.Error())
	}
	defer rows.Close()

	// Variables to store results returned by database.
	list := []*v1.Voter{}
	var createdAt time.Time
	var updatedAt time.Time

	for rows.Next() {
		voter := new(v1.Voter)
		if err := rows.Scan(&voter.ID, &createdAt, &updatedAt, &voter.Contact, &voter.Cause); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from Voter row-> "+err.Error())
		}
		// Convert time.Time from database into proto timestamp.
		voter.CreatedAt, err = ptypes.TimestampProto(createdAt)
		if err != nil {
			return nil, status.Error(codes.Unknown, "createdAt field has invalid format-> "+err.Error())
		}
		voter.UpdatedAt, err = ptypes.TimestampProto(updatedAt)
		if err != nil {
			return nil, status.Error(codes.Unknown, "updatedAt field has invalid format-> "+err.Error())
		}

		list = append(list, voter)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from Voter-> "+err.Error())
	}

	return &v1.ListVoterResponse{
		Api:   apiVersion,
		Items: list,
	}, nil
}

// Update Voter
func (s *shrikeServiceServer) UpdateVoter(ctx context.Context, req *v1.UpdateVoterRequest) (*v1.UpdateVoterResponse, error) {
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

	// update voter
	res, err := c.ExecContext(ctx, "UPDATE voter SET contact=$2, cause=$3 WHERE id=$1",
		req.Item.ID, req.Item.Contact, req.Item.Cause)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update Voter-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Voter with ID='%s' is not found",
			req.Item.ID))
	}

	return &v1.UpdateVoterResponse{
		Api:     apiVersion,
		Updated: rows,
	}, nil
}

// Delete voter
func (s *shrikeServiceServer) DeleteVoter(ctx context.Context, req *v1.DeleteVoterRequest) (*v1.DeleteVoterResponse, error) {
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

	// delete voter
	res, err := c.ExecContext(ctx, "DELETE FROM voter WHERE id=$1", req.ID)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete Voter-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Voter with ID='%s' is not found",
			req.ID))
	}

	return &v1.DeleteVoterResponse{
		Api:     apiVersion,
		Deleted: rows,
	}, nil
}
