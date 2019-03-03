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

// Create new ComponentImplementation
func (s *shrikeServiceServer) CreateComponentImplementation(ctx context.Context, req *v1.CreateComponentImplementationRequest) (*v1.CreateComponentImplementationResponse, error) {
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
	// insert ComponentImplementation entity data
	err = c.QueryRowContext(ctx, "INSERT INTO component_implementation () VALUES()  RETURNING id;").Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into ComponentImplementation-> "+err.Error())
	}

	// get ID of creates ComponentImplementation
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created ComponentImplementation-> "+err.Error())
	}

	return &v1.CreateComponentImplementationResponse{
		Api: apiVersion,
		ID:  id,
	}, nil
}

// Get component_implementation by id.
func (s *shrikeServiceServer) GetComponentImplementation(ctx context.Context, req *v1.GetComponentImplementationRequest) (*v1.GetComponentImplementationResponse, error) {
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

	// query ComponentImplementation by ID
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at FROM component_implementation WHERE id=$1",
		req.ID)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from ComponentImplementation-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from ComponentImplementation-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("ComponentImplementation with ID='%s' is not found",
			req.ID))
	}

	// scan ComponentImplementation data into protobuf model
	var componentimplementation v1.ComponentImplementation
	var createdAt time.Time
	var updatedAt time.Time

	if err := rows.Scan(&componentimplementation.ID, &createdAt, &updatedAt); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from ComponentImplementation row-> "+err.Error())
	}

	// Convert time.Time from database into proto timestamp.
	componentimplementation.CreatedAt, err = ptypes.TimestampProto(createdAt)
	if err != nil {
		return nil, status.Error(codes.Unknown, "createdAt field has invalid format-> "+err.Error())
	}
	componentimplementation.UpdatedAt, err = ptypes.TimestampProto(updatedAt)
	if err != nil {
		return nil, status.Error(codes.Unknown, "updatedAt field has invalid format-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple ComponentImplementation rows with ID='%s'",
			req.ID))
	}

	return &v1.GetComponentImplementationResponse{
		Api:  apiVersion,
		Item: &componentimplementation,
	}, nil

}

// Read all ComponentImplementation
func (s *shrikeServiceServer) ListComponentImplementation(ctx context.Context, req *v1.ListComponentImplementationRequest) (*v1.ListComponentImplementationResponse, error) {
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

	// get ComponentImplementation list
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at FROM component_implementation")
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from ComponentImplementation-> "+err.Error())
	}
	defer rows.Close()

	// Variables to store results returned by database.
	list := []*v1.ComponentImplementation{}
	var createdAt time.Time
	var updatedAt time.Time

	for rows.Next() {
		componentimplementation := new(v1.ComponentImplementation)
		if err := rows.Scan(&componentimplementation.ID, &createdAt, &updatedAt); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from ComponentImplementation row-> "+err.Error())
		}
		// Convert time.Time from database into proto timestamp.
		componentimplementation.CreatedAt, err = ptypes.TimestampProto(createdAt)
		if err != nil {
			return nil, status.Error(codes.Unknown, "createdAt field has invalid format-> "+err.Error())
		}
		componentimplementation.UpdatedAt, err = ptypes.TimestampProto(updatedAt)
		if err != nil {
			return nil, status.Error(codes.Unknown, "updatedAt field has invalid format-> "+err.Error())
		}

		list = append(list, componentimplementation)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from ComponentImplementation-> "+err.Error())
	}

	return &v1.ListComponentImplementationResponse{
		Api:   apiVersion,
		Items: list,
	}, nil
}

// Update ComponentImplementation
func (s *shrikeServiceServer) UpdateComponentImplementation(ctx context.Context, req *v1.UpdateComponentImplementationRequest) (*v1.UpdateComponentImplementationResponse, error) {
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

	// update component_implementation
	res, err := c.ExecContext(ctx, "UPDATE component_implementation SET  WHERE id=$1",
		req.Item.ID)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update ComponentImplementation-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("ComponentImplementation with ID='%s' is not found",
			req.Item.ID))
	}

	return &v1.UpdateComponentImplementationResponse{
		Api:     apiVersion,
		Updated: rows,
	}, nil
}

// Delete component_implementation
func (s *shrikeServiceServer) DeleteComponentImplementation(ctx context.Context, req *v1.DeleteComponentImplementationRequest) (*v1.DeleteComponentImplementationResponse, error) {
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

	// delete component_implementation
	res, err := c.ExecContext(ctx, "DELETE FROM component_implementation WHERE id=$1", req.ID)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete ComponentImplementation-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("ComponentImplementation with ID='%s' is not found",
			req.ID))
	}

	return &v1.DeleteComponentImplementationResponse{
		Api:     apiVersion,
		Deleted: rows,
	}, nil
}
