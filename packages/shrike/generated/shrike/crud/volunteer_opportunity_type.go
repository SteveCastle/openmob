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

// NewShrikeServiceServer creates VolunteerOpportunityType service
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

// Create new VolunteerOpportunityType
func (s *shrikeServiceServer) CreateVolunteerOpportunityType(ctx context.Context, req *v1.CreateVolunteerOpportunityTypeRequest) (*v1.CreateVolunteerOpportunityTypeResponse, error) {
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
	// insert VolunteerOpportunityType entity data
	err = c.QueryRowContext(ctx, "INSERT INTO volunteer_opportunity_type ( id  created_at  updated_at  title ) VALUES( $1 $2 $3 $4)  RETURNING id;",
		 req.ItemID  req.ItemCreatedAt  req.ItemUpdatedAt  req.ItemTitle ).Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into VolunteerOpportunityType-> "+err.Error())
	}

	// get ID of creates VolunteerOpportunityType
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created VolunteerOpportunityType-> "+err.Error())
	}

	return &v1.CreateVolunteerOpportunityTypeResponse{
		Api: apiVersion,
		Id:  id,
	}, nil
}

// Get volunteer_opportunity_type by id.
func (s *shrikeServiceServer) GetVolunteerOpportunityType(ctx context.Context, req *v1.GetVolunteerOpportunityTypeRequest) (*v1.GetVolunteerOpportunityTypeResponse, error) {
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

	// query VolunteerOpportunityType by ID
	rows, err := c.QueryContext(ctx, "SELECT id, title FROM volunteer_opportunity_type WHERE id=$1",
		req.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from VolunteerOpportunityType-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from VolunteerOpportunityType-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("VolunteerOpportunityType with ID='%d' is not found",
			req.Id))
	}

	// get VolunteerOpportunityType data
	var volunteeropportunitytype v1.VolunteerOpportunityType
	if err := rows.Scan(&volunteeropportunitytype.Id, &volunteeropportunitytype.Title); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from VolunteerOpportunityType row-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple VolunteerOpportunityType rows with ID='%d'",
			req.Id))
	}

	return &v1.GetVolunteerOpportunityTypeResponse{
		Api:  apiVersion,
		Item: &volunteeropportunitytype,
	}, nil

}

// Read all VolunteerOpportunityType
func (s *shrikeServiceServer) ListVolunteerOpportunityType(ctx context.Context, req *v1.ListVolunteerOpportunityTypeRequest) (*v1.ListVolunteerOpportunityTypeResponse, error) {
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

	// get VolunteerOpportunityType list
	rows, err := c.QueryContext(ctx, "SELECT id,title FROM volunteer_opportunity_type")
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from VolunteerOpportunityType-> "+err.Error())
	}
	defer rows.Close()

	list := []*v1.VolunteerOpportunityType{}
	for rows.Next() {
		volunteeropportunitytype := new(v1.VolunteerOpportunityType)
		if err := rows.Scan(&volunteeropportunitytype.Id, &volunteeropportunitytype.Title); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from VolunteerOpportunityType row-> "+err.Error())
		}
		list = append(list, volunteeropportunitytype)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from VolunteerOpportunityType-> "+err.Error())
	}

	return &v1.ListVolunteerOpportunityTypeResponse{
		Api:   apiVersion,
		Items: list,
	}, nil
}

// Update VolunteerOpportunityType
func (s *shrikeServiceServer) UpdateVolunteerOpportunityType(ctx context.Context, req *v1.UpdateVolunteerOpportunityTypeRequest) (*v1.UpdateVolunteerOpportunityTypeResponse, error) {
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

	// update volunteer_opportunity_type
	res, err := c.ExecContext(ctx, "UPDATE volunteer_opportunity_type SET title=$1 WHERE id=$2",
		req.Item.Title, req.Item.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update VolunteerOpportunityType-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("VolunteerOpportunityType with ID='%d' is not found",
			req.Item.Id))
	}

	return &v1.UpdateVolunteerOpportunityTypeResponse{
		Api:     apiVersion,
		Updated: rows,
	}, nil
}

// Delete volunteer_opportunity_type
func (s *shrikeServiceServer) DeleteVolunteerOpportunityType(ctx context.Context, req *v1.DeleteVolunteerOpportunityTypeRequest) (*v1.DeleteVolunteerOpportunityTypeResponse, error) {
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

	// delete volunteer_opportunity_type
	res, err := c.ExecContext(ctx, "DELETE FROM volunteer_opportunity_type WHERE id=$1", req.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete VolunteerOpportunityType-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("VolunteerOpportunityType with ID='%d' is not found",
			req.Id))
	}

	return &v1.DeleteVolunteerOpportunityTypeResponse{
		Api:     apiVersion,
		Deleted: rows,
	}, nil
}
