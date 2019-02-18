package v1

import (
	"context"
	"fmt"

	v1 "github.com/SteveCastle/openmob/packages/shrike/src/pkg/api/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Create new DonationCampaign
func (s *shrikeServiceServer) CreateDonationCampaign(ctx context.Context, req *v1.CreateDonationCampaignRequest) (*v1.CreateDonationCampaignResponse, error) {
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
	// insert DonationCampaign entity data
	err = c.QueryRowContext(ctx, "INSERT INTO donation_campaign (id, created_at, updated_at, title) VALUES($1, $2, $3, $4)  RETURNING id;",
		 req.Item.ID,  req.Item.CreatedAt,  req.Item.UpdatedAt,  req.Item.Title, ).Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into DonationCampaign-> "+err.Error())
	}

	// get ID of creates DonationCampaign
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created DonationCampaign-> "+err.Error())
	}

	return &v1.CreateDonationCampaignResponse{
		Api: apiVersion,
		ID:  id,
	}, nil
}

// Get donation_campaign by id.
func (s *shrikeServiceServer) GetDonationCampaign(ctx context.Context, req *v1.GetDonationCampaignRequest) (*v1.GetDonationCampaignResponse, error) {
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

	// query DonationCampaign by ID
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, title FROM donation_campaign WHERE id=$1",
		req.ID)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from DonationCampaign-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from DonationCampaign-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("DonationCampaign with ID='%d' is not found",
			req.ID))
	}

	// get DonationCampaign data
	var donationcampaign v1.DonationCampaign
	if err := rows.Scan( &donationcampaign.ID,  &donationcampaign.CreatedAt,  &donationcampaign.UpdatedAt,  &donationcampaign.Title, ); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from DonationCampaign row-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple DonationCampaign rows with ID='%d'",
			req.ID))
	}

	return &v1.GetDonationCampaignResponse{
		Api:  apiVersion,
		Item: &donationcampaign,
	}, nil

}

// Read all DonationCampaign
func (s *shrikeServiceServer) ListDonationCampaign(ctx context.Context, req *v1.ListDonationCampaignRequest) (*v1.ListDonationCampaignResponse, error) {
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

	// get DonationCampaign list
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, title FROM donation_campaign")
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from DonationCampaign-> "+err.Error())
	}
	defer rows.Close()

	list := []*v1.DonationCampaign{}
	for rows.Next() {
		donationcampaign := new(v1.DonationCampaign)
		if err := rows.Scan( &donationcampaign.ID,  &donationcampaign.CreatedAt,  &donationcampaign.UpdatedAt,  &donationcampaign.Title, ); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from DonationCampaign row-> "+err.Error())
		}
		list = append(list, donationcampaign)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from DonationCampaign-> "+err.Error())
	}

	return &v1.ListDonationCampaignResponse{
		Api:   apiVersion,
		Items: list,
	}, nil
}

// Update DonationCampaign
func (s *shrikeServiceServer) UpdateDonationCampaign(ctx context.Context, req *v1.UpdateDonationCampaignRequest) (*v1.UpdateDonationCampaignResponse, error) {
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

	// update donation_campaign
	res, err := c.ExecContext(ctx, "UPDATE donation_campaign SET $1 ,$2 ,$3 ,$4  WHERE id=$1",
		req.Item.ID,req.Item.CreatedAt,req.Item.UpdatedAt,req.Item.Title, )
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update DonationCampaign-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("DonationCampaign with ID='%d' is not found",
			req.Item.ID))
	}

	return &v1.UpdateDonationCampaignResponse{
		Api:     apiVersion,
		Updated: rows,
	}, nil
}

// Delete donation_campaign
func (s *shrikeServiceServer) DeleteDonationCampaign(ctx context.Context, req *v1.DeleteDonationCampaignRequest) (*v1.DeleteDonationCampaignResponse, error) {
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

	// delete donation_campaign
	res, err := c.ExecContext(ctx, "DELETE FROM donation_campaign WHERE id=$1", req.ID)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete DonationCampaign-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("DonationCampaign with ID='%d' is not found",
			req.ID))
	}

	return &v1.DeleteDonationCampaignResponse{
		Api:     apiVersion,
		Deleted: rows,
	}, nil
}
