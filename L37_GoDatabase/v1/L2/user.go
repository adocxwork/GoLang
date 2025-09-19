package main

import (
	"context"
	"database/sql"
	"time"
)

// User represents a row in the "users" table
type User struct {
	ID        int
	Name      string
	Email     string
	CreatedAt time.Time
}

// UserRepo wraps DB operations for users
type UserRepo struct {
	DB *sql.DB
}

// Create inserts a new user and returns it
func (r *UserRepo) Create(ctx context.Context, name, email string) (User, error) {
	var u User
	query := `INSERT INTO users (name, email) VALUES ($1, $2)
	          RETURNING id, name, email, created_at`
	err := r.DB.QueryRowContext(ctx, query, name, email).
		Scan(&u.ID, &u.Name, &u.Email, &u.CreatedAt)
	if err != nil {
		return User{}, err
	}
	return u, nil
}

// GetByEmail fetches a user by email. Returns empty User if not found.
func (r *UserRepo) GetByEmail(ctx context.Context, email string) (User, error) {
	var u User
	query := `SELECT id, name, email, created_at FROM users WHERE email=$1`
	err := r.DB.QueryRowContext(ctx, query, email).
		Scan(&u.ID, &u.Name, &u.Email, &u.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			// Not found
			return User{}, nil
		}
		return User{}, err
	}
	return u, nil
}

// BulkInsert demonstrates prepared statement usage
func (r *UserRepo) BulkInsert(ctx context.Context, users []User) error {
	stmt, err := r.DB.PrepareContext(ctx,
		`INSERT INTO users (name, email) VALUES ($1, $2)`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, u := range users {
		if _, err := stmt.ExecContext(ctx, u.Name, u.Email); err != nil {
			return err
		}
	}
	return nil
}
