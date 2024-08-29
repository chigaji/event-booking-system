package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

/*
*
Initializes postgres DB connection
*/
func InitDB() {
	// connStr := "postgres://postgres:ronald@postgres_event_db:8432/booking_db?sslmode=disable"
	dbURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", "localhost", "5432", "ronald", "booking_db", "ronald")
	var err error
	// dsn := "postgres://ronald:ronon$#!@localhost:5432/booking_db?sslmode=disable"

	DB, err := sql.Open("postgres", dbURL)

	if err != nil {
		log.Fatal("Failed to connect to the database:", err)

	}

	if err = DB.Ping(); err != nil {
		log.Fatal("Failed to ping DB: ", err)
	}
	fmt.Println("connect to the DB successfully")
}
