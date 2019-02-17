package v1

import (
	"context"
	"fmt"

	v1 "github.com/SteveCastle/openmob/packages/shrike/src/pkg/api/v1"

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
	var id int64
	// insert Payment entity data
	err = c.QueryRowContext(ctx, "INSERT INTO payment (id, created_at, updated_at, customer_order, ) VALUES($1, $2, $3, $4, )  RETURNING id;",
		 req.Item.ID,  req.Item.CreatedAt,  req.Item.UpdatedAt,  req.Item.CustomerOrder, ).Scan(&id)
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
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, customer_order,  FROM payment WHERE id=$1",
		req.ID)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from Payment-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from Payment-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Payment with ID='%d' is not found",
			req.ID))
	}

	// get Payment data
	var payment v1.Payment
	if err := rows.Scan( &payment.ID,  &payment.CreatedAt,  &payment.UpdatedAt,  &payment.CustomerOrder, ); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from Payment row-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple Payment rows with ID='%d'",
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

	// get Payment list
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, customer_order,  FROM payment")
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from Payment-> "+err.Error())
	}
	defer rows.Close()

	list := []*v1.Payment{}
	for rows.Next() {
		payment := new(v1.Payment)
		if err := rows.Scan( &payment.ID,  &payment.CreatedAt,  &payment.UpdatedAt,  &payment.CustomerOrder, ); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from Payment row-> "+err.Error())
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
	res, err := c.ExecContext(ctx, "UPDATE payment SET $1, $2, $3, $4,  WHERE id=$1",
		req.Item.ID,req.Item.CreatedAt,req.Item.UpdatedAt,req.Item.CustomerOrder, )
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update Payment-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Payment with ID='%d' is not found",
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
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Payment with ID='%d' is not found",
			req.ID))
	}

	return &v1.DeletePaymentResponse{
		Api:     apiVersion,
		Deleted: rows,
	}, nil
}
