package model

import (
	"context"
	"testing"
)

func TestOpenDB(t *testing.T) {
	// there should not be any panics
	db := MustOpenDB(":memory:")
	defer db.Close()
}

func TestSetup(t *testing.T) {
	db := MustOpenDB(":memory:")
	defer db.Close()

	ctx := context.Background()
	MustSetup(ctx, db)

	rows, err := db.Query("SELECT name FROM sqlite_master WHERE type='table'")
	if err != nil {
		t.Error("could not query table names", err)
	}

	expectedNames := map[string]bool{"teacher": false, "quiz": false, "question": false, "class": false,
		"class_quiz": false, "student": false, "quiz_session": false, "question_session": false,
		"sessions": false}
	for rows.Next() {
		var name string
		err = rows.Scan(&name)
		if err != nil {
			t.Error("could not scan table name", err)
		}
		if _, ok := expectedNames[name]; !ok {
			t.Errorf("found unexpected table (name: %s)", name)
		}
		expectedNames[name] = true
	}

	for name, found := range expectedNames {
		if !found {
			t.Errorf("table was not found (name: %s)", name)
		}
	}
}
