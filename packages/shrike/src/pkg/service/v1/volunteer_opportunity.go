package v1

import (
	"context"
	"fmt"

	v1 "github.com/SteveCastle/openmob/packages/shrike/src/pkg/api/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Create new VolunteerOpportunity
func (s *shrikeServiceServer) CreateVolunteerOpportunity(ctx context.Context, req *v1.CreateVolunteerOpportunityRequest) (*v1.CreateVolunteerOpportunityResponse, error) {
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
	var id int64
	// insert VolunteerOpportunity entity data
	err = c.QueryRowContext(ctx, "INSERT INTO volunteer_opportunity (id, created_at, updated_at, title, election_type) VALUES($1, $2, $3, $4, $5)  RETURNING id;",
		 req.Item.ID,  req.Item.CreatedAt,  req.Item.UpdatedAt,  req.Item.Title,  req.Item.ElectionType, ).Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into VolunteerOpportunity-> "+err.Error())
	}

	// get ID of creates VolunteerOpportunity
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created VolunteerOpportunity-> "+err.Error())
	}

	return &v1.CreateVolunteerOpportunityResponse{
		Api: apiVersion,
		ID:  id,
	}, nil
}

// Get volunteer_opportunity by id.
func (s *shrikeServiceServer) GetVolunteerOpportunity(ctx context.Context, req *v1.GetVolunteerOpportunityRequest) (*v1.GetVolunteerOpportunityResponse, error) {
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

	// query VolunteerOpportunity by ID
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, title, election_type FROM volunteer_opportunity WHERE id=$1",
		req.ID)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from VolunteerOpportunity-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from VolunteerOpportunity-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("VolunteerOpportunity with ID='%d' is not found",
			req.ID))
	}

	// get VolunteerOpportunity data
	var volunteeropportunity v1.VolunteerOpportunity
	if err := rows.Scan( &volunteeropportunity.ID,  &volunteeropportunity.CreatedAt,  &volunteeropportunity.UpdatedAt,  &volunteeropportunity.Title,  &volunteeropportunity.ElectionType, ); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from VolunteerOpportunity row-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple VolunteerOpportunity rows with ID='%d'",
			req.ID))
	}

	return &v1.GetVolunteerOpportunityResponse{
		Api:  apiVersion,
		Item: &volunteeropportunity,
	}, nil

}

// Read all VolunteerOpportunity
func (s *shrikeServiceServer) ListVolunteerOpportunity(ctx context.Context, req *v1.ListVolunteerOpportunityRequest) (*v1.ListVolunteerOpportunityResponse, error) {
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

	// get VolunteerOpportunity list
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, title, election_type FROM volunteer_opportunity")
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from VolunteerOpportunity-> "+err.Error())
	}
	defer rows.Close()

	list := []*v1.VolunteerOpportunity{}
	for rows.Next() {
		volunteeropportunity := new(v1.VolunteerOpportunity)
		if err := rows.Scan( &volunteeropportunity.ID,  &volunteeropportunity.CreatedAt,  &volunteeropportunity.UpdatedAt,  &volunteeropportunity.Title,  &volunteeropportunity.ElectionType, ); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from VolunteerOpportunity row-> "+err.Error())
		}
		list = append(list, volunteeropportunity)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from VolunteerOpportunity-> "+err.Error())
	}

	return &v1.ListVolunteerOpportunityResponse{
		Api:   apiVersion,
		Items: list,
	}, nil
}

// Update VolunteerOpportunity
func (s *shrikeServiceServer) UpdateVolunteerOpportunity(ctx context.Context, req *v1.UpdateVolunteerOpportunityRequest) (*v1.UpdateVolunteerOpportunityResponse, error) {
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

	// update volunteer_opportunity
	res, err := c.ExecContext(ctx, "UPDATE volunteer_opportunity SET id=$1, created_at=$2, updated_at=$3, title=$4, election_type=$5 WHERE id=$1",
		req.Item.ID,req.Item.CreatedAt,req.Item.UpdatedAt,req.Item.Title,req.Item.ElectionType, )
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update VolunteerOpportunity-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("VolunteerOpportunity with ID='%d' is not found",
			req.Item.ID))
	}

	return &v1.UpdateVolunteerOpportunityResponse{
		Api:     apiVersion,
		Updated: rows,
	}, nil
}

// Delete volunteer_opportunity
func (s *shrikeServiceServer) DeleteVolunteerOpportunity(ctx context.Context, req *v1.DeleteVolunteerOpportunityRequest) (*v1.DeleteVolunteerOpportunityResponse, error) {
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

	// delete volunteer_opportunity
	res, err := c.ExecContext(ctx, "DELETE FROM volunteer_opportunity WHERE id=$1", req.ID)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete VolunteerOpportunity-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("VolunteerOpportunity with ID='%d' is not found",
			req.ID))
	}

	return &v1.DeleteVolunteerOpportunityResponse{
		Api:     apiVersion,
		Deleted: rows,
	}, nil
}
