package main

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/SteveCastle/openmob/models"
	_ "github.com/lib/pq"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
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
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	var p1 models.Person
	p1.ID = 1
	err = p1.Insert(context.Background(), db, boil.Infer())
	if err != nil {
		panic(err)
	}
	doesIt, _ := models.People(qm.Where("id=?", 1)).Exists(context.Background(), db)
	fmt.Println(doesIt)
	fmt.Println("vim-go")
}
