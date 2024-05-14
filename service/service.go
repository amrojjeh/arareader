/*
Copyright © 2024 Amr Ojjeh <amrojjeh@outlook.com>
*/

package service

import (
	"context"
	"database/sql"
	"encoding/json"
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

func ApplyVowelDataToExcerpt(ds []VowelQuestionData, e *Excerpt) error {
	for _, d := range ds {
		if err := e.UnpointRef(d.Reference); err != nil {
			return err
		}
	}
	return nil
}

func MustVowelQuestionData(q model.Question) VowelQuestionData {
	var data VowelQuestionData
	err := json.Unmarshal(q.Data, &data)
	if err != nil {
		panic(fmt.Sprintf("question data is not a valid VowelQuestionData (id: %d). %s", q.ID, err.Error()))
	}
	return data
}

func MustShortAnswerQuestionData(q model.Question) ShortAnswerQuestionData {
	var data ShortAnswerQuestionData
	err := json.Unmarshal(q.Data, &data)
	if err != nil {
		panic(fmt.Sprintf("question data is not a valid ShortAnswerQuestionData (id: %d). %s", q.ID, err.Error()))
	}
	return data
}

func ExtractQuestionData(q model.Question) (QuestionData, error) {
	switch q.Type {
	case string(VowelQuestionType):
		return MustVowelQuestionData(q), nil
	case string(ShortAnswerQuestionType):
		return MustShortAnswerQuestionData(q), nil
	default:
		return nil, fmt.Errorf("question type could not be identified (type: %s)", q.Type)
	}
}

func Map[T any, V any](f func(T) V, arr []T) []V {
	vs := []V{}
	for _, el := range arr {
		vs = append(vs, f(el))
	}
	return vs
}
