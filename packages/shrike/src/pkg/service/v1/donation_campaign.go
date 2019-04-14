package v1

import (
	"context"

	v1 "github.com/SteveCastle/openmob/packages/shrike/src/pkg/api/v1"
	"github.com/SteveCastle/openmob/packages/shrike/src/pkg/models/v1"
)

// Create new DonationCampaign
func (s *shrikeServiceServer) CreateDonationCampaign(ctx context.Context, req *v1.CreateDonationCampaignRequest) (*v1.CreateDonationCampaignResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a DonationCampaign Manager
	m := models.NewDonationCampaignManager(s.db)

	// Get a list of donationCampaigns given filters, ordering, and limit rules.
	id, err := m.CreateDonationCampaign(ctx, req.Item)
	if err != nil {
		return nil, err
	}
	return &v1.CreateDonationCampaignResponse{
		Api: apiVersion,
		ID:  *id,
	}, nil
}

// Get donationCampaign by id.
func (s *shrikeServiceServer) GetDonationCampaign(ctx context.Context, req *v1.GetDonationCampaignRequest) (*v1.GetDonationCampaignResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a DonationCampaign Manager
	m := models.NewDonationCampaignManager(s.db)

	// Get a list of donationCampaigns given filters, ordering, and limit rules.
	donationCampaign, err := m.GetDonationCampaign(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	return &v1.GetDonationCampaignResponse{
		Api:  apiVersion,
		Item: m.GetProto(donationCampaign),
	}, nil

}

// Read all DonationCampaign
func (s *shrikeServiceServer) ListDonationCampaign(ctx context.Context, req *v1.ListDonationCampaignRequest) (*v1.ListDonationCampaignResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}

	// Create a DonationCampaign Manager
	m := models.NewDonationCampaignManager(s.db)

	// Get a list of donationCampaigns given filters, ordering, and limit rules.
	list, err := m.ListDonationCampaign(ctx, req.Filters, req.Ordering, req.Limit)
	if err != nil {
		return nil, err
	}

	return &v1.ListDonationCampaignResponse{
		Api:   apiVersion,
		Items: m.GetProtoList(list),
	}, nil
}

// Update DonationCampaign
func (s *shrikeServiceServer) UpdateDonationCampaign(ctx context.Context, req *v1.UpdateDonationCampaignRequest) (*v1.UpdateDonationCampaignResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a DonationCampaign Manager
	m := models.NewDonationCampaignManager(s.db)

	// Get a list of donationCampaigns given filters, ordering, and limit rules.
	rows, err := m.UpdateDonationCampaign(ctx, req.Item)
	if err != nil {
		return nil, err
	}

	return &v1.UpdateDonationCampaignResponse{
		Api:     apiVersion,
		Updated: *rows,
	}, nil
}

// Delete donationCampaign
func (s *shrikeServiceServer) DeleteDonationCampaign(ctx context.Context, req *v1.DeleteDonationCampaignRequest) (*v1.DeleteDonationCampaignResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a DonationCampaign Manager
	m := models.NewDonationCampaignManager(s.db)

	// Get a list of donationCampaigns given filters, ordering, and limit rules.
	rows, err := m.DeleteDonationCampaign(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	return &v1.DeleteDonationCampaignResponse{
		Api:     apiVersion,
		Deleted: *rows,
	}, nil
}
