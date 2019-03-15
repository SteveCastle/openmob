package queries

import (
	"fmt"

	v1 "github.com/SteveCastle/openmob/packages/shrike/src/pkg/api/v1"
	"github.com/SteveCastle/structs"
)

// BuildDistrictFilters takes a filter and ordering object for a cause.
// and returns an SQL string
func BuildDistrictFilters(f []*v1.DistrictFilterRule, o []*v1.DistrictOrdering, limit int64) string {
	var sql string
	fmt.Println("Limit: ", limit)
	for _, r := range f {
		s := structs.New(r.GetField())
		for _, f := range s.Fields() {
			fmt.Printf("Filter by field name: %+v\n", f.Name())

			if f.IsExported() {
				fmt.Printf("Filter by value   : %v\n", f.Value())
				fmt.Printf("is zero : %+v\n", f.IsZero())
			}
		}
	}
	for _, r := range o {
		s := structs.New(r.GetField())
		for _, f := range s.Fields() {
			fmt.Printf("Order by field name: %+v\n", f.Name())
			if f.IsExported() {
				fmt.Printf("Order by value   : %v\n", f.Value())
				fmt.Printf("is zero : %+v\n", f.IsZero())
			}
		}

	}
	return sql
}
