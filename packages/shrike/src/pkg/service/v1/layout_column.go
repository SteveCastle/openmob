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

// Create new LayoutColumn
func (s *shrikeServiceServer) CreateLayoutColumn(ctx context.Context, req *v1.CreateLayoutColumnRequest) (*v1.CreateLayoutColumnResponse, error) {
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
	// insert LayoutColumn entity data
	err = c.QueryRowContext(ctx, "INSERT INTO layout_column (layout_row, width) VALUES($1, $2)  RETURNING id;",
		req.Item.LayoutRow, req.Item.Width).Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into LayoutColumn-> "+err.Error())
	}

	// get ID of creates LayoutColumn
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created LayoutColumn-> "+err.Error())
	}

	return &v1.CreateLayoutColumnResponse{
		Api: apiVersion,
		ID:  id,
	}, nil
}

// Get layout_column by id.
func (s *shrikeServiceServer) GetLayoutColumn(ctx context.Context, req *v1.GetLayoutColumnRequest) (*v1.GetLayoutColumnResponse, error) {
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

	// query LayoutColumn by ID
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, layout_row, width FROM layout_column WHERE id=$1",
		req.ID)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from LayoutColumn-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from LayoutColumn-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("LayoutColumn with ID='%s' is not found",
			req.ID))
	}

	// scan LayoutColumn data into protobuf model
	var layoutcolumn v1.LayoutColumn
	var createdAt pq.NullTime
	var updatedAt pq.NullTime

	if err := rows.Scan(&layoutcolumn.ID, &createdAt, &updatedAt, &layoutcolumn.LayoutRow, &layoutcolumn.Width); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from LayoutColumn row-> "+err.Error())
	}

	// Convert pq.NullTime from database into proto timestamp.
	if createdAt.Valid {
		layoutcolumn.CreatedAt, err = ptypes.TimestampProto(createdAt.Time)
		if err != nil {
			return nil, status.Error(codes.Unknown, "createdAt field has invalid format-> "+err.Error())
		}
	}
	if updatedAt.Valid {
		layoutcolumn.UpdatedAt, err = ptypes.TimestampProto(updatedAt.Time)
		if err != nil {
			return nil, status.Error(codes.Unknown, "createdAt field has invalid format-> "+err.Error())
		}
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple LayoutColumn rows with ID='%s'",
			req.ID))
	}

	return &v1.GetLayoutColumnResponse{
		Api:  apiVersion,
		Item: &layoutcolumn,
	}, nil

}

// Read all LayoutColumn
func (s *shrikeServiceServer) ListLayoutColumn(ctx context.Context, req *v1.ListLayoutColumnRequest) (*v1.ListLayoutColumnResponse, error) {
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

	// Generate SQL to select all columns in LayoutColumn Table
	// Then generate filtering and ordering sql and finally run query.
	querySQL := queries.BuildLayoutColumnListQuery(req.Filters, req.Ordering, req.Limit)
	// Execute query and scan into return type.
	rows, err := c.QueryContext(ctx, querySQL)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from LayoutColumn-> "+err.Error())
	}
	defer rows.Close()

	// Variables to store results returned by database.
	list := []*v1.LayoutColumn{}
	var createdAt pq.NullTime
	var updatedAt pq.NullTime

	for rows.Next() {
		layoutcolumn := new(v1.LayoutColumn)
		if err := rows.Scan(&layoutcolumn.ID, &createdAt, &updatedAt, &layoutcolumn.LayoutRow, &layoutcolumn.Width); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from LayoutColumn row-> "+err.Error())
		}
		// Convert pq.NullTime from database into proto timestamp.
		if createdAt.Valid {
			layoutcolumn.CreatedAt, err = ptypes.TimestampProto(createdAt.Time)
			if err != nil {
				return nil, status.Error(codes.Unknown, "createdAt field has invalid format-> "+err.Error())
			}
		}
		if updatedAt.Valid {
			layoutcolumn.UpdatedAt, err = ptypes.TimestampProto(updatedAt.Time)
			if err != nil {
				return nil, status.Error(codes.Unknown, "createdAt field has invalid format-> "+err.Error())
			}
		}

		list = append(list, layoutcolumn)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from LayoutColumn-> "+err.Error())
	}

	return &v1.ListLayoutColumnResponse{
		Api:   apiVersion,
		Items: list,
	}, nil
}

// Update LayoutColumn
func (s *shrikeServiceServer) UpdateLayoutColumn(ctx context.Context, req *v1.UpdateLayoutColumnRequest) (*v1.UpdateLayoutColumnResponse, error) {
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

	// update layout_column
	res, err := c.ExecContext(ctx, "UPDATE layout_column SET layout_row=$2, width=$3 WHERE id=$1",
		req.Item.ID, req.Item.LayoutRow, req.Item.Width)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update LayoutColumn-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("LayoutColumn with ID='%s' is not found",
			req.Item.ID))
	}

	return &v1.UpdateLayoutColumnResponse{
		Api:     apiVersion,
		Updated: rows,
	}, nil
}

// Delete layout_column
func (s *shrikeServiceServer) DeleteLayoutColumn(ctx context.Context, req *v1.DeleteLayoutColumnRequest) (*v1.DeleteLayoutColumnResponse, error) {
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

	// delete layout_column
	res, err := c.ExecContext(ctx, "DELETE FROM layout_column WHERE id=$1", req.ID)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete LayoutColumn-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("LayoutColumn with ID='%s' is not found",
			req.ID))
	}

	return &v1.DeleteLayoutColumnResponse{
		Api:     apiVersion,
		Deleted: rows,
	}, nil
}
