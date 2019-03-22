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

// Create new HomePage
func (s *shrikeServiceServer) CreateHomePage(ctx context.Context, req *v1.CreateHomePageRequest) (*v1.CreateHomePageResponse, error) {
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
	// insert HomePage entity data
	err = c.QueryRowContext(ctx, "INSERT INTO home_page (title, layout) VALUES($1, $2)  RETURNING id;",
		req.Item.Title, req.Item.Layout).Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into HomePage-> "+err.Error())
	}

	// get ID of creates HomePage
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created HomePage-> "+err.Error())
	}

	return &v1.CreateHomePageResponse{
		Api: apiVersion,
		ID:  id,
	}, nil
}

// Get home_page by id.
func (s *shrikeServiceServer) GetHomePage(ctx context.Context, req *v1.GetHomePageRequest) (*v1.GetHomePageResponse, error) {
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

	// query HomePage by ID
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, title, layout FROM home_page WHERE id=$1",
		req.ID)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from HomePage-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from HomePage-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("HomePage with ID='%s' is not found",
			req.ID))
	}

	// scan HomePage data into protobuf model
	var homepage v1.HomePage
	var createdAt pq.NullTime
	var updatedAt pq.NullTime

	if err := rows.Scan(&homepage.ID, &createdAt, &updatedAt, &homepage.Title, &homepage.Layout); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from HomePage row-> "+err.Error())
	}

	// Convert pq.NullTime from database into proto timestamp.
	if createdAt.Valid {
		homepage.CreatedAt, err = ptypes.TimestampProto(createdAt.Time)
		if err != nil {
			return nil, status.Error(codes.Unknown, "createdAt field has invalid format-> "+err.Error())
		}
	}
	if updatedAt.Valid {
		homepage.UpdatedAt, err = ptypes.TimestampProto(updatedAt.Time)
		if err != nil {
			return nil, status.Error(codes.Unknown, "createdAt field has invalid format-> "+err.Error())
		}
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple HomePage rows with ID='%s'",
			req.ID))
	}

	return &v1.GetHomePageResponse{
		Api:  apiVersion,
		Item: &homepage,
	}, nil

}

// Read all HomePage
func (s *shrikeServiceServer) ListHomePage(ctx context.Context, req *v1.ListHomePageRequest) (*v1.ListHomePageResponse, error) {
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

	// Generate SQL to select all columns in HomePage Table
	// Then generate filtering and ordering sql and finally run query.

	baseSQL := "SELECT id, created_at, updated_at, title, layout FROM home_page"
	querySQL := queries.BuildHomePageFilters(req.Filters, req.Ordering, req.Limit)
	SQL := fmt.Sprintf("%s %s", baseSQL, querySQL)
	rows, err := c.QueryContext(ctx, SQL)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from HomePage-> "+err.Error())
	}
	defer rows.Close()

	// Variables to store results returned by database.
	list := []*v1.HomePage{}
	var createdAt pq.NullTime
	var updatedAt pq.NullTime

	for rows.Next() {
		homepage := new(v1.HomePage)
		if err := rows.Scan(&homepage.ID, &createdAt, &updatedAt, &homepage.Title, &homepage.Layout); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from HomePage row-> "+err.Error())
		}
		// Convert pq.NullTime from database into proto timestamp.
		if createdAt.Valid {
			homepage.CreatedAt, err = ptypes.TimestampProto(createdAt.Time)
			if err != nil {
				return nil, status.Error(codes.Unknown, "createdAt field has invalid format-> "+err.Error())
			}
		}
		if updatedAt.Valid {
			homepage.UpdatedAt, err = ptypes.TimestampProto(updatedAt.Time)
			if err != nil {
				return nil, status.Error(codes.Unknown, "createdAt field has invalid format-> "+err.Error())
			}
		}

		list = append(list, homepage)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from HomePage-> "+err.Error())
	}

	return &v1.ListHomePageResponse{
		Api:   apiVersion,
		Items: list,
	}, nil
}

// Update HomePage
func (s *shrikeServiceServer) UpdateHomePage(ctx context.Context, req *v1.UpdateHomePageRequest) (*v1.UpdateHomePageResponse, error) {
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

	// update home_page
	res, err := c.ExecContext(ctx, "UPDATE home_page SET title=$2, layout=$3 WHERE id=$1",
		req.Item.ID, req.Item.Title, req.Item.Layout)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update HomePage-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("HomePage with ID='%s' is not found",
			req.Item.ID))
	}

	return &v1.UpdateHomePageResponse{
		Api:     apiVersion,
		Updated: rows,
	}, nil
}

// Delete home_page
func (s *shrikeServiceServer) DeleteHomePage(ctx context.Context, req *v1.DeleteHomePageRequest) (*v1.DeleteHomePageResponse, error) {
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

	// delete home_page
	res, err := c.ExecContext(ctx, "DELETE FROM home_page WHERE id=$1", req.ID)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete HomePage-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("HomePage with ID='%s' is not found",
			req.ID))
	}

	return &v1.DeleteHomePageResponse{
		Api:     apiVersion,
		Deleted: rows,
	}, nil
}
