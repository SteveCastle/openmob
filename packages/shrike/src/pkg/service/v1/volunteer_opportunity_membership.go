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

// Create new VolunteerOpportunityMembership
func (s *shrikeServiceServer) CreateVolunteerOpportunityMembership(ctx context.Context, req *v1.CreateVolunteerOpportunityMembershipRequest) (*v1.CreateVolunteerOpportunityMembershipResponse, error) {
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
	// insert VolunteerOpportunityMembership entity data
	err = c.QueryRowContext(ctx, "INSERT INTO volunteer_opportunity_membership (cause, volunteer_opportunity) VALUES($1, $2)  RETURNING id;",
		req.Item.Cause, req.Item.VolunteerOpportunity).Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into VolunteerOpportunityMembership-> "+err.Error())
	}

	// get ID of creates VolunteerOpportunityMembership
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created VolunteerOpportunityMembership-> "+err.Error())
	}

	return &v1.CreateVolunteerOpportunityMembershipResponse{
		Api: apiVersion,
		ID:  id,
	}, nil
}

// Get volunteer_opportunity_membership by id.
func (s *shrikeServiceServer) GetVolunteerOpportunityMembership(ctx context.Context, req *v1.GetVolunteerOpportunityMembershipRequest) (*v1.GetVolunteerOpportunityMembershipResponse, error) {
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

	// query VolunteerOpportunityMembership by ID
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, cause, volunteer_opportunity FROM volunteer_opportunity_membership WHERE id=$1",
		req.ID)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from VolunteerOpportunityMembership-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from VolunteerOpportunityMembership-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("VolunteerOpportunityMembership with ID='%s' is not found",
			req.ID))
	}

	// scan VolunteerOpportunityMembership data into protobuf model
	var volunteeropportunitymembership v1.VolunteerOpportunityMembership
	var createdAt pq.NullTime
	var updatedAt pq.NullTime

	if err := rows.Scan(&volunteeropportunitymembership.ID, &createdAt, &updatedAt, &volunteeropportunitymembership.Cause, &volunteeropportunitymembership.VolunteerOpportunity); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from VolunteerOpportunityMembership row-> "+err.Error())
	}

	// Convert pq.NullTime from database into proto timestamp.
	if createdAt.Valid {
		volunteeropportunitymembership.CreatedAt, err = ptypes.TimestampProto(createdAt.Time)
		if err != nil {
			return nil, status.Error(codes.Unknown, "createdAt field has invalid format-> "+err.Error())
		}
	}
	if updatedAt.Valid {
		volunteeropportunitymembership.UpdatedAt, err = ptypes.TimestampProto(updatedAt.Time)
		if err != nil {
			return nil, status.Error(codes.Unknown, "createdAt field has invalid format-> "+err.Error())
		}
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple VolunteerOpportunityMembership rows with ID='%s'",
			req.ID))
	}

	return &v1.GetVolunteerOpportunityMembershipResponse{
		Api:  apiVersion,
		Item: &volunteeropportunitymembership,
	}, nil

}

// Read all VolunteerOpportunityMembership
func (s *shrikeServiceServer) ListVolunteerOpportunityMembership(ctx context.Context, req *v1.ListVolunteerOpportunityMembershipRequest) (*v1.ListVolunteerOpportunityMembershipResponse, error) {
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

	// Generate SQL to select all columns in VolunteerOpportunityMembership Table
	// Then generate filtering and ordering sql and finally run query.
	querySQL := queries.BuildVolunteerOpportunityMembershipListQuery(req.Filters, req.Ordering, req.Limit)
	// Execute query and scan into return type.
	rows, err := c.QueryContext(ctx, querySQL)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from VolunteerOpportunityMembership-> "+err.Error())
	}
	defer rows.Close()

	// Variables to store results returned by database.
	list := []*v1.VolunteerOpportunityMembership{}
	var createdAt pq.NullTime
	var updatedAt pq.NullTime

	for rows.Next() {
		volunteeropportunitymembership := new(v1.VolunteerOpportunityMembership)
		if err := rows.Scan(&volunteeropportunitymembership.ID, &createdAt, &updatedAt, &volunteeropportunitymembership.Cause, &volunteeropportunitymembership.VolunteerOpportunity); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from VolunteerOpportunityMembership row-> "+err.Error())
		}
		// Convert pq.NullTime from database into proto timestamp.
		if createdAt.Valid {
			volunteeropportunitymembership.CreatedAt, err = ptypes.TimestampProto(createdAt.Time)
			if err != nil {
				return nil, status.Error(codes.Unknown, "createdAt field has invalid format-> "+err.Error())
			}
		}
		if updatedAt.Valid {
			volunteeropportunitymembership.UpdatedAt, err = ptypes.TimestampProto(updatedAt.Time)
			if err != nil {
				return nil, status.Error(codes.Unknown, "createdAt field has invalid format-> "+err.Error())
			}
		}

		list = append(list, volunteeropportunitymembership)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from VolunteerOpportunityMembership-> "+err.Error())
	}

	return &v1.ListVolunteerOpportunityMembershipResponse{
		Api:   apiVersion,
		Items: list,
	}, nil
}

// Update VolunteerOpportunityMembership
func (s *shrikeServiceServer) UpdateVolunteerOpportunityMembership(ctx context.Context, req *v1.UpdateVolunteerOpportunityMembershipRequest) (*v1.UpdateVolunteerOpportunityMembershipResponse, error) {
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

	// update volunteer_opportunity_membership
	res, err := c.ExecContext(ctx, "UPDATE volunteer_opportunity_membership SET cause=$2, volunteer_opportunity=$3 WHERE id=$1",
		req.Item.ID, req.Item.Cause, req.Item.VolunteerOpportunity)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update VolunteerOpportunityMembership-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("VolunteerOpportunityMembership with ID='%s' is not found",
			req.Item.ID))
	}

	return &v1.UpdateVolunteerOpportunityMembershipResponse{
		Api:     apiVersion,
		Updated: rows,
	}, nil
}

// Delete volunteer_opportunity_membership
func (s *shrikeServiceServer) DeleteVolunteerOpportunityMembership(ctx context.Context, req *v1.DeleteVolunteerOpportunityMembershipRequest) (*v1.DeleteVolunteerOpportunityMembershipResponse, error) {
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

	// delete volunteer_opportunity_membership
	res, err := c.ExecContext(ctx, "DELETE FROM volunteer_opportunity_membership WHERE id=$1", req.ID)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete VolunteerOpportunityMembership-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("VolunteerOpportunityMembership with ID='%s' is not found",
			req.ID))
	}

	return &v1.DeleteVolunteerOpportunityMembershipResponse{
		Api:     apiVersion,
		Deleted: rows,
	}, nil
}
