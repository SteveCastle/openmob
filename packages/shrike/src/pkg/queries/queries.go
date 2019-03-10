package queries

import (
	"fmt"

	v1 "github.com/SteveCastle/openmob/packages/shrike/src/pkg/api/v1"
	"github.com/SteveCastle/structs"
)

// Filter defines a single filter struct.
type Filter struct {
}

// Order defines a single order struct.
type Order struct {
}

// BuildCauseFilters takes a filter and ordering object for a cause.
// and returns an SQL string
func BuildCauseFilters(f []*v1.CauseFilterRule, o []*v1.CauseOrdering, limit int64) string {
	var sql string
	fmt.Println("Limit: ", limit)
	for _, r := range f {
		s := structs.New(r.GetField())
		for _, f := range s.Fields() {
			fmt.Printf("field name: %+v\n", f.Name())

			if f.IsExported() {
				fmt.Printf("value   : %+v\n", f.Value())
				fmt.Printf("is zero : %+v\n", f.IsZero())
			}
		}
	}
	for _, r := range o {
		s := structs.New(r.GetField())
		for _, f := range s.Fields() {
			fmt.Printf("field name: %+v\n", f.Name())

			if f.IsExported() {
				fmt.Printf("value   : %+v\n", f.Value())
				fmt.Printf("is zero : %+v\n", f.IsZero())
			}
		}

	}
	return sql
}
