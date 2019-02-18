package v1

import (
	"context"
	"fmt"

	v1 "github.com/SteveCastle/openmob/packages/shrike/src/pkg/api/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Create new BoycottMembership
func (s *shrikeServiceServer) CreateBoycottMembership(ctx context.Context, req *v1.CreateBoycottMembershipRequest) (*v1.CreateBoycottMembershipResponse, error) {
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
	// insert BoycottMembership entity data
	err = c.QueryRowContext(ctx, "INSERT INTO boycott_membership (id, created_at, updated_at, cause, boycott) VALUES($1, $2, $3, $4, $5)  RETURNING id;",
		 req.Item.ID,  req.Item.CreatedAt,  req.Item.UpdatedAt,  req.Item.Cause,  req.Item.Boycott, ).Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into BoycottMembership-> "+err.Error())
	}

	// get ID of creates BoycottMembership
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created BoycottMembership-> "+err.Error())
	}

	return &v1.CreateBoycottMembershipResponse{
		Api: apiVersion,
		ID:  id,
	}, nil
}

// Get boycott_membership by id.
func (s *shrikeServiceServer) GetBoycottMembership(ctx context.Context, req *v1.GetBoycottMembershipRequest) (*v1.GetBoycottMembershipResponse, error) {
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

	// query BoycottMembership by ID
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, cause, boycott FROM boycott_membership WHERE id=$1",
		req.ID)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from BoycottMembership-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from BoycottMembership-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("BoycottMembership with ID='%d' is not found",
			req.ID))
	}

	// get BoycottMembership data
	var boycottmembership v1.BoycottMembership
	if err := rows.Scan( &boycottmembership.ID,  &boycottmembership.CreatedAt,  &boycottmembership.UpdatedAt,  &boycottmembership.Cause,  &boycottmembership.Boycott, ); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from BoycottMembership row-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple BoycottMembership rows with ID='%d'",
			req.ID))
	}

	return &v1.GetBoycottMembershipResponse{
		Api:  apiVersion,
		Item: &boycottmembership,
	}, nil

}

// Read all BoycottMembership
func (s *shrikeServiceServer) ListBoycottMembership(ctx context.Context, req *v1.ListBoycottMembershipRequest) (*v1.ListBoycottMembershipResponse, error) {
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

	// get BoycottMembership list
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, cause, boycott FROM boycott_membership")
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from BoycottMembership-> "+err.Error())
	}
	defer rows.Close()

	list := []*v1.BoycottMembership{}
	for rows.Next() {
		boycottmembership := new(v1.BoycottMembership)
		if err := rows.Scan( &boycottmembership.ID,  &boycottmembership.CreatedAt,  &boycottmembership.UpdatedAt,  &boycottmembership.Cause,  &boycottmembership.Boycott, ); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from BoycottMembership row-> "+err.Error())
		}
		list = append(list, boycottmembership)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from BoycottMembership-> "+err.Error())
	}

	return &v1.ListBoycottMembershipResponse{
		Api:   apiVersion,
		Items: list,
	}, nil
}

// Update BoycottMembership
func (s *shrikeServiceServer) UpdateBoycottMembership(ctx context.Context, req *v1.UpdateBoycottMembershipRequest) (*v1.UpdateBoycottMembershipResponse, error) {
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

	// update boycott_membership
	res, err := c.ExecContext(ctx, "UPDATE boycott_membership SET id=$1, created_at=$2, updated_at=$3, cause=$4, boycott=$5 WHERE id=$1",
		req.Item.ID,req.Item.CreatedAt,req.Item.UpdatedAt,req.Item.Cause,req.Item.Boycott, )
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update BoycottMembership-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("BoycottMembership with ID='%d' is not found",
			req.Item.ID))
	}

	return &v1.UpdateBoycottMembershipResponse{
		Api:     apiVersion,
		Updated: rows,
	}, nil
}

// Delete boycott_membership
func (s *shrikeServiceServer) DeleteBoycottMembership(ctx context.Context, req *v1.DeleteBoycottMembershipRequest) (*v1.DeleteBoycottMembershipResponse, error) {
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

	// delete boycott_membership
	res, err := c.ExecContext(ctx, "DELETE FROM boycott_membership WHERE id=$1", req.ID)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete BoycottMembership-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("BoycottMembership with ID='%d' is not found",
			req.ID))
	}

	return &v1.DeleteBoycottMembershipResponse{
		Api:     apiVersion,
		Deleted: rows,
	}, nil
}
