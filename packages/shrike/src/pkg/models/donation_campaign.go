package models

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	v1 "github.com/SteveCastle/openmob/packages/shrike/src/pkg/api/v1"
	"github.com/SteveCastle/structs"
	uuid "github.com/gofrs/uuid"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// DonationCampaign is a type for donation_campaign db element.
type DonationCampaign struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	Title     string
}

// DonationCampaignManager manages queries returning a donationCampaign or list of donationCampaigns.
// It is configured with a db field to contain the db driver.
type DonationCampaignManager struct {
	db *sql.DB
}

// NewDonationCampaignManager creates a donationCampaign manager
func NewDonationCampaignManager(db *sql.DB) *DonationCampaignManager {
	return &DonationCampaignManager{db: db}
}

// CRUD Methods for the DonationCampaignManager.

// Create creates a donationCampaign.
func (m *DonationCampaignManager) Create(ctx context.Context, item *v1.CreateDonationCampaign) (*string, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()
	var id string
	// Execute INSERT query and then scan the resulting id into id string.
	err = c.QueryRowContext(ctx, "INSERT INTO donation_campaign (title) VALUES($1)  RETURNING id;",
		item.Title).Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into DonationCampaign-> "+err.Error())
	}

	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created DonationCampaign-> "+err.Error())
	}
	return &id, nil
}

// Get gets a single donationCampaign from the database by ID.
func (m *DonationCampaignManager) Get(ctx context.Context, id string) (*DonationCampaign, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	// query DonationCampaign by ID
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, title FROM donation_campaign WHERE id=$1",
		id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from DonationCampaign-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from DonationCampaign-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("DonationCampaign with ID='%s' is not found", id))
	}

	// scan DonationCampaign data into protobuf model
	var donationCampaign DonationCampaign

	if err := rows.Scan(&donationCampaign.ID, &donationCampaign.CreatedAt, &donationCampaign.UpdatedAt, &donationCampaign.Title); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from DonationCampaign row-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple DonationCampaign rows with ID='%s'",
			id))
	}
	return &donationCampaign, nil
}

// List returns a slice of all donationCampaigns meeting the filter criteria.
func (m *DonationCampaignManager) List(ctx context.Context, filters []*v1.DonationCampaignFilterRule, orderings []*v1.DonationCampaignOrdering, limit int64) ([]*DonationCampaign, error) {

	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	// Generate SQL to select all columns in DonationCampaign Table
	// TODO: Allow column selection.
	// Then generate filtering and ordering sql and finally run query.
	querySQL := BuildDonationCampaignListQuery(filters, orderings, limit)

	// Execute query and scan into return type.
	rows, err := c.QueryContext(ctx, querySQL)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from DonationCampaign-> "+err.Error())
	}
	defer rows.Close()

	// Scan the results into a slice.
	list := []*DonationCampaign{}
	for rows.Next() {
		donationCampaign := new(DonationCampaign)
		if err := rows.Scan(&donationCampaign.ID, &donationCampaign.CreatedAt, &donationCampaign.UpdatedAt, &donationCampaign.Title); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from DonationCampaign row-> "+err.Error())
		}
		list = append(list, donationCampaign)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from DonationCampaign-> "+err.Error())
	}
	return list, nil
}

// Update runs an update query on the provided db and returns the rows affected as an int64.
func (m *DonationCampaignManager) Update(ctx context.Context, item *v1.DonationCampaign) (*int64, error) {

	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	res, err := c.ExecContext(ctx, "UPDATE donation_campaign SET title=$2 WHERE id=$1",
		item.ID, item.Title)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update DonationCampaign-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("DonationCampaign with ID='%s' is not found",
			item.ID))
	}
	return &rows, nil
}

//Delete creates and executes DELETE sql on a provided id and returns the number of rows affected.
func (m *DonationCampaignManager) Delete(ctx context.Context, id string) (*int64, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	res, err := c.ExecContext(ctx, "DELETE FROM donationCampaign WHERE id=$1", id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete DonationCampaign-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("DonationCampaign with ID='%s' is not found",
			id))
	}

	return &rows, nil
}

// convertToDonationCampaignProto accepts a donationCampaign struct and returns a protobuf donationCampaign struct.
func convertToDonationCampaignProto(c *DonationCampaign) *v1.DonationCampaign {
	createdAt, _ := convertTimeToProto(c.CreatedAt)
	updatedAt, _ := convertTimeToProto(c.UpdatedAt)

	return &v1.DonationCampaign{
		ID:        c.ID.String(),
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
		Title:     c.Title,
	}
}

//GetProtoList returns a slice of protobuf typed struct of a donationCampaign.
func (*DonationCampaignManager) GetProtoList(l []*DonationCampaign) []*v1.DonationCampaign {
	list := []*v1.DonationCampaign{}
	for _, v := range l {
		list = append(list, convertToDonationCampaignProto(v))
	}
	return list
}

//GetProto returns a single protobuf typed struct of a donationCampaign.
func (*DonationCampaignManager) GetProto(c *DonationCampaign) *v1.DonationCampaign {
	return convertToDonationCampaignProto(c)
}

// BuildDonationCampaignListQuery takes a filter and ordering object for a donationCampaign.
// and returns an SQL string
func BuildDonationCampaignListQuery(filters []*v1.DonationCampaignFilterRule, orderings []*v1.DonationCampaignOrdering, limit int64) string {
	baseSQL := "SELECT id, created_at, updated_at, title FROM donation_campaign"
	// Range over the provided rules and create where clauses.
	for i, r := range filters {
		if i == 0 {
			baseSQL = fmt.Sprintf("%s %s", baseSQL, "WHERE")
		} else {
			baseSQL = fmt.Sprintf("%s %s", baseSQL, r.LogicalOperator)
		}
		s := structs.New(r.GetField())
		for _, f := range s.Fields() {
			if f.IsExported() {
				baseSQL = fmt.Sprintf("%s %s %s '%s'", baseSQL, ToSnakeCase(f.Name()), Comparison[r.Rule.String()], f.Value())
			}
		}
	}
	// Range over ordering rules and create ORDER BY clauses.
	for _, r := range orderings {
		fmt.Println(r.Direction)
		s := structs.New(r.GetField())
		for _, f := range s.Fields() {
			baseSQL = fmt.Sprintf("%s %s", baseSQL, "ORDER BY")
			if f.IsExported() {
				baseSQL = fmt.Sprintf("%s %s %s", baseSQL, ToSnakeCase(f.Name()), SQLDirections[r.Direction.String()])
			}
		}

	}
	baseSQL = fmt.Sprintf("%s LIMIT %d;", baseSQL, limit)
	fmt.Printf("List SQL Executed: %v\n", baseSQL)
	return baseSQL
}

// connect returns SQL database connection from the pool.
func (m *DonationCampaignManager) connect(ctx context.Context) (*sql.Conn, error) {
	c, err := m.db.Conn(ctx)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to connect to database-> "+err.Error())
	}
	return c, nil
}
