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

// NewShrikeServiceServer creates PetitionMembership service
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
func (s *shrikeServiceServer) CreatePetitionMembership(ctx context.Context, req *v1.CreatePetitionMembershipRequest) (*v1.CreatePetitionMembershipResponse, error) {
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
	// insert PetitionMembership entity data
	err = c.QueryRowContext(ctx, "INSERT INTO petition_membership (title) VALUES($1)  RETURNING id;",
		req.Item.Title).Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into PetitionMembership-> "+err.Error())
	}

	// get ID of creates PetitionMembership
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created PetitionMembership-> "+err.Error())
	}

	return &v1.CreatePetitionMembershipResponse{
		Api: apiVersion,
		Id:  id,
	}, nil
}

// Get petition_membership by id.
func (s *shrikeServiceServer) GetPetitionMembership(ctx context.Context, req *v1.GetPetitionMembershipRequest) (*v1.GetPetitionMembershipResponse, error) {
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

	// query PetitionMembership by ID
	rows, err := c.QueryContext(ctx, "SELECT id, title FROM petition_membership WHERE id=$1",
		req.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from PetitionMembership-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from PetitionMembership-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("PetitionMembership with ID='%d' is not found",
			req.Id))
	}

	// get PetitionMembership data
	var td v1.PetitionMembership
	if err := rows.Scan(&td.Id, &td.Title); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from PetitionMembership row-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple PetitionMembership rows with ID='%d'",
			req.Id))
	}

	return &v1.GetPetitionMembershipResponse{
		Api:  apiVersion,
		Item: &td,
	}, nil

}

// Read all todo tasks
func (s *shrikeServiceServer) ListPetitionMembership(ctx context.Context, req *v1.ListPetitionMembershipRequest) (*v1.ListPetitionMembershipResponse, error) {
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

	// get PetitionMembership list
	rows, err := c.QueryContext(ctx, "SELECT id,title FROM petition_membership")
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from PetitionMembership-> "+err.Error())
	}
	defer rows.Close()

	list := []*v1.PetitionMembership{}
	for rows.Next() {
		td := new(v1.PetitionMembership)
		if err := rows.Scan(&td.Id, &td.Title); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from PetitionMembership row-> "+err.Error())
		}
		list = append(list, td)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from PetitionMembership-> "+err.Error())
	}

	return &v1.ListPetitionMembershipResponse{
		Api:   apiVersion,
		Items: list,
	}, nil
}

// Update todo task
func (s *shrikeServiceServer) UpdatePetitionMembership(ctx context.Context, req *v1.UpdatePetitionMembershipRequest) (*v1.UpdatePetitionMembershipResponse, error) {
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

	// update petition_membership
	res, err := c.ExecContext(ctx, "UPDATE petition_membership SET title=$1 WHERE id=$2",
		req.Item.Title, req.Item.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update PetitionMembership-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("PetitionMembership with ID='%d' is not found",
			req.Item.Id))
	}

	return &v1.UpdatePetitionMembershipResponse{
		Api:     apiVersion,
		Updated: rows,
	}, nil
}

// Delete petition_membership
func (s *shrikeServiceServer) DeletePetitionMembership(ctx context.Context, req *v1.DeletePetitionMembershipRequest) (*v1.DeletePetitionMembershipResponse, error) {
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

	// delete petition_membership
	res, err := c.ExecContext(ctx, "DELETE FROM petition_membership WHERE id=$1", req.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete PetitionMembership-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("PetitionMembership with ID='%d' is not found",
			req.Id))
	}

	return &v1.DeletePetitionMembershipResponse{
		Api:     apiVersion,
		Deleted: rows,
	}, nil
}
