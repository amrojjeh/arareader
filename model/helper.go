/*
Copyright © 2024 Amr Ojjeh <amrojjeh@outlook.com>
*/

package model

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
)

// MustOpenDB opens an sqlite3 database
func MustOpenDB(dsn string) *sql.DB {
	db, err := sql.Open("sqlite3", dsn)
	if err != nil {
		panic(fmt.Sprintf("could not open db (dsn: %s)", dsn))
	}

	err = db.Ping()
	if err != nil {
		panic(fmt.Sprintf("could not pingb db (dsn: %s)", dsn))
	}
	return db
}

// MustSetup initializes the database schema
func MustSetup(ctx context.Context, db *sql.DB) {
	_, err := db.ExecContext(ctx, Schema)
	if err != nil {
		panic("could not execute schema")
	}
}

func FromPlainPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	return string(hash), err
}
