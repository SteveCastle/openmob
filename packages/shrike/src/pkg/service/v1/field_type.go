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

// Create new FieldType
func (s *shrikeServiceServer) CreateFieldType(ctx context.Context, req *v1.CreateFieldTypeRequest) (*v1.CreateFieldTypeResponse, error) {
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
	// insert FieldType entity data
	err = c.QueryRowContext(ctx, "INSERT INTO field_type (title, data_type, string_value_default, int_value_default, float_value_default, boolean_value_default, date_time_value_default, component_type) VALUES($1, $2, $3, $4, $5, $6, $7, $8)  RETURNING id;",
		req.Item.Title, req.Item.DataType, req.Item.StringValueDefault, req.Item.IntValueDefault, req.Item.FloatValueDefault, req.Item.BooleanValueDefault, req.Item.DateTimeValueDefault, req.Item.ComponentType).Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into FieldType-> "+err.Error())
	}

	// get ID of creates FieldType
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created FieldType-> "+err.Error())
	}

	return &v1.CreateFieldTypeResponse{
		Api: apiVersion,
		ID:  id,
	}, nil
}

// Get field_type by id.
func (s *shrikeServiceServer) GetFieldType(ctx context.Context, req *v1.GetFieldTypeRequest) (*v1.GetFieldTypeResponse, error) {
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

	// query FieldType by ID
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, title, data_type, string_value_default, int_value_default, float_value_default, boolean_value_default, date_time_value_default, component_type FROM field_type WHERE id=$1",
		req.ID)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from FieldType-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from FieldType-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("FieldType with ID='%s' is not found",
			req.ID))
	}

	// scan FieldType data into protobuf model
	var fieldtype v1.FieldType
	var createdAt pq.NullTime
	var updatedAt pq.NullTime
	var dateTimeValueDefault pq.NullTime

	if err := rows.Scan(&fieldtype.ID, &createdAt, &updatedAt, &fieldtype.Title, &fieldtype.DataType, &fieldtype.StringValueDefault, &fieldtype.IntValueDefault, &fieldtype.FloatValueDefault, &fieldtype.BooleanValueDefault, &dateTimeValueDefault, &fieldtype.ComponentType); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from FieldType row-> "+err.Error())
	}

	// Convert pq.NullTime from database into proto timestamp.
	if createdAt.Valid {
		fieldtype.CreatedAt, err = ptypes.TimestampProto(createdAt.Time)
		if err != nil {
			return nil, status.Error(codes.Unknown, "createdAt field has invalid format-> "+err.Error())
		}
	}
	if updatedAt.Valid {
		fieldtype.UpdatedAt, err = ptypes.TimestampProto(updatedAt.Time)
		if err != nil {
			return nil, status.Error(codes.Unknown, "createdAt field has invalid format-> "+err.Error())
		}
	}
	if dateTimeValueDefault.Valid {
		fieldtype.DateTimeValueDefault, err = ptypes.TimestampProto(dateTimeValueDefault.Time)
		if err != nil {
			return nil, status.Error(codes.Unknown, "createdAt field has invalid format-> "+err.Error())
		}
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple FieldType rows with ID='%s'",
			req.ID))
	}

	return &v1.GetFieldTypeResponse{
		Api:  apiVersion,
		Item: &fieldtype,
	}, nil

}

// Read all FieldType
func (s *shrikeServiceServer) ListFieldType(ctx context.Context, req *v1.ListFieldTypeRequest) (*v1.ListFieldTypeResponse, error) {
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

	// Generate SQL to select all columns in FieldType Table
	// Then generate filtering and ordering sql and finally run query.

	baseSQL := "SELECT id, created_at, updated_at, title, data_type, string_value_default, int_value_default, float_value_default, boolean_value_default, date_time_value_default, component_type FROM field_type"
	querySQL := queries.BuildFieldTypeFilters(req.Filters, req.Ordering, req.Limit)
	SQL := fmt.Sprintf("%s %s", baseSQL, querySQL)
	rows, err := c.QueryContext(ctx, SQL)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from FieldType-> "+err.Error())
	}
	defer rows.Close()

	// Variables to store results returned by database.
	list := []*v1.FieldType{}
	var createdAt pq.NullTime
	var updatedAt pq.NullTime
	var dateTimeValueDefault pq.NullTime

	for rows.Next() {
		fieldtype := new(v1.FieldType)
		if err := rows.Scan(&fieldtype.ID, &createdAt, &updatedAt, &fieldtype.Title, &fieldtype.DataType, &fieldtype.StringValueDefault, &fieldtype.IntValueDefault, &fieldtype.FloatValueDefault, &fieldtype.BooleanValueDefault, &dateTimeValueDefault, &fieldtype.ComponentType); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from FieldType row-> "+err.Error())
		}
		// Convert pq.NullTime from database into proto timestamp.
		if createdAt.Valid {
			fieldtype.CreatedAt, err = ptypes.TimestampProto(createdAt.Time)
			if err != nil {
				return nil, status.Error(codes.Unknown, "createdAt field has invalid format-> "+err.Error())
			}
		}
		if updatedAt.Valid {
			fieldtype.UpdatedAt, err = ptypes.TimestampProto(updatedAt.Time)
			if err != nil {
				return nil, status.Error(codes.Unknown, "createdAt field has invalid format-> "+err.Error())
			}
		}
		if dateTimeValueDefault.Valid {
			fieldtype.DateTimeValueDefault, err = ptypes.TimestampProto(dateTimeValueDefault.Time)
			if err != nil {
				return nil, status.Error(codes.Unknown, "createdAt field has invalid format-> "+err.Error())
			}
		}

		list = append(list, fieldtype)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from FieldType-> "+err.Error())
	}

	return &v1.ListFieldTypeResponse{
		Api:   apiVersion,
		Items: list,
	}, nil
}

// Update FieldType
func (s *shrikeServiceServer) UpdateFieldType(ctx context.Context, req *v1.UpdateFieldTypeRequest) (*v1.UpdateFieldTypeResponse, error) {
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

	// update field_type
	res, err := c.ExecContext(ctx, "UPDATE field_type SET title=$2, data_type=$3, string_value_default=$4, int_value_default=$5, float_value_default=$6, boolean_value_default=$7, date_time_value_default=$8, component_type=$9 WHERE id=$1",
		req.Item.ID, req.Item.Title, req.Item.DataType, req.Item.StringValueDefault, req.Item.IntValueDefault, req.Item.FloatValueDefault, req.Item.BooleanValueDefault, req.Item.DateTimeValueDefault, req.Item.ComponentType)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update FieldType-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("FieldType with ID='%s' is not found",
			req.Item.ID))
	}

	return &v1.UpdateFieldTypeResponse{
		Api:     apiVersion,
		Updated: rows,
	}, nil
}

// Delete field_type
func (s *shrikeServiceServer) DeleteFieldType(ctx context.Context, req *v1.DeleteFieldTypeRequest) (*v1.DeleteFieldTypeResponse, error) {
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

	// delete field_type
	res, err := c.ExecContext(ctx, "DELETE FROM field_type WHERE id=$1", req.ID)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete FieldType-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("FieldType with ID='%s' is not found",
			req.ID))
	}

	return &v1.DeleteFieldTypeResponse{
		Api:     apiVersion,
		Deleted: rows,
	}, nil
}
