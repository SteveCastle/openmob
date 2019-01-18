package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

// CauseHandler handles http requests for a cause.
type CauseHandler struct {
	db *sql.DB
}

func (h *CauseHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var title string
	// Execute the query.
	row := h.db.QueryRow("SELECT title FROM cause")
	if err := row.Scan(&title); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	// Write it back to the client.
	fmt.Fprintf(w, "The cause is %s", title)
}

func main() {
	// Open our database connection.
	db, err := sql.Open("postgres", "postgres://tern:tern@localhost:5432/tern?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	// Register our handler.
	http.Handle("/hello", &CauseHandler{db: db})
	http.ListenAndServe(":8080", nil)
}
