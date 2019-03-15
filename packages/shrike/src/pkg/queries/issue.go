package queries

import (
	"fmt"

	v1 "github.com/SteveCastle/openmob/packages/shrike/src/pkg/api/v1"
	"github.com/SteveCastle/structs"
)

// BuildIssueFilters takes a filter and ordering object for a cause.
// and returns an SQL string
func BuildIssueFilters(filters []*v1.IssueFilterRule, orderings []*v1.IssueOrdering, limit int64) string {
	var sql string
	fmt.Println("Limit: ", limit)
	for _, r := range filters {
		s := structs.New(r.GetField())
		for _, f := range s.Fields() {
			sql = fmt.Sprintf("%s %s", sql, "WHERE")
			if f.IsExported() {
				sql = fmt.Sprintf("%s %s %s '%s'", sql, ToSnakeCase(f.Name()), Comparison["EQ"], f.Value())
				fmt.Printf("Filter generated: %v\n", sql)
			}
		}
	}
	for _, r := range orderings {
		s := structs.New(r.GetField())
		for _, f := range s.Fields() {
			fmt.Printf("Order by field name: %+v\n", ToSnakeCase(f.Name()))
			sql = fmt.Sprintf("%s %s", sql, "ORDER BY")
			if f.IsExported() {
				sql = fmt.Sprintf("%s %s ASC", sql, ToSnakeCase(f.Name()))
				fmt.Printf("Ordering generated: %v\n", sql)
			}
		}

	}
	sql = fmt.Sprintf("%s LIMIT %d;", sql, limit)
	fmt.Printf("Final SQL: %v\n", sql)
	return sql
}
