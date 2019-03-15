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

// Create new VolunteerOpportunityType
func (s *shrikeServiceServer) CreateVolunteerOpportunityType(ctx context.Context, req *v1.CreateVolunteerOpportunityTypeRequest) (*v1.CreateVolunteerOpportunityTypeResponse, error) {
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
	// insert VolunteerOpportunityType entity data
	err = c.QueryRowContext(ctx, "INSERT INTO volunteer_opportunity_type (title) VALUES($1)  RETURNING id;",
		req.Item.Title).Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into VolunteerOpportunityType-> "+err.Error())
	}

	// get ID of creates VolunteerOpportunityType
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created VolunteerOpportunityType-> "+err.Error())
	}

	return &v1.CreateVolunteerOpportunityTypeResponse{
		Api: apiVersion,
		ID:  id,
	}, nil
}

// Get volunteer_opportunity_type by id.
func (s *shrikeServiceServer) GetVolunteerOpportunityType(ctx context.Context, req *v1.GetVolunteerOpportunityTypeRequest) (*v1.GetVolunteerOpportunityTypeResponse, error) {
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

	// query VolunteerOpportunityType by ID
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, title FROM volunteer_opportunity_type WHERE id=$1",
		req.ID)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from VolunteerOpportunityType-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from VolunteerOpportunityType-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("VolunteerOpportunityType with ID='%s' is not found",
			req.ID))
	}

	// scan VolunteerOpportunityType data into protobuf model
	var volunteeropportunitytype v1.VolunteerOpportunityType
	var createdAt time.Time
	var updatedAt time.Time

	if err := rows.Scan(&volunteeropportunitytype.ID, &createdAt, &updatedAt, &volunteeropportunitytype.Title); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from VolunteerOpportunityType row-> "+err.Error())
	}

	// Convert time.Time from database into proto timestamp.
	volunteeropportunitytype.CreatedAt, err = ptypes.TimestampProto(createdAt)
	if err != nil {
		return nil, status.Error(codes.Unknown, "createdAt field has invalid format-> "+err.Error())
	}
	volunteeropportunitytype.UpdatedAt, err = ptypes.TimestampProto(updatedAt)
	if err != nil {
		return nil, status.Error(codes.Unknown, "updatedAt field has invalid format-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple VolunteerOpportunityType rows with ID='%s'",
			req.ID))
	}

	return &v1.GetVolunteerOpportunityTypeResponse{
		Api:  apiVersion,
		Item: &volunteeropportunitytype,
	}, nil

}

// Read all VolunteerOpportunityType
func (s *shrikeServiceServer) ListVolunteerOpportunityType(ctx context.Context, req *v1.ListVolunteerOpportunityTypeRequest) (*v1.ListVolunteerOpportunityTypeResponse, error) {
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

	// Generate SQL to select all columns in VolunteerOpportunityType Table
	// Then generate filtering and ordering sql and finally run query.

	baseSQL := "SELECT id, created_at, updated_at, title FROM volunteer_opportunity_type"
	querySQL := queries.BuildVolunteerOpportunityTypeFilters(req.Filters, req.Ordering, req.Limit)
	SQL := fmt.Sprintf("%s %s", baseSQL, querySQL)
	rows, err := c.QueryContext(ctx, SQL)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from VolunteerOpportunityType-> "+err.Error())
	}
	defer rows.Close()

	// Variables to store results returned by database.
	list := []*v1.VolunteerOpportunityType{}
	var createdAt time.Time
	var updatedAt time.Time

	for rows.Next() {
		volunteeropportunitytype := new(v1.VolunteerOpportunityType)
		if err := rows.Scan(&volunteeropportunitytype.ID, &createdAt, &updatedAt, &volunteeropportunitytype.Title); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from VolunteerOpportunityType row-> "+err.Error())
		}
		// Convert time.Time from database into proto timestamp.
		volunteeropportunitytype.CreatedAt, err = ptypes.TimestampProto(createdAt)
		if err != nil {
			return nil, status.Error(codes.Unknown, "createdAt field has invalid format-> "+err.Error())
		}
		volunteeropportunitytype.UpdatedAt, err = ptypes.TimestampProto(updatedAt)
		if err != nil {
			return nil, status.Error(codes.Unknown, "updatedAt field has invalid format-> "+err.Error())
		}

		list = append(list, volunteeropportunitytype)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from VolunteerOpportunityType-> "+err.Error())
	}

	return &v1.ListVolunteerOpportunityTypeResponse{
		Api:   apiVersion,
		Items: list,
	}, nil
}

// Update VolunteerOpportunityType
func (s *shrikeServiceServer) UpdateVolunteerOpportunityType(ctx context.Context, req *v1.UpdateVolunteerOpportunityTypeRequest) (*v1.UpdateVolunteerOpportunityTypeResponse, error) {
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

	// update volunteer_opportunity_type
	res, err := c.ExecContext(ctx, "UPDATE volunteer_opportunity_type SET title=$2 WHERE id=$1",
		req.Item.ID, req.Item.Title)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update VolunteerOpportunityType-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("VolunteerOpportunityType with ID='%s' is not found",
			req.Item.ID))
	}

	return &v1.UpdateVolunteerOpportunityTypeResponse{
		Api:     apiVersion,
		Updated: rows,
	}, nil
}

// Delete volunteer_opportunity_type
func (s *shrikeServiceServer) DeleteVolunteerOpportunityType(ctx context.Context, req *v1.DeleteVolunteerOpportunityTypeRequest) (*v1.DeleteVolunteerOpportunityTypeResponse, error) {
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

	// delete volunteer_opportunity_type
	res, err := c.ExecContext(ctx, "DELETE FROM volunteer_opportunity_type WHERE id=$1", req.ID)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete VolunteerOpportunityType-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("VolunteerOpportunityType with ID='%s' is not found",
			req.ID))
	}

	return &v1.DeleteVolunteerOpportunityTypeResponse{
		Api:     apiVersion,
		Deleted: rows,
	}, nil
}
