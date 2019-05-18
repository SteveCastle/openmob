package models

import (
	"database/sql"
	"regexp"
	"strings"
	"time"

	"github.com/SteveCastle/openmob/packages/shrike/src/geography"

	uuid "github.com/gofrs/uuid"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/lib/pq"
)

// Comparison maps api comparison values to sql string,
var Comparison = map[string]string{
	"EQ":      "=",
	"NE":      "!=",
	"LT":      "<",
	"GT":      ">",
	"LTE":     "<=",
	"GTE":     ">=",
	"LIKE":    "LIKE",
	"NOTLIKE": "NOT LIKE",
}

// SQLDirections maps api comparison values to sql string,
var SQLDirections = map[string]string{
	"ASCENDING":  "ASC",
	"DESCENDING": "DESC",
}

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

// ToSnakeCase converts ProtoBuf type names to database column snake case.
func ToSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

func convertTimeToProto(t time.Time) (*timestamp.Timestamp, error) {
	time, err := ptypes.TimestampProto(t)
	if err != nil {
		return nil, err
	}
	return time, nil
}
func safeNullString(ns sql.NullString) *string {
	if !ns.Valid {
		return nil
	}
	return &(ns.String)
}

func safeNullInt64(ns sql.NullInt64) *int64 {
	if !ns.Valid {
		return nil
	}
	return &(ns.Int64)
}

func safeNullFloat64(ns sql.NullFloat64) *float64 {
	if !ns.Valid {
		return nil
	}
	return &ns.Float64
}

func safeNullBool(ns sql.NullBool) *bool {
	if !ns.Valid {
		return nil
	}
	return &(ns.Bool)
}

func safeNullUUID(u uuid.NullUUID) *string {
	if !u.Valid {
		return nil
	}
	val := u.UUID.String()
	return &val
}

func safeNullTime(nt pq.NullTime) *timestamp.Timestamp {
	if !nt.Valid {
		return nil
	}
	time, err := ptypes.TimestampProto(nt.Time)
	if err != nil {
		return nil
	}
	return time
}

func safeNullRegion(g geography.NullRegion) *int64 {
	if !g.Valid {
		return nil
	}
	v := int64(6)
	return &v
}
