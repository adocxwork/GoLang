package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"
)

// User struct maps to the users table
type User struct {
	ID        int
	Name      string
	Email     string
	CreatedAt time.Time
}

// InsertUser inserts a new user and returns it
func InsertUser(ctx context.Context, db *sql.DB, name, email string) (*User, error) {
	query := `INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id, created_at`
	var u User
	u.Name = name
	u.Email = email

	err := db.QueryRowContext(ctx, query, name, email).
		Scan(&u.ID, &u.CreatedAt)
	if err != nil {
		return nil, fmt.Errorf("InsertUser: %w", err)
	}
	return &u, nil
}

// GetUserByEmail fetches a user by email
func GetUserByEmail(ctx context.Context, db *sql.DB, email string) (*User, error) {
	query := `SELECT id, name, email, created_at FROM users WHERE email=$1`
	var u User
	err := db.QueryRowContext(ctx, query, email).
		Scan(&u.ID, &u.Name, &u.Email, &u.CreatedAt)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil // not found
	}
	if err != nil {
		return nil, fmt.Errorf("GetUserByEmail: %w", err)
	}
	return &u, nil
}

// ListUsers returns all users
func ListUsers(ctx context.Context, db *sql.DB) ([]User, error) {
	query := `SELECT id, name, email, created_at FROM users ORDER BY id`
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("ListUsers: %w", err)
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var u User
		if err := rows.Scan(&u.ID, &u.Name, &u.Email, &u.CreatedAt); err != nil {
			return nil, fmt.Errorf("rows.Scan: %w", err)
		}
		users = append(users, u)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows.Err: %w", err)
	}
	return users, nil
}
