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

// Create new Note
func (s *shrikeServiceServer) CreateNote(ctx context.Context, req *v1.CreateNoteRequest) (*v1.CreateNoteResponse, error) {
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
	// insert Note entity data
	err = c.QueryRowContext(ctx, "INSERT INTO note (contact, cause, body) VALUES($1, $2, $3)  RETURNING id;",
		req.Item.Contact, req.Item.Cause, req.Item.Body).Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into Note-> "+err.Error())
	}

	// get ID of creates Note
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created Note-> "+err.Error())
	}

	return &v1.CreateNoteResponse{
		Api: apiVersion,
		ID:  id,
	}, nil
}

// Get note by id.
func (s *shrikeServiceServer) GetNote(ctx context.Context, req *v1.GetNoteRequest) (*v1.GetNoteResponse, error) {
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

	// query Note by ID
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, contact, cause, body FROM note WHERE id=$1",
		req.ID)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from Note-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from Note-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Note with ID='%s' is not found",
			req.ID))
	}

	// scan Note data into protobuf model
	var note v1.Note
	var createdAt pq.NullTime
	var updatedAt pq.NullTime

	if err := rows.Scan(&note.ID, &createdAt, &updatedAt, &note.Contact, &note.Cause, &note.Body); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from Note row-> "+err.Error())
	}

	// Convert pq.NullTime from database into proto timestamp.
	if createdAt.Valid {
		note.CreatedAt, err = ptypes.TimestampProto(createdAt.Time)
		if err != nil {
			return nil, status.Error(codes.Unknown, "createdAt field has invalid format-> "+err.Error())
		}
	}
	if updatedAt.Valid {
		note.UpdatedAt, err = ptypes.TimestampProto(updatedAt.Time)
		if err != nil {
			return nil, status.Error(codes.Unknown, "createdAt field has invalid format-> "+err.Error())
		}
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple Note rows with ID='%s'",
			req.ID))
	}

	return &v1.GetNoteResponse{
		Api:  apiVersion,
		Item: &note,
	}, nil

}

// Read all Note
func (s *shrikeServiceServer) ListNote(ctx context.Context, req *v1.ListNoteRequest) (*v1.ListNoteResponse, error) {
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

	// Generate SQL to select all columns in Note Table
	// Then generate filtering and ordering sql and finally run query.

	baseSQL := "SELECT id, created_at, updated_at, contact, cause, body FROM note"
	querySQL := queries.BuildNoteFilters(req.Filters, req.Ordering, req.Limit)
	SQL := fmt.Sprintf("%s %s", baseSQL, querySQL)
	rows, err := c.QueryContext(ctx, SQL)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from Note-> "+err.Error())
	}
	defer rows.Close()

	// Variables to store results returned by database.
	list := []*v1.Note{}
	var createdAt pq.NullTime
	var updatedAt pq.NullTime

	for rows.Next() {
		note := new(v1.Note)
		if err := rows.Scan(&note.ID, &createdAt, &updatedAt, &note.Contact, &note.Cause, &note.Body); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from Note row-> "+err.Error())
		}
		// Convert pq.NullTime from database into proto timestamp.
		if createdAt.Valid {
			note.CreatedAt, err = ptypes.TimestampProto(createdAt.Time)
			if err != nil {
				return nil, status.Error(codes.Unknown, "createdAt field has invalid format-> "+err.Error())
			}
		}
		if updatedAt.Valid {
			note.UpdatedAt, err = ptypes.TimestampProto(updatedAt.Time)
			if err != nil {
				return nil, status.Error(codes.Unknown, "createdAt field has invalid format-> "+err.Error())
			}
		}

		list = append(list, note)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from Note-> "+err.Error())
	}

	return &v1.ListNoteResponse{
		Api:   apiVersion,
		Items: list,
	}, nil
}

// Update Note
func (s *shrikeServiceServer) UpdateNote(ctx context.Context, req *v1.UpdateNoteRequest) (*v1.UpdateNoteResponse, error) {
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

	// update note
	res, err := c.ExecContext(ctx, "UPDATE note SET contact=$2, cause=$3, body=$4 WHERE id=$1",
		req.Item.ID, req.Item.Contact, req.Item.Cause, req.Item.Body)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update Note-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Note with ID='%s' is not found",
			req.Item.ID))
	}

	return &v1.UpdateNoteResponse{
		Api:     apiVersion,
		Updated: rows,
	}, nil
}

// Delete note
func (s *shrikeServiceServer) DeleteNote(ctx context.Context, req *v1.DeleteNoteRequest) (*v1.DeleteNoteResponse, error) {
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

	// delete note
	res, err := c.ExecContext(ctx, "DELETE FROM note WHERE id=$1", req.ID)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete Note-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Note with ID='%s' is not found",
			req.ID))
	}

	return &v1.DeleteNoteResponse{
		Api:     apiVersion,
		Deleted: rows,
	}, nil
}
