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

// causeServiceServer is implementation of v1.CauseServiceServer proto interface
type causeServiceServer struct {
	db *sql.DB
}

// NewCauseServiceServer creates Cause service
func NewCauseServiceServer(db *sql.DB) v1.CauseServiceServer {
	return &causeServiceServer{db: db}
}

// connect returns SQL database connection from the pool
func (s *causeServiceServer) connect(ctx context.Context) (*sql.Conn, error) {
	c, err := s.db.Conn(ctx)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to connect to database-> "+err.Error())
	}
	return c, nil
}

// Create new todo task
func (s *causeServiceServer) CreateCause(ctx context.Context, req *v1.CreateCauseRequest) (*v1.CreateCauseResponse, error) {

	// get SQL connection from pool
	c, err := s.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	// insert Cause entity data
	res, err := c.ExecContext(ctx, "INSERT INTO Cause(`Title`) VALUES(?)",
		req.Item.Title)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into Cause-> "+err.Error())
	}

	// get ID of creates Cause
	id, err := res.LastInsertId()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created Cause-> "+err.Error())
	}

	return &v1.CreateCauseResponse{
		Id: id,
	}, nil
}

// Read todo task
func (s *causeServiceServer) GetCause(ctx context.Context, req *v1.GetCauseRequest) (*v1.GetCauseResponse, error) {

	// get SQL connection from pool
	c, err := s.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	// query Cause by ID
	rows, err := c.QueryContext(ctx, "SELECT `id`, `title` FROM cause WHERE `id`=?",
		req.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from Cause-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from Cause-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Cause with ID='%d' is not found",
			req.Id))
	}

	// get Cause data
	var td v1.Cause
	if err := rows.Scan(&td.Id, &td.Title); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from Cause row-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple Cause rows with ID='%d'",
			req.Id))
	}

	return &v1.GetCauseResponse{
		Item: &td,
	}, nil

}
