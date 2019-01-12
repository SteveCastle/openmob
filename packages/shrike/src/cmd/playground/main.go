package main

import (
	"database/sql"
	"fmt"

	_ "github.com/cridenour/go-postgis"
	postgis "github.com/cridenour/go-postgis"
	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "postgres://tern:tern@localhost:5432/tern?sslmode=disable")
	if err != nil {
		fmt.Println(err)
	}
	point := postgis.PointS{4326, -84.5014, 39.1064}
	var newPoint postgis.PointS

	// Ensure we have PostGIS on the table
	db.Exec("CREATE EXTENSION IF NOT EXISTS postgis")

	// Demonstrate both driver.Valuer and sql.Scanner support
	db.QueryRow("SELECT GeomFromEWKB($1);", point).Scan(&newPoint)

	if point == newPoint {
		fmt.Println("Point returned equal from PostGIS!")
	}
}
