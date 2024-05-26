/*
Copyright © 2024 Amr Ojjeh <amrojjeh@outlook.com>
*/

package routes

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"runtime/debug"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/amrojjeh/arareader/model"
	"github.com/amrojjeh/arareader/must"
	"github.com/amrojjeh/arareader/ui/components"
	"github.com/amrojjeh/arareader/ui/page"
	"github.com/amrojjeh/arareader/ui/static"
)

func Server(handler http.Handler, addr string) http.Server {
	return http.Server{
		Addr:              addr,
		Handler:           handler,
		ReadTimeout:       15 * time.Second,
		ReadHeaderTimeout: 15 * time.Second,
		WriteTimeout:      15 * time.Second,
		IdleTimeout:       15 * time.Second,
		ErrorLog:          log.Default(),
	}
}

func NewRootHandler(db *sql.DB) http.Handler {
	sm := scs.New()
	var handler http.Handler
	handler = rootHandler{
		db:      db,
		queries: model.New(db),
		sm:      sm,
	}
	handler = sm.LoadAndSave(handler)
	handler = gracefulRecovery(handler)
	handler = logRequest(handler)
	return handler
}

type clientError int

func gracefulRecovery(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rec := recover(); rec != nil {
				if clientErr, ok := rec.(clientError); ok {
					http.Error(w, http.StatusText(int(clientErr)), int(clientErr))
					return
				}
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				log.Printf("INTERNAL SERVER ERROR: %+v", rec)
				debug.PrintStack()
			}
		}()
		h.ServeHTTP(w, r)
	})
}

func logRequest(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL.Path)
		h.ServeHTTP(w, r)
	})
}

type rootHandler struct {
	db      *sql.DB
	queries *model.Queries
	sm      *scs.SessionManager
}

// TODO(Amr Ojjeh): Redirect if HTTP to HTTPS
func (rh rootHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch head := shiftPath(r); head {
	case "static":
		allowMethods(r, http.MethodGet)
		http.FileServer(http.FS(static.Files)).ServeHTTP(w, r)
	case ".":
		allowMethods(r, http.MethodGet)
		http.Redirect(w, r, "/quiz/1/0", http.StatusSeeOther)
	case "quiz":
		quiz := shiftQuiz(r, rh.queries)
		newQuizHandler(r, rh, quiz).ServeHTTP(w, r)
	default:
		http.NotFound(w, r)
	}
}

// quizHandler serves a particualr quiz
type quizHandler struct {
	rootHandler
	quiz        model.Quiz
	quizSession model.QuizSession
	excerpt     *model.Excerpt
	questions   []model.Question
}

func newQuizHandler(r *http.Request, rh rootHandler, quiz model.Quiz) http.Handler {
	excerpt := must.Get(model.ExcerptFromQuiz(quiz))
	qs := must.Get(rh.queries.ListQuestionsByQuiz(r.Context(), quiz.ID))
	sqs, _ := rh.fetchQuizSession(r, quiz.ID, 1) // TEMP(Amr Ojjeh): Temporary until there's class management
	return quizHandler{
		rootHandler: rh,
		quiz:        quiz,
		excerpt:     excerpt,
		questions:   qs,
		quizSession: sqs,
	}
}

func (qh quizHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	allowMethods(r, http.MethodGet, http.MethodPost)
	switch r.Method {
	case http.MethodGet:
		qh.getMethod(w, r)
	case http.MethodPost:
		qh.postMethod(w, r)
	}
}

func (qh quizHandler) getMethod(w http.ResponseWriter, r *http.Request) {
	// TODO(Amr Ojjeh): Send feedback (with animation?)
	qh.applyVowelQuestions()
	qh.applyVowelAnswers(r)
	question := shiftQuestion(r, qh.questions)
	qd := model.MustParseQuestionData(question)
	inputMethod := components.QuestionToInputMethod(components.QuestionToInputMethodParams{
		SubmitURL:       fmt.Sprintf("/quiz/1/%d", question.Position), // TEMP(Amr Ojjeh): Turn into a function
		Question:        question,
		QuestionData:    qd,
		QuestionSession: qh.questionSession(r, question),
	})
	page.QuestionPage(page.QuestionParams{
		Excerpt:        qh.excerpt,
		HighlightedRef: qd.Reference,
		QuizTitle:      qh.quiz.Title,
		Prompt:         qd.Prompt,
		InputMethod:    inputMethod(),
		QuestionNavProps: components.QuestionNavProps{
			CurrentQuestion: question.Position + 1,
			TotalQuestions:  len(qh.questions),
			SkipForwardURL:  "",
			SkipBackwardURL: "",
			NextURL:         "",
			PrevURL:         "",
		},
	}).Render(w)

}

func (qh quizHandler) postMethod(w http.ResponseWriter, r *http.Request) {
	question := shiftQuestion(r, qh.questions)
	qh.mustParseForm(r)
	ans := r.Form.Get("ans")
	data := model.MustParseQuestionData(question)
	if model.ValidateQuestion(data, ans) {
		qh.submitAnswer(r, question, ans, model.CorrectQuestionStatus)
	} else {
		qh.submitAnswer(r, question, ans, model.IncorrectQuestionStatus)
	}
	http.Redirect(w, r, "/quiz/1/0", http.StatusSeeOther)
}
