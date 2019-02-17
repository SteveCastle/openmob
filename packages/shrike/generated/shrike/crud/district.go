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

// NewShrikeServiceServer creates District service
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

// Create new District
func (s *shrikeServiceServer) CreateDistrict(ctx context.Context, req *v1.CreateDistrictRequest) (*v1.CreateDistrictResponse, error) {
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
	// insert District entity data
	err = c.QueryRowContext(ctx, "INSERT INTO district ( id  created_at  updated_at  geom  title  district_type ) VALUES( $1 $2 $3 $4 $5 $6)  RETURNING id;",
		 req.ItemID  req.ItemCreatedAt  req.ItemUpdatedAt  req.ItemGeom  req.ItemTitle  req.ItemDistrictType ).Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into District-> "+err.Error())
	}

	// get ID of creates District
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created District-> "+err.Error())
	}

	return &v1.CreateDistrictResponse{
		Api: apiVersion,
		Id:  id,
	}, nil
}

// Get district by id.
func (s *shrikeServiceServer) GetDistrict(ctx context.Context, req *v1.GetDistrictRequest) (*v1.GetDistrictResponse, error) {
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

	// query District by ID
	rows, err := c.QueryContext(ctx, "SELECT id, title FROM district WHERE id=$1",
		req.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from District-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from District-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("District with ID='%d' is not found",
			req.Id))
	}

	// get District data
	var district v1.District
	if err := rows.Scan(&district.Id, &district.Title); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from District row-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple District rows with ID='%d'",
			req.Id))
	}

	return &v1.GetDistrictResponse{
		Api:  apiVersion,
		Item: &district,
	}, nil

}

// Read all District
func (s *shrikeServiceServer) ListDistrict(ctx context.Context, req *v1.ListDistrictRequest) (*v1.ListDistrictResponse, error) {
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

	// get District list
	rows, err := c.QueryContext(ctx, "SELECT id,title FROM district")
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from District-> "+err.Error())
	}
	defer rows.Close()

	list := []*v1.District{}
	for rows.Next() {
		district := new(v1.District)
		if err := rows.Scan(&district.Id, &district.Title); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from District row-> "+err.Error())
		}
		list = append(list, district)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from District-> "+err.Error())
	}

	return &v1.ListDistrictResponse{
		Api:   apiVersion,
		Items: list,
	}, nil
}

// Update District
func (s *shrikeServiceServer) UpdateDistrict(ctx context.Context, req *v1.UpdateDistrictRequest) (*v1.UpdateDistrictResponse, error) {
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

	// update district
	res, err := c.ExecContext(ctx, "UPDATE district SET title=$1 WHERE id=$2",
		req.Item.Title, req.Item.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update District-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("District with ID='%d' is not found",
			req.Item.Id))
	}

	return &v1.UpdateDistrictResponse{
		Api:     apiVersion,
		Updated: rows,
	}, nil
}

// Delete district
func (s *shrikeServiceServer) DeleteDistrict(ctx context.Context, req *v1.DeleteDistrictRequest) (*v1.DeleteDistrictResponse, error) {
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

	// delete district
	res, err := c.ExecContext(ctx, "DELETE FROM district WHERE id=$1", req.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete District-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("District with ID='%d' is not found",
			req.Id))
	}

	return &v1.DeleteDistrictResponse{
		Api:     apiVersion,
		Deleted: rows,
	}, nil
}
