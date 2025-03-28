package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() *sql.DB {
	connStr := "host=localhost port=5432 user=admin password=admin dbname=shortener sslmode=disable"

	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Database connected!")

	createTables()

	return DB
}

func createTables() {
	query := `
	CREATE TABLE IF NOT EXISTS urls (
		id SERIAL PRIMARY KEY,
		original_url TEXT NOT NULL,
		short_url TEXT NOT NULL
	);
	`

	if _, err := DB.Exec(query); err != nil {
		log.Fatalf("Error on creating tables: %v", err)
	}

	fmt.Println("🚀 Tables created!")
}
