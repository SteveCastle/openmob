package v1

import (
	"context"
	"fmt"

	v1 "github.com/SteveCastle/openmob/packages/shrike/src/pkg/api/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Create new Company
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
	err = c.QueryRowContext(ctx, "INSERT INTO company (id, created_at, updated_at, title) VALUES($1, $2, $3, $4)  RETURNING id;",
		 req.Item.ID,  req.Item.CreatedAt,  req.Item.UpdatedAt,  req.Item.Title, ).Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into Company-> "+err.Error())
	}

	// get ID of creates Company
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created Company-> "+err.Error())
	}

	return &v1.CreateCompanyResponse{
		Api: apiVersion,
		ID:  id,
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
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, title FROM company WHERE id=$1",
		req.ID)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from Company-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from Company-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Company with ID='%d' is not found",
			req.ID))
	}

	// get Company data
	var company v1.Company
	if err := rows.Scan( &company.ID,  &company.CreatedAt,  &company.UpdatedAt,  &company.Title, ); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from Company row-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple Company rows with ID='%d'",
			req.ID))
	}

	return &v1.GetCompanyResponse{
		Api:  apiVersion,
		Item: &company,
	}, nil

}

// Read all Company
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
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, title FROM company")
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from Company-> "+err.Error())
	}
	defer rows.Close()

	list := []*v1.Company{}
	for rows.Next() {
		company := new(v1.Company)
		if err := rows.Scan( &company.ID,  &company.CreatedAt,  &company.UpdatedAt,  &company.Title, ); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from Company row-> "+err.Error())
		}
		list = append(list, company)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from Company-> "+err.Error())
	}

	return &v1.ListCompanyResponse{
		Api:   apiVersion,
		Items: list,
	}, nil
}

// Update Company
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
	res, err := c.ExecContext(ctx, "UPDATE company SET id=$1, created_at=$2, updated_at=$3, title=$4 WHERE id=$1",
		req.Item.ID,req.Item.CreatedAt,req.Item.UpdatedAt,req.Item.Title, )
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update Company-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Company with ID='%d' is not found",
			req.Item.ID))
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
	res, err := c.ExecContext(ctx, "DELETE FROM company WHERE id=$1", req.ID)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete Company-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Company with ID='%d' is not found",
			req.ID))
	}

	return &v1.DeleteCompanyResponse{
		Api:     apiVersion,
		Deleted: rows,
	}, nil
}