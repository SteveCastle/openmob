package v1

import (
	"context"
	"fmt"

	v1 "github.com/SteveCastle/openmob/packages/shrike/src/pkg/api/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Create new Delivery
func (s *shrikeServiceServer) CreateDelivery(ctx context.Context, req *v1.CreateDeliveryRequest) (*v1.CreateDeliveryResponse, error) {
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
	// insert Delivery entity data
	err = c.QueryRowContext(ctx, "INSERT INTO delivery (id, created_at, updated_at) VALUES($1, $2, $3)  RETURNING id;",
		 req.Item.ID,  req.Item.CreatedAt,  req.Item.UpdatedAt, ).Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into Delivery-> "+err.Error())
	}

	// get ID of creates Delivery
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created Delivery-> "+err.Error())
	}

	return &v1.CreateDeliveryResponse{
		Api: apiVersion,
		ID:  id,
	}, nil
}

// Get delivery by id.
func (s *shrikeServiceServer) GetDelivery(ctx context.Context, req *v1.GetDeliveryRequest) (*v1.GetDeliveryResponse, error) {
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

	// query Delivery by ID
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at FROM delivery WHERE id=$1",
		req.ID)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from Delivery-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from Delivery-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Delivery with ID='%d' is not found",
			req.ID))
	}

	// get Delivery data
	var delivery v1.Delivery
	if err := rows.Scan( &delivery.ID,  &delivery.CreatedAt,  &delivery.UpdatedAt, ); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from Delivery row-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple Delivery rows with ID='%d'",
			req.ID))
	}

	return &v1.GetDeliveryResponse{
		Api:  apiVersion,
		Item: &delivery,
	}, nil

}

// Read all Delivery
func (s *shrikeServiceServer) ListDelivery(ctx context.Context, req *v1.ListDeliveryRequest) (*v1.ListDeliveryResponse, error) {
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

	// get Delivery list
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at FROM delivery")
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from Delivery-> "+err.Error())
	}
	defer rows.Close()

	list := []*v1.Delivery{}
	for rows.Next() {
		delivery := new(v1.Delivery)
		if err := rows.Scan( &delivery.ID,  &delivery.CreatedAt,  &delivery.UpdatedAt, ); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from Delivery row-> "+err.Error())
		}
		list = append(list, delivery)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from Delivery-> "+err.Error())
	}

	return &v1.ListDeliveryResponse{
		Api:   apiVersion,
		Items: list,
	}, nil
}

// Update Delivery
func (s *shrikeServiceServer) UpdateDelivery(ctx context.Context, req *v1.UpdateDeliveryRequest) (*v1.UpdateDeliveryResponse, error) {
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

	// update delivery
	res, err := c.ExecContext(ctx, "UPDATE delivery SET id=$1, created_at=$2, updated_at=$3 WHERE id=$1",
		req.Item.ID,req.Item.CreatedAt,req.Item.UpdatedAt, )
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update Delivery-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Delivery with ID='%d' is not found",
			req.Item.ID))
	}

	return &v1.UpdateDeliveryResponse{
		Api:     apiVersion,
		Updated: rows,
	}, nil
}

// Delete delivery
func (s *shrikeServiceServer) DeleteDelivery(ctx context.Context, req *v1.DeleteDeliveryRequest) (*v1.DeleteDeliveryResponse, error) {
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

	// delete delivery
	res, err := c.ExecContext(ctx, "DELETE FROM delivery WHERE id=$1", req.ID)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete Delivery-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Delivery with ID='%d' is not found",
			req.ID))
	}

	return &v1.DeleteDeliveryResponse{
		Api:     apiVersion,
		Deleted: rows,
	}, nil
}
