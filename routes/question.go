package routes

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/amrojjeh/arareader/model"
	"github.com/amrojjeh/arareader/ui/components"
	"github.com/amrojjeh/arareader/ui/page"
	"github.com/go-chi/chi/v5"
)

type questionResource struct{}

const questionIDParam = "questionID"

func (qr questionResource) Routes() chi.Router {
	r := chi.NewRouter()
	r.Get("/", qr.List)
	r.Post("/", qr.Post)
	r.Route(fmt.Sprintf("/{%s:[0-9]+}", questionIDParam), func(r chi.Router) {
		r.Use(qr.QuestionID)
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
	// TODO(Amr Ojjeh): Finish page
	page.QuestionPage(page.QuestionParams{
		Excerpt:          &model.Excerpt{},
		HighlightedRef:   0,
		QuizTitle:        "",
		Prompt:           "",
		QuestionNavProps: components.QuestionNavProps{},
		InputMethod:      nil,
	}).Render(w)
}

func (qr questionResource) Put(w http.ResponseWriter, r *http.Request) {
	http.Error(w, http.StatusText(http.StatusNotImplemented), http.StatusNotImplemented)
}

func (qr questionResource) Delete(w http.ResponseWriter, r *http.Request) {
	http.Error(w, http.StatusText(http.StatusNotImplemented), http.StatusNotImplemented)
}

func (qr questionResource) QuestionID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, questionIDParam)
		id, err := strconv.Atoi(idStr)
		if err != nil {
			err = fmt.Errorf("%s is not an int (%v)", questionIDParam, err)
			panic(err)
		}

		ctx := context.WithValue(r.Context(), questionIDParam, id)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
