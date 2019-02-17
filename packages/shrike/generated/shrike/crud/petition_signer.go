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

// NewShrikeServiceServer creates PetitionSigner service
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

// Create new PetitionSigner
func (s *shrikeServiceServer) CreatePetitionSigner(ctx context.Context, req *v1.CreatePetitionSignerRequest) (*v1.CreatePetitionSignerResponse, error) {
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
	// insert PetitionSigner entity data
	err = c.QueryRowContext(ctx, "INSERT INTO petition_signer (id, created_at, updated_at, petition, contact, cause, ) VALUES($1, $2, $3, $4, $5, $6, )  RETURNING id;",
		 req.ItemID  req.ItemCreatedAt  req.ItemUpdatedAt  req.ItemPetition  req.ItemContact  req.ItemCause ).Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into PetitionSigner-> "+err.Error())
	}

	// get ID of creates PetitionSigner
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created PetitionSigner-> "+err.Error())
	}

	return &v1.CreatePetitionSignerResponse{
		Api: apiVersion,
		Id:  id,
	}, nil
}

// Get petition_signer by id.
func (s *shrikeServiceServer) GetPetitionSigner(ctx context.Context, req *v1.GetPetitionSignerRequest) (*v1.GetPetitionSignerResponse, error) {
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

	// query PetitionSigner by ID
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, petition, contact, cause,  FROM petition_signer WHERE id=$1",
		req.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from PetitionSigner-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from PetitionSigner-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("PetitionSigner with ID='%d' is not found",
			req.Id))
	}

	// get PetitionSigner data
	var petitionsigner v1.PetitionSigner
	if err := rows.Scan( &petitionsigner.ID,  &petitionsigner.CreatedAt,  &petitionsigner.UpdatedAt,  &petitionsigner.Petition,  &petitionsigner.Contact,  &petitionsigner.Cause, ); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from PetitionSigner row-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple PetitionSigner rows with ID='%d'",
			req.Id))
	}

	return &v1.GetPetitionSignerResponse{
		Api:  apiVersion,
		Item: &petitionsigner,
	}, nil

}

// Read all PetitionSigner
func (s *shrikeServiceServer) ListPetitionSigner(ctx context.Context, req *v1.ListPetitionSignerRequest) (*v1.ListPetitionSignerResponse, error) {
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

	// get PetitionSigner list
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, petition, contact, cause,  FROM petition_signer")
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from PetitionSigner-> "+err.Error())
	}
	defer rows.Close()

	list := []*v1.PetitionSigner{}
	for rows.Next() {
		petitionsigner := new(v1.PetitionSigner)
		if err := rows.Scan( &petitionsigner.ID,  &petitionsigner.CreatedAt,  &petitionsigner.UpdatedAt,  &petitionsigner.Petition,  &petitionsigner.Contact,  &petitionsigner.Cause, ); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from PetitionSigner row-> "+err.Error())
		}
		list = append(list, petitionsigner)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from PetitionSigner-> "+err.Error())
	}

	return &v1.ListPetitionSignerResponse{
		Api:   apiVersion,
		Items: list,
	}, nil
}

// Update PetitionSigner
func (s *shrikeServiceServer) UpdatePetitionSigner(ctx context.Context, req *v1.UpdatePetitionSignerRequest) (*v1.UpdatePetitionSignerResponse, error) {
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

	// update petition_signer
	res, err := c.ExecContext(ctx, "UPDATE petition_signer SET title=$1 WHERE id=$2",
		req.Item.Title, req.Item.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update PetitionSigner-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("PetitionSigner with ID='%d' is not found",
			req.Item.Id))
	}

	return &v1.UpdatePetitionSignerResponse{
		Api:     apiVersion,
		Updated: rows,
	}, nil
}

// Delete petition_signer
func (s *shrikeServiceServer) DeletePetitionSigner(ctx context.Context, req *v1.DeletePetitionSignerRequest) (*v1.DeletePetitionSignerResponse, error) {
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

	// delete petition_signer
	res, err := c.ExecContext(ctx, "DELETE FROM petition_signer WHERE id=$1", req.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete PetitionSigner-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("PetitionSigner with ID='%d' is not found",
			req.Id))
	}

	return &v1.DeletePetitionSignerResponse{
		Api:     apiVersion,
		Deleted: rows,
	}, nil
}
