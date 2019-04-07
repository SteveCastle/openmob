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

// Create new ExperimentVariant
func (s *shrikeServiceServer) CreateExperimentVariant(ctx context.Context, req *v1.CreateExperimentVariantRequest) (*v1.CreateExperimentVariantResponse, error) {
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
	// insert ExperimentVariant entity data
	err = c.QueryRowContext(ctx, "INSERT INTO experiment_variant (title, variant_type, experiment, landing_page, field, component) VALUES($1, $2, $3, $4, $5, $6)  RETURNING id;",
		req.Item.Title, req.Item.VariantType, req.Item.Experiment, req.Item.LandingPage, req.Item.Field, req.Item.Component).Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into ExperimentVariant-> "+err.Error())
	}

	// get ID of creates ExperimentVariant
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created ExperimentVariant-> "+err.Error())
	}

	return &v1.CreateExperimentVariantResponse{
		Api: apiVersion,
		ID:  id,
	}, nil
}

// Get experiment_variant by id.
func (s *shrikeServiceServer) GetExperimentVariant(ctx context.Context, req *v1.GetExperimentVariantRequest) (*v1.GetExperimentVariantResponse, error) {
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

	// query ExperimentVariant by ID
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, title, variant_type, experiment, landing_page, field, component FROM experiment_variant WHERE id=$1",
		req.ID)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from ExperimentVariant-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from ExperimentVariant-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("ExperimentVariant with ID='%s' is not found",
			req.ID))
	}

	// scan ExperimentVariant data into protobuf model
	var experimentvariant v1.ExperimentVariant
	var createdAt pq.NullTime
	var updatedAt pq.NullTime

	if err := rows.Scan(&experimentvariant.ID, &createdAt, &updatedAt, &experimentvariant.Title, &experimentvariant.VariantType, &experimentvariant.Experiment, &experimentvariant.LandingPage, &experimentvariant.Field, &experimentvariant.Component); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from ExperimentVariant row-> "+err.Error())
	}

	// Convert pq.NullTime from database into proto timestamp.
	if createdAt.Valid {
		experimentvariant.CreatedAt, err = ptypes.TimestampProto(createdAt.Time)
		if err != nil {
			return nil, status.Error(codes.Unknown, "createdAt field has invalid format-> "+err.Error())
		}
	}
	if updatedAt.Valid {
		experimentvariant.UpdatedAt, err = ptypes.TimestampProto(updatedAt.Time)
		if err != nil {
			return nil, status.Error(codes.Unknown, "createdAt field has invalid format-> "+err.Error())
		}
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple ExperimentVariant rows with ID='%s'",
			req.ID))
	}

	return &v1.GetExperimentVariantResponse{
		Api:  apiVersion,
		Item: &experimentvariant,
	}, nil

}

// Read all ExperimentVariant
func (s *shrikeServiceServer) ListExperimentVariant(ctx context.Context, req *v1.ListExperimentVariantRequest) (*v1.ListExperimentVariantResponse, error) {
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

	// Generate SQL to select all columns in ExperimentVariant Table
	// Then generate filtering and ordering sql and finally run query.
	querySQL := queries.BuildExperimentVariantListQuery(req.Filters, req.Ordering, req.Limit)
	// Execute query and scan into return type.
	rows, err := c.QueryContext(ctx, querySQL)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from ExperimentVariant-> "+err.Error())
	}
	defer rows.Close()

	// Variables to store results returned by database.
	list := []*v1.ExperimentVariant{}
	var createdAt pq.NullTime
	var updatedAt pq.NullTime

	for rows.Next() {
		experimentvariant := new(v1.ExperimentVariant)
		if err := rows.Scan(&experimentvariant.ID, &createdAt, &updatedAt, &experimentvariant.Title, &experimentvariant.VariantType, &experimentvariant.Experiment, &experimentvariant.LandingPage, &experimentvariant.Field, &experimentvariant.Component); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from ExperimentVariant row-> "+err.Error())
		}
		// Convert pq.NullTime from database into proto timestamp.
		if createdAt.Valid {
			experimentvariant.CreatedAt, err = ptypes.TimestampProto(createdAt.Time)
			if err != nil {
				return nil, status.Error(codes.Unknown, "createdAt field has invalid format-> "+err.Error())
			}
		}
		if updatedAt.Valid {
			experimentvariant.UpdatedAt, err = ptypes.TimestampProto(updatedAt.Time)
			if err != nil {
				return nil, status.Error(codes.Unknown, "createdAt field has invalid format-> "+err.Error())
			}
		}

		list = append(list, experimentvariant)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from ExperimentVariant-> "+err.Error())
	}

	return &v1.ListExperimentVariantResponse{
		Api:   apiVersion,
		Items: list,
	}, nil
}

// Update ExperimentVariant
func (s *shrikeServiceServer) UpdateExperimentVariant(ctx context.Context, req *v1.UpdateExperimentVariantRequest) (*v1.UpdateExperimentVariantResponse, error) {
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

	// update experiment_variant
	res, err := c.ExecContext(ctx, "UPDATE experiment_variant SET title=$2, variant_type=$3, experiment=$4, landing_page=$5, field=$6, component=$7 WHERE id=$1",
		req.Item.ID, req.Item.Title, req.Item.VariantType, req.Item.Experiment, req.Item.LandingPage, req.Item.Field, req.Item.Component)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update ExperimentVariant-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("ExperimentVariant with ID='%s' is not found",
			req.Item.ID))
	}

	return &v1.UpdateExperimentVariantResponse{
		Api:     apiVersion,
		Updated: rows,
	}, nil
}

// Delete experiment_variant
func (s *shrikeServiceServer) DeleteExperimentVariant(ctx context.Context, req *v1.DeleteExperimentVariantRequest) (*v1.DeleteExperimentVariantResponse, error) {
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

	// delete experiment_variant
	res, err := c.ExecContext(ctx, "DELETE FROM experiment_variant WHERE id=$1", req.ID)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete ExperimentVariant-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("ExperimentVariant with ID='%s' is not found",
			req.ID))
	}

	return &v1.DeleteExperimentVariantResponse{
		Api:     apiVersion,
		Deleted: rows,
	}, nil
}
