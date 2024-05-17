/*
Copyright © 2024 Amr Ojjeh <amrojjeh@outlook.com>
*/

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

func (rh rootHandler) mustParseForm(r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		panic(clientError(http.StatusBadRequest))
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

func (rh rootHandler) quizSession(r *http.Request, quizID, studentID int) (model.QuizSession, bool) {
	quizSession, err := rh.queries.GetQuizSession(r.Context(), model.GetQuizSessionParams{
		StudentID: studentID,
		QuizID:    quizID,
	})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.QuizSession{}, false
		}
		panic(fmt.Sprintf("could not retrieve quiz session: %s", err))
	}
	return quizSession, true
}

func (qh quizHandler) applyVowelQuestions() {
	vowelQuestions := service.Filter(isVowelQuestionType, qh.questions)
	questionData := service.Map(model.MustParseQuestionData, vowelQuestions)
	refs := service.Map(excerptRef, questionData)
	qh.excerpt.UnpointRefs(refs)
}

func (qh quizHandler) applyVowelAnswers(r *http.Request) {
	vowelSessions := must.Get(qh.queries.ListQuestionSessionByType(r.Context(), model.ListQuestionSessionByTypeParams{
		QuizSessionID: qh.quizSession.ID,
		Type:          model.VowelQuestionType,
	}))
	ids := map[int]bool{}
	for _, vs := range vowelSessions {
		if vs.Status.IsSubmitted() {
			ids[vs.QuestionID] = true
		}
	}
	for _, q := range qh.questions {
		if ids[q.ID] {
			data := model.MustParseQuestionData(q)
			qh.excerpt.Ref(data.Reference).ReplaceWithText(data.Answer)
		}
	}
}

// Useful for map
func excerptRef(qd model.QuestionData) int {
	return qd.Reference
}

func isVowelQuestionType(q model.Question) bool {
	return q.Type == model.VowelQuestionType
}

func (qh quizHandler) submitAnswer(r *http.Request, q model.Question, answer string, status model.QuestionStatus) {
	_, err := qh.queries.SubmitAnswer(r.Context(), model.SubmitAnswerParams{
		Answer:        answer,
		Status:        status,
		QuizSessionID: qh.quizSession.ID,
		QuestionID:    q.ID,
	})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			qh.questionSession(r, q)
			// TEST(Amr Ojjeh): VERY IMPORTANT TO TEST THIS BRANCH
			qh.submitAnswer(r, q, answer, status)
		}
		panic(err)
	}
}

func (qh quizHandler) questionSession(r *http.Request, q model.Question) model.QuestionSession {
	questionSession, err := qh.queries.GetQuestionSession(r.Context(), model.GetQuestionSessionParams{
		QuizSessionID: qh.quizSession.ID,
		QuestionID:    q.ID,
	})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return must.Get(qh.queries.CreateQuestionSession(r.Context(), model.CreateQuestionSessionParams{
				QuizSessionID: qh.quizSession.ID,
				QuestionID:    q.ID,
				Answer:        "",
				Status:        model.UnattemptedQuestionStatus,
			}))
		}
		panic(err)
	}
	return questionSession
}
