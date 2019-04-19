package v1

import (
	"context"

	v1 "github.com/SteveCastle/openmob/packages/shrike/src/pkg/api/v1"
	"github.com/SteveCastle/openmob/packages/shrike/src/pkg/models/v1"
)

// Create new LiveEventMembership
func (s *shrikeServiceServer) CreateLiveEventMembership(ctx context.Context, req *v1.CreateLiveEventMembershipRequest) (*v1.CreateLiveEventMembershipResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a LiveEventMembership Manager
	m := models.NewLiveEventMembershipManager(s.db)

	// Get a list of liveEventMemberships given filters, ordering, and limit rules.
	id, err := m.Create(ctx, req.Item)
	if err != nil {
		return nil, err
	}
	return &v1.CreateLiveEventMembershipResponse{
		Api: apiVersion,
		ID:  *id,
	}, nil
}

// Get liveEventMembership by id.
func (s *shrikeServiceServer) GetLiveEventMembership(ctx context.Context, req *v1.GetLiveEventMembershipRequest) (*v1.GetLiveEventMembershipResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a LiveEventMembership Manager
	m := models.NewLiveEventMembershipManager(s.db)

	// Get a list of liveEventMemberships given filters, ordering, and limit rules.
	liveEventMembership, err := m.Get(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	return &v1.GetLiveEventMembershipResponse{
		Api:  apiVersion,
		Item: m.GetProto(liveEventMembership),
	}, nil

}

// Read all LiveEventMembership
func (s *shrikeServiceServer) ListLiveEventMembership(ctx context.Context, req *v1.ListLiveEventMembershipRequest) (*v1.ListLiveEventMembershipResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}

	// Create a LiveEventMembership Manager
	m := models.NewLiveEventMembershipManager(s.db)

	// Get a list of liveEventMemberships given filters, ordering, and limit rules.
	list, err := m.List(ctx, req.Filters, req.Ordering, req.Limit)
	if err != nil {
		return nil, err
	}

	return &v1.ListLiveEventMembershipResponse{
		Api:   apiVersion,
		Items: m.GetProtoList(list),
	}, nil
}

// Update LiveEventMembership
func (s *shrikeServiceServer) UpdateLiveEventMembership(ctx context.Context, req *v1.UpdateLiveEventMembershipRequest) (*v1.UpdateLiveEventMembershipResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a LiveEventMembership Manager
	m := models.NewLiveEventMembershipManager(s.db)

	// Get a list of liveEventMemberships given filters, ordering, and limit rules.
	rows, err := m.Update(ctx, req.Item)
	if err != nil {
		return nil, err
	}

	return &v1.UpdateLiveEventMembershipResponse{
		Api:     apiVersion,
		Updated: *rows,
	}, nil
}

// Delete liveEventMembership
func (s *shrikeServiceServer) DeleteLiveEventMembership(ctx context.Context, req *v1.DeleteLiveEventMembershipRequest) (*v1.DeleteLiveEventMembershipResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a LiveEventMembership Manager
	m := models.NewLiveEventMembershipManager(s.db)

	// Get a list of liveEventMemberships given filters, ordering, and limit rules.
	rows, err := m.Delete(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	return &v1.DeleteLiveEventMembershipResponse{
		Api:     apiVersion,
		Deleted: *rows,
	}, nil
}
