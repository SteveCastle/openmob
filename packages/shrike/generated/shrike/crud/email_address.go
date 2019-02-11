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

// NewShrikeServiceServer creates EmailAddress service
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
func (s *shrikeServiceServer) CreateEmailAddress(ctx context.Context, req *v1.CreateEmailAddressRequest) (*v1.CreateEmailAddressResponse, error) {
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
	// insert EmailAddress entity data
	err = c.QueryRowContext(ctx, "INSERT INTO emailaddress (title) VALUES($1)  RETURNING id;",
		req.Item.Title).Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into EmailAddress-> "+err.Error())
	}

	// get ID of creates EmailAddress
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created EmailAddress-> "+err.Error())
	}

	return &v1.CreateEmailAddressResponse{
		Api: apiVersion,
		Id:  id,
	}, nil
}

// Get emailaddress by id.
func (s *shrikeServiceServer) GetEmailAddress(ctx context.Context, req *v1.GetEmailAddressRequest) (*v1.GetEmailAddressResponse, error) {
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

	// query EmailAddress by ID
	rows, err := c.QueryContext(ctx, "SELECT id, title FROM emailaddress WHERE id=$1",
		req.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from EmailAddress-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from EmailAddress-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("EmailAddress with ID='%d' is not found",
			req.Id))
	}

	// get EmailAddress data
	var td v1.EmailAddress
	if err := rows.Scan(&td.Id, &td.Title); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from EmailAddress row-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple EmailAddress rows with ID='%d'",
			req.Id))
	}

	return &v1.GetEmailAddressResponse{
		Api:  apiVersion,
		Item: &td,
	}, nil

}

// Read all todo tasks
func (s *shrikeServiceServer) ListEmailAddress(ctx context.Context, req *v1.ListEmailAddressRequest) (*v1.ListEmailAddressResponse, error) {
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

	// get EmailAddress list
	rows, err := c.QueryContext(ctx, "SELECT id,title FROM EmailAddress")
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from EmailAddress-> "+err.Error())
	}
	defer rows.Close()

	list := []*v1.EmailAddress{}
	for rows.Next() {
		td := new(v1.EmailAddress)
		if err := rows.Scan(&td.Id, &td.Title); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from EmailAddress row-> "+err.Error())
		}
		list = append(list, td)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from EmailAddress-> "+err.Error())
	}

	return &v1.ListEmailAddressResponse{
		Api:   apiVersion,
		Items: list,
	}, nil
}

// Update todo task
func (s *shrikeServiceServer) UpdateEmailAddress(ctx context.Context, req *v1.UpdateEmailAddressRequest) (*v1.UpdateEmailAddressResponse, error) {
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

	// update emailaddress
	res, err := c.ExecContext(ctx, "UPDATE emailaddress SET title=$1 WHERE id=$2",
		req.Item.Title, req.Item.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update emailaddress-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("emailaddress with ID='%d' is not found",
			req.Item.Id))
	}

	return &v1.UpdateEmailAddressResponse{
		Api:     apiVersion,
		Updated: rows,
	}, nil
}

// Delete emailaddress
func (s *shrikeServiceServer) DeleteEmailAddress(ctx context.Context, req *v1.DeleteEmailAddressRequest) (*v1.DeleteEmailAddressResponse, error) {
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

	// delete emailaddress
	res, err := c.ExecContext(ctx, "DELETE FROM emailaddress WHERE id=$1", req.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete emailaddress-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("emailaddress with ID='%d' is not found",
			req.Id))
	}

	return &v1.DeleteEmailAddressResponse{
		Api:     apiVersion,
		Deleted: rows,
	}, nil
}
