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

// NewShrikeServiceServer creates Contact service
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

// Create new Contact
func (s *shrikeServiceServer) CreateContact(ctx context.Context, req *v1.CreateContactRequest) (*v1.CreateContactResponse, error) {
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
	// insert Contact entity data
	err = c.QueryRowContext(ctx, "INSERT INTO contact ( id  created_at  updated_at ) VALUES( $1 $2 $3)  RETURNING id;",
		 req.ItemID  req.ItemCreatedAt  req.ItemUpdatedAt ).Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into Contact-> "+err.Error())
	}

	// get ID of creates Contact
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created Contact-> "+err.Error())
	}

	return &v1.CreateContactResponse{
		Api: apiVersion,
		Id:  id,
	}, nil
}

// Get contact by id.
func (s *shrikeServiceServer) GetContact(ctx context.Context, req *v1.GetContactRequest) (*v1.GetContactResponse, error) {
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

	// query Contact by ID
	rows, err := c.QueryContext(ctx, "SELECT id, title FROM contact WHERE id=$1",
		req.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from Contact-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from Contact-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Contact with ID='%d' is not found",
			req.Id))
	}

	// get Contact data
	var contact v1.Contact
	if err := rows.Scan(&contact.Id, &contact.Title); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from Contact row-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple Contact rows with ID='%d'",
			req.Id))
	}

	return &v1.GetContactResponse{
		Api:  apiVersion,
		Item: &contact,
	}, nil

}

// Read all Contact
func (s *shrikeServiceServer) ListContact(ctx context.Context, req *v1.ListContactRequest) (*v1.ListContactResponse, error) {
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

	// get Contact list
	rows, err := c.QueryContext(ctx, "SELECT id,title FROM contact")
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from Contact-> "+err.Error())
	}
	defer rows.Close()

	list := []*v1.Contact{}
	for rows.Next() {
		contact := new(v1.Contact)
		if err := rows.Scan(&contact.Id, &contact.Title); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from Contact row-> "+err.Error())
		}
		list = append(list, contact)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from Contact-> "+err.Error())
	}

	return &v1.ListContactResponse{
		Api:   apiVersion,
		Items: list,
	}, nil
}

// Update Contact
func (s *shrikeServiceServer) UpdateContact(ctx context.Context, req *v1.UpdateContactRequest) (*v1.UpdateContactResponse, error) {
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

	// update contact
	res, err := c.ExecContext(ctx, "UPDATE contact SET title=$1 WHERE id=$2",
		req.Item.Title, req.Item.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update Contact-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Contact with ID='%d' is not found",
			req.Item.Id))
	}

	return &v1.UpdateContactResponse{
		Api:     apiVersion,
		Updated: rows,
	}, nil
}

// Delete contact
func (s *shrikeServiceServer) DeleteContact(ctx context.Context, req *v1.DeleteContactRequest) (*v1.DeleteContactResponse, error) {
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

	// delete contact
	res, err := c.ExecContext(ctx, "DELETE FROM contact WHERE id=$1", req.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete Contact-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Contact with ID='%d' is not found",
			req.Id))
	}

	return &v1.DeleteContactResponse{
		Api:     apiVersion,
		Deleted: rows,
	}, nil
}
