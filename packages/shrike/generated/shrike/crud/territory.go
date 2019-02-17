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

// NewShrikeServiceServer creates Territory service
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

// Create new Territory
func (s *shrikeServiceServer) CreateTerritory(ctx context.Context, req *v1.CreateTerritoryRequest) (*v1.CreateTerritoryResponse, error) {
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
	// insert Territory entity data
	err = c.QueryRowContext(ctx, "INSERT INTO territory ( id  created_at  updated_at  title ) VALUES( $1 $2 $3 $4)  RETURNING id;",
		 req.ItemID  req.ItemCreatedAt  req.ItemUpdatedAt  req.ItemTitle ).Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into Territory-> "+err.Error())
	}

	// get ID of creates Territory
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created Territory-> "+err.Error())
	}

	return &v1.CreateTerritoryResponse{
		Api: apiVersion,
		Id:  id,
	}, nil
}

// Get territory by id.
func (s *shrikeServiceServer) GetTerritory(ctx context.Context, req *v1.GetTerritoryRequest) (*v1.GetTerritoryResponse, error) {
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

	// query Territory by ID
	rows, err := c.QueryContext(ctx, "SELECT id, title FROM territory WHERE id=$1",
		req.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from Territory-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from Territory-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Territory with ID='%d' is not found",
			req.Id))
	}

	// get Territory data
	var territory v1.Territory
	if err := rows.Scan(&territory.Id, &territory.Title); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from Territory row-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple Territory rows with ID='%d'",
			req.Id))
	}

	return &v1.GetTerritoryResponse{
		Api:  apiVersion,
		Item: &territory,
	}, nil

}

// Read all Territory
func (s *shrikeServiceServer) ListTerritory(ctx context.Context, req *v1.ListTerritoryRequest) (*v1.ListTerritoryResponse, error) {
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

	// get Territory list
	rows, err := c.QueryContext(ctx, "SELECT id,title FROM territory")
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from Territory-> "+err.Error())
	}
	defer rows.Close()

	list := []*v1.Territory{}
	for rows.Next() {
		territory := new(v1.Territory)
		if err := rows.Scan(&territory.Id, &territory.Title); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from Territory row-> "+err.Error())
		}
		list = append(list, territory)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from Territory-> "+err.Error())
	}

	return &v1.ListTerritoryResponse{
		Api:   apiVersion,
		Items: list,
	}, nil
}

// Update Territory
func (s *shrikeServiceServer) UpdateTerritory(ctx context.Context, req *v1.UpdateTerritoryRequest) (*v1.UpdateTerritoryResponse, error) {
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

	// update territory
	res, err := c.ExecContext(ctx, "UPDATE territory SET title=$1 WHERE id=$2",
		req.Item.Title, req.Item.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update Territory-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Territory with ID='%d' is not found",
			req.Item.Id))
	}

	return &v1.UpdateTerritoryResponse{
		Api:     apiVersion,
		Updated: rows,
	}, nil
}

// Delete territory
func (s *shrikeServiceServer) DeleteTerritory(ctx context.Context, req *v1.DeleteTerritoryRequest) (*v1.DeleteTerritoryResponse, error) {
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

	// delete territory
	res, err := c.ExecContext(ctx, "DELETE FROM territory WHERE id=$1", req.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete Territory-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Territory with ID='%d' is not found",
			req.Id))
	}

	return &v1.DeleteTerritoryResponse{
		Api:     apiVersion,
		Deleted: rows,
	}, nil
}
