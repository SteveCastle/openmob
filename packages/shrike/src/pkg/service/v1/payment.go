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

// Create new Payment
func (s *shrikeServiceServer) CreatePayment(ctx context.Context, req *v1.CreatePaymentRequest) (*v1.CreatePaymentResponse, error) {
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
	// insert Payment entity data
	err = c.QueryRowContext(ctx, "INSERT INTO payment (customer_order) VALUES($1)  RETURNING id;",
		req.Item.CustomerOrder).Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into Payment-> "+err.Error())
	}

	// get ID of creates Payment
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created Payment-> "+err.Error())
	}

	return &v1.CreatePaymentResponse{
		Api: apiVersion,
		ID:  id,
	}, nil
}

// Get payment by id.
func (s *shrikeServiceServer) GetPayment(ctx context.Context, req *v1.GetPaymentRequest) (*v1.GetPaymentResponse, error) {
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

	// query Payment by ID
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, customer_order FROM payment WHERE id=$1",
		req.ID)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from Payment-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from Payment-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Payment with ID='%s' is not found",
			req.ID))
	}

	// scan Payment data into protobuf model
	var payment v1.Payment
	var createdAt pq.NullTime
	var updatedAt pq.NullTime

	if err := rows.Scan(&payment.ID, &createdAt, &updatedAt, &payment.CustomerOrder); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from Payment row-> "+err.Error())
	}

	// Convert pq.NullTime from database into proto timestamp.
	if createdAt.Valid {
		payment.CreatedAt, err = ptypes.TimestampProto(createdAt.Time)
		if err != nil {
			return nil, status.Error(codes.Unknown, "createdAt field has invalid format-> "+err.Error())
		}
	}
	if updatedAt.Valid {
		payment.UpdatedAt, err = ptypes.TimestampProto(updatedAt.Time)
		if err != nil {
			return nil, status.Error(codes.Unknown, "createdAt field has invalid format-> "+err.Error())
		}
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple Payment rows with ID='%s'",
			req.ID))
	}

	return &v1.GetPaymentResponse{
		Api:  apiVersion,
		Item: &payment,
	}, nil

}

// Read all Payment
func (s *shrikeServiceServer) ListPayment(ctx context.Context, req *v1.ListPaymentRequest) (*v1.ListPaymentResponse, error) {
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

	// Generate SQL to select all columns in Payment Table
	// Then generate filtering and ordering sql and finally run query.

	baseSQL := "SELECT id, created_at, updated_at, customer_order FROM payment"
	querySQL := queries.BuildPaymentFilters(req.Filters, req.Ordering, req.Limit)
	SQL := fmt.Sprintf("%s %s", baseSQL, querySQL)
	rows, err := c.QueryContext(ctx, SQL)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from Payment-> "+err.Error())
	}
	defer rows.Close()

	// Variables to store results returned by database.
	list := []*v1.Payment{}
	var createdAt pq.NullTime
	var updatedAt pq.NullTime

	for rows.Next() {
		payment := new(v1.Payment)
		if err := rows.Scan(&payment.ID, &createdAt, &updatedAt, &payment.CustomerOrder); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from Payment row-> "+err.Error())
		}
		// Convert pq.NullTime from database into proto timestamp.
		if createdAt.Valid {
			payment.CreatedAt, err = ptypes.TimestampProto(createdAt.Time)
			if err != nil {
				return nil, status.Error(codes.Unknown, "createdAt field has invalid format-> "+err.Error())
			}
		}
		if updatedAt.Valid {
			payment.UpdatedAt, err = ptypes.TimestampProto(updatedAt.Time)
			if err != nil {
				return nil, status.Error(codes.Unknown, "createdAt field has invalid format-> "+err.Error())
			}
		}

		list = append(list, payment)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from Payment-> "+err.Error())
	}

	return &v1.ListPaymentResponse{
		Api:   apiVersion,
		Items: list,
	}, nil
}

// Update Payment
func (s *shrikeServiceServer) UpdatePayment(ctx context.Context, req *v1.UpdatePaymentRequest) (*v1.UpdatePaymentResponse, error) {
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

	// update payment
	res, err := c.ExecContext(ctx, "UPDATE payment SET customer_order=$2 WHERE id=$1",
		req.Item.ID, req.Item.CustomerOrder)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update Payment-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Payment with ID='%s' is not found",
			req.Item.ID))
	}

	return &v1.UpdatePaymentResponse{
		Api:     apiVersion,
		Updated: rows,
	}, nil
}

// Delete payment
func (s *shrikeServiceServer) DeletePayment(ctx context.Context, req *v1.DeletePaymentRequest) (*v1.DeletePaymentResponse, error) {
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

	// delete payment
	res, err := c.ExecContext(ctx, "DELETE FROM payment WHERE id=$1", req.ID)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete Payment-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Payment with ID='%s' is not found",
			req.ID))
	}

	return &v1.DeletePaymentResponse{
		Api:     apiVersion,
		Deleted: rows,
	}, nil
}
