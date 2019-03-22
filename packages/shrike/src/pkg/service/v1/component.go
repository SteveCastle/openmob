package v1

import (
	"context"
	"fmt"

	v1 "github.com/SteveCastle/openmob/packages/shrike/src/pkg/api/v1"
	"github.com/SteveCastle/openmob/packages/shrike/src/pkg/queries"
	"github.com/golang/protobuf/ptypes"
	"github.com/lib/pq"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Create new Component
func (s *shrikeServiceServer) CreateComponent(ctx context.Context, req *v1.CreateComponentRequest) (*v1.CreateComponentResponse, error) {
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
	// insert Component entity data
	err = c.QueryRowContext(ctx, "INSERT INTO component (component_type, component_implementation, layout_column) VALUES($1, $2, $3)  RETURNING id;",
		req.Item.ComponentType, req.Item.ComponentImplementation, req.Item.LayoutColumn).Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into Component-> "+err.Error())
	}

	// get ID of creates Component
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created Component-> "+err.Error())
	}

	return &v1.CreateComponentResponse{
		Api: apiVersion,
		ID:  id,
	}, nil
}

// Get component by id.
func (s *shrikeServiceServer) GetComponent(ctx context.Context, req *v1.GetComponentRequest) (*v1.GetComponentResponse, error) {
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

	// query Component by ID
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, component_type, component_implementation, layout_column FROM component WHERE id=$1",
		req.ID)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from Component-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from Component-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Component with ID='%s' is not found",
			req.ID))
	}

	// scan Component data into protobuf model
	var component v1.Component
	var createdAt pq.NullTime
	var updatedAt pq.NullTime

	if err := rows.Scan(&component.ID, &createdAt, &updatedAt, &component.ComponentType, &component.ComponentImplementation, &component.LayoutColumn); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from Component row-> "+err.Error())
	}

	// Convert pq.NullTime from database into proto timestamp.
	if createdAt.Valid {
		component.CreatedAt, err = ptypes.TimestampProto(createdAt.Time)
		if err != nil {
			return nil, status.Error(codes.Unknown, "createdAt field has invalid format-> "+err.Error())
		}
	}
	if updatedAt.Valid {
		component.UpdatedAt, err = ptypes.TimestampProto(updatedAt.Time)
		if err != nil {
			return nil, status.Error(codes.Unknown, "createdAt field has invalid format-> "+err.Error())
		}
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple Component rows with ID='%s'",
			req.ID))
	}

	return &v1.GetComponentResponse{
		Api:  apiVersion,
		Item: &component,
	}, nil

}

// Read all Component
func (s *shrikeServiceServer) ListComponent(ctx context.Context, req *v1.ListComponentRequest) (*v1.ListComponentResponse, error) {
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

	// Generate SQL to select all columns in Component Table
	// Then generate filtering and ordering sql and finally run query.

	baseSQL := "SELECT id, created_at, updated_at, component_type, component_implementation, layout_column FROM component"
	querySQL := queries.BuildComponentFilters(req.Filters, req.Ordering, req.Limit)
	SQL := fmt.Sprintf("%s %s", baseSQL, querySQL)
	rows, err := c.QueryContext(ctx, SQL)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from Component-> "+err.Error())
	}
	defer rows.Close()

	// Variables to store results returned by database.
	list := []*v1.Component{}
	var createdAt pq.NullTime
	var updatedAt pq.NullTime

	for rows.Next() {
		component := new(v1.Component)
		if err := rows.Scan(&component.ID, &createdAt, &updatedAt, &component.ComponentType, &component.ComponentImplementation, &component.LayoutColumn); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from Component row-> "+err.Error())
		}
		// Convert pq.NullTime from database into proto timestamp.
		if createdAt.Valid {
			component.CreatedAt, err = ptypes.TimestampProto(createdAt.Time)
			if err != nil {
				return nil, status.Error(codes.Unknown, "createdAt field has invalid format-> "+err.Error())
			}
		}
		if updatedAt.Valid {
			component.UpdatedAt, err = ptypes.TimestampProto(updatedAt.Time)
			if err != nil {
				return nil, status.Error(codes.Unknown, "createdAt field has invalid format-> "+err.Error())
			}
		}

		list = append(list, component)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from Component-> "+err.Error())
	}

	return &v1.ListComponentResponse{
		Api:   apiVersion,
		Items: list,
	}, nil
}

// Update Component
func (s *shrikeServiceServer) UpdateComponent(ctx context.Context, req *v1.UpdateComponentRequest) (*v1.UpdateComponentResponse, error) {
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

	// update component
	res, err := c.ExecContext(ctx, "UPDATE component SET component_type=$2, component_implementation=$3, layout_column=$4 WHERE id=$1",
		req.Item.ID, req.Item.ComponentType, req.Item.ComponentImplementation, req.Item.LayoutColumn)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update Component-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Component with ID='%s' is not found",
			req.Item.ID))
	}

	return &v1.UpdateComponentResponse{
		Api:     apiVersion,
		Updated: rows,
	}, nil
}

// Delete component
func (s *shrikeServiceServer) DeleteComponent(ctx context.Context, req *v1.DeleteComponentRequest) (*v1.DeleteComponentResponse, error) {
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

	// delete component
	res, err := c.ExecContext(ctx, "DELETE FROM component WHERE id=$1", req.ID)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete Component-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Component with ID='%s' is not found",
			req.ID))
	}

	return &v1.DeleteComponentResponse{
		Api:     apiVersion,
		Deleted: rows,
	}, nil
}
