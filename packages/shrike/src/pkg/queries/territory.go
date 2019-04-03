package queries

import (
	"fmt"

	v1 "github.com/SteveCastle/openmob/packages/shrike/src/pkg/api/v1"
	"github.com/SteveCastle/structs"
)

// BuildTerritoryListQuery takes a filter and ordering object for a cause.
// and returns an SQL string
func BuildTerritoryListQuery(filters []*v1.TerritoryFilterRule, orderings []*v1.TerritoryOrdering, limit int64) string {
	// SQL to get all Territorys and all columns.
	baseSQL := "SELECT id, created_at, updated_at, title FROM territory"
	// Generate WHERE clause from filters passed in request.
	for i, r := range filters {
		// Insert where clause before the first filter.
		// And the Logical operator of each successive filter.
		if i == 0 {
			baseSQL = fmt.Sprintf("%s %s", baseSQL, "WHERE")
		} else {
			baseSQL = fmt.Sprintf("%s %s", baseSQL, "AND")
		}
		s := structs.New(r.GetField())
		for _, f := range s.Fields() {
			if f.IsExported() {
				baseSQL = fmt.Sprintf("%s %s %s '%s'", baseSQL, ToSnakeCase(f.Name()), Comparison["EQ"], f.Value())
			}
		}
	}
	// Generate ORDER BY clause from ordering passed in request.
	for _, r := range orderings {
		s := structs.New(r.GetField())
		for _, f := range s.Fields() {
			baseSQL = fmt.Sprintf("%s %s", baseSQL, "ORDER BY")
			if f.IsExported() {
				baseSQL = fmt.Sprintf("%s %s ASC", baseSQL, ToSnakeCase(f.Name()))
			}
		}

	}
	baseSQL = fmt.Sprintf("%s LIMIT %d;", baseSQL, limit)
	fmt.Printf("List SQL Executed: %v\n", baseSQL)
	return baseSQL
}
