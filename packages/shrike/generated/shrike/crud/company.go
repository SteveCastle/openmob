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

// NewShrikeServiceServer creates Company service
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
func (s *shrikeServiceServer) CreateCompany(ctx context.Context, req *v1.CreateCompanyRequest) (*v1.CreateCompanyResponse, error) {
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
	// insert Company entity data
	err = c.QueryRowContext(ctx, "INSERT INTO company (title) VALUES($1)  RETURNING id;",
		req.Item.Title).Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into Company-> "+err.Error())
	}

	// get ID of creates Company
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created Company-> "+err.Error())
	}

	return &v1.CreateCompanyResponse{
		Api: apiVersion,
		Id:  id,
	}, nil
}

// Get company by id.
func (s *shrikeServiceServer) GetCompany(ctx context.Context, req *v1.GetCompanyRequest) (*v1.GetCompanyResponse, error) {
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

	// query Company by ID
	rows, err := c.QueryContext(ctx, "SELECT id, title FROM company WHERE id=$1",
		req.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from Company-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from Company-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Company with ID='%d' is not found",
			req.Id))
	}

	// get Company data
	var td v1.Company
	if err := rows.Scan(&td.Id, &td.Title); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from Company row-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple Company rows with ID='%d'",
			req.Id))
	}

	return &v1.GetCompanyResponse{
		Api:  apiVersion,
		Item: &td,
	}, nil

}

// Read all todo tasks
func (s *shrikeServiceServer) ListCompany(ctx context.Context, req *v1.ListCompanyRequest) (*v1.ListCompanyResponse, error) {
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

	// get Company list
	rows, err := c.QueryContext(ctx, "SELECT id,title FROM Company")
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from Company-> "+err.Error())
	}
	defer rows.Close()

	list := []*v1.Company{}
	for rows.Next() {
		td := new(v1.Company)
		if err := rows.Scan(&td.Id, &td.Title); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from Company row-> "+err.Error())
		}
		list = append(list, td)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from Company-> "+err.Error())
	}

	return &v1.ListCompanyResponse{
		Api:   apiVersion,
		Items: list,
	}, nil
}

// Update todo task
func (s *shrikeServiceServer) UpdateCompany(ctx context.Context, req *v1.UpdateCompanyRequest) (*v1.UpdateCompanyResponse, error) {
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

	// update company
	res, err := c.ExecContext(ctx, "UPDATE company SET title=$1 WHERE id=$2",
		req.Item.Title, req.Item.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update company-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("company with ID='%d' is not found",
			req.Item.Id))
	}

	return &v1.UpdateCompanyResponse{
		Api:     apiVersion,
		Updated: rows,
	}, nil
}

// Delete company
func (s *shrikeServiceServer) DeleteCompany(ctx context.Context, req *v1.DeleteCompanyRequest) (*v1.DeleteCompanyResponse, error) {
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

	// delete company
	res, err := c.ExecContext(ctx, "DELETE FROM company WHERE id=$1", req.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete company-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("company with ID='%d' is not found",
			req.Id))
	}

	return &v1.DeleteCompanyResponse{
		Api:     apiVersion,
		Deleted: rows,
	}, nil
}
