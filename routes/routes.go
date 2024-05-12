package routes

import (
	"bytes"
	"database/sql"
	"log"
	"net/http"
	"runtime/debug"
	"strconv"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/amrojjeh/arareader/model"
	"github.com/amrojjeh/arareader/must"
	"github.com/amrojjeh/arareader/service"
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
	switch head := shiftURL(r); head {
	case "static":
		http.FileServer(http.FS(static.Files)).ServeHTTP(w, r)
	case ".":
		// NOTE(Amr Ojjeh): Not an issue since this is assuming demo data
		quiz := must.Get(rh.queries.GetQuiz(r.Context(), 1))
		qs := must.Get(rh.queries.ListQuestionsByQuizAndType(r.Context(), model.ListQuestionsByQuizAndTypeParams{
			QuizID: quiz.ID,
			Type:   string(service.VowelQuestionType),
		}))
		e := must.Get(service.ExcerptFromXML(bytes.NewReader(quiz.Excerpt)))
		service.ApplyVowelQuestionsToExcerpt(qs, e)
		page.SVowel(page.SVowelParams{
			Excerpt:         e,
			HighlightedRef:  1,
			QuizTitle:       quiz.Title,
			CurrentQuestion: 1,
			TotalQuestions:  4,
		}).Render(w)
	default:
		// TEMP(Amr Ojjeh): Just for fun
		if i, err := strconv.Atoi(head); err == nil {
			quiz := must.Get(rh.queries.GetQuiz(r.Context(), 1))
			qs := must.Get(rh.queries.ListQuestionsByQuizAndType(r.Context(), model.ListQuestionsByQuizAndTypeParams{
				QuizID: quiz.ID,
				Type:   string(service.VowelQuestionType),
			}))
			e := must.Get(service.ExcerptFromXML(bytes.NewReader(quiz.Excerpt)))
			service.ApplyVowelQuestionsToExcerpt(qs, e)
			page.SVowel(page.SVowelParams{
				Excerpt:         e,
				HighlightedRef:  i,
				QuizTitle:       quiz.Title,
				CurrentQuestion: 1,
				TotalQuestions:  4,
			}).Render(w)
			return
		}
		http.NotFound(w, r)
	}
}
