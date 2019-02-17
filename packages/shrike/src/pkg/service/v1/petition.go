package v1

import (
	"context"
	"fmt"

	v1 "github.com/SteveCastle/openmob/packages/shrike/src/pkg/api/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)


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
		 req.Item.ID,  req.Item.CreatedAt,  req.Item.UpdatedAt,  req.Item.Title, ).Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into Petition-> "+err.Error())
	}

	// get ID of creates Petition
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created Petition-> "+err.Error())
	}

	return &v1.CreatePetitionResponse{
		Api: apiVersion,
		ID:  id,
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
		req.ID)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from Petition-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from Petition-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Petition with ID='%d' is not found",
			req.ID))
	}

	// get Petition data
	var petition v1.Petition
	if err := rows.Scan( &petition.ID,  &petition.CreatedAt,  &petition.UpdatedAt,  &petition.Title, ); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from Petition row-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple Petition rows with ID='%d'",
			req.ID))
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
	res, err := c.ExecContext(ctx, "UPDATE petition SET $1, $2, $3, $4,  WHERE id=$1",
		req.Item.ID,req.Item.CreatedAt,req.Item.UpdatedAt,req.Item.Title, )
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update Petition-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Petition with ID='%d' is not found",
			req.Item.ID))
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
	res, err := c.ExecContext(ctx, "DELETE FROM petition WHERE id=$1", req.ID)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete Petition-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Petition with ID='%d' is not found",
			req.ID))
	}

	return &v1.DeletePetitionResponse{
		Api:     apiVersion,
		Deleted: rows,
	}, nil
}
