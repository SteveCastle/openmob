package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	v1 "github.com/SteveCastle/openmob/packages/shrike/src/pkg/api/v1"
)

func TestNewDistrictManager(t *testing.T) {
	type args struct {
		db *sql.DB
	}
	db, _, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	tests := []struct {
		name string
		args args
		want *DistrictManager
	}{
		{"Returns a new district manager.", args{db: db}, &DistrictManager{db: db}},
	}
	for _, tt := range tests {
		fmt.Println(tt.name)
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDistrictManager(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDistrictManager() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDistrictManager_Create(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	type args struct {
		ctx  context.Context
		item *v1.CreateDistrict
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &DistrictManager{
				db: tt.fields.db,
			}
			got, err := m.Create(tt.args.ctx, tt.args.item)
			if (err != nil) != tt.wantErr {
				t.Errorf("DistrictManager.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("DistrictManager.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDistrictManager_Get(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	type args struct {
		ctx context.Context
		id  string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *District
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &DistrictManager{
				db: tt.fields.db,
			}
			got, err := m.Get(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("DistrictManager.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DistrictManager.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDistrictManager_List(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	type args struct {
		ctx       context.Context
		filters   []*v1.DistrictFilterRule
		orderings []*v1.DistrictOrdering
		limit     int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*District
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &DistrictManager{
				db: tt.fields.db,
			}
			got, err := m.List(tt.args.ctx, tt.args.filters, tt.args.orderings, tt.args.limit)
			if (err != nil) != tt.wantErr {
				t.Errorf("DistrictManager.List() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DistrictManager.List() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDistrictManager_Update(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	type args struct {
		ctx  context.Context
		item *v1.District
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *int64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &DistrictManager{
				db: tt.fields.db,
			}
			got, err := m.Update(tt.args.ctx, tt.args.item)
			if (err != nil) != tt.wantErr {
				t.Errorf("DistrictManager.Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("DistrictManager.Update() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDistrictManager_Delete(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	type args struct {
		ctx context.Context
		id  string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *int64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &DistrictManager{
				db: tt.fields.db,
			}
			got, err := m.Delete(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("DistrictManager.Delete() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("DistrictManager.Delete() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_convertToDistrictProto(t *testing.T) {
	type args struct {
		c *District
	}
	tests := []struct {
		name string
		args args
		want *v1.District
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertToDistrictProto(tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("convertToDistrictProto() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDistrictManager_GetProto(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	type args struct {
		c *District
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *v1.District
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &DistrictManager{
				db: tt.fields.db,
			}
			if got := c.GetProto(tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DistrictManager.GetProto() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBuildDistrictListQuery(t *testing.T) {
	type args struct {
		filters   []*v1.DistrictFilterRule
		orderings []*v1.DistrictOrdering
		limit     int64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BuildDistrictListQuery(tt.args.filters, tt.args.orderings, tt.args.limit); got != tt.want {
				t.Errorf("BuildDistrictListQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDistrictManager_connect(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *sql.Conn
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &DistrictManager{
				db: tt.fields.db,
			}
			got, err := m.connect(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("DistrictManager.connect() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DistrictManager.connect() = %v, want %v", got, tt.want)
			}
		})
	}
}
