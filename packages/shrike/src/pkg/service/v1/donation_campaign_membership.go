package v1

import (
	"context"

	v1 "github.com/SteveCastle/openmob/packages/shrike/src/pkg/api/v1"
	"github.com/SteveCastle/openmob/packages/shrike/src/pkg/models/v1"
)

// Create new DonationCampaignMembership
func (s *shrikeServiceServer) CreateDonationCampaignMembership(ctx context.Context, req *v1.CreateDonationCampaignMembershipRequest) (*v1.CreateDonationCampaignMembershipResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a DonationCampaignMembership Manager
	m := models.NewDonationCampaignMembershipManager(s.db)

	// Get a list of donationCampaignMemberships given filters, ordering, and limit rules.
	id, err := m.CreateDonationCampaignMembership(ctx, req.Item)
	if err != nil {
		return nil, err
	}
	return &v1.CreateDonationCampaignMembershipResponse{
		Api: apiVersion,
		ID:  *id,
	}, nil
}

// Get donationCampaignMembership by id.
func (s *shrikeServiceServer) GetDonationCampaignMembership(ctx context.Context, req *v1.GetDonationCampaignMembershipRequest) (*v1.GetDonationCampaignMembershipResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a DonationCampaignMembership Manager
	m := models.NewDonationCampaignMembershipManager(s.db)

	// Get a list of donationCampaignMemberships given filters, ordering, and limit rules.
	donationCampaignMembership, err := m.GetDonationCampaignMembership(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	return &v1.GetDonationCampaignMembershipResponse{
		Api:  apiVersion,
		Item: m.GetProto(donationCampaignMembership),
	}, nil

}

// Read all DonationCampaignMembership
func (s *shrikeServiceServer) ListDonationCampaignMembership(ctx context.Context, req *v1.ListDonationCampaignMembershipRequest) (*v1.ListDonationCampaignMembershipResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}

	// Create a DonationCampaignMembership Manager
	m := models.NewDonationCampaignMembershipManager(s.db)

	// Get a list of donationCampaignMemberships given filters, ordering, and limit rules.
	list, err := m.ListDonationCampaignMembership(ctx, req.Filters, req.Ordering, req.Limit)
	if err != nil {
		return nil, err
	}

	return &v1.ListDonationCampaignMembershipResponse{
		Api:   apiVersion,
		Items: m.GetProtoList(list),
	}, nil
}

// Update DonationCampaignMembership
func (s *shrikeServiceServer) UpdateDonationCampaignMembership(ctx context.Context, req *v1.UpdateDonationCampaignMembershipRequest) (*v1.UpdateDonationCampaignMembershipResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a DonationCampaignMembership Manager
	m := models.NewDonationCampaignMembershipManager(s.db)

	// Get a list of donationCampaignMemberships given filters, ordering, and limit rules.
	rows, err := m.UpdateDonationCampaignMembership(ctx, req.Item)
	if err != nil {
		return nil, err
	}

	return &v1.UpdateDonationCampaignMembershipResponse{
		Api:     apiVersion,
		Updated: *rows,
	}, nil
}

// Delete donationCampaignMembership
func (s *shrikeServiceServer) DeleteDonationCampaignMembership(ctx context.Context, req *v1.DeleteDonationCampaignMembershipRequest) (*v1.DeleteDonationCampaignMembershipResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a DonationCampaignMembership Manager
	m := models.NewDonationCampaignMembershipManager(s.db)

	// Get a list of donationCampaignMemberships given filters, ordering, and limit rules.
	rows, err := m.DeleteDonationCampaignMembership(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	return &v1.DeleteDonationCampaignMembershipResponse{
		Api:     apiVersion,
		Deleted: *rows,
	}, nil
}
