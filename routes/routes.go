/*
Copyright © 2024 Amr Ojjeh <amrojjeh@outlook.com>
*/

package routes

import (
	"context"
	"database/sql"
	"net/http"
	"time"

	"github.com/alexedwards/scs/sqlite3store"
	"github.com/alexedwards/scs/v2"
	"github.com/amrojjeh/arareader/model"
	"github.com/amrojjeh/arareader/ui/static"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type rootResource struct {
	sm *scs.SessionManager
	db *sql.DB
}

func Routes(db *sql.DB) http.Handler {
	sm := scs.New()
	sm.Lifetime = time.Hour * 24
	sm.Store = sqlite3store.New(db)

	rs := rootResource{
		sm,
		db,
	}

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.CleanPath)
	r.Use(middleware.StripSlashes)
	r.Use(sm.LoadAndSave)
	r.Use(rs.htmxVary)
	r.Use(rs.Auth)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/quiz/2/question/0", http.StatusSeeOther)
	})
	r.Post("/restart", rs.Restart)
	r.Get("/static*", func(w http.ResponseWriter, r *http.Request) {
		http.StripPrefix("/static", http.FileServer(http.FS(static.Files))).ServeHTTP(w, r)
	})
	r.Mount("/quiz", quizResource{db}.Routes())
	return r
}

// TEMP(Amr Ojjeh): Temporary until there's class management
func (rs rootResource) Restart(w http.ResponseWriter, r *http.Request) {
	q := model.New(rs.db)
	err := q.DeleteQuizSessions(r.Context())
	if err != nil {
		panic(err)
	}
	w.Header().Add("HX-Redirect", "/quiz/2/question/0")
}

// TODO(Amr Ojjeh): Actually retrieve studentID from session manager once login is supported
func (rs rootResource) Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), studentIDKey, 1)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (rs rootResource) htmxVary(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Vary", "HX-Request")
		w.Header().Add("Vary", "HX-Target")
		next.ServeHTTP(w, r)
	})
}
