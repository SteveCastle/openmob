package v1

import (
	"context"
	"fmt"

	v1 "github.com/SteveCastle/openmob/packages/shrike/src/pkg/api/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Create new Experiment
func (s *shrikeServiceServer) CreateExperiment(ctx context.Context, req *v1.CreateExperimentRequest) (*v1.CreateExperimentResponse, error) {
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
	// insert Experiment entity data
	err = c.QueryRowContext(ctx, "INSERT INTO experiment (id, created_at, updated_at, title, landing_page) VALUES($1, $2, $3, $4, $5)  RETURNING id;",
		 req.Item.ID,  req.Item.CreatedAt,  req.Item.UpdatedAt,  req.Item.Title,  req.Item.LandingPage, ).Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into Experiment-> "+err.Error())
	}

	// get ID of creates Experiment
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created Experiment-> "+err.Error())
	}

	return &v1.CreateExperimentResponse{
		Api: apiVersion,
		ID:  id,
	}, nil
}

// Get experiment by id.
func (s *shrikeServiceServer) GetExperiment(ctx context.Context, req *v1.GetExperimentRequest) (*v1.GetExperimentResponse, error) {
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

	// query Experiment by ID
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, title, landing_page FROM experiment WHERE id=$1",
		req.ID)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from Experiment-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from Experiment-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Experiment with ID='%d' is not found",
			req.ID))
	}

	// get Experiment data
	var experiment v1.Experiment
	if err := rows.Scan( &experiment.ID,  &experiment.CreatedAt,  &experiment.UpdatedAt,  &experiment.Title,  &experiment.LandingPage, ); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from Experiment row-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple Experiment rows with ID='%d'",
			req.ID))
	}

	return &v1.GetExperimentResponse{
		Api:  apiVersion,
		Item: &experiment,
	}, nil

}

// Read all Experiment
func (s *shrikeServiceServer) ListExperiment(ctx context.Context, req *v1.ListExperimentRequest) (*v1.ListExperimentResponse, error) {
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

	// get Experiment list
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, title, landing_page FROM experiment")
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from Experiment-> "+err.Error())
	}
	defer rows.Close()

	list := []*v1.Experiment{}
	for rows.Next() {
		experiment := new(v1.Experiment)
		if err := rows.Scan( &experiment.ID,  &experiment.CreatedAt,  &experiment.UpdatedAt,  &experiment.Title,  &experiment.LandingPage, ); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from Experiment row-> "+err.Error())
		}
		list = append(list, experiment)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from Experiment-> "+err.Error())
	}

	return &v1.ListExperimentResponse{
		Api:   apiVersion,
		Items: list,
	}, nil
}

// Update Experiment
func (s *shrikeServiceServer) UpdateExperiment(ctx context.Context, req *v1.UpdateExperimentRequest) (*v1.UpdateExperimentResponse, error) {
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

	// update experiment
	res, err := c.ExecContext(ctx, "UPDATE experiment SET id=$1, created_at=$2, updated_at=$3, title=$4, landing_page=$5 WHERE id=$1",
		req.Item.ID,req.Item.CreatedAt,req.Item.UpdatedAt,req.Item.Title,req.Item.LandingPage, )
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update Experiment-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Experiment with ID='%d' is not found",
			req.Item.ID))
	}

	return &v1.UpdateExperimentResponse{
		Api:     apiVersion,
		Updated: rows,
	}, nil
}

// Delete experiment
func (s *shrikeServiceServer) DeleteExperiment(ctx context.Context, req *v1.DeleteExperimentRequest) (*v1.DeleteExperimentResponse, error) {
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

	// delete experiment
	res, err := c.ExecContext(ctx, "DELETE FROM experiment WHERE id=$1", req.ID)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete Experiment-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Experiment with ID='%d' is not found",
			req.ID))
	}

	return &v1.DeleteExperimentResponse{
		Api:     apiVersion,
		Deleted: rows,
	}, nil
}