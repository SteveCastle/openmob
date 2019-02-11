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

// NewShrikeServiceServer creates BoycottMembership service
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
func (s *shrikeServiceServer) CreateBoycottMembership(ctx context.Context, req *v1.CreateBoycottMembershipRequest) (*v1.CreateBoycottMembershipResponse, error) {
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
	// insert BoycottMembership entity data
	err = c.QueryRowContext(ctx, "INSERT INTO boycottmembership (title) VALUES($1)  RETURNING id;",
		req.Item.Title).Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into BoycottMembership-> "+err.Error())
	}

	// get ID of creates BoycottMembership
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created BoycottMembership-> "+err.Error())
	}

	return &v1.CreateBoycottMembershipResponse{
		Api: apiVersion,
		Id:  id,
	}, nil
}

// Get boycottmembership by id.
func (s *shrikeServiceServer) GetBoycottMembership(ctx context.Context, req *v1.GetBoycottMembershipRequest) (*v1.GetBoycottMembershipResponse, error) {
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

	// query BoycottMembership by ID
	rows, err := c.QueryContext(ctx, "SELECT id, title FROM boycottmembership WHERE id=$1",
		req.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from BoycottMembership-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from BoycottMembership-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("BoycottMembership with ID='%d' is not found",
			req.Id))
	}

	// get BoycottMembership data
	var td v1.BoycottMembership
	if err := rows.Scan(&td.Id, &td.Title); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from BoycottMembership row-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple BoycottMembership rows with ID='%d'",
			req.Id))
	}

	return &v1.GetBoycottMembershipResponse{
		Api:  apiVersion,
		Item: &td,
	}, nil

}

// Read all todo tasks
func (s *shrikeServiceServer) ListBoycottMembership(ctx context.Context, req *v1.ListBoycottMembershipRequest) (*v1.ListBoycottMembershipResponse, error) {
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

	// get BoycottMembership list
	rows, err := c.QueryContext(ctx, "SELECT id,title FROM BoycottMembership")
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from BoycottMembership-> "+err.Error())
	}
	defer rows.Close()

	list := []*v1.BoycottMembership{}
	for rows.Next() {
		td := new(v1.BoycottMembership)
		if err := rows.Scan(&td.Id, &td.Title); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from BoycottMembership row-> "+err.Error())
		}
		list = append(list, td)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from BoycottMembership-> "+err.Error())
	}

	return &v1.ListBoycottMembershipResponse{
		Api:   apiVersion,
		Items: list,
	}, nil
}

// Update todo task
func (s *shrikeServiceServer) UpdateBoycottMembership(ctx context.Context, req *v1.UpdateBoycottMembershipRequest) (*v1.UpdateBoycottMembershipResponse, error) {
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

	// update boycottmembership
	res, err := c.ExecContext(ctx, "UPDATE boycottmembership SET title=$1 WHERE id=$2",
		req.Item.Title, req.Item.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update boycottmembership-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("boycottmembership with ID='%d' is not found",
			req.Item.Id))
	}

	return &v1.UpdateBoycottMembershipResponse{
		Api:     apiVersion,
		Updated: rows,
	}, nil
}

// Delete boycottmembership
func (s *shrikeServiceServer) DeleteBoycottMembership(ctx context.Context, req *v1.DeleteBoycottMembershipRequest) (*v1.DeleteBoycottMembershipResponse, error) {
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

	// delete boycottmembership
	res, err := c.ExecContext(ctx, "DELETE FROM boycottmembership WHERE id=$1", req.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete boycottmembership-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("boycottmembership with ID='%d' is not found",
			req.Id))
	}

	return &v1.DeleteBoycottMembershipResponse{
		Api:     apiVersion,
		Deleted: rows,
	}, nil
}
