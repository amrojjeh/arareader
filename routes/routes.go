/*
Copyright © 2024 Amr Ojjeh <amrojjeh@outlook.com>
*/

package routes

import (
	"database/sql"
	"net/http"

	"github.com/amrojjeh/arareader/ui/static"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Routes(db *sql.DB) http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.CleanPath)
	r.Use(middleware.StripSlashes)
	r.Get("/static*", func(w http.ResponseWriter, r *http.Request) {
		http.StripPrefix("/static", http.FileServer(http.FS(static.Files))).ServeHTTP(w, r)
	})
	r.Mount("/quiz", quizResource{}.Routes())
	return r
}

// func newQuizHandler(r *http.Request, rh rootHandler, quiz model.Quiz) http.Handler {
// 	excerpt := must.Get(model.ExcerptFromQuiz(quiz))
// 	qs := must.Get(rh.queries.ListQuestionsByQuiz(r.Context(), quiz.ID))
// 	sqs, _ := rh.fetchQuizSession(r, quiz.ID, 1) // TEMP(Amr Ojjeh): Temporary until there's class management
// 	return quizHandler{
// 		rootHandler: rh,
// 		quiz:        quiz,
// 		excerpt:     excerpt,
// 		questions:   qs,
// 		quizSession: sqs,
// 	}
// }

// func (qh quizHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	allowMethods(r, http.MethodGet, http.MethodPost)
// 	switch r.Method {
// 	case http.MethodGet:
// 		qh.getMethod(w, r)
// 	case http.MethodPost:
// 		qh.postMethod(w, r)
// 	}
// }

// func (qh quizHandler) getMethod(w http.ResponseWriter, r *http.Request) {
// 	// TODO(Amr Ojjeh): Send feedback (with animation?)
// 	qh.applyVowelQuestions()
// 	qh.applyVowelAnswers(r)
// 	question := shiftQuestion(r, qh.questions)
// 	qd := model.MustParseQuestionData(question)
// 	inputMethod := components.QuestionToInputMethod(components.QuestionToInputMethodParams{
// 		SubmitURL:       qh.submitURL(question.Position),
// 		NextURL:         qh.nextURL(question.Position),
// 		Question:        question,
// 		QuestionData:    qd,
// 		QuestionSession: qh.questionSession(r, question),
// 	})
// 	page.QuestionPage(page.QuestionParams{
// 		Excerpt:        qh.excerpt,
// 		HighlightedRef: qd.Reference,
// 		QuizTitle:      qh.quiz.Title,
// 		Prompt:         qd.Prompt,
// 		InputMethod:    inputMethod(),
// 		QuestionNavProps: components.QuestionNavProps{
// 			CurrentQuestion: question.Position + 1,
// 			TotalQuestions:  len(qh.questions),
// 			SkipForwardURL:  "",
// 			SkipBackwardURL: "",
// 			NextURL:         qh.nextURL(question.Position),
// 			PrevURL:         qh.prevURL(question.Position),
// 		},
// 	}).Render(w)
// }

// func (qh quizHandler) submitURL(questionPos int) string {
// 	return fmt.Sprintf("/quiz/%d/%d", qh.quiz.ID, questionPos)
// }

// func (qh quizHandler) nextURL(questionPos int) string {
// 	if questionPos+1 < len(qh.questions) {
// 		return qh.questionURL(questionPos + 1)
// 	}
// 	return ""
// }

// func (qh quizHandler) prevURL(questionPos int) string {
// 	if questionPos-1 >= 0 {
// 		return qh.questionURL(questionPos - 1)
// 	}
// 	return ""
// }

// func (qh quizHandler) questionURL(questionPos int) string {
// 	return fmt.Sprintf("/quiz/%d/%d", qh.quiz.ID, questionPos)
// }

// func (qh quizHandler) postMethod(w http.ResponseWriter, r *http.Request) {
// 	question := shiftQuestion(r, qh.questions)
// 	qh.mustParseForm(r)
// 	ans := r.Form.Get("ans")
// 	data := model.MustParseQuestionData(question)
// 	if model.ValidateQuestion(data, ans) {
// 		qh.submitAnswer(r, question, ans, model.CorrectQuestionStatus)
// 	} else {
// 		qh.submitAnswer(r, question, ans, model.IncorrectQuestionStatus)
// 	}
// 	http.Redirect(w, r, qh.questionURL(question.Position), http.StatusSeeOther)
// }
