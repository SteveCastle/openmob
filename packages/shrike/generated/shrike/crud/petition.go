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

// NewShrikeServiceServer creates Petition service
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

// Create new Petition
func (s *shrikeServiceServer) CreatePetition(ctx context.Context, req *v1.CreatePetitionRequest) (*v1.CreatePetitionResponse, error) {
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
	// insert Petition entity data
	err = c.QueryRowContext(ctx, "INSERT INTO petition (id, created_at, updated_at, title, ) VALUES($1, $2, $3, $4, )  RETURNING id;",
		 req.ItemID  req.ItemCreatedAt  req.ItemUpdatedAt  req.ItemTitle ).Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into Petition-> "+err.Error())
	}

	// get ID of creates Petition
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created Petition-> "+err.Error())
	}

	return &v1.CreatePetitionResponse{
		Api: apiVersion,
		Id:  id,
	}, nil
}

// Get petition by id.
func (s *shrikeServiceServer) GetPetition(ctx context.Context, req *v1.GetPetitionRequest) (*v1.GetPetitionResponse, error) {
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

	// query Petition by ID
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, title,  FROM petition WHERE id=$1",
		req.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from Petition-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from Petition-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Petition with ID='%d' is not found",
			req.Id))
	}

	// get Petition data
	var petition v1.Petition
	if err := rows.Scan( &petition.ID,  &petition.CreatedAt,  &petition.UpdatedAt,  &petition.Title, ); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from Petition row-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple Petition rows with ID='%d'",
			req.Id))
	}

	return &v1.GetPetitionResponse{
		Api:  apiVersion,
		Item: &petition,
	}, nil

}

// Read all Petition
func (s *shrikeServiceServer) ListPetition(ctx context.Context, req *v1.ListPetitionRequest) (*v1.ListPetitionResponse, error) {
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

	// get Petition list
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, title,  FROM petition")
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from Petition-> "+err.Error())
	}
	defer rows.Close()

	list := []*v1.Petition{}
	for rows.Next() {
		petition := new(v1.Petition)
		if err := rows.Scan( &petition.ID,  &petition.CreatedAt,  &petition.UpdatedAt,  &petition.Title, ); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from Petition row-> "+err.Error())
		}
		list = append(list, petition)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from Petition-> "+err.Error())
	}

	return &v1.ListPetitionResponse{
		Api:   apiVersion,
		Items: list,
	}, nil
}

// Update Petition
func (s *shrikeServiceServer) UpdatePetition(ctx context.Context, req *v1.UpdatePetitionRequest) (*v1.UpdatePetitionResponse, error) {
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

	// update petition
	res, err := c.ExecContext(ctx, "UPDATE petition SET title=$1 WHERE id=$2",
		req.Item.Title, req.Item.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update Petition-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Petition with ID='%d' is not found",
			req.Item.Id))
	}

	return &v1.UpdatePetitionResponse{
		Api:     apiVersion,
		Updated: rows,
	}, nil
}

// Delete petition
func (s *shrikeServiceServer) DeletePetition(ctx context.Context, req *v1.DeletePetitionRequest) (*v1.DeletePetitionResponse, error) {
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

	// delete petition
	res, err := c.ExecContext(ctx, "DELETE FROM petition WHERE id=$1", req.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete Petition-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Petition with ID='%d' is not found",
			req.Id))
	}

	return &v1.DeletePetitionResponse{
		Api:     apiVersion,
		Deleted: rows,
	}, nil
}
