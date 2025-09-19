package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib" // pgx driver for database/sql
)

func main() {
	// Connection string (DSN)
	dsn := "postgresql://appuser:pass@localhost:5432/appdb?sslmode=disable"

	// Open connection pool
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

	// Verify connection
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := db.PingContext(ctx); err != nil {
		log.Fatalf("db.PingContext: %v", err)
	}
	fmt.Println("✅ Connected to Postgres")

	// Create users table if not exists
	createTable := `
	CREATE TABLE IF NOT EXISTS users (
	  id SERIAL PRIMARY KEY,
	  name TEXT NOT NULL,
	  email TEXT UNIQUE NOT NULL,
	  created_at TIMESTAMPTZ NOT NULL DEFAULT now()
	);`
	if _, err := db.ExecContext(ctx, createTable); err != nil {
		log.Fatalf("create table: %v", err)
	}
	fmt.Println("✅ Ensured users table exists")

	// Insert Alice (single insert example)
	var newID int
	var createdAt time.Time
	insertQ := `INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id, created_at`
	if err := db.QueryRowContext(ctx, insertQ, "Alice", "alice@example.com").Scan(&newID, &createdAt); err != nil {
		log.Fatalf("insert user: %v", err)
	}
	fmt.Printf("Inserted user id=%d created_at=%s\n", newID, createdAt.UTC().Format(time.RFC3339))

	// Transaction example: insert Bob and Carol together
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		log.Fatalf("BeginTx: %v", err)
	}
	defer tx.Rollback() // rollback if commit is not called

	if _, err := tx.ExecContext(ctx, `INSERT INTO users (name, email) VALUES ($1, $2)`, "Bob", "bob@example.com"); err != nil {
		log.Fatalf("insert bob: %v", err)
	}
	if _, err := tx.ExecContext(ctx, `INSERT INTO users (name, email) VALUES ($1, $2)`, "Carol", "carol@example.com"); err != nil {
		log.Fatalf("insert carol: %v", err)
	}

	if err := tx.Commit(); err != nil {
		log.Fatalf("commit: %v", err)
	}
	fmt.Println("✅ Transaction committed (Bob & Carol added)")

	// Query and print all users
	rows, err := db.QueryContext(ctx, `SELECT id, name, email, created_at FROM users`)
	if err != nil {
		log.Fatalf("query users: %v", err)
	}
	defer rows.Close()

	fmt.Println("All users:")
	for rows.Next() {
		var id int
		var name, email string
		var ts time.Time
		if err := rows.Scan(&id, &name, &email, &ts); err != nil {
			log.Fatalf("scan: %v", err)
		}
		fmt.Printf(" - id=%d name=%s email=%s created_at=%s\n", id, name, email, ts.UTC().Format(time.RFC3339))
	}
	if err := rows.Err(); err != nil {
		log.Fatalf("rows iteration: %v", err)
	}
}
