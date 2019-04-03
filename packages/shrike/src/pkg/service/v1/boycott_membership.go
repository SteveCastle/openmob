package v1

import (
	"context"
	"fmt"

	v1 "github.com/SteveCastle/openmob/packages/shrike/src/pkg/api/v1"
	"github.com/SteveCastle/openmob/packages/shrike/src/pkg/queries"
	"github.com/golang/protobuf/ptypes"
	"github.com/lib/pq"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Create new BoycottMembership
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
	var id string
	// insert BoycottMembership entity data
	err = c.QueryRowContext(ctx, "INSERT INTO boycott_membership (cause, boycott) VALUES($1, $2)  RETURNING id;",
		req.Item.Cause, req.Item.Boycott).Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into BoycottMembership-> "+err.Error())
	}

	// get ID of creates BoycottMembership
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created BoycottMembership-> "+err.Error())
	}

	return &v1.CreateBoycottMembershipResponse{
		Api: apiVersion,
		ID:  id,
	}, nil
}

// Get boycott_membership by id.
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
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, cause, boycott FROM boycott_membership WHERE id=$1",
		req.ID)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from BoycottMembership-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from BoycottMembership-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("BoycottMembership with ID='%s' is not found",
			req.ID))
	}

	// scan BoycottMembership data into protobuf model
	var boycottmembership v1.BoycottMembership
	var createdAt pq.NullTime
	var updatedAt pq.NullTime

	if err := rows.Scan(&boycottmembership.ID, &createdAt, &updatedAt, &boycottmembership.Cause, &boycottmembership.Boycott); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from BoycottMembership row-> "+err.Error())
	}

	// Convert pq.NullTime from database into proto timestamp.
	if createdAt.Valid {
		boycottmembership.CreatedAt, err = ptypes.TimestampProto(createdAt.Time)
		if err != nil {
			return nil, status.Error(codes.Unknown, "createdAt field has invalid format-> "+err.Error())
		}
	}
	if updatedAt.Valid {
		boycottmembership.UpdatedAt, err = ptypes.TimestampProto(updatedAt.Time)
		if err != nil {
			return nil, status.Error(codes.Unknown, "createdAt field has invalid format-> "+err.Error())
		}
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple BoycottMembership rows with ID='%s'",
			req.ID))
	}

	return &v1.GetBoycottMembershipResponse{
		Api:  apiVersion,
		Item: &boycottmembership,
	}, nil

}

// Read all BoycottMembership
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

	// Generate SQL to select all columns in BoycottMembership Table
	// Then generate filtering and ordering sql and finally run query.
	querySQL := queries.BuildBoycottMembershipListQuery(req.Filters, req.Ordering, req.Limit)
	// Execute query and scan into return type.
	rows, err := c.QueryContext(ctx, querySQL)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from BoycottMembership-> "+err.Error())
	}
	defer rows.Close()

	// Variables to store results returned by database.
	list := []*v1.BoycottMembership{}
	var createdAt pq.NullTime
	var updatedAt pq.NullTime

	for rows.Next() {
		boycottmembership := new(v1.BoycottMembership)
		if err := rows.Scan(&boycottmembership.ID, &createdAt, &updatedAt, &boycottmembership.Cause, &boycottmembership.Boycott); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from BoycottMembership row-> "+err.Error())
		}
		// Convert pq.NullTime from database into proto timestamp.
		if createdAt.Valid {
			boycottmembership.CreatedAt, err = ptypes.TimestampProto(createdAt.Time)
			if err != nil {
				return nil, status.Error(codes.Unknown, "createdAt field has invalid format-> "+err.Error())
			}
		}
		if updatedAt.Valid {
			boycottmembership.UpdatedAt, err = ptypes.TimestampProto(updatedAt.Time)
			if err != nil {
				return nil, status.Error(codes.Unknown, "createdAt field has invalid format-> "+err.Error())
			}
		}

		list = append(list, boycottmembership)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from BoycottMembership-> "+err.Error())
	}

	return &v1.ListBoycottMembershipResponse{
		Api:   apiVersion,
		Items: list,
	}, nil
}

// Update BoycottMembership
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

	// update boycott_membership
	res, err := c.ExecContext(ctx, "UPDATE boycott_membership SET cause=$2, boycott=$3 WHERE id=$1",
		req.Item.ID, req.Item.Cause, req.Item.Boycott)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update BoycottMembership-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("BoycottMembership with ID='%s' is not found",
			req.Item.ID))
	}

	return &v1.UpdateBoycottMembershipResponse{
		Api:     apiVersion,
		Updated: rows,
	}, nil
}

// Delete boycott_membership
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

	// delete boycott_membership
	res, err := c.ExecContext(ctx, "DELETE FROM boycott_membership WHERE id=$1", req.ID)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete BoycottMembership-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("BoycottMembership with ID='%s' is not found",
			req.ID))
	}

	return &v1.DeleteBoycottMembershipResponse{
		Api:     apiVersion,
		Deleted: rows,
	}, nil
}
