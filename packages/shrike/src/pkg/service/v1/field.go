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

// Create new Field
func (s *shrikeServiceServer) CreateField(ctx context.Context, req *v1.CreateFieldRequest) (*v1.CreateFieldResponse, error) {
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
	// insert Field entity data
	err = c.QueryRowContext(ctx, "INSERT INTO field (field_type, string_value, int_value, float_value, boolean_value, date_time_value, data_path, component) VALUES($1, $2, $3, $4, $5, $6, $7, $8)  RETURNING id;",
		req.Item.FieldType, req.Item.StringValue, req.Item.IntValue, req.Item.FloatValue, req.Item.BooleanValue, req.Item.DateTimeValue, req.Item.DataPath, req.Item.Component).Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into Field-> "+err.Error())
	}

	// get ID of creates Field
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created Field-> "+err.Error())
	}

	return &v1.CreateFieldResponse{
		Api: apiVersion,
		ID:  id,
	}, nil
}

// Get field by id.
func (s *shrikeServiceServer) GetField(ctx context.Context, req *v1.GetFieldRequest) (*v1.GetFieldResponse, error) {
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

	// query Field by ID
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, field_type, string_value, int_value, float_value, boolean_value, date_time_value, data_path, component FROM field WHERE id=$1",
		req.ID)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from Field-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from Field-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Field with ID='%s' is not found",
			req.ID))
	}

	// scan Field data into protobuf model
	var field v1.Field
	var createdAt pq.NullTime
	var updatedAt pq.NullTime
	var dateTimeValue pq.NullTime

	if err := rows.Scan(&field.ID, &createdAt, &updatedAt, &field.FieldType, &field.StringValue, &field.IntValue, &field.FloatValue, &field.BooleanValue, &dateTimeValue, &field.DataPath, &field.Component); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from Field row-> "+err.Error())
	}

	// Convert pq.NullTime from database into proto timestamp.
	if createdAt.Valid {
		field.CreatedAt, err = ptypes.TimestampProto(createdAt.Time)
		if err != nil {
			return nil, status.Error(codes.Unknown, "createdAt field has invalid format-> "+err.Error())
		}
	}
	if updatedAt.Valid {
		field.UpdatedAt, err = ptypes.TimestampProto(updatedAt.Time)
		if err != nil {
			return nil, status.Error(codes.Unknown, "createdAt field has invalid format-> "+err.Error())
		}
	}
	if dateTimeValue.Valid {
		field.DateTimeValue, err = ptypes.TimestampProto(dateTimeValue.Time)
		if err != nil {
			return nil, status.Error(codes.Unknown, "createdAt field has invalid format-> "+err.Error())
		}
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple Field rows with ID='%s'",
			req.ID))
	}

	return &v1.GetFieldResponse{
		Api:  apiVersion,
		Item: &field,
	}, nil

}

// Read all Field
func (s *shrikeServiceServer) ListField(ctx context.Context, req *v1.ListFieldRequest) (*v1.ListFieldResponse, error) {
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

	// Generate SQL to select all columns in Field Table
	// Then generate filtering and ordering sql and finally run query.
	querySQL := queries.BuildFieldListQuery(req.Filters, req.Ordering, req.Limit)
	// Execute query and scan into return type.
	rows, err := c.QueryContext(ctx, querySQL)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from Field-> "+err.Error())
	}
	defer rows.Close()

	// Variables to store results returned by database.
	list := []*v1.Field{}
	var createdAt pq.NullTime
	var updatedAt pq.NullTime
	var dateTimeValue pq.NullTime

	for rows.Next() {
		field := new(v1.Field)
		if err := rows.Scan(&field.ID, &createdAt, &updatedAt, &field.FieldType, &field.StringValue, &field.IntValue, &field.FloatValue, &field.BooleanValue, &dateTimeValue, &field.DataPath, &field.Component); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from Field row-> "+err.Error())
		}
		// Convert pq.NullTime from database into proto timestamp.
		if createdAt.Valid {
			field.CreatedAt, err = ptypes.TimestampProto(createdAt.Time)
			if err != nil {
				return nil, status.Error(codes.Unknown, "createdAt field has invalid format-> "+err.Error())
			}
		}
		if updatedAt.Valid {
			field.UpdatedAt, err = ptypes.TimestampProto(updatedAt.Time)
			if err != nil {
				return nil, status.Error(codes.Unknown, "createdAt field has invalid format-> "+err.Error())
			}
		}
		if dateTimeValue.Valid {
			field.DateTimeValue, err = ptypes.TimestampProto(dateTimeValue.Time)
			if err != nil {
				return nil, status.Error(codes.Unknown, "createdAt field has invalid format-> "+err.Error())
			}
		}

		list = append(list, field)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from Field-> "+err.Error())
	}

	return &v1.ListFieldResponse{
		Api:   apiVersion,
		Items: list,
	}, nil
}

// Update Field
func (s *shrikeServiceServer) UpdateField(ctx context.Context, req *v1.UpdateFieldRequest) (*v1.UpdateFieldResponse, error) {
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

	// update field
	res, err := c.ExecContext(ctx, "UPDATE field SET field_type=$2, string_value=$3, int_value=$4, float_value=$5, boolean_value=$6, date_time_value=$7, data_path=$8, component=$9 WHERE id=$1",
		req.Item.ID, req.Item.FieldType, req.Item.StringValue, req.Item.IntValue, req.Item.FloatValue, req.Item.BooleanValue, req.Item.DateTimeValue, req.Item.DataPath, req.Item.Component)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update Field-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Field with ID='%s' is not found",
			req.Item.ID))
	}

	return &v1.UpdateFieldResponse{
		Api:     apiVersion,
		Updated: rows,
	}, nil
}

// Delete field
func (s *shrikeServiceServer) DeleteField(ctx context.Context, req *v1.DeleteFieldRequest) (*v1.DeleteFieldResponse, error) {
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

	// delete field
	res, err := c.ExecContext(ctx, "DELETE FROM field WHERE id=$1", req.ID)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete Field-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Field with ID='%s' is not found",
			req.ID))
	}

	return &v1.DeleteFieldResponse{
		Api:     apiVersion,
		Deleted: rows,
	}, nil
}
