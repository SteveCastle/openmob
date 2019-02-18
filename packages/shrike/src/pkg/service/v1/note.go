package v1

import (
	"context"
	"fmt"

	v1 "github.com/SteveCastle/openmob/packages/shrike/src/pkg/api/v1"

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
	var id int64
	// insert Note entity data
	err = c.QueryRowContext(ctx, "INSERT INTO note (id, created_at, updated_at, contact, cause, body) VALUES($1, $2, $3, $4, $5, $6)  RETURNING id;",
		 req.Item.ID,  req.Item.CreatedAt,  req.Item.UpdatedAt,  req.Item.Contact,  req.Item.Cause,  req.Item.Body, ).Scan(&id)
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
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Note with ID='%d' is not found",
			req.ID))
	}

	// get Note data
	var note v1.Note
	if err := rows.Scan( &note.ID,  &note.CreatedAt,  &note.UpdatedAt,  &note.Contact,  &note.Cause,  &note.Body, ); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from Note row-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple Note rows with ID='%d'",
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

	// get Note list
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, contact, cause, body FROM note")
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from Note-> "+err.Error())
	}
	defer rows.Close()

	list := []*v1.Note{}
	for rows.Next() {
		note := new(v1.Note)
		if err := rows.Scan( &note.ID,  &note.CreatedAt,  &note.UpdatedAt,  &note.Contact,  &note.Cause,  &note.Body, ); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from Note row-> "+err.Error())
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
	res, err := c.ExecContext(ctx, "UPDATE note SET id=$1, created_at=$2, updated_at=$3, contact=$4, cause=$5, body=$6 WHERE id=$1",
		req.Item.ID,req.Item.CreatedAt,req.Item.UpdatedAt,req.Item.Contact,req.Item.Cause,req.Item.Body, )
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update Note-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Note with ID='%d' is not found",
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
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Note with ID='%d' is not found",
			req.ID))
	}

	return &v1.DeleteNoteResponse{
		Api:     apiVersion,
		Deleted: rows,
	}, nil
}
