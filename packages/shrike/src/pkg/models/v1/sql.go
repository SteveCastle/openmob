package models

import (
	"regexp"
	"strings"
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

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

// ToSnakeCase converts ProtoBuf type names to database column snake case.
func ToSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}
