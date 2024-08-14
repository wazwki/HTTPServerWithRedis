package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

var DB *sql.DB

const connStr string = "user=admin dbname=gotest sslmode=disable password=1234"

func Conn() error {
	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		return err
	}
	_, err = DB.Exec(`CREATE TABLE IF NOT EXISTS users(
		id SERIAL PRIMARY KEY,
		firstname TEXT,
		lastname TEXT,
	)`)
	if err != nil {
		return err
	}
	return nil
}
