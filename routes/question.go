/*
Copyright © 2024 Amr Ojjeh <amrojjeh@outlook.com>
*/

package routes

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/amrojjeh/arareader/model"
	"github.com/amrojjeh/arareader/ui/components"
	"github.com/amrojjeh/arareader/ui/page"
)

func (rs rootResource) questionPost(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	studentID, found := studentIDFromRequest(r)
	if !found {
		panic("getting student id")
	}

	quiz := rs.quiz(r)
	questions := rs.questions(r, quiz.ID)
	question := rs.selectedQuestion(r, questions)

	ans := r.Form.Get("ans")
	if strings.TrimSpace(ans) == "" {
		http.Redirect(w, r, questionURLUnsafe(quiz.ID, question.Position), http.StatusSeeOther)
		return
	}

	quizSession, err := rs.q.GetQuizSession(r.Context(), model.GetQuizSessionParams{
		StudentID: studentID,
		QuizID:    quiz.ID,
	})
	if err != nil {
		err = fmt.Errorf("retrieving quiz session: %v", err)
		panic(err)
	}

	if !model.ValidateQuestionInput(question, ans) {
		http.Redirect(w, r, questionURL(quiz.ID, question.Position, len(questions)), http.StatusSeeOther)
		return
	}

	questionSession, err := rs.q.GetQuestionSession(r.Context(), model.GetQuestionSessionParams{
		QuizSessionID: quizSession.ID,
		QuestionID:    question.ID,
	})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			questionSession, err = rs.q.CreateQuestionSession(r.Context(), model.CreateQuestionSessionParams{
				QuizSessionID: quizSession.ID,
				QuestionID:    question.ID,
				Answer:        ans,
				Status:        model.ValidateQuestion(question, ans),
			})
			if err != nil {
				err = fmt.Errorf("creating question: %v", err)
				panic(err)
			}
		}
	} else {
		_, err = rs.q.SubmitAnswer(r.Context(), model.SubmitAnswerParams{
			Answer: ans,
			Status: model.ValidateQuestion(question, ans),
			ID:     questionSession.ID,
		})
		if err != nil {
			err = fmt.Errorf("submitting answer: %v", err)
			panic(err)
		}
	}

	if r.Header.Get("HX-Target") != "" {
		rs.htmxSelect(w, r)
		return
	}
	http.Redirect(w, r, questionURL(quiz.ID, question.Position, len(questions)), http.StatusSeeOther)
}

func (rs rootResource) questionGet(w http.ResponseWriter, r *http.Request) {
	studentID, found := studentIDFromRequest(r)
	if !found {
		panic("getting student id")
	}

	quiz := rs.quiz(r)
	questions := rs.questions(r, quiz.ID)
	question := rs.selectedQuestion(r, questions)

	quizSession, err := rs.q.GetQuizSession(r.Context(), model.GetQuizSessionParams{
		StudentID: studentID,
		QuizID:    quiz.ID,
	})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			quizSession, err = rs.q.CreateQuizSession(r.Context(), model.CreateQuizSessionParams{
				StudentID: studentID,
				QuizID:    quiz.ID,
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

	questionSession := model.QuestionSession{
		Status: model.UnattemptedQuestionStatus,
	}

	questionSessions := rs.questionSessions(r, quizSession.ID)
	excerpt := rs.excerpt(quiz, questions, questionSessions)

	for _, qs := range questionSessions {
		if qs.QuestionID == question.ID {
			questionSession = qs
		}
	}

	submitURL, feedback := "", ""
	if !questionSession.Status.IsSubmitted() {
		submitURL = questionURL(quiz.ID, question.Position, len(questions))
	} else {
		feedback = question.Feedback
	}

	page.QuestionPage(page.QuestionParams{
		Excerpt:          components.Excerpt(false, excerpt, question.Reference),
		Prompt:           question.Prompt,
		InputMethod:      components.QuestionToInputMethod(question, questionSession),
		SidebarQuestions: sidebar(quiz.ID, questionSessions, questions),
		NextURL:          questionURLHTMX(quiz.ID, question.Position+1, len(questions)),
		PrevURL:          questionURLHTMX(quiz.ID, question.Position-1, len(questions)),
		SubmitURL:        submitURL,
		SummaryURL:       summaryURL(quiz.ID),
		Feedback:         feedback,
	}).Render(w)
}

func (rs rootResource) htmxSelect(w http.ResponseWriter, r *http.Request) {
	studentID, found := studentIDFromRequest(r)
	if !found {
		panic("expected studentID")
	}

	quiz := rs.quiz(r)

	questions := rs.questions(r, quiz.ID)
	question := rs.selectedQuestion(r, questions)
	quizSession := rs.quizSession(r, studentID, quiz.ID)
	questionSessions := rs.questionSessions(r, quizSession.ID)
	var questionSession model.QuestionSession

	for _, qs := range questionSessions {
		if question.ID == qs.QuestionID {
			questionSession = qs
		}
	}

	excerpt := rs.excerpt(quiz, questions, questionSessions)

	feedback := ""
	if questionSession.Status.IsSubmitted() {
		feedback = question.Feedback
	}

	w.Header().Add("HX-Push-Url", questionURL(quiz.ID, question.Position, len(questions)))
	page.Question(question.Prompt,
		components.QuestionToInputMethod(question, questionSession),
		feedback,
		questionURLHTMX(quiz.ID, question.Position-1, len(questions)),
		questionURLHTMX(quiz.ID, question.Position+1, len(questions)),
		questionURL(quiz.ID, question.Position, len(questions)),
	).Render(w)
	page.Sidebar(true, sidebar(quiz.ID, questionSessions, questions), summaryURL(quiz.ID)).Render(w)
	components.Excerpt(true, excerpt, question.Reference).Render(w)
}

func (rs rootResource) quiz(r *http.Request) model.Quiz {
	quizID, found := quizIDFromRequest(r)
	if !found {
		panic("getting quiz id")
	}
	quiz, err := rs.q.GetQuiz(r.Context(), quizID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			rs.clientError(http.StatusNotFound)
		}
		panic(err)
	}
	return quiz
}

func (rs rootResource) questions(r *http.Request, quizID int) []model.Question {
	all, err := rs.q.ListQuestionsByQuiz(r.Context(), quizID)
	if err != nil {
		err = fmt.Errorf("retrieving questions: %v", err)
		panic(err)
	}
	return all
}

func (rs rootResource) selectedQuestion(r *http.Request, questions []model.Question) model.Question {
	questionPos, found := questionPositionFromRequest(r)
	if !found {
		panic("getting question position")
	}

	if questionPos >= len(questions) || questionPos < 0 {
		rs.clientError(http.StatusBadRequest)
	}

	return questions[questionPos]
}

func (rs rootResource) excerpt(quiz model.Quiz, questions []model.Question, questionSessions []model.QuestionSession) *model.ReferenceNode {
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

	for _, qs := range questionSessions {
		if qs.Status.IsSubmitted() && questionWithID(questions, qs.QuestionID).Type == model.VowelQuestionType {
			question := questionWithID(questions, qs.QuestionID)
			excerpt.Ref(question.Reference).ReplaceWithText(question.Solution)
		}
	}

	return excerpt
}

func (rs rootResource) quizSession(r *http.Request, studentID, quizID int) model.QuizSession {
	quizSession, err := rs.q.GetQuizSession(r.Context(), model.GetQuizSessionParams{
		StudentID: studentID,
		QuizID:    quizID,
	})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			rs.clientError(http.StatusNotFound)
		}
		panic(err)
	}
	return quizSession
}

func (rs rootResource) questionSessions(r *http.Request, quizSessionID int) []model.QuestionSession {
	questionSessions, err := rs.q.ListQuestionSessions(r.Context(), quizSessionID)
	if err != nil {
		err = fmt.Errorf("getting vowel question sessions: %v", err)
		panic(err)
	}
	return questionSessions
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

func questionURLHTMX(quizID, questionPos, totalQuestions int) string {
	url := questionURL(quizID, questionPos, totalQuestions)
	if url == "" {
		return url
	}
	return url + "/htmx/select"
}

func summaryURL(quizID int) string {
	return fmt.Sprintf("/quiz/%d/summary", quizID)
}

func questionURLUnsafe(quizID, questionPos int) string {
	return fmt.Sprintf("/quiz/%d/question/%d", quizID, questionPos)
}

func sidebar(quizID int, questionSessions []model.QuestionSession, questions []model.Question) []page.SidebarQuestion {
	sqs := make([]page.SidebarQuestion, 0, len(questions))

	for i, question := range questions {
		sq := page.SidebarQuestion{
			Prompt: question.Prompt,
			Status: model.UnattemptedQuestionStatus,
			URL:    questionURLHTMX(quizID, i, len(questions)),
			Target: true,
		}
		for _, session := range questionSessions {
			if session.QuestionID == question.ID {
				sq.Status = session.Status
			}
		}
		sqs = append(sqs, sq)
	}
	return sqs
}
