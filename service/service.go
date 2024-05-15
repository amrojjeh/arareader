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
	_, err := db.ExecContext(ctx, model.Schema)
	if err != nil {
		panic("could not execute schema")
	}
}

func FromPlainPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	return string(hash), err
}

func Map[T any, V any](f func(T) V, arr []T) []V {
	vs := []V{}
	for _, el := range arr {
		vs = append(vs, f(el))
	}
	return vs
}

func Filter[T any](pred func(T) bool, arr []T) []T {
	vs := []T{}
	for _, el := range arr {
		if pred(el) {
			vs = append(vs, el)
		}
	}
	return vs
}
