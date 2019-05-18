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

func TestNewLiveEventMembershipManager(t *testing.T) {
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
		want *LiveEventMembershipManager
	}{
		{"Returns a new liveEventMembership manager.", args{db: db}, &LiveEventMembershipManager{db: db}},
	}
	for _, tt := range tests {
		fmt.Println(tt.name)
		t.Run(tt.name, func(t *testing.T) {
			if got := NewLiveEventMembershipManager(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewLiveEventMembershipManager() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLiveEventMembershipManager_Create(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	type args struct {
		ctx  context.Context
		item *v1.CreateLiveEventMembership
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
			m := &LiveEventMembershipManager{
				db: tt.fields.db,
			}
			got, err := m.Create(tt.args.ctx, tt.args.item)
			if (err != nil) != tt.wantErr {
				t.Errorf("LiveEventMembershipManager.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("LiveEventMembershipManager.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLiveEventMembershipManager_Get(t *testing.T) {
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
		want    *LiveEventMembership
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &LiveEventMembershipManager{
				db: tt.fields.db,
			}
			got, err := m.Get(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("LiveEventMembershipManager.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LiveEventMembershipManager.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLiveEventMembershipManager_List(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	type args struct {
		ctx       context.Context
		filters   []*v1.LiveEventMembershipFilterRule
		orderings []*v1.LiveEventMembershipOrdering
		limit     int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*LiveEventMembership
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &LiveEventMembershipManager{
				db: tt.fields.db,
			}
			got, err := m.List(tt.args.ctx, tt.args.filters, tt.args.orderings, tt.args.limit)
			if (err != nil) != tt.wantErr {
				t.Errorf("LiveEventMembershipManager.List() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LiveEventMembershipManager.List() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLiveEventMembershipManager_Update(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	type args struct {
		ctx  context.Context
		item *v1.LiveEventMembership
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
			m := &LiveEventMembershipManager{
				db: tt.fields.db,
			}
			got, err := m.Update(tt.args.ctx, tt.args.item)
			if (err != nil) != tt.wantErr {
				t.Errorf("LiveEventMembershipManager.Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("LiveEventMembershipManager.Update() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLiveEventMembershipManager_Delete(t *testing.T) {
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
			m := &LiveEventMembershipManager{
				db: tt.fields.db,
			}
			got, err := m.Delete(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("LiveEventMembershipManager.Delete() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("LiveEventMembershipManager.Delete() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_convertToLiveEventMembershipProto(t *testing.T) {
	type args struct {
		c *LiveEventMembership
	}
	tests := []struct {
		name string
		args args
		want *v1.LiveEventMembership
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertToLiveEventMembershipProto(tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("convertToLiveEventMembershipProto() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLiveEventMembershipManager_GetProto(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	type args struct {
		c *LiveEventMembership
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *v1.LiveEventMembership
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &LiveEventMembershipManager{
				db: tt.fields.db,
			}
			if got := c.GetProto(tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LiveEventMembershipManager.GetProto() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBuildLiveEventMembershipListQuery(t *testing.T) {
	type args struct {
		filters   []*v1.LiveEventMembershipFilterRule
		orderings []*v1.LiveEventMembershipOrdering
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
			if got := BuildLiveEventMembershipListQuery(tt.args.filters, tt.args.orderings, tt.args.limit); got != tt.want {
				t.Errorf("BuildLiveEventMembershipListQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLiveEventMembershipManager_connect(t *testing.T) {
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
			m := &LiveEventMembershipManager{
				db: tt.fields.db,
			}
			got, err := m.connect(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("LiveEventMembershipManager.connect() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LiveEventMembershipManager.connect() = %v, want %v", got, tt.want)
			}
		})
	}
}
