package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	// 1. Connect
	dsn := "postgresql://appuser:pass@localhost:5432/appdb?sslmode=disable"
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		log.Fatalf("sql.Open: %v", err)
	}
	defer db.Close()

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(time.Hour)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		log.Fatalf("db.PingContext: %v", err)
	}
	fmt.Println("✅ Connected to Postgres")

	// 2. Insert users
	alice, err := InsertUser(ctx, db, "Alice", "alice@example.com")
	if err != nil {
		log.Fatalf("InsertUser: %v", err)
	}
	fmt.Printf("Inserted user: %+v\n", alice)

	// 3. Get existing user
	found, err := GetUserByEmail(ctx, db, "alice@example.com")
	if err != nil {
		log.Fatalf("GetUserByEmail: %v", err)
	}
	if found != nil {
		fmt.Printf("Found user: %+v\n", found)
	} else {
		fmt.Println("No user found with that email")
	}

	// 4. Try to get non-existent user
	notFound, err := GetUserByEmail(ctx, db, "nobody@example.com")
	if err != nil {
		log.Fatalf("GetUserByEmail: %v", err)
	}
	if notFound == nil {
		fmt.Println("Correctly handled: user not found")
	}

	// 5. List all users
	users, err := ListUsers(ctx, db)
	if err != nil {
		log.Fatalf("ListUsers: %v", err)
	}
	fmt.Println("All users:")
	for _, u := range users {
		fmt.Printf("  %+v\n", u)
	}
}
