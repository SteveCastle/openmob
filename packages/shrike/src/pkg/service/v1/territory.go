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

// Create new Territory
func (s *shrikeServiceServer) CreateTerritory(ctx context.Context, req *v1.CreateTerritoryRequest) (*v1.CreateTerritoryResponse, error) {
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
	// insert Territory entity data
	err = c.QueryRowContext(ctx, "INSERT INTO territory (title) VALUES($1)  RETURNING id;",
		req.Item.Title).Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into Territory-> "+err.Error())
	}

	// get ID of creates Territory
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created Territory-> "+err.Error())
	}

	return &v1.CreateTerritoryResponse{
		Api: apiVersion,
		ID:  id,
	}, nil
}

// Get territory by id.
func (s *shrikeServiceServer) GetTerritory(ctx context.Context, req *v1.GetTerritoryRequest) (*v1.GetTerritoryResponse, error) {
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

	// query Territory by ID
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, title FROM territory WHERE id=$1",
		req.ID)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from Territory-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from Territory-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Territory with ID='%d' is not found",
			req.ID))
	}

	// scan Territory data into protobuf model
	var territory v1.Territory
	var createdAt time.Time
	var updatedAt time.Time

	if err := rows.Scan(&territory.ID, &createdAt, &updatedAt, &territory.Title); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from Territory row-> "+err.Error())
	}

	// Convert time.Time from database into proto timestamp.
	territory.CreatedAt, err = ptypes.TimestampProto(createdAt)
	if err != nil {
		return nil, status.Error(codes.Unknown, "createdAt field has invalid format-> "+err.Error())
	}
	territory.UpdatedAt, err = ptypes.TimestampProto(updatedAt)
	if err != nil {
		return nil, status.Error(codes.Unknown, "updatedAt field has invalid format-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple Territory rows with ID='%d'",
			req.ID))
	}

	return &v1.GetTerritoryResponse{
		Api:  apiVersion,
		Item: &territory,
	}, nil

}

// Read all Territory
func (s *shrikeServiceServer) ListTerritory(ctx context.Context, req *v1.ListTerritoryRequest) (*v1.ListTerritoryResponse, error) {
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

	// get Territory list
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, title FROM territory")
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from Territory-> "+err.Error())
	}
	defer rows.Close()

	// Variables to store results returned by database.
	list := []*v1.Territory{}
	var createdAt time.Time
	var updatedAt time.Time

	for rows.Next() {
		territory := new(v1.Territory)
		if err := rows.Scan(&territory.ID, &createdAt, &updatedAt, &territory.Title); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from Territory row-> "+err.Error())
		}
		// Convert time.Time from database into proto timestamp.
		territory.CreatedAt, err = ptypes.TimestampProto(createdAt)
		if err != nil {
			return nil, status.Error(codes.Unknown, "createdAt field has invalid format-> "+err.Error())
		}
		territory.UpdatedAt, err = ptypes.TimestampProto(updatedAt)
		if err != nil {
			return nil, status.Error(codes.Unknown, "updatedAt field has invalid format-> "+err.Error())
		}

		list = append(list, territory)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from Territory-> "+err.Error())
	}

	return &v1.ListTerritoryResponse{
		Api:   apiVersion,
		Items: list,
	}, nil
}

// Update Territory
func (s *shrikeServiceServer) UpdateTerritory(ctx context.Context, req *v1.UpdateTerritoryRequest) (*v1.UpdateTerritoryResponse, error) {
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

	// update territory
	res, err := c.ExecContext(ctx, "UPDATE territory SET title=$2 WHERE id=$1",
		req.Item.ID, req.Item.Title)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update Territory-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Territory with ID='%d' is not found",
			req.Item.ID))
	}

	return &v1.UpdateTerritoryResponse{
		Api:     apiVersion,
		Updated: rows,
	}, nil
}

// Delete territory
func (s *shrikeServiceServer) DeleteTerritory(ctx context.Context, req *v1.DeleteTerritoryRequest) (*v1.DeleteTerritoryResponse, error) {
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

	// delete territory
	res, err := c.ExecContext(ctx, "DELETE FROM territory WHERE id=$1", req.ID)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete Territory-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Territory with ID='%d' is not found",
			req.ID))
	}

	return &v1.DeleteTerritoryResponse{
		Api:     apiVersion,
		Deleted: rows,
	}, nil
}
