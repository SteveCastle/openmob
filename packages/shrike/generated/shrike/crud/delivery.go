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

// NewShrikeServiceServer creates Delivery service
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
	err = c.QueryRowContext(ctx, "INSERT INTO delivery ( id  created_at  updated_at ) VALUES( $1 $2 $3)  RETURNING id;",
		 req.ItemID  req.ItemCreatedAt  req.ItemUpdatedAt ).Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into Delivery-> "+err.Error())
	}

	// get ID of creates Delivery
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created Delivery-> "+err.Error())
	}

	return &v1.CreateDeliveryResponse{
		Api: apiVersion,
		Id:  id,
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
	rows, err := c.QueryContext(ctx, "SELECT id, title FROM delivery WHERE id=$1",
		req.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from Delivery-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from Delivery-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Delivery with ID='%d' is not found",
			req.Id))
	}

	// get Delivery data
	var delivery v1.Delivery
	if err := rows.Scan(&delivery.Id, &delivery.Title); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from Delivery row-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple Delivery rows with ID='%d'",
			req.Id))
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
	rows, err := c.QueryContext(ctx, "SELECT id,title FROM delivery")
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from Delivery-> "+err.Error())
	}
	defer rows.Close()

	list := []*v1.Delivery{}
	for rows.Next() {
		delivery := new(v1.Delivery)
		if err := rows.Scan(&delivery.Id, &delivery.Title); err != nil {
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
	res, err := c.ExecContext(ctx, "UPDATE delivery SET title=$1 WHERE id=$2",
		req.Item.Title, req.Item.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update Delivery-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Delivery with ID='%d' is not found",
			req.Item.Id))
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
	res, err := c.ExecContext(ctx, "DELETE FROM delivery WHERE id=$1", req.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete Delivery-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Delivery with ID='%d' is not found",
			req.Id))
	}

	return &v1.DeleteDeliveryResponse{
		Api:     apiVersion,
		Deleted: rows,
	}, nil
}
