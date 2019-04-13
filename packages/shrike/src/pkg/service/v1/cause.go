package v1

import (
	"context"
	"fmt"

	v1 "github.com/SteveCastle/openmob/packages/shrike/src/pkg/api/v1"
	"github.com/SteveCastle/openmob/packages/shrike/src/pkg/models/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Create new Cause
func (s *shrikeServiceServer) CreateCause(ctx context.Context, req *v1.CreateCauseRequest) (*v1.CreateCauseResponse, error) {
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
	// insert Cause entity data
	err = c.QueryRowContext(ctx, "INSERT INTO cause (title, slug, summary, home_page, photo) VALUES($1, $2, $3, $4, $5)  RETURNING id;",
		req.Item.Title, req.Item.Slug, req.Item.Summary, req.Item.HomePage, req.Item.Photo).Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into Cause-> "+err.Error())
	}

	// get ID of creates Cause
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created Cause-> "+err.Error())
	}

	return &v1.CreateCauseResponse{
		Api: apiVersion,
		ID:  id,
	}, nil
}

// Get cause by id.
func (s *shrikeServiceServer) GetCause(ctx context.Context, req *v1.GetCauseRequest) (*v1.GetCauseResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a Cause Manager
	m := models.NewCauseManager(s.db)

	// Get a list of causes given filters, ordering, and limit rules.
	cause, err := m.GetCause(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	return &v1.GetCauseResponse{
		Api:  apiVersion,
		Item: m.GetProto(cause),
	}, nil

}

// Read all Cause
func (s *shrikeServiceServer) ListCause(ctx context.Context, req *v1.ListCauseRequest) (*v1.ListCauseResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}

	// Create a Cause Manager
	m := models.NewCauseManager(s.db)

	// Get a list of causes given filters, ordering, and limit rules.
	list, err := m.ListCause(ctx, req.Filters, req.Ordering, req.Limit)
	if err != nil {
		return nil, err
	}

	return &v1.ListCauseResponse{
		Api:   apiVersion,
		Items: m.GetProtoList(list),
	}, nil
}

// Update Cause
func (s *shrikeServiceServer) UpdateCause(ctx context.Context, req *v1.UpdateCauseRequest) (*v1.UpdateCauseResponse, error) {
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

	// update cause
	res, err := c.ExecContext(ctx, "UPDATE cause SET title=$2, slug=$3, summary=$4, home_page=$5, photo=$6 WHERE id=$1",
		req.Item.ID, req.Item.Title, req.Item.Slug, req.Item.Summary, req.Item.HomePage, req.Item.Photo)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update Cause-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Cause with ID='%s' is not found",
			req.Item.ID))
	}

	return &v1.UpdateCauseResponse{
		Api:     apiVersion,
		Updated: rows,
	}, nil
}

// Delete cause
func (s *shrikeServiceServer) DeleteCause(ctx context.Context, req *v1.DeleteCauseRequest) (*v1.DeleteCauseResponse, error) {
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

	// delete cause
	res, err := c.ExecContext(ctx, "DELETE FROM cause WHERE id=$1", req.ID)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete Cause-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Cause with ID='%s' is not found",
			req.ID))
	}

	return &v1.DeleteCauseResponse{
		Api:     apiVersion,
		Deleted: rows,
	}, nil
}
