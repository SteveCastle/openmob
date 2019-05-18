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

func TestNewAgentMembershipManager(t *testing.T) {
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
		want *AgentMembershipManager
	}{
		{"Returns a new agentMembership manager.", args{db: db}, &AgentMembershipManager{db: db}},
	}
	for _, tt := range tests {
		fmt.Println(tt.name)
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAgentMembershipManager(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAgentMembershipManager() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAgentMembershipManager_Create(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	type args struct {
		ctx  context.Context
		item *v1.CreateAgentMembership
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
			m := &AgentMembershipManager{
				db: tt.fields.db,
			}
			got, err := m.Create(tt.args.ctx, tt.args.item)
			if (err != nil) != tt.wantErr {
				t.Errorf("AgentMembershipManager.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("AgentMembershipManager.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAgentMembershipManager_Get(t *testing.T) {
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
		want    *AgentMembership
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &AgentMembershipManager{
				db: tt.fields.db,
			}
			got, err := m.Get(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("AgentMembershipManager.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AgentMembershipManager.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAgentMembershipManager_List(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	type args struct {
		ctx       context.Context
		filters   []*v1.AgentMembershipFilterRule
		orderings []*v1.AgentMembershipOrdering
		limit     int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*AgentMembership
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &AgentMembershipManager{
				db: tt.fields.db,
			}
			got, err := m.List(tt.args.ctx, tt.args.filters, tt.args.orderings, tt.args.limit)
			if (err != nil) != tt.wantErr {
				t.Errorf("AgentMembershipManager.List() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AgentMembershipManager.List() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAgentMembershipManager_Update(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	type args struct {
		ctx  context.Context
		item *v1.AgentMembership
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
			m := &AgentMembershipManager{
				db: tt.fields.db,
			}
			got, err := m.Update(tt.args.ctx, tt.args.item)
			if (err != nil) != tt.wantErr {
				t.Errorf("AgentMembershipManager.Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("AgentMembershipManager.Update() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAgentMembershipManager_Delete(t *testing.T) {
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
			m := &AgentMembershipManager{
				db: tt.fields.db,
			}
			got, err := m.Delete(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("AgentMembershipManager.Delete() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("AgentMembershipManager.Delete() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_convertToAgentMembershipProto(t *testing.T) {
	type args struct {
		c *AgentMembership
	}
	tests := []struct {
		name string
		args args
		want *v1.AgentMembership
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertToAgentMembershipProto(tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("convertToAgentMembershipProto() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAgentMembershipManager_GetProto(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	type args struct {
		c *AgentMembership
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *v1.AgentMembership
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &AgentMembershipManager{
				db: tt.fields.db,
			}
			if got := c.GetProto(tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AgentMembershipManager.GetProto() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBuildAgentMembershipListQuery(t *testing.T) {
	type args struct {
		filters   []*v1.AgentMembershipFilterRule
		orderings []*v1.AgentMembershipOrdering
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
			if got := BuildAgentMembershipListQuery(tt.args.filters, tt.args.orderings, tt.args.limit); got != tt.want {
				t.Errorf("BuildAgentMembershipListQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAgentMembershipManager_connect(t *testing.T) {
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
			m := &AgentMembershipManager{
				db: tt.fields.db,
			}
			got, err := m.connect(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("AgentMembershipManager.connect() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AgentMembershipManager.connect() = %v, want %v", got, tt.want)
			}
		})
	}
}
