/*
Copyright © 2024 Amr Ojjeh <amrojjeh@outlook.com>
*/

package routes

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"

	"github.com/amrojjeh/arareader/arabic"
	"github.com/amrojjeh/arareader/model"
	"github.com/amrojjeh/arareader/ui/components"
	"github.com/amrojjeh/arareader/ui/page"
)

type questionSessionResource struct {
	db *sql.DB
}

func (qr questionSessionResource) Post(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	var quizID, studentID, questionPos int
	var found bool
	if studentID, found = studentIDFromRequest(r); !found {
		panic("getting student id")
	}
	if quizID, found = quizIDFromRequest(r); !found {
		panic("getting quiz id")
	}
	if questionPos, found = questionPositionFromRequest(r); !found {
		panic("getting question position")
	}

	q := model.New(qr.db)

	quizSession, err := q.GetQuizSession(r.Context(), model.GetQuizSessionParams{
		StudentID: studentID,
		QuizID:    quizID,
	})
	if err != nil {
		err = fmt.Errorf("retrieving quiz session: %v", err)
		panic(err)
	}

	questions, err := q.ListQuestionsByQuiz(r.Context(), quizID)
	if err != nil {
		err = fmt.Errorf("retrieving questions: %v", err)
		panic(err)
	}
	question := questions[questionPos]

	questionSession, err := q.GetQuestionSession(r.Context(), model.GetQuestionSessionParams{
		QuizSessionID: quizSession.ID,
		QuestionID:    question.ID,
	})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			questionSession, err = q.CreateQuestionSession(r.Context(), model.CreateQuestionSessionParams{
				QuizSessionID: quizSession.ID,
				QuestionID:    question.ID,
				Answer:        r.Form.Get("ans"),
				Status:        model.ValidateQuestion(question, r.Form.Get("ans")),
			})
			if err != nil {
				err = fmt.Errorf("creating question: %v", err)
				panic(err)
			}
		}
	} else {
		_, err = q.SubmitAnswer(r.Context(), model.SubmitAnswerParams{
			Answer: r.Form.Get("ans"),
			Status: model.ValidateQuestion(question, r.Form.Get("ans")),
			ID:     questionSession.ID,
		})
		if err != nil {
			err = fmt.Errorf("submitting answer: %v", err)
			panic(err)
		}
	}

	http.Redirect(w, r, questionURL(quizID, questionPos, len(questions)), http.StatusSeeOther)
}

func (qr questionSessionResource) List(w http.ResponseWriter, r *http.Request) {
	http.Error(w, http.StatusText(http.StatusNotImplemented), http.StatusNotImplemented)
}

func (qr questionSessionResource) Get(w http.ResponseWriter, r *http.Request) {
	var quizID, studentID, questionPos int
	var found bool
	if studentID, found = studentIDFromRequest(r); !found {
		panic("getting student id")
	}
	if quizID, found = quizIDFromRequest(r); !found {
		panic("getting quiz id")
	}
	if questionPos, found = questionPositionFromRequest(r); !found {
		panic("getting question position")
	}

	q := model.New(qr.db)
	quiz, err := q.GetQuiz(r.Context(), quizID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			http.NotFound(w, r)
			return
		}
		err = fmt.Errorf("retrieving quiz: %v", err)
		panic(err)
	}
	questions, err := q.ListQuestionsByQuiz(r.Context(), quizID)
	if err != nil {
		err = fmt.Errorf("retrieving questions: %v", err)
		panic(err)
	}

	if questionPos < 0 || questionPos >= len(questions) {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	question := questions[questionPos]

	excerpt, err := model.ExcerptFromQuiz(quiz)
	if err != nil {
		err = fmt.Errorf("parsing excerpt from quiz: %v", err)
		panic(err)
	}

	for _, question := range questions {
		if question.Type == model.VowelQuestionType {
			excerpt.UnpointRef(question.Reference)
		}
	}

	quizSession, err := q.GetQuizSession(r.Context(), model.GetQuizSessionParams{
		StudentID: studentID,
		QuizID:    quizID,
	})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			quizSession, err = q.CreateQuizSession(r.Context(), model.CreateQuizSessionParams{
				StudentID: studentID,
				QuizID:    quizID,
				Status:    model.UnsubmittedQuizStatus,
			})
			if err != nil {
				err = fmt.Errorf("creating quiz session: %v", err)
				panic(err)
			}
		} else {
			err = fmt.Errorf("getting quiz session: %v", err)
			panic(err)
		}
	}

	questionSessions, err := q.ListQuestionSessions(r.Context(), quizSession.ID)
	if err != nil {
		err = fmt.Errorf("getting vowel question sessions: %v", err)
		panic(err)
	}

	for _, qs := range questionSessions {
		if qs.Status.IsSubmitted() {
			question := questionWithID(questions, qs.QuestionID)
			excerpt.Ref(question.Reference).ReplaceWithText(question.Solution)
		}
	}

	page.QuestionPage(page.QuestionParams{
		Excerpt:     components.Excerpt(excerpt, question.Reference),
		Prompt:      question.Prompt,
		InputMethod: components.VowelInputMethodUnsubmitted(arabic.Unpointed(question.Solution)),
		NextURL:     questionURL(quiz.ID, questionPos+1, len(questions)),
		PrevURL:     questionURL(quiz.ID, questionPos-1, len(questions)),
	}).Render(w)
}

func (qr questionSessionResource) Put(w http.ResponseWriter, r *http.Request) {
	http.Error(w, http.StatusText(http.StatusNotImplemented), http.StatusNotImplemented)
}

func (qr questionSessionResource) Delete(w http.ResponseWriter, r *http.Request) {
	http.Error(w, http.StatusText(http.StatusNotImplemented), http.StatusNotImplemented)
}

func questionWithID(qs []model.Question, id int) model.Question {
	for _, q := range qs {
		if q.ID == id {
			return q
		}
	}
	err := fmt.Errorf("getting question with id %d", id)
	panic(err)
}

func questionURL(quizID, questionPos, totalQuestions int) string {
	if questionPos < 0 || questionPos >= totalQuestions {
		return ""
	}
	return fmt.Sprintf("/quiz/%d/question/%d", quizID, questionPos)
}
