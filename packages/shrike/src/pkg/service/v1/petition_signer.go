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

// Create new PetitionSigner
func (s *shrikeServiceServer) CreatePetitionSigner(ctx context.Context, req *v1.CreatePetitionSignerRequest) (*v1.CreatePetitionSignerResponse, error) {
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
	// insert PetitionSigner entity data
	err = c.QueryRowContext(ctx, "INSERT INTO petition_signer (petition, contact, cause) VALUES($1, $2, $3)  RETURNING id;",
		req.Item.Petition, req.Item.Contact, req.Item.Cause).Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into PetitionSigner-> "+err.Error())
	}

	// get ID of creates PetitionSigner
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created PetitionSigner-> "+err.Error())
	}

	return &v1.CreatePetitionSignerResponse{
		Api: apiVersion,
		ID:  id,
	}, nil
}

// Get petition_signer by id.
func (s *shrikeServiceServer) GetPetitionSigner(ctx context.Context, req *v1.GetPetitionSignerRequest) (*v1.GetPetitionSignerResponse, error) {
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

	// query PetitionSigner by ID
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, petition, contact, cause FROM petition_signer WHERE id=$1",
		req.ID)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from PetitionSigner-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from PetitionSigner-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("PetitionSigner with ID='%s' is not found",
			req.ID))
	}

	// scan PetitionSigner data into protobuf model
	var petitionsigner v1.PetitionSigner
	var createdAt pq.NullTime
	var updatedAt pq.NullTime

	if err := rows.Scan(&petitionsigner.ID, &createdAt, &updatedAt, &petitionsigner.Petition, &petitionsigner.Contact, &petitionsigner.Cause); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from PetitionSigner row-> "+err.Error())
	}

	// Convert pq.NullTime from database into proto timestamp.
	if createdAt.Valid {
		petitionsigner.CreatedAt, err = ptypes.TimestampProto(createdAt.Time)
		if err != nil {
			return nil, status.Error(codes.Unknown, "createdAt field has invalid format-> "+err.Error())
		}
	}
	if updatedAt.Valid {
		petitionsigner.UpdatedAt, err = ptypes.TimestampProto(updatedAt.Time)
		if err != nil {
			return nil, status.Error(codes.Unknown, "createdAt field has invalid format-> "+err.Error())
		}
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple PetitionSigner rows with ID='%s'",
			req.ID))
	}

	return &v1.GetPetitionSignerResponse{
		Api:  apiVersion,
		Item: &petitionsigner,
	}, nil

}

// Read all PetitionSigner
func (s *shrikeServiceServer) ListPetitionSigner(ctx context.Context, req *v1.ListPetitionSignerRequest) (*v1.ListPetitionSignerResponse, error) {
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

	// Generate SQL to select all columns in PetitionSigner Table
	// Then generate filtering and ordering sql and finally run query.
	querySQL := queries.BuildPetitionSignerListQuery(req.Filters, req.Ordering, req.Limit)
	// Execute query and scan into return type.
	rows, err := c.QueryContext(ctx, querySQL)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from PetitionSigner-> "+err.Error())
	}
	defer rows.Close()

	// Variables to store results returned by database.
	list := []*v1.PetitionSigner{}
	var createdAt pq.NullTime
	var updatedAt pq.NullTime

	for rows.Next() {
		petitionsigner := new(v1.PetitionSigner)
		if err := rows.Scan(&petitionsigner.ID, &createdAt, &updatedAt, &petitionsigner.Petition, &petitionsigner.Contact, &petitionsigner.Cause); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from PetitionSigner row-> "+err.Error())
		}
		// Convert pq.NullTime from database into proto timestamp.
		if createdAt.Valid {
			petitionsigner.CreatedAt, err = ptypes.TimestampProto(createdAt.Time)
			if err != nil {
				return nil, status.Error(codes.Unknown, "createdAt field has invalid format-> "+err.Error())
			}
		}
		if updatedAt.Valid {
			petitionsigner.UpdatedAt, err = ptypes.TimestampProto(updatedAt.Time)
			if err != nil {
				return nil, status.Error(codes.Unknown, "createdAt field has invalid format-> "+err.Error())
			}
		}

		list = append(list, petitionsigner)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from PetitionSigner-> "+err.Error())
	}

	return &v1.ListPetitionSignerResponse{
		Api:   apiVersion,
		Items: list,
	}, nil
}

// Update PetitionSigner
func (s *shrikeServiceServer) UpdatePetitionSigner(ctx context.Context, req *v1.UpdatePetitionSignerRequest) (*v1.UpdatePetitionSignerResponse, error) {
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

	// update petition_signer
	res, err := c.ExecContext(ctx, "UPDATE petition_signer SET petition=$2, contact=$3, cause=$4 WHERE id=$1",
		req.Item.ID, req.Item.Petition, req.Item.Contact, req.Item.Cause)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update PetitionSigner-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("PetitionSigner with ID='%s' is not found",
			req.Item.ID))
	}

	return &v1.UpdatePetitionSignerResponse{
		Api:     apiVersion,
		Updated: rows,
	}, nil
}

// Delete petition_signer
func (s *shrikeServiceServer) DeletePetitionSigner(ctx context.Context, req *v1.DeletePetitionSignerRequest) (*v1.DeletePetitionSignerResponse, error) {
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

	// delete petition_signer
	res, err := c.ExecContext(ctx, "DELETE FROM petition_signer WHERE id=$1", req.ID)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete PetitionSigner-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("PetitionSigner with ID='%s' is not found",
			req.ID))
	}

	return &v1.DeletePetitionSignerResponse{
		Api:     apiVersion,
		Deleted: rows,
	}, nil
}
