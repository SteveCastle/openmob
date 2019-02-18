package v1

import (
	"context"
	"fmt"

	v1 "github.com/SteveCastle/openmob/packages/shrike/src/pkg/api/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Create new Candidate
func (s *shrikeServiceServer) CreateCandidate(ctx context.Context, req *v1.CreateCandidateRequest) (*v1.CreateCandidateResponse, error) {
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
	// insert Candidate entity data
	err = c.QueryRowContext(ctx, "INSERT INTO candidate (id, created_at, updated_at, election) VALUES($1, $2, $3, $4)  RETURNING id;",
		 req.Item.ID,  req.Item.CreatedAt,  req.Item.UpdatedAt,  req.Item.Election, ).Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into Candidate-> "+err.Error())
	}

	// get ID of creates Candidate
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created Candidate-> "+err.Error())
	}

	return &v1.CreateCandidateResponse{
		Api: apiVersion,
		ID:  id,
	}, nil
}

// Get candidate by id.
func (s *shrikeServiceServer) GetCandidate(ctx context.Context, req *v1.GetCandidateRequest) (*v1.GetCandidateResponse, error) {
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

	// query Candidate by ID
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, election FROM candidate WHERE id=$1",
		req.ID)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from Candidate-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from Candidate-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Candidate with ID='%d' is not found",
			req.ID))
	}

	// get Candidate data
	var candidate v1.Candidate
	if err := rows.Scan( &candidate.ID,  &candidate.CreatedAt,  &candidate.UpdatedAt,  &candidate.Election, ); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from Candidate row-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple Candidate rows with ID='%d'",
			req.ID))
	}

	return &v1.GetCandidateResponse{
		Api:  apiVersion,
		Item: &candidate,
	}, nil

}

// Read all Candidate
func (s *shrikeServiceServer) ListCandidate(ctx context.Context, req *v1.ListCandidateRequest) (*v1.ListCandidateResponse, error) {
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

	// get Candidate list
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, election FROM candidate")
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from Candidate-> "+err.Error())
	}
	defer rows.Close()

	list := []*v1.Candidate{}
	for rows.Next() {
		candidate := new(v1.Candidate)
		if err := rows.Scan( &candidate.ID,  &candidate.CreatedAt,  &candidate.UpdatedAt,  &candidate.Election, ); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from Candidate row-> "+err.Error())
		}
		list = append(list, candidate)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from Candidate-> "+err.Error())
	}

	return &v1.ListCandidateResponse{
		Api:   apiVersion,
		Items: list,
	}, nil
}

// Update Candidate
func (s *shrikeServiceServer) UpdateCandidate(ctx context.Context, req *v1.UpdateCandidateRequest) (*v1.UpdateCandidateResponse, error) {
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

	// update candidate
	res, err := c.ExecContext(ctx, "UPDATE candidate SET $1 ,$2 ,$3 ,$4  WHERE id=$1",
		req.Item.ID,req.Item.CreatedAt,req.Item.UpdatedAt,req.Item.Election, )
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update Candidate-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Candidate with ID='%d' is not found",
			req.Item.ID))
	}

	return &v1.UpdateCandidateResponse{
		Api:     apiVersion,
		Updated: rows,
	}, nil
}

// Delete candidate
func (s *shrikeServiceServer) DeleteCandidate(ctx context.Context, req *v1.DeleteCandidateRequest) (*v1.DeleteCandidateResponse, error) {
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

	// delete candidate
	res, err := c.ExecContext(ctx, "DELETE FROM candidate WHERE id=$1", req.ID)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete Candidate-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Candidate with ID='%d' is not found",
			req.ID))
	}

	return &v1.DeleteCandidateResponse{
		Api:     apiVersion,
		Deleted: rows,
	}, nil
}
