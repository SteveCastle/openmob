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

// NewShrikeServiceServer creates Candidate service
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

// Create new todo task
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
	err = c.QueryRowContext(ctx, "INSERT INTO candidate (title) VALUES($1)  RETURNING id;",
		req.Item.Title).Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into Candidate-> "+err.Error())
	}

	// get ID of creates Candidate
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created Candidate-> "+err.Error())
	}

	return &v1.CreateCandidateResponse{
		Api: apiVersion,
		Id:  id,
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
	rows, err := c.QueryContext(ctx, "SELECT id, title FROM candidate WHERE id=$1",
		req.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from Candidate-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from Candidate-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Candidate with ID='%d' is not found",
			req.Id))
	}

	// get Candidate data
	var td v1.Candidate
	if err := rows.Scan(&td.Id, &td.Title); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from Candidate row-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple Candidate rows with ID='%d'",
			req.Id))
	}

	return &v1.GetCandidateResponse{
		Api:  apiVersion,
		Item: &td,
	}, nil

}

// Read all todo tasks
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
	rows, err := c.QueryContext(ctx, "SELECT id,title FROM candidate")
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from Candidate-> "+err.Error())
	}
	defer rows.Close()

	list := []*v1.Candidate{}
	for rows.Next() {
		td := new(v1.Candidate)
		if err := rows.Scan(&td.Id, &td.Title); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from Candidate row-> "+err.Error())
		}
		list = append(list, td)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from Candidate-> "+err.Error())
	}

	return &v1.ListCandidateResponse{
		Api:   apiVersion,
		Items: list,
	}, nil
}

// Update todo task
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
	res, err := c.ExecContext(ctx, "UPDATE candidate SET title=$1 WHERE id=$2",
		req.Item.Title, req.Item.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update Candidate-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Candidate with ID='%d' is not found",
			req.Item.Id))
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
	res, err := c.ExecContext(ctx, "DELETE FROM candidate WHERE id=$1", req.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete Candidate-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Candidate with ID='%d' is not found",
			req.Id))
	}

	return &v1.DeleteCandidateResponse{
		Api:     apiVersion,
		Deleted: rows,
	}, nil
}
