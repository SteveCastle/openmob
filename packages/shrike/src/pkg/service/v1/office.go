package v1

import (
	"context"
	"fmt"
	"time"

	v1 "github.com/SteveCastle/openmob/packages/shrike/src/pkg/api/v1"
	"github.com/SteveCastle/openmob/packages/shrike/src/pkg/queries"
	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Create new Office
func (s *shrikeServiceServer) CreateOffice(ctx context.Context, req *v1.CreateOfficeRequest) (*v1.CreateOfficeResponse, error) {
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
	// insert Office entity data
	err = c.QueryRowContext(ctx, "INSERT INTO office (title, election) VALUES($1, $2)  RETURNING id;",
		req.Item.Title, req.Item.Election).Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into Office-> "+err.Error())
	}

	// get ID of creates Office
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created Office-> "+err.Error())
	}

	return &v1.CreateOfficeResponse{
		Api: apiVersion,
		ID:  id,
	}, nil
}

// Get office by id.
func (s *shrikeServiceServer) GetOffice(ctx context.Context, req *v1.GetOfficeRequest) (*v1.GetOfficeResponse, error) {
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

	// query Office by ID
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, title, election FROM office WHERE id=$1",
		req.ID)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from Office-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from Office-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Office with ID='%s' is not found",
			req.ID))
	}

	// scan Office data into protobuf model
	var office v1.Office
	var createdAt time.Time
	var updatedAt time.Time

	if err := rows.Scan(&office.ID, &createdAt, &updatedAt, &office.Title, &office.Election); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from Office row-> "+err.Error())
	}

	// Convert time.Time from database into proto timestamp.
	office.CreatedAt, err = ptypes.TimestampProto(createdAt)
	if err != nil {
		return nil, status.Error(codes.Unknown, "createdAt field has invalid format-> "+err.Error())
	}
	office.UpdatedAt, err = ptypes.TimestampProto(updatedAt)
	if err != nil {
		return nil, status.Error(codes.Unknown, "updatedAt field has invalid format-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple Office rows with ID='%s'",
			req.ID))
	}

	return &v1.GetOfficeResponse{
		Api:  apiVersion,
		Item: &office,
	}, nil

}

// Read all Office
func (s *shrikeServiceServer) ListOffice(ctx context.Context, req *v1.ListOfficeRequest) (*v1.ListOfficeResponse, error) {
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

	// Generate SQL to select all columns in Office Table
	// Then generate filtering and ordering sql and finally run query.

	baseSQL := "SELECT id, created_at, updated_at, title, election FROM office"
	querySQL := queries.BuildOfficeFilters(req.Filters, req.Ordering, req.Limit)
	SQL := fmt.Sprintf("%s %s", baseSQL, querySQL)
	rows, err := c.QueryContext(ctx, SQL)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from Office-> "+err.Error())
	}
	defer rows.Close()

	// Variables to store results returned by database.
	list := []*v1.Office{}
	var createdAt time.Time
	var updatedAt time.Time

	for rows.Next() {
		office := new(v1.Office)
		if err := rows.Scan(&office.ID, &createdAt, &updatedAt, &office.Title, &office.Election); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from Office row-> "+err.Error())
		}
		// Convert time.Time from database into proto timestamp.
		office.CreatedAt, err = ptypes.TimestampProto(createdAt)
		if err != nil {
			return nil, status.Error(codes.Unknown, "createdAt field has invalid format-> "+err.Error())
		}
		office.UpdatedAt, err = ptypes.TimestampProto(updatedAt)
		if err != nil {
			return nil, status.Error(codes.Unknown, "updatedAt field has invalid format-> "+err.Error())
		}

		list = append(list, office)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from Office-> "+err.Error())
	}

	return &v1.ListOfficeResponse{
		Api:   apiVersion,
		Items: list,
	}, nil
}

// Update Office
func (s *shrikeServiceServer) UpdateOffice(ctx context.Context, req *v1.UpdateOfficeRequest) (*v1.UpdateOfficeResponse, error) {
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

	// update office
	res, err := c.ExecContext(ctx, "UPDATE office SET title=$2, election=$3 WHERE id=$1",
		req.Item.ID, req.Item.Title, req.Item.Election)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update Office-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Office with ID='%s' is not found",
			req.Item.ID))
	}

	return &v1.UpdateOfficeResponse{
		Api:     apiVersion,
		Updated: rows,
	}, nil
}

// Delete office
func (s *shrikeServiceServer) DeleteOffice(ctx context.Context, req *v1.DeleteOfficeRequest) (*v1.DeleteOfficeResponse, error) {
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

	// delete office
	res, err := c.ExecContext(ctx, "DELETE FROM office WHERE id=$1", req.ID)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete Office-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Office with ID='%s' is not found",
			req.ID))
	}

	return &v1.DeleteOfficeResponse{
		Api:     apiVersion,
		Deleted: rows,
	}, nil
}
