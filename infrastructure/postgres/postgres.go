package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // database adapter
)

// DB is the Database handle
var DB *sql.DB

// InitDB initializes the database
func InitDB(name string, user string, password string, port string) error {
	var err error
	DB, err = sql.Open("postgres", fmt.Sprintf("user=%s dbname=%s password=%s port=%s sslmode=disable", user, name, password, port))
	return err
}
