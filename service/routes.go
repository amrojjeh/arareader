package service

import (
	"bytes"
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/amrojjeh/arareader/model"
	"github.com/amrojjeh/arareader/must"
	"github.com/amrojjeh/arareader/ui/page"
	"github.com/amrojjeh/arareader/ui/static"
)

func Server(logger *log.Logger, handler http.Handler, addr string) http.Server {
	return http.Server{
		Addr:              addr,
		Handler:           handler,
		ReadTimeout:       15 * time.Second,
		ReadHeaderTimeout: 15 * time.Second,
		WriteTimeout:      15 * time.Second,
		IdleTimeout:       15 * time.Second,
		ErrorLog:          logger,
	}
}

func NewRootHandler(logger *log.Logger, db *sql.DB) http.Handler {
	sm := scs.New()
	var handler http.Handler
	handler = rootHandler{
		db:      db,
		queries: model.New(db),
		sm:      sm,
		logger:  logger,
	}
	handler = sm.LoadAndSave(handler)
	handler = logRequest(logger, handler)
	return handler
}

func logRequest(logger *log.Logger, h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL.Path)
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
	switch shiftURL(r) {
	case "static":
		http.FileServer(http.FS(static.Files)).ServeHTTP(w, r)
	case ".":
		// FIXME(Amr Ojjeh): handle with a clientError panic recover
		quiz := must.Get(rh.queries.GetQuiz(r.Context(), 1))
		page.SVowel(page.SVowelParams{
			Excerpt:        bytes.NewReader(quiz.Excerpt),
			HighlightedRef: "1",
		}).Render(w)
	default:
		http.NotFound(w, r)
	}
}
