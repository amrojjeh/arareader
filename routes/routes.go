package routes

import (
	"database/sql"
	"log"
	"net/http"
	"runtime/debug"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/amrojjeh/arareader/model"
	"github.com/amrojjeh/arareader/service"
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
	logger  *log.Logger
}

func (rh rootHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch head := shiftPath(r); head {
	case "static":
		http.FileServer(http.FS(static.Files)).ServeHTTP(w, r)
	case ".":
		http.Redirect(w, r, "/1", http.StatusSeeOther)
	case "question":
		rh.serveQuiz(w, r)
	default:
		http.NotFound(w, r)
	}
}

// TODO(Amr Ojjeh): Turn into its own handler to store quiz and excerpt data and make API clearer
func (rh rootHandler) serveQuiz(w http.ResponseWriter, r *http.Request) {
	questionNum := shiftInteger(r)
	quiz := rh.quiz(r, 1) // TEMP(Amr Ojjeh): Until we add more quizzes
	e := rh.applyVowelQuestions(r, quiz)
	qs, _ := rh.queries.ListQuestionsByQuiz(r.Context(), quiz.ID) // TEMP(Amr Ojjeh): Awful code. Change ASAP
	q := qs[questionNum]
	qd, _ := service.ExtractQuestionData(q) // TEMP(Amr Ojjeh): More awful code
	page.SVowel(page.SVowelParams{
		Excerpt:        e,
		HighlightedRef: qd.RefID(),
		QuizTitle:      quiz.Title,
		QuestionNavProps: components.QuestionNavProps{
			CurrentQuestion: questionNum + 1,
			TotalQuestions:  len(qs),
			SkipForwardURL:  "",
			SkipBackwardURL: "",
			NextURL:         "",
			PrevURL:         "",
		},
	}).Render(w)
	return
}
