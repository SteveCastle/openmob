package main

import (
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "tern"
	password = "tern"
	dbname   = "tern"
	sslmode  = "disable"
)

func main() {
	fmt.Println("vim-go")
}
