package routes

import (
	"bytes"
	"fmt"
	"net/http"
	"path"
	"strconv"
	"strings"

	"github.com/amrojjeh/arareader/model"
	"github.com/amrojjeh/arareader/service"
)

func setURL(r *http.Request, val string) {
	r.URL.Path = val
	r.URL.RawPath = val
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

func (rh rootHandler) quiz(r *http.Request, id int) model.Quiz {
	quiz, err := rh.queries.GetQuiz(r.Context(), id)
	if err != nil {
		panic(clientError(http.StatusNotFound))
	}
	return quiz
}

func (rh rootHandler) vowelQuestions(r *http.Request, quiz model.Quiz) []model.Question {
	qs, err := rh.queries.ListQuestionsByQuizAndType(r.Context(), model.ListQuestionsByQuizAndTypeParams{
		QuizID: quiz.ID,
		Type:   string(service.VowelQuestionType),
	})
	if err != nil {
		panic(fmt.Sprintf("expected quiz to exist but it does not (id: %d). %s", quiz.ID, err.Error()))
	}
	return qs
}

func (rh rootHandler) excerpt(quiz model.Quiz) *service.Excerpt {
	e, err := service.ExcerptFromXML(bytes.NewReader(quiz.Excerpt))
	if err != nil {
		panic(fmt.Sprintf("quiz excerpt cannot be parsed (id: %d). %s", quiz.ID, err.Error()))
	}
	return e
}

func (rh rootHandler) applyVowelQuestions(r *http.Request, quiz model.Quiz) *service.Excerpt {
	e := rh.excerpt(quiz)
	vowelQuestions := rh.vowelQuestions(r, quiz)
	vowelData := service.Map(service.MustVowelQuestionData, vowelQuestions)
	// NOTE(Amr Ojjeh): This error should probably be logged at least
	// ignoring error since we can't do anything to recover from that
	_ = service.ApplyVowelDataToExcerpt(vowelData, e)
	return e
}
