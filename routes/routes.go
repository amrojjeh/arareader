/*
Copyright © 2024 Amr Ojjeh <amrojjeh@outlook.com>
*/

package routes

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
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
	q  *model.Queries
}

func Routes(db *sql.DB) http.Handler {
	sm := scs.New()
	sm.Store = sqlite3store.New(db)
	sm.Lifetime = time.Hour * 24

	rs := rootResource{
		sm,
		db,
		model.New(db),
	}

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(clientErrorRecoverer)
	r.Use(middleware.CleanPath)
	r.Use(middleware.StripSlashes)
	r.Use(sm.LoadAndSave)
	r.Use(rs.htmxVary)
	r.Use(rs.Auth)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/quiz/2/question/0", http.StatusSeeOther)
	})
	r.Get("/static*", func(w http.ResponseWriter, r *http.Request) {
		http.StripPrefix("/static", http.FileServer(http.FS(static.Files))).ServeHTTP(w, r)
	})

	r.Route(fmt.Sprintf("/quiz/{%s:[0-9]+}/question/{%s:[0-9]+}", quizIDKey, questionPosKey), func(r chi.Router) {
		r.Use(quizID, questionPosition)
		r.Get("/", rs.questionGet)
		r.Post("/", rs.questionPost)

		r.Get("/htmx/select", rs.htmxSelect)
	})

	r.With(quizID).
		Get(fmt.Sprintf("/quiz/{%s:[0-9]+}/summary", quizIDKey), rs.summaryGet)

	r.With(quizID).Delete(fmt.Sprintf("/quiz/{%s:[0-9]+}", quizIDKey), rs.restart)

	return r
}

func (rs rootResource) clientError(code int) {
	panic(clientErr(code))
}

type clientErr int

func clientErrorRecoverer(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				ce, ok := r.(clientErr)
				if !ok {
					panic(r)
				}
				http.Error(w, http.StatusText(int(ce)), int(ce))
			}
		}()
		h.ServeHTTP(w, r)
	})
}

func quizID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, string(quizIDKey))
		id, err := strconv.Atoi(idStr)
		if err != nil {
			err = fmt.Errorf("%s is not an int (%v)", quizIDKey, err)
			panic(err)
		}

		ctx := context.WithValue(r.Context(), quizIDKey, id)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func questionPosition(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, string(questionPosKey))
		id, err := strconv.Atoi(idStr)
		if err != nil {
			err = fmt.Errorf("%s is not an int (%v)", questionPosKey, err)
			panic(err)
		}

		ctx := context.WithValue(r.Context(), questionPosKey, id)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// TEMP(Amr Ojjeh): Temporary until there's class management
func (rs rootResource) restart(w http.ResponseWriter, r *http.Request) {
	studentID, found := studentIDFromRequest(r)
	if !found {
		panic("getting student id")
	}

	quiz := rs.quiz(r)
	err := rs.q.DeleteQuizSessions(r.Context(), model.DeleteQuizSessionsParams{
		StudentID: studentID,
		QuizID:    quiz.ID,
	})
	if err != nil {
		panic(err)
	}

	w.Header().Add("HX-Redirect", summaryURL(quiz.ID))
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
