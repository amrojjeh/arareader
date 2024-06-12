package routes

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type quizResource struct {
	db *sql.DB
}

const quizIDParam = "quizID"

func (qr quizResource) Routes() chi.Router {
	r := chi.NewRouter()
	r.Get("/", qr.List)
	r.Post("/", qr.Post)
	r.Route(fmt.Sprintf("/{%s:[0-9]+}", quizIDParam), func(r chi.Router) {
		r.Use(qr.QuizID)
		r.Get("/", qr.Get)
		r.Put("/", qr.Put)
		r.Delete("/", qr.Delete)
		r.Mount("/question", questionResource{qr.db}.Routes())
	})
	return r
}

func (qr quizResource) Post(w http.ResponseWriter, r *http.Request) {
	http.Error(w, http.StatusText(http.StatusNotImplemented), http.StatusNotImplemented)
}

func (qr quizResource) List(w http.ResponseWriter, r *http.Request) {
	http.Error(w, http.StatusText(http.StatusNotImplemented), http.StatusNotImplemented)
}

func (qr quizResource) Get(w http.ResponseWriter, r *http.Request) {
	http.Error(w, http.StatusText(http.StatusNotImplemented), http.StatusNotImplemented)
}

func (qr quizResource) Put(w http.ResponseWriter, r *http.Request) {
	http.Error(w, http.StatusText(http.StatusNotImplemented), http.StatusNotImplemented)
}

func (qr quizResource) Delete(w http.ResponseWriter, r *http.Request) {
	http.Error(w, http.StatusText(http.StatusNotImplemented), http.StatusNotImplemented)
}

func (qr quizResource) QuizID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, quizIDParam)
		id, err := strconv.Atoi(idStr)
		if err != nil {
			err = fmt.Errorf("%s is not an int (%v)", quizIDParam, err)
			panic(err)
		}

		ctx := context.WithValue(r.Context(), quizIDParam, id)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func quizIDFromRequest(r *http.Request) int {
	return r.Context().Value(quizIDParam).(int)
}
