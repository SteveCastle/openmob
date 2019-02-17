package v1

import (
	"context"
	"fmt"

	v1 "github.com/SteveCastle/openmob/packages/shrike/src/pkg/api/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)


// Create new VolunteerOpportunityMembership
func (s *shrikeServiceServer) CreateVolunteerOpportunityMembership(ctx context.Context, req *v1.CreateVolunteerOpportunityMembershipRequest) (*v1.CreateVolunteerOpportunityMembershipResponse, error) {
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
	// insert VolunteerOpportunityMembership entity data
	err = c.QueryRowContext(ctx, "INSERT INTO volunteer_opportunity_membership (id, created_at, updated_at, cause, volunteer_opportunity, ) VALUES($1, $2, $3, $4, $5, )  RETURNING id;",
		 req.Item.ID,  req.Item.CreatedAt,  req.Item.UpdatedAt,  req.Item.Cause,  req.Item.VolunteerOpportunity, ).Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into VolunteerOpportunityMembership-> "+err.Error())
	}

	// get ID of creates VolunteerOpportunityMembership
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created VolunteerOpportunityMembership-> "+err.Error())
	}

	return &v1.CreateVolunteerOpportunityMembershipResponse{
		Api: apiVersion,
		ID:  id,
	}, nil
}

// Get volunteer_opportunity_membership by id.
func (s *shrikeServiceServer) GetVolunteerOpportunityMembership(ctx context.Context, req *v1.GetVolunteerOpportunityMembershipRequest) (*v1.GetVolunteerOpportunityMembershipResponse, error) {
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

	// query VolunteerOpportunityMembership by ID
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, cause, volunteer_opportunity,  FROM volunteer_opportunity_membership WHERE id=$1",
		req.ID)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from VolunteerOpportunityMembership-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from VolunteerOpportunityMembership-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("VolunteerOpportunityMembership with ID='%d' is not found",
			req.ID))
	}

	// get VolunteerOpportunityMembership data
	var volunteeropportunitymembership v1.VolunteerOpportunityMembership
	if err := rows.Scan( &volunteeropportunitymembership.ID,  &volunteeropportunitymembership.CreatedAt,  &volunteeropportunitymembership.UpdatedAt,  &volunteeropportunitymembership.Cause,  &volunteeropportunitymembership.VolunteerOpportunity, ); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from VolunteerOpportunityMembership row-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple VolunteerOpportunityMembership rows with ID='%d'",
			req.ID))
	}

	return &v1.GetVolunteerOpportunityMembershipResponse{
		Api:  apiVersion,
		Item: &volunteeropportunitymembership,
	}, nil

}

// Read all VolunteerOpportunityMembership
func (s *shrikeServiceServer) ListVolunteerOpportunityMembership(ctx context.Context, req *v1.ListVolunteerOpportunityMembershipRequest) (*v1.ListVolunteerOpportunityMembershipResponse, error) {
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

	// get VolunteerOpportunityMembership list
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, cause, volunteer_opportunity,  FROM volunteer_opportunity_membership")
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from VolunteerOpportunityMembership-> "+err.Error())
	}
	defer rows.Close()

	list := []*v1.VolunteerOpportunityMembership{}
	for rows.Next() {
		volunteeropportunitymembership := new(v1.VolunteerOpportunityMembership)
		if err := rows.Scan( &volunteeropportunitymembership.ID,  &volunteeropportunitymembership.CreatedAt,  &volunteeropportunitymembership.UpdatedAt,  &volunteeropportunitymembership.Cause,  &volunteeropportunitymembership.VolunteerOpportunity, ); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from VolunteerOpportunityMembership row-> "+err.Error())
		}
		list = append(list, volunteeropportunitymembership)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from VolunteerOpportunityMembership-> "+err.Error())
	}

	return &v1.ListVolunteerOpportunityMembershipResponse{
		Api:   apiVersion,
		Items: list,
	}, nil
}

// Update VolunteerOpportunityMembership
func (s *shrikeServiceServer) UpdateVolunteerOpportunityMembership(ctx context.Context, req *v1.UpdateVolunteerOpportunityMembershipRequest) (*v1.UpdateVolunteerOpportunityMembershipResponse, error) {
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

	// update volunteer_opportunity_membership
	res, err := c.ExecContext(ctx, "UPDATE volunteer_opportunity_membership SET $1, $2, $3, $4, $5,  WHERE id=$1",
		req.Item.ID,req.Item.CreatedAt,req.Item.UpdatedAt,req.Item.Cause,req.Item.VolunteerOpportunity, )
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update VolunteerOpportunityMembership-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("VolunteerOpportunityMembership with ID='%d' is not found",
			req.Item.ID))
	}

	return &v1.UpdateVolunteerOpportunityMembershipResponse{
		Api:     apiVersion,
		Updated: rows,
	}, nil
}

// Delete volunteer_opportunity_membership
func (s *shrikeServiceServer) DeleteVolunteerOpportunityMembership(ctx context.Context, req *v1.DeleteVolunteerOpportunityMembershipRequest) (*v1.DeleteVolunteerOpportunityMembershipResponse, error) {
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

	// delete volunteer_opportunity_membership
	res, err := c.ExecContext(ctx, "DELETE FROM volunteer_opportunity_membership WHERE id=$1", req.ID)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete VolunteerOpportunityMembership-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("VolunteerOpportunityMembership with ID='%d' is not found",
			req.ID))
	}

	return &v1.DeleteVolunteerOpportunityMembershipResponse{
		Api:     apiVersion,
		Deleted: rows,
	}, nil
}
