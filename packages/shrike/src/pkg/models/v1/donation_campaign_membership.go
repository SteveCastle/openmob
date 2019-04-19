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

// DonationCampaignMembership is a type for donation_campaign_membership db element.
type DonationCampaignMembership struct {
	ID               uuid.UUID
	CreatedAt        time.Time
	UpdatedAt        time.Time
	Cause            uuid.UUID
	DonationCampaign uuid.UUID
}

// DonationCampaignMembershipManager manages queries returning a donationCampaignMembership or list of donationCampaignMemberships.
// It is configured with a db field to contain the db driver.
type DonationCampaignMembershipManager struct {
	db *sql.DB
}

// NewDonationCampaignMembershipManager creates a donationCampaignMembership manager
func NewDonationCampaignMembershipManager(db *sql.DB) *DonationCampaignMembershipManager {
	return &DonationCampaignMembershipManager{db: db}
}

// CRUD Methods for the DonationCampaignMembershipManager.

// Create creates a donationCampaignMembership.
func (m *DonationCampaignMembershipManager) Create(ctx context.Context, item *v1.CreateDonationCampaignMembership) (*string, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()
	var id string
	// Execute INSERT query and then scan the resulting id into id string.
	err = c.QueryRowContext(ctx, "INSERT INTO donation_campaign_membership (cause, donation_campaign) VALUES($1, $2)  RETURNING id;",
		item.Cause, item.DonationCampaign).Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into DonationCampaignMembership-> "+err.Error())
	}

	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created DonationCampaignMembership-> "+err.Error())
	}
	return &id, nil
}

// Get gets a single donationCampaignMembership from the database by ID.
func (m *DonationCampaignMembershipManager) Get(ctx context.Context, id string) (*DonationCampaignMembership, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	// query DonationCampaignMembership by ID
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, cause, donation_campaign FROM donation_campaign_membership WHERE id=$1",
		id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from DonationCampaignMembership-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from DonationCampaignMembership-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("DonationCampaignMembership with ID='%s' is not found", id))
	}

	// scan DonationCampaignMembership data into protobuf model
	var donationCampaignMembership DonationCampaignMembership

	if err := rows.Scan(&donationCampaignMembership.ID, &donationCampaignMembership.CreatedAt, &donationCampaignMembership.UpdatedAt, &donationCampaignMembership.Cause, &donationCampaignMembership.DonationCampaign); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from DonationCampaignMembership row-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple DonationCampaignMembership rows with ID='%s'",
			id))
	}
	return &donationCampaignMembership, nil
}

// List returns a slice of all donationCampaignMemberships meeting the filter criteria.
func (m *DonationCampaignMembershipManager) List(ctx context.Context, filters []*v1.DonationCampaignMembershipFilterRule, orderings []*v1.DonationCampaignMembershipOrdering, limit int64) ([]*DonationCampaignMembership, error) {

	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	// Generate SQL to select all columns in DonationCampaignMembership Table
	// TODO: Allow column selection.
	// Then generate filtering and ordering sql and finally run query.
	querySQL := BuildDonationCampaignMembershipListQuery(filters, orderings, limit)

	// Execute query and scan into return type.
	rows, err := c.QueryContext(ctx, querySQL)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from DonationCampaignMembership-> "+err.Error())
	}
	defer rows.Close()

	// Scan the results into a slice.
	list := []*DonationCampaignMembership{}
	for rows.Next() {
		donationCampaignMembership := new(DonationCampaignMembership)
		if err := rows.Scan(&donationCampaignMembership.ID, &donationCampaignMembership.CreatedAt, &donationCampaignMembership.UpdatedAt, &donationCampaignMembership.Cause, &donationCampaignMembership.DonationCampaign); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from DonationCampaignMembership row-> "+err.Error())
		}
		list = append(list, donationCampaignMembership)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from DonationCampaignMembership-> "+err.Error())
	}
	return list, nil
}

// Update runs an update query on the provided db and returns the rows affected as an int64.
func (m *DonationCampaignMembershipManager) Update(ctx context.Context, item *v1.DonationCampaignMembership) (*int64, error) {

	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	res, err := c.ExecContext(ctx, "UPDATE donation_campaign_membership SET cause=$2, donation_campaign=$3 WHERE id=$1",
		item.ID, item.Cause, item.DonationCampaign)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update DonationCampaignMembership-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("DonationCampaignMembership with ID='%s' is not found",
			item.ID))
	}
	return &rows, nil
}

//Delete creates and executes DELETE sql on a provided id and returns the number of rows affected.
func (m *DonationCampaignMembershipManager) Delete(ctx context.Context, id string) (*int64, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	res, err := c.ExecContext(ctx, "DELETE FROM donationCampaignMembership WHERE id=$1", id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete DonationCampaignMembership-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("DonationCampaignMembership with ID='%s' is not found",
			id))
	}

	return &rows, nil
}

// convertToDonationCampaignMembershipProto accepts a donationCampaignMembership struct and returns a protobuf donationCampaignMembership struct.
func convertToDonationCampaignMembershipProto(c *DonationCampaignMembership) *v1.DonationCampaignMembership {
	createdAt, _ := convertTimeToProto(c.CreatedAt)
	updatedAt, _ := convertTimeToProto(c.UpdatedAt)

	return &v1.DonationCampaignMembership{
		ID:               c.ID.String(),
		CreatedAt:        createdAt,
		UpdatedAt:        updatedAt,
		Cause:            c.Cause.String(),
		DonationCampaign: c.DonationCampaign.String(),
	}
}

//GetProtoList returns a slice of protobuf typed struct of a donationCampaignMembership.
func (*DonationCampaignMembershipManager) GetProtoList(l []*DonationCampaignMembership) []*v1.DonationCampaignMembership {
	list := []*v1.DonationCampaignMembership{}
	for _, v := range l {
		list = append(list, convertToDonationCampaignMembershipProto(v))
	}
	return list
}

//GetProto returns a single protobuf typed struct of a donationCampaignMembership.
func (*DonationCampaignMembershipManager) GetProto(c *DonationCampaignMembership) *v1.DonationCampaignMembership {
	return convertToDonationCampaignMembershipProto(c)
}

// BuildDonationCampaignMembershipListQuery takes a filter and ordering object for a donationCampaignMembership.
// and returns an SQL string
func BuildDonationCampaignMembershipListQuery(filters []*v1.DonationCampaignMembershipFilterRule, orderings []*v1.DonationCampaignMembershipOrdering, limit int64) string {
	baseSQL := "SELECT id, created_at, updated_at, cause, donation_campaign FROM donation_campaign_membership"
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
func (m *DonationCampaignMembershipManager) connect(ctx context.Context) (*sql.Conn, error) {
	c, err := m.db.Conn(ctx)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to connect to database-> "+err.Error())
	}
	return c, nil
}
