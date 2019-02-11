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

// NewShrikeServiceServer creates DonationCampaign service
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
	err = c.QueryRowContext(ctx, "INSERT INTO donation_campaign (title) VALUES($1)  RETURNING id;",
		req.Item.Title).Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into DonationCampaign-> "+err.Error())
	}

	// get ID of creates DonationCampaign
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created DonationCampaign-> "+err.Error())
	}

	return &v1.CreateDonationCampaignResponse{
		Api: apiVersion,
		Id:  id,
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
	rows, err := c.QueryContext(ctx, "SELECT id, title FROM donation_campaign WHERE id=$1",
		req.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from DonationCampaign-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from DonationCampaign-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("DonationCampaign with ID='%d' is not found",
			req.Id))
	}

	// get DonationCampaign data
	var td v1.DonationCampaign
	if err := rows.Scan(&td.Id, &td.Title); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from DonationCampaign row-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple DonationCampaign rows with ID='%d'",
			req.Id))
	}

	return &v1.GetDonationCampaignResponse{
		Api:  apiVersion,
		Item: &td,
	}, nil

}

// Read all todo tasks
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
	rows, err := c.QueryContext(ctx, "SELECT id,title FROM donation_campaign")
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from DonationCampaign-> "+err.Error())
	}
	defer rows.Close()

	list := []*v1.DonationCampaign{}
	for rows.Next() {
		td := new(v1.DonationCampaign)
		if err := rows.Scan(&td.Id, &td.Title); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from DonationCampaign row-> "+err.Error())
		}
		list = append(list, td)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from DonationCampaign-> "+err.Error())
	}

	return &v1.ListDonationCampaignResponse{
		Api:   apiVersion,
		Items: list,
	}, nil
}

// Update todo task
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
	res, err := c.ExecContext(ctx, "UPDATE donation_campaign SET title=$1 WHERE id=$2",
		req.Item.Title, req.Item.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update DonationCampaign-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("DonationCampaign with ID='%d' is not found",
			req.Item.Id))
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
	res, err := c.ExecContext(ctx, "DELETE FROM donation_campaign WHERE id=$1", req.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete DonationCampaign-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("DonationCampaign with ID='%d' is not found",
			req.Id))
	}

	return &v1.DeleteDonationCampaignResponse{
		Api:     apiVersion,
		Deleted: rows,
	}, nil
}
