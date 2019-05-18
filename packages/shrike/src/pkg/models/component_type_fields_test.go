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

func TestNewComponentTypeFieldsManager(t *testing.T) {
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
		want *ComponentTypeFieldsManager
	}{
		{"Returns a new componentTypeFields manager.", args{db: db}, &ComponentTypeFieldsManager{db: db}},
	}
	for _, tt := range tests {
		fmt.Println(tt.name)
		t.Run(tt.name, func(t *testing.T) {
			if got := NewComponentTypeFieldsManager(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewComponentTypeFieldsManager() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComponentTypeFieldsManager_Create(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	type args struct {
		ctx  context.Context
		item *v1.CreateComponentTypeFields
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
			m := &ComponentTypeFieldsManager{
				db: tt.fields.db,
			}
			got, err := m.Create(tt.args.ctx, tt.args.item)
			if (err != nil) != tt.wantErr {
				t.Errorf("ComponentTypeFieldsManager.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ComponentTypeFieldsManager.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComponentTypeFieldsManager_Get(t *testing.T) {
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
		want    *ComponentTypeFields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &ComponentTypeFieldsManager{
				db: tt.fields.db,
			}
			got, err := m.Get(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("ComponentTypeFieldsManager.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ComponentTypeFieldsManager.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComponentTypeFieldsManager_List(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	type args struct {
		ctx       context.Context
		filters   []*v1.ComponentTypeFieldsFilterRule
		orderings []*v1.ComponentTypeFieldsOrdering
		limit     int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*ComponentTypeFields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &ComponentTypeFieldsManager{
				db: tt.fields.db,
			}
			got, err := m.List(tt.args.ctx, tt.args.filters, tt.args.orderings, tt.args.limit)
			if (err != nil) != tt.wantErr {
				t.Errorf("ComponentTypeFieldsManager.List() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ComponentTypeFieldsManager.List() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComponentTypeFieldsManager_Update(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	type args struct {
		ctx  context.Context
		item *v1.ComponentTypeFields
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
			m := &ComponentTypeFieldsManager{
				db: tt.fields.db,
			}
			got, err := m.Update(tt.args.ctx, tt.args.item)
			if (err != nil) != tt.wantErr {
				t.Errorf("ComponentTypeFieldsManager.Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ComponentTypeFieldsManager.Update() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComponentTypeFieldsManager_Delete(t *testing.T) {
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
			m := &ComponentTypeFieldsManager{
				db: tt.fields.db,
			}
			got, err := m.Delete(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("ComponentTypeFieldsManager.Delete() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ComponentTypeFieldsManager.Delete() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_convertToComponentTypeFieldsProto(t *testing.T) {
	type args struct {
		c *ComponentTypeFields
	}
	tests := []struct {
		name string
		args args
		want *v1.ComponentTypeFields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertToComponentTypeFieldsProto(tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("convertToComponentTypeFieldsProto() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComponentTypeFieldsManager_GetProto(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	type args struct {
		c *ComponentTypeFields
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *v1.ComponentTypeFields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &ComponentTypeFieldsManager{
				db: tt.fields.db,
			}
			if got := c.GetProto(tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ComponentTypeFieldsManager.GetProto() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBuildComponentTypeFieldsListQuery(t *testing.T) {
	type args struct {
		filters   []*v1.ComponentTypeFieldsFilterRule
		orderings []*v1.ComponentTypeFieldsOrdering
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
			if got := BuildComponentTypeFieldsListQuery(tt.args.filters, tt.args.orderings, tt.args.limit); got != tt.want {
				t.Errorf("BuildComponentTypeFieldsListQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComponentTypeFieldsManager_connect(t *testing.T) {
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
			m := &ComponentTypeFieldsManager{
				db: tt.fields.db,
			}
			got, err := m.connect(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("ComponentTypeFieldsManager.connect() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ComponentTypeFieldsManager.connect() = %v, want %v", got, tt.want)
			}
		})
	}
}
