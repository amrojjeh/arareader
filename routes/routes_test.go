/*
Copyright © 2024 Amr Ojjeh <amrojjeh@outlook.com>
*/

package routes

import (
	"context"
	"database/sql"
	"net/http"
	"net/url"
	"testing"

	"github.com/amrojjeh/arareader/arabic"
	"github.com/amrojjeh/arareader/demo"
	"github.com/amrojjeh/arareader/model"
	"github.com/stretchr/testify/assert"
)

func TestHTTPRouteServeHTTP(t *testing.T) {
	tests := []struct {
		name string
		path string
		code int
		body string
	}{
		{
			name: "Root",
			path: "/",
			code: http.StatusSeeOther,
		},
		{
			name: "Static",
			path: "/static/main.css",
			code: http.StatusOK,
			body: "body",
		},
		{
			name: "Does not exist",
			path: "/doesnotexist",
			code: http.StatusNotFound,
			body: "404",
		},
		{
			name: "Question",
			path: "/quiz/1/question/0",
			code: http.StatusOK,
			body: "html",
		},
		{
			name: "Negative question",
			path: "/quiz/1/question/-1",
			code: http.StatusNotFound,
		},
		{
			name: "Too many questions",
			path: "/quiz/1/question/5",
			code: http.StatusBadRequest,
		},
		{
			name: "Quiz does not exist",
			path: "/quiz/0/question/2",
			code: http.StatusNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			db := model.MustOpenDB(":memory:")
			model.MustSetup(ctx, db)
			demo.Demo(ctx, db)

			handler := Routes(db)
			assert.HTTPStatusCode(t, handler.ServeHTTP, http.MethodGet, tt.path, nil, tt.code)
			assert.HTTPBodyContains(t, handler.ServeHTTP, http.MethodGet, tt.path, nil, tt.body)
		})
	}
}

func TestShortVowel(t *testing.T) {
	r := Routes(demoDB(t))
	assert.HTTPStatusCode(t, r.ServeHTTP, http.MethodGet, "/quiz/1/question/0", nil, http.StatusOK)
	assert.HTTPBodyContains(t, r.ServeHTTP, http.MethodGet, "/quiz/1/question/0", nil, arabic.FromBuckwalter("a"))
	assert.HTTPBodyContains(t, r.ServeHTTP, http.MethodGet, "/quiz/1/question/0", nil, "</form>")
	assert.HTTPRedirect(t, r.ServeHTTP, http.MethodPost, "/quiz/1/question/0", url.Values{
		"ans": []string{arabic.FromBuckwalter("lo")},
	})
	assert.HTTPBodyNotContains(t, r.ServeHTTP, http.MethodGet, "/quiz/1/question/0", nil, "</form>")
}

func TestShortVowelEmpty(t *testing.T) {
	r := Routes(demoDB(t))
	assert.HTTPStatusCode(t, r.ServeHTTP, http.MethodGet, "/quiz/1/question/0", nil, http.StatusOK)
	assert.HTTPBodyContains(t, r.ServeHTTP, http.MethodGet, "/quiz/1/question/0", nil, arabic.FromBuckwalter("a"))
	assert.HTTPBodyContains(t, r.ServeHTTP, http.MethodGet, "/quiz/1/question/0", nil, "</form>")
	assert.HTTPRedirect(t, r.ServeHTTP, http.MethodPost, "/quiz/1/question/0", url.Values{
		"ans": []string{" "},
	})
	assert.HTTPBodyContains(t, r.ServeHTTP, http.MethodGet, "/quiz/1/question/0", nil, "</form>")
}

func TestShortVowelInvalid(t *testing.T) {
	r := Routes(demoDB(t))
	assert.HTTPStatusCode(t, r.ServeHTTP, http.MethodGet, "/quiz/1/question/0", nil, http.StatusOK)
	assert.HTTPBodyContains(t, r.ServeHTTP, http.MethodGet, "/quiz/1/question/0", nil, arabic.FromBuckwalter("a"))
	assert.HTTPBodyContains(t, r.ServeHTTP, http.MethodGet, "/quiz/1/question/0", nil, "</form>")
	assert.HTTPRedirect(t, r.ServeHTTP, http.MethodPost, "/quiz/1/question/0", url.Values{
		"ans": []string{arabic.FromBuckwalter("to")},
	})
	assert.HTTPBodyContains(t, r.ServeHTTP, http.MethodGet, "/quiz/1/question/0", nil, "</form>")
}

func TestShortAnswer(t *testing.T) {
	r := Routes(demoDB(t))
	assert.HTTPStatusCode(t, r.ServeHTTP, http.MethodGet, "/quiz/1/question/4", nil, http.StatusOK)
	assert.HTTPBodyContains(t, r.ServeHTTP, http.MethodGet, "/quiz/1/question/4", nil, "What is the meaning of the word")
	assert.HTTPBodyContains(t, r.ServeHTTP, http.MethodGet, "/quiz/1/question/4", nil, "</form>")

	assert.HTTPRedirect(t, r.ServeHTTP, http.MethodPost, "/quiz/1/question/4", url.Values{
		"ans": []string{"intention"},
	})
	assert.HTTPBodyNotContains(t, r.ServeHTTP, http.MethodGet, "/quiz/1/question/4", nil, "</form>")
	assert.HTTPBodyNotContains(t, r.ServeHTTP, http.MethodGet, "/quiz/1/question/4", nil, "circle-xmark")
	assert.HTTPBodyContains(t, r.ServeHTTP, http.MethodGet, "/quiz/1/question/4", nil, "check")
}

func TestShortAnswerInvalid(t *testing.T) {
	r := Routes(demoDB(t))
	assert.HTTPStatusCode(t, r.ServeHTTP, http.MethodGet, "/quiz/1/question/4", nil, http.StatusOK)
	assert.HTTPBodyContains(t, r.ServeHTTP, http.MethodGet, "/quiz/1/question/4", nil, "What is the meaning of the word")
	assert.HTTPBodyContains(t, r.ServeHTTP, http.MethodGet, "/quiz/1/question/4", nil, "</form>")

	assert.HTTPRedirect(t, r.ServeHTTP, http.MethodPost, "/quiz/1/question/4", url.Values{
		"ans": []string{"inten"},
	})
	assert.HTTPBodyNotContains(t, r.ServeHTTP, http.MethodGet, "/quiz/1/question/4", nil, "</form>")
	assert.HTTPBodyContains(t, r.ServeHTTP, http.MethodGet, "/quiz/1/question/4", nil, "circle-xmark")
	assert.HTTPBodyContains(t, r.ServeHTTP, http.MethodGet, "/quiz/1/question/4", nil, "check")
}

func TestShortAnswerEmpty(t *testing.T) {
	r := Routes(demoDB(t))
	assert.HTTPStatusCode(t, r.ServeHTTP, http.MethodGet, "/quiz/1/question/4", nil, http.StatusOK)
	assert.HTTPBodyContains(t, r.ServeHTTP, http.MethodGet, "/quiz/1/question/4", nil, "What is the meaning of the word")
	assert.HTTPBodyContains(t, r.ServeHTTP, http.MethodGet, "/quiz/1/question/4", nil, "</form>")

	assert.HTTPRedirect(t, r.ServeHTTP, http.MethodPost, "/quiz/1/question/4", url.Values{
		"ans": []string{""},
	})
	assert.HTTPBodyContains(t, r.ServeHTTP, http.MethodGet, "/quiz/1/question/4", nil, "</form>")
	assert.HTTPBodyNotContains(t, r.ServeHTTP, http.MethodGet, "/quiz/1/question/4", nil, "circle-xmark")
	assert.HTTPBodyNotContains(t, r.ServeHTTP, http.MethodGet, "/quiz/1/question/4", nil, "check")
}

func demoDB(t *testing.T) *sql.DB {
	t.Helper()
	ctx := context.Background()
	db := model.MustOpenDB(":memory:")
	model.MustSetup(ctx, db)
	demo.Demo(ctx, db)
	return db
}
