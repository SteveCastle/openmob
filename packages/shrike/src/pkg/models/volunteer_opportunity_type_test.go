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

func TestNewVolunteerOpportunityTypeManager(t *testing.T) {
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
		want *VolunteerOpportunityTypeManager
	}{
		{"Returns a new volunteerOpportunityType manager.", args{db: db}, &VolunteerOpportunityTypeManager{db: db}},
	}
	for _, tt := range tests {
		fmt.Println(tt.name)
		t.Run(tt.name, func(t *testing.T) {
			if got := NewVolunteerOpportunityTypeManager(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewVolunteerOpportunityTypeManager() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVolunteerOpportunityTypeManager_Create(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	type args struct {
		ctx  context.Context
		item *v1.CreateVolunteerOpportunityType
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
			m := &VolunteerOpportunityTypeManager{
				db: tt.fields.db,
			}
			got, err := m.Create(tt.args.ctx, tt.args.item)
			if (err != nil) != tt.wantErr {
				t.Errorf("VolunteerOpportunityTypeManager.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("VolunteerOpportunityTypeManager.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVolunteerOpportunityTypeManager_Get(t *testing.T) {
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
		want    *VolunteerOpportunityType
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &VolunteerOpportunityTypeManager{
				db: tt.fields.db,
			}
			got, err := m.Get(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("VolunteerOpportunityTypeManager.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("VolunteerOpportunityTypeManager.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVolunteerOpportunityTypeManager_List(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	type args struct {
		ctx       context.Context
		filters   []*v1.VolunteerOpportunityTypeFilterRule
		orderings []*v1.VolunteerOpportunityTypeOrdering
		limit     int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*VolunteerOpportunityType
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &VolunteerOpportunityTypeManager{
				db: tt.fields.db,
			}
			got, err := m.List(tt.args.ctx, tt.args.filters, tt.args.orderings, tt.args.limit)
			if (err != nil) != tt.wantErr {
				t.Errorf("VolunteerOpportunityTypeManager.List() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("VolunteerOpportunityTypeManager.List() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVolunteerOpportunityTypeManager_Update(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	type args struct {
		ctx  context.Context
		item *v1.VolunteerOpportunityType
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
			m := &VolunteerOpportunityTypeManager{
				db: tt.fields.db,
			}
			got, err := m.Update(tt.args.ctx, tt.args.item)
			if (err != nil) != tt.wantErr {
				t.Errorf("VolunteerOpportunityTypeManager.Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("VolunteerOpportunityTypeManager.Update() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVolunteerOpportunityTypeManager_Delete(t *testing.T) {
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
			m := &VolunteerOpportunityTypeManager{
				db: tt.fields.db,
			}
			got, err := m.Delete(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("VolunteerOpportunityTypeManager.Delete() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("VolunteerOpportunityTypeManager.Delete() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_convertToVolunteerOpportunityTypeProto(t *testing.T) {
	type args struct {
		c *VolunteerOpportunityType
	}
	tests := []struct {
		name string
		args args
		want *v1.VolunteerOpportunityType
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertToVolunteerOpportunityTypeProto(tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("convertToVolunteerOpportunityTypeProto() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVolunteerOpportunityTypeManager_GetProto(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	type args struct {
		c *VolunteerOpportunityType
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *v1.VolunteerOpportunityType
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &VolunteerOpportunityTypeManager{
				db: tt.fields.db,
			}
			if got := c.GetProto(tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("VolunteerOpportunityTypeManager.GetProto() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBuildVolunteerOpportunityTypeListQuery(t *testing.T) {
	type args struct {
		filters   []*v1.VolunteerOpportunityTypeFilterRule
		orderings []*v1.VolunteerOpportunityTypeOrdering
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
			if got := BuildVolunteerOpportunityTypeListQuery(tt.args.filters, tt.args.orderings, tt.args.limit); got != tt.want {
				t.Errorf("BuildVolunteerOpportunityTypeListQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVolunteerOpportunityTypeManager_connect(t *testing.T) {
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
			m := &VolunteerOpportunityTypeManager{
				db: tt.fields.db,
			}
			got, err := m.connect(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("VolunteerOpportunityTypeManager.connect() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("VolunteerOpportunityTypeManager.connect() = %v, want %v", got, tt.want)
			}
		})
	}
}
