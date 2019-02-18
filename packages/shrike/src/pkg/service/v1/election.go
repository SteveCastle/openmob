package v1

import (
	"context"
	"fmt"

	v1 "github.com/SteveCastle/openmob/packages/shrike/src/pkg/api/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Create new Election
func (s *shrikeServiceServer) CreateElection(ctx context.Context, req *v1.CreateElectionRequest) (*v1.CreateElectionResponse, error) {
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
	// insert Election entity data
	err = c.QueryRowContext(ctx, "INSERT INTO election (id, created_at, updated_at, title) VALUES($1, $2, $3, $4)  RETURNING id;",
		 req.Item.ID,  req.Item.CreatedAt,  req.Item.UpdatedAt,  req.Item.Title, ).Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into Election-> "+err.Error())
	}

	// get ID of creates Election
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created Election-> "+err.Error())
	}

	return &v1.CreateElectionResponse{
		Api: apiVersion,
		ID:  id,
	}, nil
}

// Get election by id.
func (s *shrikeServiceServer) GetElection(ctx context.Context, req *v1.GetElectionRequest) (*v1.GetElectionResponse, error) {
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

	// query Election by ID
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, title FROM election WHERE id=$1",
		req.ID)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from Election-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from Election-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Election with ID='%d' is not found",
			req.ID))
	}

	// get Election data
	var election v1.Election
	if err := rows.Scan( &election.ID,  &election.CreatedAt,  &election.UpdatedAt,  &election.Title, ); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from Election row-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple Election rows with ID='%d'",
			req.ID))
	}

	return &v1.GetElectionResponse{
		Api:  apiVersion,
		Item: &election,
	}, nil

}

// Read all Election
func (s *shrikeServiceServer) ListElection(ctx context.Context, req *v1.ListElectionRequest) (*v1.ListElectionResponse, error) {
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

	// get Election list
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, title FROM election")
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from Election-> "+err.Error())
	}
	defer rows.Close()

	list := []*v1.Election{}
	for rows.Next() {
		election := new(v1.Election)
		if err := rows.Scan( &election.ID,  &election.CreatedAt,  &election.UpdatedAt,  &election.Title, ); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from Election row-> "+err.Error())
		}
		list = append(list, election)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from Election-> "+err.Error())
	}

	return &v1.ListElectionResponse{
		Api:   apiVersion,
		Items: list,
	}, nil
}

// Update Election
func (s *shrikeServiceServer) UpdateElection(ctx context.Context, req *v1.UpdateElectionRequest) (*v1.UpdateElectionResponse, error) {
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

	// update election
	res, err := c.ExecContext(ctx, "UPDATE election SET $1 ,$2 ,$3 ,$4  WHERE id=$1",
		req.Item.ID,req.Item.CreatedAt,req.Item.UpdatedAt,req.Item.Title, )
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update Election-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Election with ID='%d' is not found",
			req.Item.ID))
	}

	return &v1.UpdateElectionResponse{
		Api:     apiVersion,
		Updated: rows,
	}, nil
}

// Delete election
func (s *shrikeServiceServer) DeleteElection(ctx context.Context, req *v1.DeleteElectionRequest) (*v1.DeleteElectionResponse, error) {
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

	// delete election
	res, err := c.ExecContext(ctx, "DELETE FROM election WHERE id=$1", req.ID)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete Election-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Election with ID='%d' is not found",
			req.ID))
	}

	return &v1.DeleteElectionResponse{
		Api:     apiVersion,
		Deleted: rows,
	}, nil
}
