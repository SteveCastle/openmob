package v1

import (
	"context"
	"fmt"

	v1 "github.com/SteveCastle/openmob/packages/shrike/src/pkg/api/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Create new Purchaser
func (s *shrikeServiceServer) CreatePurchaser(ctx context.Context, req *v1.CreatePurchaserRequest) (*v1.CreatePurchaserResponse, error) {
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
	// insert Purchaser entity data
	err = c.QueryRowContext(ctx, "INSERT INTO purchaser (id, created_at, updated_at, customer_order, contact, cause) VALUES($1, $2, $3, $4, $5, $6)  RETURNING id;",
		 req.Item.ID,  req.Item.CreatedAt,  req.Item.UpdatedAt,  req.Item.CustomerOrder,  req.Item.Contact,  req.Item.Cause, ).Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into Purchaser-> "+err.Error())
	}

	// get ID of creates Purchaser
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created Purchaser-> "+err.Error())
	}

	return &v1.CreatePurchaserResponse{
		Api: apiVersion,
		ID:  id,
	}, nil
}

// Get purchaser by id.
func (s *shrikeServiceServer) GetPurchaser(ctx context.Context, req *v1.GetPurchaserRequest) (*v1.GetPurchaserResponse, error) {
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

	// query Purchaser by ID
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, customer_order, contact, cause FROM purchaser WHERE id=$1",
		req.ID)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from Purchaser-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from Purchaser-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Purchaser with ID='%d' is not found",
			req.ID))
	}

	// get Purchaser data
	var purchaser v1.Purchaser
	if err := rows.Scan( &purchaser.ID,  &purchaser.CreatedAt,  &purchaser.UpdatedAt,  &purchaser.CustomerOrder,  &purchaser.Contact,  &purchaser.Cause, ); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from Purchaser row-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple Purchaser rows with ID='%d'",
			req.ID))
	}

	return &v1.GetPurchaserResponse{
		Api:  apiVersion,
		Item: &purchaser,
	}, nil

}

// Read all Purchaser
func (s *shrikeServiceServer) ListPurchaser(ctx context.Context, req *v1.ListPurchaserRequest) (*v1.ListPurchaserResponse, error) {
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

	// get Purchaser list
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, customer_order, contact, cause FROM purchaser")
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from Purchaser-> "+err.Error())
	}
	defer rows.Close()

	list := []*v1.Purchaser{}
	for rows.Next() {
		purchaser := new(v1.Purchaser)
		if err := rows.Scan( &purchaser.ID,  &purchaser.CreatedAt,  &purchaser.UpdatedAt,  &purchaser.CustomerOrder,  &purchaser.Contact,  &purchaser.Cause, ); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from Purchaser row-> "+err.Error())
		}
		list = append(list, purchaser)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from Purchaser-> "+err.Error())
	}

	return &v1.ListPurchaserResponse{
		Api:   apiVersion,
		Items: list,
	}, nil
}

// Update Purchaser
func (s *shrikeServiceServer) UpdatePurchaser(ctx context.Context, req *v1.UpdatePurchaserRequest) (*v1.UpdatePurchaserResponse, error) {
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

	// update purchaser
	res, err := c.ExecContext(ctx, "UPDATE purchaser SET id=$1, created_at=$2, updated_at=$3, customer_order=$4, contact=$5, cause=$6 WHERE id=$1",
		req.Item.ID,req.Item.CreatedAt,req.Item.UpdatedAt,req.Item.CustomerOrder,req.Item.Contact,req.Item.Cause, )
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update Purchaser-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Purchaser with ID='%d' is not found",
			req.Item.ID))
	}

	return &v1.UpdatePurchaserResponse{
		Api:     apiVersion,
		Updated: rows,
	}, nil
}

// Delete purchaser
func (s *shrikeServiceServer) DeletePurchaser(ctx context.Context, req *v1.DeletePurchaserRequest) (*v1.DeletePurchaserResponse, error) {
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

	// delete purchaser
	res, err := c.ExecContext(ctx, "DELETE FROM purchaser WHERE id=$1", req.ID)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete Purchaser-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Purchaser with ID='%d' is not found",
			req.ID))
	}

	return &v1.DeletePurchaserResponse{
		Api:     apiVersion,
		Deleted: rows,
	}, nil
}
