package v1

import (
	"context"

	v1 "github.com/SteveCastle/openmob/packages/shrike/src/pkg/api/v1"
	"github.com/SteveCastle/openmob/packages/shrike/src/pkg/models/v1"
)

// Create new Note
func (s *shrikeServiceServer) CreateNote(ctx context.Context, req *v1.CreateNoteRequest) (*v1.CreateNoteResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a Note Manager
	m := models.NewNoteManager(s.db)

	// Get a list of notes given filters, ordering, and limit rules.
	id, err := m.CreateNote(ctx, req.Item)
	if err != nil {
		return nil, err
	}
	return &v1.CreateNoteResponse{
		Api: apiVersion,
		ID:  *id,
	}, nil
}

// Get note by id.
func (s *shrikeServiceServer) GetNote(ctx context.Context, req *v1.GetNoteRequest) (*v1.GetNoteResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a Note Manager
	m := models.NewNoteManager(s.db)

	// Get a list of notes given filters, ordering, and limit rules.
	note, err := m.GetNote(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	return &v1.GetNoteResponse{
		Api:  apiVersion,
		Item: m.GetProto(note),
	}, nil

}

// Read all Note
func (s *shrikeServiceServer) ListNote(ctx context.Context, req *v1.ListNoteRequest) (*v1.ListNoteResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}

	// Create a Note Manager
	m := models.NewNoteManager(s.db)

	// Get a list of notes given filters, ordering, and limit rules.
	list, err := m.ListNote(ctx, req.Filters, req.Ordering, req.Limit)
	if err != nil {
		return nil, err
	}

	return &v1.ListNoteResponse{
		Api:   apiVersion,
		Items: m.GetProtoList(list),
	}, nil
}

// Update Note
func (s *shrikeServiceServer) UpdateNote(ctx context.Context, req *v1.UpdateNoteRequest) (*v1.UpdateNoteResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a Note Manager
	m := models.NewNoteManager(s.db)

	// Get a list of notes given filters, ordering, and limit rules.
	rows, err := m.UpdateNote(ctx, req.Item)
	if err != nil {
		return nil, err
	}

	return &v1.UpdateNoteResponse{
		Api:     apiVersion,
		Updated: *rows,
	}, nil
}

// Delete note
func (s *shrikeServiceServer) DeleteNote(ctx context.Context, req *v1.DeleteNoteRequest) (*v1.DeleteNoteResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a Note Manager
	m := models.NewNoteManager(s.db)

	// Get a list of notes given filters, ordering, and limit rules.
	rows, err := m.DeleteNote(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	return &v1.DeleteNoteResponse{
		Api:     apiVersion,
		Deleted: *rows,
	}, nil
}
