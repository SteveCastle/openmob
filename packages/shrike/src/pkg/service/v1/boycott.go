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

// Create new Boycott
func (s *shrikeServiceServer) CreateBoycott(ctx context.Context, req *v1.CreateBoycottRequest) (*v1.CreateBoycottResponse, error) {
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
	// insert Boycott entity data
	err = c.QueryRowContext(ctx, "INSERT INTO boycott (title) VALUES($1)  RETURNING id;",
		req.Item.Title).Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into Boycott-> "+err.Error())
	}

	// get ID of creates Boycott
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created Boycott-> "+err.Error())
	}

	return &v1.CreateBoycottResponse{
		Api: apiVersion,
		ID:  id,
	}, nil
}

// Get boycott by id.
func (s *shrikeServiceServer) GetBoycott(ctx context.Context, req *v1.GetBoycottRequest) (*v1.GetBoycottResponse, error) {
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

	// query Boycott by ID
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, title FROM boycott WHERE id=$1",
		req.ID)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from Boycott-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from Boycott-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Boycott with ID='%d' is not found",
			req.ID))
	}

	// scan Boycott data into protobuf model
	var boycott v1.Boycott
	var createdAt time.Time
	var updatedAt time.Time

	if err := rows.Scan(&boycott.ID, &createdAt, &updatedAt, &boycott.Title); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from Boycott row-> "+err.Error())
	}

	// Convert time.Time from database into proto timestamp.
	boycott.CreatedAt, err = ptypes.TimestampProto(createdAt)
	if err != nil {
		return nil, status.Error(codes.Unknown, "createdAt field has invalid format-> "+err.Error())
	}
	boycott.UpdatedAt, err = ptypes.TimestampProto(updatedAt)
	if err != nil {
		return nil, status.Error(codes.Unknown, "updatedAt field has invalid format-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple Boycott rows with ID='%d'",
			req.ID))
	}

	return &v1.GetBoycottResponse{
		Api:  apiVersion,
		Item: &boycott,
	}, nil

}

// Read all Boycott
func (s *shrikeServiceServer) ListBoycott(ctx context.Context, req *v1.ListBoycottRequest) (*v1.ListBoycottResponse, error) {
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

	// get Boycott list
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, title FROM boycott")
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from Boycott-> "+err.Error())
	}
	defer rows.Close()

	// Variables to store results returned by database.
	list := []*v1.Boycott{}
	var createdAt time.Time
	var updatedAt time.Time

	for rows.Next() {
		boycott := new(v1.Boycott)
		if err := rows.Scan(&boycott.ID, &createdAt, &updatedAt, &boycott.Title); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from Boycott row-> "+err.Error())
		}
		// Convert time.Time from database into proto timestamp.
		boycott.CreatedAt, err = ptypes.TimestampProto(createdAt)
		if err != nil {
			return nil, status.Error(codes.Unknown, "createdAt field has invalid format-> "+err.Error())
		}
		boycott.UpdatedAt, err = ptypes.TimestampProto(updatedAt)
		if err != nil {
			return nil, status.Error(codes.Unknown, "updatedAt field has invalid format-> "+err.Error())
		}

		list = append(list, boycott)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from Boycott-> "+err.Error())
	}

	return &v1.ListBoycottResponse{
		Api:   apiVersion,
		Items: list,
	}, nil
}

// Update Boycott
func (s *shrikeServiceServer) UpdateBoycott(ctx context.Context, req *v1.UpdateBoycottRequest) (*v1.UpdateBoycottResponse, error) {
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

	// update boycott
	res, err := c.ExecContext(ctx, "UPDATE boycott SET title=$2 WHERE id=$1",
		req.Item.ID, req.Item.Title)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update Boycott-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Boycott with ID='%d' is not found",
			req.Item.ID))
	}

	return &v1.UpdateBoycottResponse{
		Api:     apiVersion,
		Updated: rows,
	}, nil
}

// Delete boycott
func (s *shrikeServiceServer) DeleteBoycott(ctx context.Context, req *v1.DeleteBoycottRequest) (*v1.DeleteBoycottResponse, error) {
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

	// delete boycott
	res, err := c.ExecContext(ctx, "DELETE FROM boycott WHERE id=$1", req.ID)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete Boycott-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Boycott with ID='%d' is not found",
			req.ID))
	}

	return &v1.DeleteBoycottResponse{
		Api:     apiVersion,
		Deleted: rows,
	}, nil
}
