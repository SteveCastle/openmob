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

// NewShrikeServiceServer creates Office service
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
	var id int64
	// insert Office entity data
	err = c.QueryRowContext(ctx, "INSERT INTO office (title) VALUES($1)  RETURNING id;",
		req.Item.Title).Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into Office-> "+err.Error())
	}

	// get ID of creates Office
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created Office-> "+err.Error())
	}

	return &v1.CreateOfficeResponse{
		Api: apiVersion,
		Id:  id,
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
	rows, err := c.QueryContext(ctx, "SELECT id, title FROM office WHERE id=$1",
		req.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from Office-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from Office-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Office with ID='%d' is not found",
			req.Id))
	}

	// get Office data
	var td v1.Office
	if err := rows.Scan(&td.Id, &td.Title); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from Office row-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple Office rows with ID='%d'",
			req.Id))
	}

	return &v1.GetOfficeResponse{
		Api:  apiVersion,
		Item: &td,
	}, nil

}

// Read all todo tasks
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

	// get Office list
	rows, err := c.QueryContext(ctx, "SELECT id,title FROM Office")
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from Office-> "+err.Error())
	}
	defer rows.Close()

	list := []*v1.Office{}
	for rows.Next() {
		td := new(v1.Office)
		if err := rows.Scan(&td.Id, &td.Title); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from Office row-> "+err.Error())
		}
		list = append(list, td)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from Office-> "+err.Error())
	}

	return &v1.ListOfficeResponse{
		Api:   apiVersion,
		Items: list,
	}, nil
}

// Update todo task
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
	res, err := c.ExecContext(ctx, "UPDATE office SET title=$1 WHERE id=$2",
		req.Item.Title, req.Item.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update office-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("office with ID='%d' is not found",
			req.Item.Id))
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
	res, err := c.ExecContext(ctx, "DELETE FROM office WHERE id=$1", req.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete office-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("office with ID='%d' is not found",
			req.Id))
	}

	return &v1.DeleteOfficeResponse{
		Api:     apiVersion,
		Deleted: rows,
	}, nil
}
