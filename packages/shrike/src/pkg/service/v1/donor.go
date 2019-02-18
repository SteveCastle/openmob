package v1

import (
	"context"
	"fmt"

	v1 "github.com/SteveCastle/openmob/packages/shrike/src/pkg/api/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Create new Donor
func (s *shrikeServiceServer) CreateDonor(ctx context.Context, req *v1.CreateDonorRequest) (*v1.CreateDonorResponse, error) {
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
	// insert Donor entity data
	err = c.QueryRowContext(ctx, "INSERT INTO donor (id, created_at, updated_at, customer_order, contact, cause) VALUES($1, $2, $3, $4, $5, $6)  RETURNING id;",
		 req.Item.ID,  req.Item.CreatedAt,  req.Item.UpdatedAt,  req.Item.CustomerOrder,  req.Item.Contact,  req.Item.Cause, ).Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into Donor-> "+err.Error())
	}

	// get ID of creates Donor
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created Donor-> "+err.Error())
	}

	return &v1.CreateDonorResponse{
		Api: apiVersion,
		ID:  id,
	}, nil
}

// Get donor by id.
func (s *shrikeServiceServer) GetDonor(ctx context.Context, req *v1.GetDonorRequest) (*v1.GetDonorResponse, error) {
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

	// query Donor by ID
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, customer_order, contact, cause FROM donor WHERE id=$1",
		req.ID)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from Donor-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from Donor-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Donor with ID='%d' is not found",
			req.ID))
	}

	// get Donor data
	var donor v1.Donor
	if err := rows.Scan( &donor.ID,  &donor.CreatedAt,  &donor.UpdatedAt,  &donor.CustomerOrder,  &donor.Contact,  &donor.Cause, ); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from Donor row-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple Donor rows with ID='%d'",
			req.ID))
	}

	return &v1.GetDonorResponse{
		Api:  apiVersion,
		Item: &donor,
	}, nil

}

// Read all Donor
func (s *shrikeServiceServer) ListDonor(ctx context.Context, req *v1.ListDonorRequest) (*v1.ListDonorResponse, error) {
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

	// get Donor list
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, customer_order, contact, cause FROM donor")
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from Donor-> "+err.Error())
	}
	defer rows.Close()

	list := []*v1.Donor{}
	for rows.Next() {
		donor := new(v1.Donor)
		if err := rows.Scan( &donor.ID,  &donor.CreatedAt,  &donor.UpdatedAt,  &donor.CustomerOrder,  &donor.Contact,  &donor.Cause, ); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from Donor row-> "+err.Error())
		}
		list = append(list, donor)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from Donor-> "+err.Error())
	}

	return &v1.ListDonorResponse{
		Api:   apiVersion,
		Items: list,
	}, nil
}

// Update Donor
func (s *shrikeServiceServer) UpdateDonor(ctx context.Context, req *v1.UpdateDonorRequest) (*v1.UpdateDonorResponse, error) {
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

	// update donor
	res, err := c.ExecContext(ctx, "UPDATE donor SET id=$1, created_at=$2, updated_at=$3, customer_order=$4, contact=$5, cause=$6 WHERE id=$1",
		req.Item.ID,req.Item.CreatedAt,req.Item.UpdatedAt,req.Item.CustomerOrder,req.Item.Contact,req.Item.Cause, )
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update Donor-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Donor with ID='%d' is not found",
			req.Item.ID))
	}

	return &v1.UpdateDonorResponse{
		Api:     apiVersion,
		Updated: rows,
	}, nil
}

// Delete donor
func (s *shrikeServiceServer) DeleteDonor(ctx context.Context, req *v1.DeleteDonorRequest) (*v1.DeleteDonorResponse, error) {
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

	// delete donor
	res, err := c.ExecContext(ctx, "DELETE FROM donor WHERE id=$1", req.ID)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete Donor-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Donor with ID='%d' is not found",
			req.ID))
	}

	return &v1.DeleteDonorResponse{
		Api:     apiVersion,
		Deleted: rows,
	}, nil
}