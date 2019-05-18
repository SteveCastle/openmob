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

func TestNewElectionMembershipManager(t *testing.T) {
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
		want *ElectionMembershipManager
	}{
		{"Returns a new electionMembership manager.", args{db: db}, &ElectionMembershipManager{db: db}},
	}
	for _, tt := range tests {
		fmt.Println(tt.name)
		t.Run(tt.name, func(t *testing.T) {
			if got := NewElectionMembershipManager(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewElectionMembershipManager() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestElectionMembershipManager_Create(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	type args struct {
		ctx  context.Context
		item *v1.CreateElectionMembership
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
			m := &ElectionMembershipManager{
				db: tt.fields.db,
			}
			got, err := m.Create(tt.args.ctx, tt.args.item)
			if (err != nil) != tt.wantErr {
				t.Errorf("ElectionMembershipManager.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ElectionMembershipManager.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestElectionMembershipManager_Get(t *testing.T) {
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
		want    *ElectionMembership
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &ElectionMembershipManager{
				db: tt.fields.db,
			}
			got, err := m.Get(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("ElectionMembershipManager.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ElectionMembershipManager.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestElectionMembershipManager_List(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	type args struct {
		ctx       context.Context
		filters   []*v1.ElectionMembershipFilterRule
		orderings []*v1.ElectionMembershipOrdering
		limit     int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*ElectionMembership
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &ElectionMembershipManager{
				db: tt.fields.db,
			}
			got, err := m.List(tt.args.ctx, tt.args.filters, tt.args.orderings, tt.args.limit)
			if (err != nil) != tt.wantErr {
				t.Errorf("ElectionMembershipManager.List() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ElectionMembershipManager.List() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestElectionMembershipManager_Update(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	type args struct {
		ctx  context.Context
		item *v1.ElectionMembership
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
			m := &ElectionMembershipManager{
				db: tt.fields.db,
			}
			got, err := m.Update(tt.args.ctx, tt.args.item)
			if (err != nil) != tt.wantErr {
				t.Errorf("ElectionMembershipManager.Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ElectionMembershipManager.Update() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestElectionMembershipManager_Delete(t *testing.T) {
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
			m := &ElectionMembershipManager{
				db: tt.fields.db,
			}
			got, err := m.Delete(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("ElectionMembershipManager.Delete() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ElectionMembershipManager.Delete() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_convertToElectionMembershipProto(t *testing.T) {
	type args struct {
		c *ElectionMembership
	}
	tests := []struct {
		name string
		args args
		want *v1.ElectionMembership
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertToElectionMembershipProto(tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("convertToElectionMembershipProto() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestElectionMembershipManager_GetProto(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	type args struct {
		c *ElectionMembership
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *v1.ElectionMembership
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &ElectionMembershipManager{
				db: tt.fields.db,
			}
			if got := c.GetProto(tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ElectionMembershipManager.GetProto() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBuildElectionMembershipListQuery(t *testing.T) {
	type args struct {
		filters   []*v1.ElectionMembershipFilterRule
		orderings []*v1.ElectionMembershipOrdering
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
			if got := BuildElectionMembershipListQuery(tt.args.filters, tt.args.orderings, tt.args.limit); got != tt.want {
				t.Errorf("BuildElectionMembershipListQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestElectionMembershipManager_connect(t *testing.T) {
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
			m := &ElectionMembershipManager{
				db: tt.fields.db,
			}
			got, err := m.connect(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("ElectionMembershipManager.connect() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ElectionMembershipManager.connect() = %v, want %v", got, tt.want)
			}
		})
	}
}
