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

func TestNewVolunteerOpportunityManager(t *testing.T) {
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
		want *VolunteerOpportunityManager
	}{
		{"Returns a new volunteerOpportunity manager.", args{db: db}, &VolunteerOpportunityManager{db: db}},
	}
	for _, tt := range tests {
		fmt.Println(tt.name)
		t.Run(tt.name, func(t *testing.T) {
			if got := NewVolunteerOpportunityManager(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewVolunteerOpportunityManager() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVolunteerOpportunityManager_Create(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	type args struct {
		ctx  context.Context
		item *v1.CreateVolunteerOpportunity
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
			m := &VolunteerOpportunityManager{
				db: tt.fields.db,
			}
			got, err := m.Create(tt.args.ctx, tt.args.item)
			if (err != nil) != tt.wantErr {
				t.Errorf("VolunteerOpportunityManager.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("VolunteerOpportunityManager.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVolunteerOpportunityManager_Get(t *testing.T) {
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
		want    *VolunteerOpportunity
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &VolunteerOpportunityManager{
				db: tt.fields.db,
			}
			got, err := m.Get(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("VolunteerOpportunityManager.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("VolunteerOpportunityManager.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVolunteerOpportunityManager_List(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	type args struct {
		ctx       context.Context
		filters   []*v1.VolunteerOpportunityFilterRule
		orderings []*v1.VolunteerOpportunityOrdering
		limit     int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*VolunteerOpportunity
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &VolunteerOpportunityManager{
				db: tt.fields.db,
			}
			got, err := m.List(tt.args.ctx, tt.args.filters, tt.args.orderings, tt.args.limit)
			if (err != nil) != tt.wantErr {
				t.Errorf("VolunteerOpportunityManager.List() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("VolunteerOpportunityManager.List() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVolunteerOpportunityManager_Update(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	type args struct {
		ctx  context.Context
		item *v1.VolunteerOpportunity
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
			m := &VolunteerOpportunityManager{
				db: tt.fields.db,
			}
			got, err := m.Update(tt.args.ctx, tt.args.item)
			if (err != nil) != tt.wantErr {
				t.Errorf("VolunteerOpportunityManager.Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("VolunteerOpportunityManager.Update() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVolunteerOpportunityManager_Delete(t *testing.T) {
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
			m := &VolunteerOpportunityManager{
				db: tt.fields.db,
			}
			got, err := m.Delete(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("VolunteerOpportunityManager.Delete() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("VolunteerOpportunityManager.Delete() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_convertToVolunteerOpportunityProto(t *testing.T) {
	type args struct {
		c *VolunteerOpportunity
	}
	tests := []struct {
		name string
		args args
		want *v1.VolunteerOpportunity
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertToVolunteerOpportunityProto(tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("convertToVolunteerOpportunityProto() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVolunteerOpportunityManager_GetProto(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	type args struct {
		c *VolunteerOpportunity
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *v1.VolunteerOpportunity
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &VolunteerOpportunityManager{
				db: tt.fields.db,
			}
			if got := c.GetProto(tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("VolunteerOpportunityManager.GetProto() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBuildVolunteerOpportunityListQuery(t *testing.T) {
	type args struct {
		filters   []*v1.VolunteerOpportunityFilterRule
		orderings []*v1.VolunteerOpportunityOrdering
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
			if got := BuildVolunteerOpportunityListQuery(tt.args.filters, tt.args.orderings, tt.args.limit); got != tt.want {
				t.Errorf("BuildVolunteerOpportunityListQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVolunteerOpportunityManager_connect(t *testing.T) {
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
			m := &VolunteerOpportunityManager{
				db: tt.fields.db,
			}
			got, err := m.connect(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("VolunteerOpportunityManager.connect() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("VolunteerOpportunityManager.connect() = %v, want %v", got, tt.want)
			}
		})
	}
}
