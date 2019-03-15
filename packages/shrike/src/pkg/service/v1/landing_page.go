package v1

import (
	"context"
	"fmt"
	"time"

	v1 "github.com/SteveCastle/openmob/packages/shrike/src/pkg/api/v1"
	"github.com/SteveCastle/openmob/packages/shrike/src/pkg/queries"
	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Create new LandingPage
func (s *shrikeServiceServer) CreateLandingPage(ctx context.Context, req *v1.CreateLandingPageRequest) (*v1.CreateLandingPageResponse, error) {
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
	// insert LandingPage entity data
	err = c.QueryRowContext(ctx, "INSERT INTO landing_page (title, cause, layout) VALUES($1, $2, $3)  RETURNING id;",
		req.Item.Title, req.Item.Cause, req.Item.Layout).Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into LandingPage-> "+err.Error())
	}

	// get ID of creates LandingPage
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created LandingPage-> "+err.Error())
	}

	return &v1.CreateLandingPageResponse{
		Api: apiVersion,
		ID:  id,
	}, nil
}

// Get landing_page by id.
func (s *shrikeServiceServer) GetLandingPage(ctx context.Context, req *v1.GetLandingPageRequest) (*v1.GetLandingPageResponse, error) {
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

	// query LandingPage by ID
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, title, cause, layout FROM landing_page WHERE id=$1",
		req.ID)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from LandingPage-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from LandingPage-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("LandingPage with ID='%s' is not found",
			req.ID))
	}

	// scan LandingPage data into protobuf model
	var landingpage v1.LandingPage
	var createdAt time.Time
	var updatedAt time.Time

	if err := rows.Scan(&landingpage.ID, &createdAt, &updatedAt, &landingpage.Title, &landingpage.Cause, &landingpage.Layout); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from LandingPage row-> "+err.Error())
	}

	// Convert time.Time from database into proto timestamp.
	landingpage.CreatedAt, err = ptypes.TimestampProto(createdAt)
	if err != nil {
		return nil, status.Error(codes.Unknown, "createdAt field has invalid format-> "+err.Error())
	}
	landingpage.UpdatedAt, err = ptypes.TimestampProto(updatedAt)
	if err != nil {
		return nil, status.Error(codes.Unknown, "updatedAt field has invalid format-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple LandingPage rows with ID='%s'",
			req.ID))
	}

	return &v1.GetLandingPageResponse{
		Api:  apiVersion,
		Item: &landingpage,
	}, nil

}

// Read all LandingPage
func (s *shrikeServiceServer) ListLandingPage(ctx context.Context, req *v1.ListLandingPageRequest) (*v1.ListLandingPageResponse, error) {
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

	// Generate SQL to select all columns in LandingPage Table
	// Then generate filtering and ordering sql and finally run query.

	baseSQL := "SELECT id, created_at, updated_at, title, cause, layout FROM landing_page"
	querySQL := queries.BuildLandingPageFilters(req.Filters, req.Ordering, req.Limit)
	SQL := fmt.Sprintf("%s %s", baseSQL, querySQL)
	rows, err := c.QueryContext(ctx, SQL)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from LandingPage-> "+err.Error())
	}
	defer rows.Close()

	// Variables to store results returned by database.
	list := []*v1.LandingPage{}
	var createdAt time.Time
	var updatedAt time.Time

	for rows.Next() {
		landingpage := new(v1.LandingPage)
		if err := rows.Scan(&landingpage.ID, &createdAt, &updatedAt, &landingpage.Title, &landingpage.Cause, &landingpage.Layout); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from LandingPage row-> "+err.Error())
		}
		// Convert time.Time from database into proto timestamp.
		landingpage.CreatedAt, err = ptypes.TimestampProto(createdAt)
		if err != nil {
			return nil, status.Error(codes.Unknown, "createdAt field has invalid format-> "+err.Error())
		}
		landingpage.UpdatedAt, err = ptypes.TimestampProto(updatedAt)
		if err != nil {
			return nil, status.Error(codes.Unknown, "updatedAt field has invalid format-> "+err.Error())
		}

		list = append(list, landingpage)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from LandingPage-> "+err.Error())
	}

	return &v1.ListLandingPageResponse{
		Api:   apiVersion,
		Items: list,
	}, nil
}

// Update LandingPage
func (s *shrikeServiceServer) UpdateLandingPage(ctx context.Context, req *v1.UpdateLandingPageRequest) (*v1.UpdateLandingPageResponse, error) {
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

	// update landing_page
	res, err := c.ExecContext(ctx, "UPDATE landing_page SET title=$2, cause=$3, layout=$4 WHERE id=$1",
		req.Item.ID, req.Item.Title, req.Item.Cause, req.Item.Layout)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update LandingPage-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("LandingPage with ID='%s' is not found",
			req.Item.ID))
	}

	return &v1.UpdateLandingPageResponse{
		Api:     apiVersion,
		Updated: rows,
	}, nil
}

// Delete landing_page
func (s *shrikeServiceServer) DeleteLandingPage(ctx context.Context, req *v1.DeleteLandingPageRequest) (*v1.DeleteLandingPageResponse, error) {
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

	// delete landing_page
	res, err := c.ExecContext(ctx, "DELETE FROM landing_page WHERE id=$1", req.ID)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete LandingPage-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("LandingPage with ID='%s' is not found",
			req.ID))
	}

	return &v1.DeleteLandingPageResponse{
		Api:     apiVersion,
		Deleted: rows,
	}, nil
}
