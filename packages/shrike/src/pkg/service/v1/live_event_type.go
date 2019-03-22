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

// Create new LiveEventType
func (s *shrikeServiceServer) CreateLiveEventType(ctx context.Context, req *v1.CreateLiveEventTypeRequest) (*v1.CreateLiveEventTypeResponse, error) {
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
	// insert LiveEventType entity data
	err = c.QueryRowContext(ctx, "INSERT INTO live_event_type (title) VALUES($1)  RETURNING id;",
		req.Item.Title).Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into LiveEventType-> "+err.Error())
	}

	// get ID of creates LiveEventType
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created LiveEventType-> "+err.Error())
	}

	return &v1.CreateLiveEventTypeResponse{
		Api: apiVersion,
		ID:  id,
	}, nil
}

// Get live_event_type by id.
func (s *shrikeServiceServer) GetLiveEventType(ctx context.Context, req *v1.GetLiveEventTypeRequest) (*v1.GetLiveEventTypeResponse, error) {
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

	// query LiveEventType by ID
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, title FROM live_event_type WHERE id=$1",
		req.ID)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from LiveEventType-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from LiveEventType-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("LiveEventType with ID='%s' is not found",
			req.ID))
	}

	// scan LiveEventType data into protobuf model
	var liveeventtype v1.LiveEventType
	var createdAt pq.NullTime
	var updatedAt pq.NullTime

	if err := rows.Scan(&liveeventtype.ID, &createdAt, &updatedAt, &liveeventtype.Title); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from LiveEventType row-> "+err.Error())
	}

	// Convert pq.NullTime from database into proto timestamp.
	if createdAt.Valid {
		liveeventtype.CreatedAt, err = ptypes.TimestampProto(createdAt.Time)
		if err != nil {
			return nil, status.Error(codes.Unknown, "createdAt field has invalid format-> "+err.Error())
		}
	}
	if updatedAt.Valid {
		liveeventtype.UpdatedAt, err = ptypes.TimestampProto(updatedAt.Time)
		if err != nil {
			return nil, status.Error(codes.Unknown, "createdAt field has invalid format-> "+err.Error())
		}
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple LiveEventType rows with ID='%s'",
			req.ID))
	}

	return &v1.GetLiveEventTypeResponse{
		Api:  apiVersion,
		Item: &liveeventtype,
	}, nil

}

// Read all LiveEventType
func (s *shrikeServiceServer) ListLiveEventType(ctx context.Context, req *v1.ListLiveEventTypeRequest) (*v1.ListLiveEventTypeResponse, error) {
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

	// Generate SQL to select all columns in LiveEventType Table
	// Then generate filtering and ordering sql and finally run query.

	baseSQL := "SELECT id, created_at, updated_at, title FROM live_event_type"
	querySQL := queries.BuildLiveEventTypeFilters(req.Filters, req.Ordering, req.Limit)
	SQL := fmt.Sprintf("%s %s", baseSQL, querySQL)
	rows, err := c.QueryContext(ctx, SQL)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from LiveEventType-> "+err.Error())
	}
	defer rows.Close()

	// Variables to store results returned by database.
	list := []*v1.LiveEventType{}
	var createdAt pq.NullTime
	var updatedAt pq.NullTime

	for rows.Next() {
		liveeventtype := new(v1.LiveEventType)
		if err := rows.Scan(&liveeventtype.ID, &createdAt, &updatedAt, &liveeventtype.Title); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from LiveEventType row-> "+err.Error())
		}
		// Convert pq.NullTime from database into proto timestamp.
		if createdAt.Valid {
			liveeventtype.CreatedAt, err = ptypes.TimestampProto(createdAt.Time)
			if err != nil {
				return nil, status.Error(codes.Unknown, "createdAt field has invalid format-> "+err.Error())
			}
		}
		if updatedAt.Valid {
			liveeventtype.UpdatedAt, err = ptypes.TimestampProto(updatedAt.Time)
			if err != nil {
				return nil, status.Error(codes.Unknown, "createdAt field has invalid format-> "+err.Error())
			}
		}

		list = append(list, liveeventtype)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from LiveEventType-> "+err.Error())
	}

	return &v1.ListLiveEventTypeResponse{
		Api:   apiVersion,
		Items: list,
	}, nil
}

// Update LiveEventType
func (s *shrikeServiceServer) UpdateLiveEventType(ctx context.Context, req *v1.UpdateLiveEventTypeRequest) (*v1.UpdateLiveEventTypeResponse, error) {
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

	// update live_event_type
	res, err := c.ExecContext(ctx, "UPDATE live_event_type SET title=$2 WHERE id=$1",
		req.Item.ID, req.Item.Title)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update LiveEventType-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("LiveEventType with ID='%s' is not found",
			req.Item.ID))
	}

	return &v1.UpdateLiveEventTypeResponse{
		Api:     apiVersion,
		Updated: rows,
	}, nil
}

// Delete live_event_type
func (s *shrikeServiceServer) DeleteLiveEventType(ctx context.Context, req *v1.DeleteLiveEventTypeRequest) (*v1.DeleteLiveEventTypeResponse, error) {
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

	// delete live_event_type
	res, err := c.ExecContext(ctx, "DELETE FROM live_event_type WHERE id=$1", req.ID)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete LiveEventType-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("LiveEventType with ID='%s' is not found",
			req.ID))
	}

	return &v1.DeleteLiveEventTypeResponse{
		Api:     apiVersion,
		Deleted: rows,
	}, nil
}
