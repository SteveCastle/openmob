package v1

import (
	"context"
	"fmt"

	v1 "github.com/SteveCastle/openmob/packages/shrike/src/pkg/api/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)


// Create new DonationCampaignMembership
func (s *shrikeServiceServer) CreateDonationCampaignMembership(ctx context.Context, req *v1.CreateDonationCampaignMembershipRequest) (*v1.CreateDonationCampaignMembershipResponse, error) {
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
	// insert DonationCampaignMembership entity data
	err = c.QueryRowContext(ctx, "INSERT INTO donation_campaign_membership (id, created_at, updated_at, cause, donation_campaign, ) VALUES($1, $2, $3, $4, $5, )  RETURNING id;",
		 req.Item.ID,  req.Item.CreatedAt,  req.Item.UpdatedAt,  req.Item.Cause,  req.Item.DonationCampaign, ).Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into DonationCampaignMembership-> "+err.Error())
	}

	// get ID of creates DonationCampaignMembership
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created DonationCampaignMembership-> "+err.Error())
	}

	return &v1.CreateDonationCampaignMembershipResponse{
		Api: apiVersion,
		ID:  id,
	}, nil
}

// Get donation_campaign_membership by id.
func (s *shrikeServiceServer) GetDonationCampaignMembership(ctx context.Context, req *v1.GetDonationCampaignMembershipRequest) (*v1.GetDonationCampaignMembershipResponse, error) {
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

	// query DonationCampaignMembership by ID
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, cause, donation_campaign,  FROM donation_campaign_membership WHERE id=$1",
		req.ID)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from DonationCampaignMembership-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from DonationCampaignMembership-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("DonationCampaignMembership with ID='%d' is not found",
			req.ID))
	}

	// get DonationCampaignMembership data
	var donationcampaignmembership v1.DonationCampaignMembership
	if err := rows.Scan( &donationcampaignmembership.ID,  &donationcampaignmembership.CreatedAt,  &donationcampaignmembership.UpdatedAt,  &donationcampaignmembership.Cause,  &donationcampaignmembership.DonationCampaign, ); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from DonationCampaignMembership row-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple DonationCampaignMembership rows with ID='%d'",
			req.ID))
	}

	return &v1.GetDonationCampaignMembershipResponse{
		Api:  apiVersion,
		Item: &donationcampaignmembership,
	}, nil

}

// Read all DonationCampaignMembership
func (s *shrikeServiceServer) ListDonationCampaignMembership(ctx context.Context, req *v1.ListDonationCampaignMembershipRequest) (*v1.ListDonationCampaignMembershipResponse, error) {
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

	// get DonationCampaignMembership list
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, cause, donation_campaign,  FROM donation_campaign_membership")
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from DonationCampaignMembership-> "+err.Error())
	}
	defer rows.Close()

	list := []*v1.DonationCampaignMembership{}
	for rows.Next() {
		donationcampaignmembership := new(v1.DonationCampaignMembership)
		if err := rows.Scan( &donationcampaignmembership.ID,  &donationcampaignmembership.CreatedAt,  &donationcampaignmembership.UpdatedAt,  &donationcampaignmembership.Cause,  &donationcampaignmembership.DonationCampaign, ); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from DonationCampaignMembership row-> "+err.Error())
		}
		list = append(list, donationcampaignmembership)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from DonationCampaignMembership-> "+err.Error())
	}

	return &v1.ListDonationCampaignMembershipResponse{
		Api:   apiVersion,
		Items: list,
	}, nil
}

// Update DonationCampaignMembership
func (s *shrikeServiceServer) UpdateDonationCampaignMembership(ctx context.Context, req *v1.UpdateDonationCampaignMembershipRequest) (*v1.UpdateDonationCampaignMembershipResponse, error) {
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

	// update donation_campaign_membership
	res, err := c.ExecContext(ctx, "UPDATE donation_campaign_membership SET $1, $2, $3, $4, $5,  WHERE id=$1",
		req.Item.ID,req.Item.CreatedAt,req.Item.UpdatedAt,req.Item.Cause,req.Item.DonationCampaign, )
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update DonationCampaignMembership-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("DonationCampaignMembership with ID='%d' is not found",
			req.Item.ID))
	}

	return &v1.UpdateDonationCampaignMembershipResponse{
		Api:     apiVersion,
		Updated: rows,
	}, nil
}

// Delete donation_campaign_membership
func (s *shrikeServiceServer) DeleteDonationCampaignMembership(ctx context.Context, req *v1.DeleteDonationCampaignMembershipRequest) (*v1.DeleteDonationCampaignMembershipResponse, error) {
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

	// delete donation_campaign_membership
	res, err := c.ExecContext(ctx, "DELETE FROM donation_campaign_membership WHERE id=$1", req.ID)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete DonationCampaignMembership-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("DonationCampaignMembership with ID='%d' is not found",
			req.ID))
	}

	return &v1.DeleteDonationCampaignMembershipResponse{
		Api:     apiVersion,
		Deleted: rows,
	}, nil
}
