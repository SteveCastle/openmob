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

func TestNewCustomerCartManager(t *testing.T) {
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
		want *CustomerCartManager
	}{
		{"Returns a new customerCart manager.", args{db: db}, &CustomerCartManager{db: db}},
	}
	for _, tt := range tests {
		fmt.Println(tt.name)
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCustomerCartManager(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCustomerCartManager() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCustomerCartManager_Create(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	type args struct {
		ctx  context.Context
		item *v1.CreateCustomerCart
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
			m := &CustomerCartManager{
				db: tt.fields.db,
			}
			got, err := m.Create(tt.args.ctx, tt.args.item)
			if (err != nil) != tt.wantErr {
				t.Errorf("CustomerCartManager.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CustomerCartManager.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCustomerCartManager_Get(t *testing.T) {
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
		want    *CustomerCart
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &CustomerCartManager{
				db: tt.fields.db,
			}
			got, err := m.Get(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("CustomerCartManager.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CustomerCartManager.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCustomerCartManager_List(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	type args struct {
		ctx       context.Context
		filters   []*v1.CustomerCartFilterRule
		orderings []*v1.CustomerCartOrdering
		limit     int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*CustomerCart
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &CustomerCartManager{
				db: tt.fields.db,
			}
			got, err := m.List(tt.args.ctx, tt.args.filters, tt.args.orderings, tt.args.limit)
			if (err != nil) != tt.wantErr {
				t.Errorf("CustomerCartManager.List() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CustomerCartManager.List() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCustomerCartManager_Update(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	type args struct {
		ctx  context.Context
		item *v1.CustomerCart
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
			m := &CustomerCartManager{
				db: tt.fields.db,
			}
			got, err := m.Update(tt.args.ctx, tt.args.item)
			if (err != nil) != tt.wantErr {
				t.Errorf("CustomerCartManager.Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CustomerCartManager.Update() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCustomerCartManager_Delete(t *testing.T) {
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
			m := &CustomerCartManager{
				db: tt.fields.db,
			}
			got, err := m.Delete(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("CustomerCartManager.Delete() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CustomerCartManager.Delete() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_convertToCustomerCartProto(t *testing.T) {
	type args struct {
		c *CustomerCart
	}
	tests := []struct {
		name string
		args args
		want *v1.CustomerCart
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertToCustomerCartProto(tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("convertToCustomerCartProto() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCustomerCartManager_GetProto(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	type args struct {
		c *CustomerCart
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *v1.CustomerCart
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &CustomerCartManager{
				db: tt.fields.db,
			}
			if got := c.GetProto(tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CustomerCartManager.GetProto() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBuildCustomerCartListQuery(t *testing.T) {
	type args struct {
		filters   []*v1.CustomerCartFilterRule
		orderings []*v1.CustomerCartOrdering
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
			if got := BuildCustomerCartListQuery(tt.args.filters, tt.args.orderings, tt.args.limit); got != tt.want {
				t.Errorf("BuildCustomerCartListQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCustomerCartManager_connect(t *testing.T) {
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
			m := &CustomerCartManager{
				db: tt.fields.db,
			}
			got, err := m.connect(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("CustomerCartManager.connect() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CustomerCartManager.connect() = %v, want %v", got, tt.want)
			}
		})
	}
}
