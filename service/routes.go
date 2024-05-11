package service

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/amrojjeh/arareader/arabic"
	"github.com/amrojjeh/arareader/model"
	"github.com/amrojjeh/arareader/ui/components"
	"github.com/amrojjeh/arareader/ui/page"
	"github.com/amrojjeh/arareader/ui/static"
)

func Server(logger *log.Logger, handler HTTPRoute, addr string) http.Server {
	return http.Server{
		Addr:              addr,
		Handler:           HTTPRoute{},
		ReadTimeout:       15 * time.Second,
		ReadHeaderTimeout: 15 * time.Second,
		WriteTimeout:      15 * time.Second,
		IdleTimeout:       15 * time.Second,
		ErrorLog:          logger,
	}
}

type HTTPRoute struct {
	DB      *sql.DB
	Queries model.Queries
}

func (hr HTTPRoute) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch shiftURL(r) {
	case "static":
		http.FileServer(http.FS(static.Files)).ServeHTTP(w, r)
	case ".":
		// TODO(Amr Ojjeh): load an actual excerpt from DB
		page.SVowel(page.SVowelParams{
			Excerpt: components.Excerpt(arabic.FromBuckwalter("h*A bytN.")),
			Prompt:  "Hm?",
		}).Render(w)
	default:
		http.NotFound(w, r)
	}
}
