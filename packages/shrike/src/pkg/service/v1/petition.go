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
	var id string
	// insert Petition entity data
	err = c.QueryRowContext(ctx, "INSERT INTO petition (title) VALUES($1)  RETURNING id;",
		req.Item.Title).Scan(&id)
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
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, title FROM petition WHERE id=$1",
		req.ID)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from Petition-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from Petition-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Petition with ID='%s' is not found",
			req.ID))
	}

	// scan Petition data into protobuf model
	var petition v1.Petition
	var createdAt time.Time
	var updatedAt time.Time

	if err := rows.Scan(&petition.ID, &createdAt, &updatedAt, &petition.Title); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from Petition row-> "+err.Error())
	}

	// Convert time.Time from database into proto timestamp.
	petition.CreatedAt, err = ptypes.TimestampProto(createdAt)
	if err != nil {
		return nil, status.Error(codes.Unknown, "createdAt field has invalid format-> "+err.Error())
	}
	petition.UpdatedAt, err = ptypes.TimestampProto(updatedAt)
	if err != nil {
		return nil, status.Error(codes.Unknown, "createdAt field has invalid format-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple Petition rows with ID='%s'",
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

	// Generate SQL to select all columns in Petition Table
	// Then generate filtering and ordering sql and finally run query.

	baseSQL := "SELECT id, created_at, updated_at, title FROM petition"
	querySQL := queries.BuildPetitionFilters(req.Filters, req.Ordering, req.Limit)
	SQL := fmt.Sprintf("%s %s", baseSQL, querySQL)
	rows, err := c.QueryContext(ctx, SQL)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from Petition-> "+err.Error())
	}
	defer rows.Close()

	// Variables to store results returned by database.
	list := []*v1.Petition{}
	var createdAt time.Time
	var updatedAt time.Time

	for rows.Next() {
		petition := new(v1.Petition)
		if err := rows.Scan(&petition.ID, &createdAt, &updatedAt, &petition.Title); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from Petition row-> "+err.Error())
		}
		// Convert time.Time from database into proto timestamp.
		petition.CreatedAt, err = ptypes.TimestampProto(createdAt)
		if err != nil {
			return nil, status.Error(codes.Unknown, "createdAt field has invalid format-> "+err.Error())
		}
		petition.UpdatedAt, err = ptypes.TimestampProto(updatedAt)
		if err != nil {
			return nil, status.Error(codes.Unknown, "createdAt field has invalid format-> "+err.Error())
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
	res, err := c.ExecContext(ctx, "UPDATE petition SET title=$2 WHERE id=$1",
		req.Item.ID, req.Item.Title)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update Petition-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Petition with ID='%s' is not found",
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
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Petition with ID='%s' is not found",
			req.ID))
	}

	return &v1.DeletePetitionResponse{
		Api:     apiVersion,
		Deleted: rows,
	}, nil
}
