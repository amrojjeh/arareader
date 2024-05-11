/*
Copyright © 2024 Amr Ojjeh <amrojjeh@outlook.com>
*/

package components

import (
	"fmt"

	g "github.com/maragudk/gomponents"
	h "github.com/maragudk/gomponents/html"
)

func Icon() g.Node {
	return h.Div(h.Class("nahwapp-icon"),
		g.Text("NA"),
	)
}

func CSS(path string) g.Node {
	return h.Link(h.Rel("stylesheet"), h.Href(path))
}

func roboto() string {
	return "family=Roboto:ital,wght@0,100;0,300;0,400;0,500;0,700;0,900;1,100;1,300;1,400;1,500;1,700;1,900&"
}

func notoNaskh() string {
	return "family=Noto+Naskh+Arabic:wght@400..700&"
}

func permanentMarker() string {
	return "family=Permanent+Marker&"
}

func Fonts() g.Node {
	fonts := fmt.Sprintf("https://fonts.googleapis.com/css2?%vdisplay=swap", notoNaskh()+roboto()+permanentMarker())
	return g.Group([]g.Node{
		h.Link(h.Rel("preconnect"), h.Href("https://fonts.googleapis.com")),
		h.Link(h.Rel("preconnect"), h.Href("https://fonts.gstatic.com"), g.Attr("crossorigin")),
		h.Link(h.Href(fonts), h.Rel("stylesheet")),
	})
}
