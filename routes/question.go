package routes

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type questionResource struct {
	db *sql.DB
}

func (qr questionResource) Routes() chi.Router {
	r := chi.NewRouter()
	r.Get("/", qr.List)
	r.Post("/", qr.Post)
	r.Route(fmt.Sprintf("/{%s:[0-9]+}", questionPosKey), func(r chi.Router) {
		r.Use(qr.QuestionPosition)
		r.Get("/", qr.Get)
		r.Put("/", qr.Put)
		r.Delete("/", qr.Delete)
	})
	return r
}

func (qr questionResource) Post(w http.ResponseWriter, r *http.Request) {
	http.Error(w, http.StatusText(http.StatusNotImplemented), http.StatusNotImplemented)
}

func (qr questionResource) List(w http.ResponseWriter, r *http.Request) {
	http.Error(w, http.StatusText(http.StatusNotImplemented), http.StatusNotImplemented)
}

func (qr questionResource) Get(w http.ResponseWriter, r *http.Request) {
	if r.Context().Value(studentIDKey) != nil {
		questionSessionResource{qr.db}.Get(w, r)
		return
	}
	http.Error(w, http.StatusText(http.StatusNotImplemented), http.StatusNotImplemented)
}

func (qr questionResource) Put(w http.ResponseWriter, r *http.Request) {
	http.Error(w, http.StatusText(http.StatusNotImplemented), http.StatusNotImplemented)
}

func (qr questionResource) Delete(w http.ResponseWriter, r *http.Request) {
	http.Error(w, http.StatusText(http.StatusNotImplemented), http.StatusNotImplemented)
}

func (qr questionResource) QuestionPosition(next http.Handler) http.Handler {
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
