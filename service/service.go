/*
Copyright © 2024 Amr Ojjeh <amrojjeh@outlook.com>
*/

package service

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	"github.com/amrojjeh/arareader/model"
	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
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

func FromPlainPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	return string(hash), err
}

func ApplyVowelQuestionsToExcerpt(qs []model.Question, e *Excerpt) error {
	for _, q := range qs {
		var data VowelQuestionData
		if err := json.Unmarshal(q.Data, &data); err != nil {
			return err
		}
		log.Println("unpointed one reference")
		if err := e.UnpointRef(data.Reference); err != nil {
			return err
		}
	}
	return nil
}
