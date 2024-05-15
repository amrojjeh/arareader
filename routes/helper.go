package routes

import (
	"bytes"
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"path"
	"strconv"
	"strings"

	"github.com/amrojjeh/arareader/model"
	"github.com/amrojjeh/arareader/must"
	"github.com/amrojjeh/arareader/service"
)

func setURL(r *http.Request, val string) {
	r.URL.Path = val
	r.URL.RawPath = val
}

func shiftQuestion(r *http.Request, qs []model.Question) model.Question {
	questionIndex := shiftInteger(r)
	if questionIndex < 0 || questionIndex >= len(qs) {
		panic(clientError(http.StatusBadRequest))
	}
	return qs[questionIndex]
}

func shiftQuiz(r *http.Request, q *model.Queries) model.Quiz {
	id := shiftInteger(r)
	quiz, err := q.GetQuiz(r.Context(), id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			panic(clientError(http.StatusNotFound))
		}
		panic(err)
	}
	return quiz
}

func shiftInteger(r *http.Request) int {
	head := shiftPath(r)
	num, err := strconv.Atoi(head)
	if err != nil {
		panic(clientError(http.StatusBadRequest))
	}
	return num
}

func shiftPath(r *http.Request) string {
	head, tail := shiftPathHelper(r.URL.Path)
	setURL(r, tail)
	return head
}

// p is an absolute path
func shiftPathHelper(p string) (head string, tail string) {
	if p[0] != '/' {
		panic("routes.shiftPath: p must be an absolute path")
	}
	p = path.Clean(p[1:])
	si := strings.Index(p, "/")
	if si == -1 {
		return p, "/"
	}
	return p[:si], p[si:]
}

func allowMethods(r *http.Request, methods ...string) {
	violated := true
	for _, m := range methods {
		r.Header.Add("Allow", m)
		if r.Method == m {
			violated = false
		}
	}
	if violated {
		panic(clientError(http.StatusMethodNotAllowed))
	}
}

func (rh rootHandler) excerpt(quiz model.Quiz) *model.Excerpt {
	e, err := model.ExcerptFromXML(bytes.NewReader(quiz.Excerpt))
	if err != nil {
		panic(fmt.Sprintf("quiz excerpt cannot be parsed (id: %d). %s", quiz.ID, err.Error()))
	}
	return e
}

func (rh rootHandler) questions(r *http.Request, quiz model.Quiz) []model.Question {
	return must.Get(rh.queries.ListQuestionsByQuiz(r.Context(), quiz.ID))
}

func (qh quizHandler) applyVowelQuestions() {
	vowelQuestions := service.Filter(isVowelQuestionType, qh.questions)
	questionData := service.Map(model.MustParseQuestionData, vowelQuestions)
	refs := service.Map(excerptRef, questionData)
	qh.excerpt.UnpointRefs(refs)
}

func (qh quizHandler) vowelQuestions(r *http.Request) []model.Question {
	qs, err := qh.queries.ListQuestionsByQuizAndType(r.Context(), model.ListQuestionsByQuizAndTypeParams{
		QuizID: qh.quiz.ID,
		Type:   model.VowelQuestionType,
	})
	if err != nil {
		panic(fmt.Sprintf("expected quiz to exist but it does not (id: %d). %s", qh.quiz.ID, err.Error()))
	}
	return qs
}

// Useful for map
func excerptRef(qd model.QuestionData) int {
	return qd.Reference
}

func isVowelQuestionType(q model.Question) bool {
	return q.Type == model.VowelQuestionType
}
