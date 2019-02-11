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

// NewShrikeServiceServer creates MailingAddress service
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
func (s *shrikeServiceServer) CreateMailingAddress(ctx context.Context, req *v1.CreateMailingAddressRequest) (*v1.CreateMailingAddressResponse, error) {
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
	// insert MailingAddress entity data
	err = c.QueryRowContext(ctx, "INSERT INTO mailingaddress (title) VALUES($1)  RETURNING id;",
		req.Item.Title).Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into MailingAddress-> "+err.Error())
	}

	// get ID of creates MailingAddress
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created MailingAddress-> "+err.Error())
	}

	return &v1.CreateMailingAddressResponse{
		Api: apiVersion,
		Id:  id,
	}, nil
}

// Get mailingaddress by id.
func (s *shrikeServiceServer) GetMailingAddress(ctx context.Context, req *v1.GetMailingAddressRequest) (*v1.GetMailingAddressResponse, error) {
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

	// query MailingAddress by ID
	rows, err := c.QueryContext(ctx, "SELECT id, title FROM mailingaddress WHERE id=$1",
		req.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from MailingAddress-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from MailingAddress-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("MailingAddress with ID='%d' is not found",
			req.Id))
	}

	// get MailingAddress data
	var td v1.MailingAddress
	if err := rows.Scan(&td.Id, &td.Title); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from MailingAddress row-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple MailingAddress rows with ID='%d'",
			req.Id))
	}

	return &v1.GetMailingAddressResponse{
		Api:  apiVersion,
		Item: &td,
	}, nil

}

// Read all todo tasks
func (s *shrikeServiceServer) ListMailingAddress(ctx context.Context, req *v1.ListMailingAddressRequest) (*v1.ListMailingAddressResponse, error) {
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

	// get MailingAddress list
	rows, err := c.QueryContext(ctx, "SELECT id,title FROM MailingAddress")
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from MailingAddress-> "+err.Error())
	}
	defer rows.Close()

	list := []*v1.MailingAddress{}
	for rows.Next() {
		td := new(v1.MailingAddress)
		if err := rows.Scan(&td.Id, &td.Title); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from MailingAddress row-> "+err.Error())
		}
		list = append(list, td)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from MailingAddress-> "+err.Error())
	}

	return &v1.ListMailingAddressResponse{
		Api:   apiVersion,
		Items: list,
	}, nil
}

// Update todo task
func (s *shrikeServiceServer) UpdateMailingAddress(ctx context.Context, req *v1.UpdateMailingAddressRequest) (*v1.UpdateMailingAddressResponse, error) {
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

	// update mailingaddress
	res, err := c.ExecContext(ctx, "UPDATE mailingaddress SET title=$1 WHERE id=$2",
		req.Item.Title, req.Item.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update mailingaddress-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("mailingaddress with ID='%d' is not found",
			req.Item.Id))
	}

	return &v1.UpdateMailingAddressResponse{
		Api:     apiVersion,
		Updated: rows,
	}, nil
}

// Delete mailingaddress
func (s *shrikeServiceServer) DeleteMailingAddress(ctx context.Context, req *v1.DeleteMailingAddressRequest) (*v1.DeleteMailingAddressResponse, error) {
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

	// delete mailingaddress
	res, err := c.ExecContext(ctx, "DELETE FROM mailingaddress WHERE id=$1", req.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete mailingaddress-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("mailingaddress with ID='%d' is not found",
			req.Id))
	}

	return &v1.DeleteMailingAddressResponse{
		Api:     apiVersion,
		Deleted: rows,
	}, nil
}
