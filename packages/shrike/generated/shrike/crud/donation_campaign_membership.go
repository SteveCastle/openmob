package v1

import (
	"context"
	"database/sql"
	"fmt"

	v1 "github.com/SteveCastle/openmob/packages/shrike/src/pkg/api/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	// apiVersion is version of API is provided by server
	apiVersion = "v1"
)

// shrikeServiceServer is implementation of v1.ShrikeServiceServer proto interface
type shrikeServiceServer struct {
	db *sql.DB
}

// NewShrikeServiceServer creates DonationCampaignMembership service
func NewShrikeServiceServer(db *sql.DB) v1.ShrikeServiceServer {
	return &shrikeServiceServer{db: db}
}

// checkAPI checks if the API version requested by client is supported by server
func (s *shrikeServiceServer) checkAPI(api string) error {
	// API version is "" means use current version of the service
	if len(api) > 0 {
		if apiVersion != api {
			return status.Errorf(codes.Unimplemented,
				"unsupported API version: service implements API version '%s', but asked for '%s'", apiVersion, api)
		}
	}
	return nil
}

// connect returns SQL database connection from the pool
func (s *shrikeServiceServer) connect(ctx context.Context) (*sql.Conn, error) {
	c, err := s.db.Conn(ctx)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to connect to database-> "+err.Error())
	}
	return c, nil
}

// Create new todo task
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
	err = c.QueryRowContext(ctx, "INSERT INTO donation_campaign_membership (title) VALUES($1)  RETURNING id;",
		req.Item.Title).Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into DonationCampaignMembership-> "+err.Error())
	}

	// get ID of creates DonationCampaignMembership
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created DonationCampaignMembership-> "+err.Error())
	}

	return &v1.CreateDonationCampaignMembershipResponse{
		Api: apiVersion,
		Id:  id,
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
	rows, err := c.QueryContext(ctx, "SELECT id, title FROM donation_campaign_membership WHERE id=$1",
		req.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from DonationCampaignMembership-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from DonationCampaignMembership-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("DonationCampaignMembership with ID='%d' is not found",
			req.Id))
	}

	// get DonationCampaignMembership data
	var td v1.DonationCampaignMembership
	if err := rows.Scan(&td.Id, &td.Title); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from DonationCampaignMembership row-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple DonationCampaignMembership rows with ID='%d'",
			req.Id))
	}

	return &v1.GetDonationCampaignMembershipResponse{
		Api:  apiVersion,
		Item: &td,
	}, nil

}

// Read all todo tasks
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
	rows, err := c.QueryContext(ctx, "SELECT id,title FROM donation_campaign_membership")
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from DonationCampaignMembership-> "+err.Error())
	}
	defer rows.Close()

	list := []*v1.DonationCampaignMembership{}
	for rows.Next() {
		td := new(v1.DonationCampaignMembership)
		if err := rows.Scan(&td.Id, &td.Title); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from DonationCampaignMembership row-> "+err.Error())
		}
		list = append(list, td)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from DonationCampaignMembership-> "+err.Error())
	}

	return &v1.ListDonationCampaignMembershipResponse{
		Api:   apiVersion,
		Items: list,
	}, nil
}

// Update todo task
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
	res, err := c.ExecContext(ctx, "UPDATE donation_campaign_membership SET title=$1 WHERE id=$2",
		req.Item.Title, req.Item.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update DonationCampaignMembership-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("DonationCampaignMembership with ID='%d' is not found",
			req.Item.Id))
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
	res, err := c.ExecContext(ctx, "DELETE FROM donation_campaign_membership WHERE id=$1", req.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete DonationCampaignMembership-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("DonationCampaignMembership with ID='%d' is not found",
			req.Id))
	}

	return &v1.DeleteDonationCampaignMembershipResponse{
		Api:     apiVersion,
		Deleted: rows,
	}, nil
}
