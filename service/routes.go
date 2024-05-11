package service

import (
	"net/http"

	"github.com/amrojjeh/arareader/arabic"
	"github.com/amrojjeh/arareader/ui/components"
	"github.com/amrojjeh/arareader/ui/page"
	"github.com/amrojjeh/arareader/ui/static"
)

type HTTPRoute struct {
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
