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

// Create new ComponentTypeFields
func (s *shrikeServiceServer) CreateComponentTypeFields(ctx context.Context, req *v1.CreateComponentTypeFieldsRequest) (*v1.CreateComponentTypeFieldsResponse, error) {
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
	// insert ComponentTypeFields entity data
	err = c.QueryRowContext(ctx, "INSERT INTO component_type_fields (component_type, field_type, weight, required) VALUES($1, $2, $3, $4)  RETURNING id;",
		req.Item.ComponentType, req.Item.FieldType, req.Item.Weight, req.Item.Required).Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into ComponentTypeFields-> "+err.Error())
	}

	// get ID of creates ComponentTypeFields
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created ComponentTypeFields-> "+err.Error())
	}

	return &v1.CreateComponentTypeFieldsResponse{
		Api: apiVersion,
		ID:  id,
	}, nil
}

// Get component_type_fields by id.
func (s *shrikeServiceServer) GetComponentTypeFields(ctx context.Context, req *v1.GetComponentTypeFieldsRequest) (*v1.GetComponentTypeFieldsResponse, error) {
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

	// query ComponentTypeFields by ID
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, component_type, field_type, weight, required FROM component_type_fields WHERE id=$1",
		req.ID)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from ComponentTypeFields-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from ComponentTypeFields-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("ComponentTypeFields with ID='%s' is not found",
			req.ID))
	}

	// scan ComponentTypeFields data into protobuf model
	var componenttypefields v1.ComponentTypeFields
	var createdAt pq.NullTime
	var updatedAt pq.NullTime

	if err := rows.Scan(&componenttypefields.ID, &createdAt, &updatedAt, &componenttypefields.ComponentType, &componenttypefields.FieldType, &componenttypefields.Weight, &componenttypefields.Required); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from ComponentTypeFields row-> "+err.Error())
	}

	// Convert pq.NullTime from database into proto timestamp.
	if createdAt.Valid {
		componenttypefields.CreatedAt, err = ptypes.TimestampProto(createdAt.Time)
		if err != nil {
			return nil, status.Error(codes.Unknown, "createdAt field has invalid format-> "+err.Error())
		}
	}
	if updatedAt.Valid {
		componenttypefields.UpdatedAt, err = ptypes.TimestampProto(updatedAt.Time)
		if err != nil {
			return nil, status.Error(codes.Unknown, "createdAt field has invalid format-> "+err.Error())
		}
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple ComponentTypeFields rows with ID='%s'",
			req.ID))
	}

	return &v1.GetComponentTypeFieldsResponse{
		Api:  apiVersion,
		Item: &componenttypefields,
	}, nil

}

// Read all ComponentTypeFields
func (s *shrikeServiceServer) ListComponentTypeFields(ctx context.Context, req *v1.ListComponentTypeFieldsRequest) (*v1.ListComponentTypeFieldsResponse, error) {
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

	// Generate SQL to select all columns in ComponentTypeFields Table
	// Then generate filtering and ordering sql and finally run query.
	querySQL := queries.BuildComponentTypeFieldsListQuery(req.Filters, req.Ordering, req.Limit)
	// Execute query and scan into return type.
	rows, err := c.QueryContext(ctx, querySQL)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from ComponentTypeFields-> "+err.Error())
	}
	defer rows.Close()

	// Variables to store results returned by database.
	list := []*v1.ComponentTypeFields{}
	var createdAt pq.NullTime
	var updatedAt pq.NullTime

	for rows.Next() {
		componenttypefields := new(v1.ComponentTypeFields)
		if err := rows.Scan(&componenttypefields.ID, &createdAt, &updatedAt, &componenttypefields.ComponentType, &componenttypefields.FieldType, &componenttypefields.Weight, &componenttypefields.Required); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from ComponentTypeFields row-> "+err.Error())
		}
		// Convert pq.NullTime from database into proto timestamp.
		if createdAt.Valid {
			componenttypefields.CreatedAt, err = ptypes.TimestampProto(createdAt.Time)
			if err != nil {
				return nil, status.Error(codes.Unknown, "createdAt field has invalid format-> "+err.Error())
			}
		}
		if updatedAt.Valid {
			componenttypefields.UpdatedAt, err = ptypes.TimestampProto(updatedAt.Time)
			if err != nil {
				return nil, status.Error(codes.Unknown, "createdAt field has invalid format-> "+err.Error())
			}
		}

		list = append(list, componenttypefields)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from ComponentTypeFields-> "+err.Error())
	}

	return &v1.ListComponentTypeFieldsResponse{
		Api:   apiVersion,
		Items: list,
	}, nil
}

// Update ComponentTypeFields
func (s *shrikeServiceServer) UpdateComponentTypeFields(ctx context.Context, req *v1.UpdateComponentTypeFieldsRequest) (*v1.UpdateComponentTypeFieldsResponse, error) {
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

	// update component_type_fields
	res, err := c.ExecContext(ctx, "UPDATE component_type_fields SET component_type=$2, field_type=$3, weight=$4, required=$5 WHERE id=$1",
		req.Item.ID, req.Item.ComponentType, req.Item.FieldType, req.Item.Weight, req.Item.Required)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update ComponentTypeFields-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("ComponentTypeFields with ID='%s' is not found",
			req.Item.ID))
	}

	return &v1.UpdateComponentTypeFieldsResponse{
		Api:     apiVersion,
		Updated: rows,
	}, nil
}

// Delete component_type_fields
func (s *shrikeServiceServer) DeleteComponentTypeFields(ctx context.Context, req *v1.DeleteComponentTypeFieldsRequest) (*v1.DeleteComponentTypeFieldsResponse, error) {
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

	// delete component_type_fields
	res, err := c.ExecContext(ctx, "DELETE FROM component_type_fields WHERE id=$1", req.ID)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete ComponentTypeFields-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("ComponentTypeFields with ID='%s' is not found",
			req.ID))
	}

	return &v1.DeleteComponentTypeFieldsResponse{
		Api:     apiVersion,
		Deleted: rows,
	}, nil
}
