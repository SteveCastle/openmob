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

// NewShrikeServiceServer creates OwnerMembership service
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

// Create new OwnerMembership
func (s *shrikeServiceServer) CreateOwnerMembership(ctx context.Context, req *v1.CreateOwnerMembershipRequest) (*v1.CreateOwnerMembershipResponse, error) {
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
	// insert OwnerMembership entity data
	err = c.QueryRowContext(ctx, "INSERT INTO owner_membership ( id  created_at  updated_at  cause  account ) VALUES( $1 $2 $3 $4 $5)  RETURNING id;",
		 req.ItemID  req.ItemCreatedAt  req.ItemUpdatedAt  req.ItemCause  req.ItemAccount ).Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into OwnerMembership-> "+err.Error())
	}

	// get ID of creates OwnerMembership
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created OwnerMembership-> "+err.Error())
	}

	return &v1.CreateOwnerMembershipResponse{
		Api: apiVersion,
		Id:  id,
	}, nil
}

// Get owner_membership by id.
func (s *shrikeServiceServer) GetOwnerMembership(ctx context.Context, req *v1.GetOwnerMembershipRequest) (*v1.GetOwnerMembershipResponse, error) {
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

	// query OwnerMembership by ID
	rows, err := c.QueryContext(ctx, "SELECT id, title FROM owner_membership WHERE id=$1",
		req.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from OwnerMembership-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from OwnerMembership-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("OwnerMembership with ID='%d' is not found",
			req.Id))
	}

	// get OwnerMembership data
	var ownermembership v1.OwnerMembership
	if err := rows.Scan(&ownermembership.Id, &ownermembership.Title); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from OwnerMembership row-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple OwnerMembership rows with ID='%d'",
			req.Id))
	}

	return &v1.GetOwnerMembershipResponse{
		Api:  apiVersion,
		Item: &ownermembership,
	}, nil

}

// Read all OwnerMembership
func (s *shrikeServiceServer) ListOwnerMembership(ctx context.Context, req *v1.ListOwnerMembershipRequest) (*v1.ListOwnerMembershipResponse, error) {
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

	// get OwnerMembership list
	rows, err := c.QueryContext(ctx, "SELECT id,title FROM owner_membership")
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from OwnerMembership-> "+err.Error())
	}
	defer rows.Close()

	list := []*v1.OwnerMembership{}
	for rows.Next() {
		ownermembership := new(v1.OwnerMembership)
		if err := rows.Scan(&ownermembership.Id, &ownermembership.Title); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from OwnerMembership row-> "+err.Error())
		}
		list = append(list, ownermembership)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from OwnerMembership-> "+err.Error())
	}

	return &v1.ListOwnerMembershipResponse{
		Api:   apiVersion,
		Items: list,
	}, nil
}

// Update OwnerMembership
func (s *shrikeServiceServer) UpdateOwnerMembership(ctx context.Context, req *v1.UpdateOwnerMembershipRequest) (*v1.UpdateOwnerMembershipResponse, error) {
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

	// update owner_membership
	res, err := c.ExecContext(ctx, "UPDATE owner_membership SET title=$1 WHERE id=$2",
		req.Item.Title, req.Item.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update OwnerMembership-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("OwnerMembership with ID='%d' is not found",
			req.Item.Id))
	}

	return &v1.UpdateOwnerMembershipResponse{
		Api:     apiVersion,
		Updated: rows,
	}, nil
}

// Delete owner_membership
func (s *shrikeServiceServer) DeleteOwnerMembership(ctx context.Context, req *v1.DeleteOwnerMembershipRequest) (*v1.DeleteOwnerMembershipResponse, error) {
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

	// delete owner_membership
	res, err := c.ExecContext(ctx, "DELETE FROM owner_membership WHERE id=$1", req.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete OwnerMembership-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("OwnerMembership with ID='%d' is not found",
			req.Id))
	}

	return &v1.DeleteOwnerMembershipResponse{
		Api:     apiVersion,
		Deleted: rows,
	}, nil
}
