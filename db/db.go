package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/h00s/todo-api/config"
	_ "github.com/lib/pq" //for a postgres
)

// Database handles DB connections
type Database struct {
	Conn *sql.DB
}

// Connect create new Database struct and connects to DB
func Connect(db config.Database) (*Database, error) {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s", db.Host, db.Port, db.User, db.Password, db.Name)
	conn, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	err = conn.Ping()
	if err != nil {
		log.Fatal(err)
	}

	return &Database{Conn: conn}, nil
}
