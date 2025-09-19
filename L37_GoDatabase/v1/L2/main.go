package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib" // Postgres driver for database/sql
)

func main() {
	// DSN: user=appuser, password=pass, db=appdb, host=localhost, port=5432
	dsn := "postgresql://appuser:pass@localhost:5432/appdb?sslmode=disable"

	// Open DB connection pool
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		log.Fatalf("sql.Open: %v", err)
	}
	defer db.Close()

	// Pool settings
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(30 * time.Minute)

	// Verify DB connection
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := db.PingContext(ctx); err != nil {
		log.Fatalf("db.PingContext: %v", err)
	}
	fmt.Println("✅ Connected to Postgres")

	// Create repository
	repo := &UserRepo{DB: db}

	// 1. Create a new user
	alice, err := repo.Create(ctx, "Alice2", "alice2@example.com")
	if err != nil {
		log.Fatalf("repo.Create: %v", err)
	}
	fmt.Printf("Created user: %+v\n", alice)

	// 2. Fetch by email (exists)
	found, err := repo.GetByEmail(ctx, "alice2@example.com")
	if err != nil {
		log.Fatalf("repo.GetByEmail: %v", err)
	}
	if found.ID == 0 {
		fmt.Println("No user found with that email")
	} else {
		fmt.Printf("Fetched user: %+v\n", found)
	}

	// 3. Fetch by email (does not exist → sql.ErrNoRows case)
	missing, err := repo.GetByEmail(ctx, "notfound@example.com")
	if err != nil {
		log.Fatalf("repo.GetByEmail missing: %v", err)
	}
	if missing.ID == 0 {
		fmt.Println("Correctly handled missing user (sql.ErrNoRows)")
	}

	// 4. Bulk insert multiple users (prepared statement)
	users := []User{
		{Name: "Bob", Email: "bob2@example.com"},
		{Name: "Carol", Email: "carol2@example.com"},
	}
	if err := repo.BulkInsert(ctx, users); err != nil {
		log.Fatalf("repo.BulkInsert: %v", err)
	}
	fmt.Println("✅ Bulk insert done")
}
