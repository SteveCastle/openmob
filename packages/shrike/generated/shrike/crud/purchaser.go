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

// NewShrikeServiceServer creates Purchaser service
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

// Create new todo task
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
	err = c.QueryRowContext(ctx, "INSERT INTO purchaser (title) VALUES($1)  RETURNING id;",
		req.Item.Title).Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into Purchaser-> "+err.Error())
	}

	// get ID of creates Purchaser
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created Purchaser-> "+err.Error())
	}

	return &v1.CreatePurchaserResponse{
		Api: apiVersion,
		Id:  id,
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
	rows, err := c.QueryContext(ctx, "SELECT id, title FROM purchaser WHERE id=$1",
		req.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from Purchaser-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from Purchaser-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Purchaser with ID='%d' is not found",
			req.Id))
	}

	// get Purchaser data
	var td v1.Purchaser
	if err := rows.Scan(&td.Id, &td.Title); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from Purchaser row-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple Purchaser rows with ID='%d'",
			req.Id))
	}

	return &v1.GetPurchaserResponse{
		Api:  apiVersion,
		Item: &td,
	}, nil

}

// Read all todo tasks
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
	rows, err := c.QueryContext(ctx, "SELECT id,title FROM Purchaser")
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from Purchaser-> "+err.Error())
	}
	defer rows.Close()

	list := []*v1.Purchaser{}
	for rows.Next() {
		td := new(v1.Purchaser)
		if err := rows.Scan(&td.Id, &td.Title); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from Purchaser row-> "+err.Error())
		}
		list = append(list, td)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from Purchaser-> "+err.Error())
	}

	return &v1.ListPurchaserResponse{
		Api:   apiVersion,
		Items: list,
	}, nil
}

// Update todo task
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
	res, err := c.ExecContext(ctx, "UPDATE purchaser SET title=$1 WHERE id=$2",
		req.Item.Title, req.Item.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update purchaser-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("purchaser with ID='%d' is not found",
			req.Item.Id))
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
	res, err := c.ExecContext(ctx, "DELETE FROM purchaser WHERE id=$1", req.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete purchaser-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("purchaser with ID='%d' is not found",
			req.Id))
	}

	return &v1.DeletePurchaserResponse{
		Api:     apiVersion,
		Deleted: rows,
	}, nil
}
