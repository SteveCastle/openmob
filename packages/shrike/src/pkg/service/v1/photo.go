package v1

import (
	"context"
	"fmt"

	v1 "github.com/SteveCastle/openmob/packages/shrike/src/pkg/api/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Create new Photo
func (s *shrikeServiceServer) CreatePhoto(ctx context.Context, req *v1.CreatePhotoRequest) (*v1.CreatePhotoResponse, error) {
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
	// insert Photo entity data
	err = c.QueryRowContext(ctx, "INSERT INTO photo (id, created_at, updated_at, img_url) VALUES($1, $2, $3, $4)  RETURNING id;",
		 req.Item.ID,  req.Item.CreatedAt,  req.Item.UpdatedAt,  req.Item.ImgURL, ).Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into Photo-> "+err.Error())
	}

	// get ID of creates Photo
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created Photo-> "+err.Error())
	}

	return &v1.CreatePhotoResponse{
		Api: apiVersion,
		ID:  id,
	}, nil
}

// Get photo by id.
func (s *shrikeServiceServer) GetPhoto(ctx context.Context, req *v1.GetPhotoRequest) (*v1.GetPhotoResponse, error) {
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

	// query Photo by ID
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, img_url FROM photo WHERE id=$1",
		req.ID)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from Photo-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from Photo-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Photo with ID='%d' is not found",
			req.ID))
	}

	// get Photo data
	var photo v1.Photo
	if err := rows.Scan( &photo.ID,  &photo.CreatedAt,  &photo.UpdatedAt,  &photo.ImgURL, ); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from Photo row-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple Photo rows with ID='%d'",
			req.ID))
	}

	return &v1.GetPhotoResponse{
		Api:  apiVersion,
		Item: &photo,
	}, nil

}

// Read all Photo
func (s *shrikeServiceServer) ListPhoto(ctx context.Context, req *v1.ListPhotoRequest) (*v1.ListPhotoResponse, error) {
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

	// get Photo list
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, img_url FROM photo")
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from Photo-> "+err.Error())
	}
	defer rows.Close()

	list := []*v1.Photo{}
	for rows.Next() {
		photo := new(v1.Photo)
		if err := rows.Scan( &photo.ID,  &photo.CreatedAt,  &photo.UpdatedAt,  &photo.ImgURL, ); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from Photo row-> "+err.Error())
		}
		list = append(list, photo)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from Photo-> "+err.Error())
	}

	return &v1.ListPhotoResponse{
		Api:   apiVersion,
		Items: list,
	}, nil
}

// Update Photo
func (s *shrikeServiceServer) UpdatePhoto(ctx context.Context, req *v1.UpdatePhotoRequest) (*v1.UpdatePhotoResponse, error) {
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

	// update photo
	res, err := c.ExecContext(ctx, "UPDATE photo SET id=$1, created_at=$2, updated_at=$3, img_url=$4 WHERE id=$1",
		req.Item.ID,req.Item.CreatedAt,req.Item.UpdatedAt,req.Item.ImgURL, )
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update Photo-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Photo with ID='%d' is not found",
			req.Item.ID))
	}

	return &v1.UpdatePhotoResponse{
		Api:     apiVersion,
		Updated: rows,
	}, nil
}

// Delete photo
func (s *shrikeServiceServer) DeletePhoto(ctx context.Context, req *v1.DeletePhotoRequest) (*v1.DeletePhotoResponse, error) {
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

	// delete photo
	res, err := c.ExecContext(ctx, "DELETE FROM photo WHERE id=$1", req.ID)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete Photo-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Photo with ID='%d' is not found",
			req.ID))
	}

	return &v1.DeletePhotoResponse{
		Api:     apiVersion,
		Deleted: rows,
	}, nil
}