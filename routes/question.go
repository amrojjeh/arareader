package routes

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/amrojjeh/arareader/model"
	"github.com/amrojjeh/arareader/ui/components"
	"github.com/amrojjeh/arareader/ui/page"
	"github.com/go-chi/chi/v5"
)

type questionResource struct {
	db *sql.DB
}

const questionPositionParam = "questionID"

func (qr questionResource) Routes() chi.Router {
	r := chi.NewRouter()
	r.Get("/", qr.List)
	r.Post("/", qr.Post)
	r.Route(fmt.Sprintf("/{%s:[0-9]+}", questionPositionParam), func(r chi.Router) {
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
	quizID := quizIDFromRequest(r)
	q := model.New(qr.db)
	quiz, err := q.GetQuiz(r.Context(), quizID)
	if err != nil {
		err = fmt.Errorf("retrieving quiz: %v", err)
		panic(err)
	}
	questionPos := questionPositionFromRequest(r)
	questions, err := q.ListQuestionsByQuiz(r.Context(), quizID)
	if err != nil {
		err = fmt.Errorf("retrieving questions: %v", err)
		panic(err)
	}
	data := model.MustParseQuestionData(questions[questionPos])
	excerpt, err := model.ExcerptFromQuiz(quiz)
	if err != nil {
		err = fmt.Errorf("parsing excerpt from quiz: %v", err)
		panic(err)
	}
	page.QuestionPage(page.QuestionParams{
		Excerpt:     components.Excerpt(excerpt, data.Reference),
		QuizTitle:   quiz.Title,
		Prompt:      data.Prompt,
		InputMethod: nil,
	}).Render(w)
}

func (qr questionResource) Put(w http.ResponseWriter, r *http.Request) {
	http.Error(w, http.StatusText(http.StatusNotImplemented), http.StatusNotImplemented)
}

func (qr questionResource) Delete(w http.ResponseWriter, r *http.Request) {
	http.Error(w, http.StatusText(http.StatusNotImplemented), http.StatusNotImplemented)
}

func (qr questionResource) QuestionPosition(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, questionPositionParam)
		id, err := strconv.Atoi(idStr)
		if err != nil {
			err = fmt.Errorf("%s is not an int (%v)", questionPositionParam, err)
			panic(err)
		}

		ctx := context.WithValue(r.Context(), questionPositionParam, id)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func questionPositionFromRequest(r *http.Request) int {
	return r.Context().Value(questionPositionParam).(int)
}
