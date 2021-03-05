package dbservice

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq" //postgres driver
)

var globalDb *sql.DB

// Connect : connect database
func Connect() (*sql.DB, error) {

	if globalDb != nil {
		return globalDb, nil
	}

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_DATABASE")

	connectString := fmt.Sprintf(
		`host=%s port=%s user=%s dbname=%s password=%s sslmode=disable`,
		host,
		port,
		user,
		dbName,
		password,
	)

	globalDb, err := sql.Open("postgres", connectString)
	return globalDb, err
}

// Query : query
func Query(query string) (rows *sql.Rows, err error) {
	db, err := Connect()
	if err != nil {
		return
	}
	rows, err = db.Query(query)
	return
}

// QueryRow : query single row
func QueryRow(query string) (*sql.Row, error) {
	db, err := Connect()
	if err != nil {
		return nil, err
	}
	row := db.QueryRow(query)
	return row, nil
}

// Exec : executes a query without returning any row
func Exec(query string, args ...interface{}) (sql.Result, error) {
	db, err := Connect()
	if err != nil {
		return nil, err
	}
	result, err := db.Exec(query, args...)
	return result, err
}
