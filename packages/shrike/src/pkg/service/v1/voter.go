package v1

import (
	"context"
	"fmt"

	v1 "github.com/SteveCastle/openmob/packages/shrike/src/pkg/api/v1"

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
	var id int64
	// insert Voter entity data
	err = c.QueryRowContext(ctx, "INSERT INTO voter (id, created_at, updated_at, contact, cause) VALUES($1, $2, $3, $4, $5)  RETURNING id;",
		 req.Item.ID,  req.Item.CreatedAt,  req.Item.UpdatedAt,  req.Item.Contact,  req.Item.Cause, ).Scan(&id)
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
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Voter with ID='%d' is not found",
			req.ID))
	}

	// get Voter data
	var voter v1.Voter
	if err := rows.Scan( &voter.ID,  &voter.CreatedAt,  &voter.UpdatedAt,  &voter.Contact,  &voter.Cause, ); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from Voter row-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple Voter rows with ID='%d'",
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

	list := []*v1.Voter{}
	for rows.Next() {
		voter := new(v1.Voter)
		if err := rows.Scan( &voter.ID,  &voter.CreatedAt,  &voter.UpdatedAt,  &voter.Contact,  &voter.Cause, ); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from Voter row-> "+err.Error())
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
	res, err := c.ExecContext(ctx, "UPDATE voter SET id=$1, created_at=$2, updated_at=$3, contact=$4, cause=$5 WHERE id=$1",
		req.Item.ID,req.Item.CreatedAt,req.Item.UpdatedAt,req.Item.Contact,req.Item.Cause, )
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update Voter-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Voter with ID='%d' is not found",
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
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Voter with ID='%d' is not found",
			req.ID))
	}

	return &v1.DeleteVoterResponse{
		Api:     apiVersion,
		Deleted: rows,
	}, nil
}
