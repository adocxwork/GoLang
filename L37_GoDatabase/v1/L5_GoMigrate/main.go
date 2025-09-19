package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/golang-migrate/migrate/v4"
)

func main() {
	dsn := "postgres://appuser:pass@localhost:5432/appdb?sslmode=disable"

	// Run migrations
	m, err := migrate.New(
		"file://migrations",
		dsn,
	)
	if err != nil {
		log.Fatalf("migrate.New: %v", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("m.Up: %v", err)
	}

	fmt.Println("✅ Migrations applied successfully")

	// Now connect and verify
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("sql.Open: %v", err)
	}
	defer db.Close()

	// Verify connection
	if err := db.Ping(); err != nil {
		log.Fatalf("db.Ping: %v", err)
	}
	fmt.Println("✅ DB connection working")
}
