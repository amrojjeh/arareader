/*
Copyright © 2024 Amr Ojjeh <amrojjeh@outlook.com>
*/

package service

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/amrojjeh/arareader/model"
	_ "github.com/mattn/go-sqlite3"
)

// OpenDB opens an sqlite3 database
func OpenDB(dsn string) *sql.DB {
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

// Setup initializes the database schema
func Setup(ctx context.Context, db *sql.DB) {
	_, err := db.ExecContext(ctx, model.Schema)
	if err != nil {
		panic("could not execute schema")
	}
}
