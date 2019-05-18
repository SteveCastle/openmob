package v1

import (
	"context"
	"database/sql"
	"reflect"
	"testing"

	v1 "github.com/SteveCastle/openmob/packages/shrike/src/pkg/api/v1"
)

func Test_shrikeServiceServer_CreateCause(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	type args struct {
		ctx context.Context
		req *v1.CreateCauseRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *v1.CreateCauseResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &shrikeServiceServer{
				db: tt.fields.db,
			}
			got, err := s.CreateCause(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("shrikeServiceServer.CreateCause() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("shrikeServiceServer.CreateCause() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_shrikeServiceServer_GetCause(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	type args struct {
		ctx context.Context
		req *v1.GetCauseRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *v1.GetCauseResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &shrikeServiceServer{
				db: tt.fields.db,
			}
			got, err := s.GetCause(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("shrikeServiceServer.GetCause() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("shrikeServiceServer.GetCause() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_shrikeServiceServer_ListCause(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	type args struct {
		ctx context.Context
		req *v1.ListCauseRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *v1.ListCauseResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &shrikeServiceServer{
				db: tt.fields.db,
			}
			got, err := s.ListCause(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("shrikeServiceServer.ListCause() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("shrikeServiceServer.ListCause() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_shrikeServiceServer_UpdateCause(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	type args struct {
		ctx context.Context
		req *v1.UpdateCauseRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *v1.UpdateCauseResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &shrikeServiceServer{
				db: tt.fields.db,
			}
			got, err := s.UpdateCause(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("shrikeServiceServer.UpdateCause() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("shrikeServiceServer.UpdateCause() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_shrikeServiceServer_DeleteCause(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	type args struct {
		ctx context.Context
		req *v1.DeleteCauseRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *v1.DeleteCauseResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &shrikeServiceServer{
				db: tt.fields.db,
			}
			got, err := s.DeleteCause(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("shrikeServiceServer.DeleteCause() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("shrikeServiceServer.DeleteCause() = %v, want %v", got, tt.want)
			}
		})
	}
}
