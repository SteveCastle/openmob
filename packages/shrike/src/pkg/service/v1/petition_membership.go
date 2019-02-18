package v1

import (
	"context"
	"fmt"

	v1 "github.com/SteveCastle/openmob/packages/shrike/src/pkg/api/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Create new PetitionMembership
func (s *shrikeServiceServer) CreatePetitionMembership(ctx context.Context, req *v1.CreatePetitionMembershipRequest) (*v1.CreatePetitionMembershipResponse, error) {
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
	// insert PetitionMembership entity data
	err = c.QueryRowContext(ctx, "INSERT INTO petition_membership (id, created_at, updated_at, cause, petition) VALUES($1, $2, $3, $4, $5)  RETURNING id;",
		 req.Item.ID,  req.Item.CreatedAt,  req.Item.UpdatedAt,  req.Item.Cause,  req.Item.Petition, ).Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into PetitionMembership-> "+err.Error())
	}

	// get ID of creates PetitionMembership
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created PetitionMembership-> "+err.Error())
	}

	return &v1.CreatePetitionMembershipResponse{
		Api: apiVersion,
		ID:  id,
	}, nil
}

// Get petition_membership by id.
func (s *shrikeServiceServer) GetPetitionMembership(ctx context.Context, req *v1.GetPetitionMembershipRequest) (*v1.GetPetitionMembershipResponse, error) {
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

	// query PetitionMembership by ID
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, cause, petition FROM petition_membership WHERE id=$1",
		req.ID)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from PetitionMembership-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from PetitionMembership-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("PetitionMembership with ID='%d' is not found",
			req.ID))
	}

	// get PetitionMembership data
	var petitionmembership v1.PetitionMembership
	if err := rows.Scan( &petitionmembership.ID,  &petitionmembership.CreatedAt,  &petitionmembership.UpdatedAt,  &petitionmembership.Cause,  &petitionmembership.Petition, ); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from PetitionMembership row-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple PetitionMembership rows with ID='%d'",
			req.ID))
	}

	return &v1.GetPetitionMembershipResponse{
		Api:  apiVersion,
		Item: &petitionmembership,
	}, nil

}

// Read all PetitionMembership
func (s *shrikeServiceServer) ListPetitionMembership(ctx context.Context, req *v1.ListPetitionMembershipRequest) (*v1.ListPetitionMembershipResponse, error) {
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

	// get PetitionMembership list
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, cause, petition FROM petition_membership")
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from PetitionMembership-> "+err.Error())
	}
	defer rows.Close()

	list := []*v1.PetitionMembership{}
	for rows.Next() {
		petitionmembership := new(v1.PetitionMembership)
		if err := rows.Scan( &petitionmembership.ID,  &petitionmembership.CreatedAt,  &petitionmembership.UpdatedAt,  &petitionmembership.Cause,  &petitionmembership.Petition, ); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from PetitionMembership row-> "+err.Error())
		}
		list = append(list, petitionmembership)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from PetitionMembership-> "+err.Error())
	}

	return &v1.ListPetitionMembershipResponse{
		Api:   apiVersion,
		Items: list,
	}, nil
}

// Update PetitionMembership
func (s *shrikeServiceServer) UpdatePetitionMembership(ctx context.Context, req *v1.UpdatePetitionMembershipRequest) (*v1.UpdatePetitionMembershipResponse, error) {
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

	// update petition_membership
	res, err := c.ExecContext(ctx, "UPDATE petition_membership SET id=$1, created_at=$2, updated_at=$3, cause=$4, petition=$5 WHERE id=$1",
		req.Item.ID,req.Item.CreatedAt,req.Item.UpdatedAt,req.Item.Cause,req.Item.Petition, )
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update PetitionMembership-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("PetitionMembership with ID='%d' is not found",
			req.Item.ID))
	}

	return &v1.UpdatePetitionMembershipResponse{
		Api:     apiVersion,
		Updated: rows,
	}, nil
}

// Delete petition_membership
func (s *shrikeServiceServer) DeletePetitionMembership(ctx context.Context, req *v1.DeletePetitionMembershipRequest) (*v1.DeletePetitionMembershipResponse, error) {
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

	// delete petition_membership
	res, err := c.ExecContext(ctx, "DELETE FROM petition_membership WHERE id=$1", req.ID)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete PetitionMembership-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("PetitionMembership with ID='%d' is not found",
			req.ID))
	}

	return &v1.DeletePetitionMembershipResponse{
		Api:     apiVersion,
		Deleted: rows,
	}, nil
}