package queries

import (
	"fmt"

	v1 "github.com/SteveCastle/openmob/packages/shrike/src/pkg/api/v1"
	"github.com/SteveCastle/structs"
)

// BuildMailingAddressFilters takes a filter and ordering object for a cause.
// and returns an SQL string
func BuildMailingAddressFilters(filters []*v1.MailingAddressFilterRule, orderings []*v1.MailingAddressOrdering, limit int64) string {
	var sql string
	for i, r := range filters {
		// Insert where clause before the first filter.
		// And the Logical operator of each successive filter.
		if i == 0 {
			sql = fmt.Sprintf("%s %s", sql, "WHERE")
		} else {
			sql = fmt.Sprintf("%s %s", sql, "AND")
		}
		s := structs.New(r.GetField())
		for _, f := range s.Fields() {
			if f.IsExported() {
				sql = fmt.Sprintf("%s %s %s '%s'", sql, ToSnakeCase(f.Name()), Comparison["EQ"], f.Value())
			}
		}
	}
	for _, r := range orderings {
		s := structs.New(r.GetField())
		for _, f := range s.Fields() {
			sql = fmt.Sprintf("%s %s", sql, "ORDER BY")
			if f.IsExported() {
				sql = fmt.Sprintf("%s %s ASC", sql, ToSnakeCase(f.Name()))
			}
		}

	}
	sql = fmt.Sprintf("%s LIMIT %d;", sql, limit)
	fmt.Printf("List SQL Executed: %v\n", sql)
	return sql
}
