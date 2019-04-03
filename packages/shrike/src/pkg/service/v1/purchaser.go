package v1

import (
	"context"
	"fmt"

	v1 "github.com/SteveCastle/openmob/packages/shrike/src/pkg/api/v1"
	"github.com/SteveCastle/openmob/packages/shrike/src/pkg/queries"
	"github.com/golang/protobuf/ptypes"
	"github.com/lib/pq"
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
	var id string
	// insert Purchaser entity data
	err = c.QueryRowContext(ctx, "INSERT INTO purchaser (customer_order, contact, cause) VALUES($1, $2, $3)  RETURNING id;",
		req.Item.CustomerOrder, req.Item.Contact, req.Item.Cause).Scan(&id)
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
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Purchaser with ID='%s' is not found",
			req.ID))
	}

	// scan Purchaser data into protobuf model
	var purchaser v1.Purchaser
	var createdAt pq.NullTime
	var updatedAt pq.NullTime

	if err := rows.Scan(&purchaser.ID, &createdAt, &updatedAt, &purchaser.CustomerOrder, &purchaser.Contact, &purchaser.Cause); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from Purchaser row-> "+err.Error())
	}

	// Convert pq.NullTime from database into proto timestamp.
	if createdAt.Valid {
		purchaser.CreatedAt, err = ptypes.TimestampProto(createdAt.Time)
		if err != nil {
			return nil, status.Error(codes.Unknown, "createdAt field has invalid format-> "+err.Error())
		}
	}
	if updatedAt.Valid {
		purchaser.UpdatedAt, err = ptypes.TimestampProto(updatedAt.Time)
		if err != nil {
			return nil, status.Error(codes.Unknown, "createdAt field has invalid format-> "+err.Error())
		}
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple Purchaser rows with ID='%s'",
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

	// Generate SQL to select all columns in Purchaser Table
	// Then generate filtering and ordering sql and finally run query.
	querySQL := queries.BuildPurchaserListQuery(req.Filters, req.Ordering, req.Limit)
	// Execute query and scan into return type.
	rows, err := c.QueryContext(ctx, querySQL)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from Purchaser-> "+err.Error())
	}
	defer rows.Close()

	// Variables to store results returned by database.
	list := []*v1.Purchaser{}
	var createdAt pq.NullTime
	var updatedAt pq.NullTime

	for rows.Next() {
		purchaser := new(v1.Purchaser)
		if err := rows.Scan(&purchaser.ID, &createdAt, &updatedAt, &purchaser.CustomerOrder, &purchaser.Contact, &purchaser.Cause); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from Purchaser row-> "+err.Error())
		}
		// Convert pq.NullTime from database into proto timestamp.
		if createdAt.Valid {
			purchaser.CreatedAt, err = ptypes.TimestampProto(createdAt.Time)
			if err != nil {
				return nil, status.Error(codes.Unknown, "createdAt field has invalid format-> "+err.Error())
			}
		}
		if updatedAt.Valid {
			purchaser.UpdatedAt, err = ptypes.TimestampProto(updatedAt.Time)
			if err != nil {
				return nil, status.Error(codes.Unknown, "createdAt field has invalid format-> "+err.Error())
			}
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
	res, err := c.ExecContext(ctx, "UPDATE purchaser SET customer_order=$2, contact=$3, cause=$4 WHERE id=$1",
		req.Item.ID, req.Item.CustomerOrder, req.Item.Contact, req.Item.Cause)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update Purchaser-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Purchaser with ID='%s' is not found",
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
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Purchaser with ID='%s' is not found",
			req.ID))
	}

	return &v1.DeletePurchaserResponse{
		Api:     apiVersion,
		Deleted: rows,
	}, nil
}
