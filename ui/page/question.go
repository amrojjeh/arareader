/*
Copyright © 2024 Amr Ojjeh <amrojjeh@outlook.com>
*/

package page

import (
	ar "github.com/amrojjeh/arareader/ui/components"
	g "github.com/maragudk/gomponents"
	htmx "github.com/maragudk/gomponents-htmx"
	c "github.com/maragudk/gomponents/components"
	. "github.com/maragudk/gomponents/html"
)

type QuestionParams struct {
	Excerpt     g.Node
	Prompt      string
	InputMethod g.Node
	NextURL     string
	PrevURL     string
}

func QuestionPage(p QuestionParams) g.Node {
	return c.HTML5(c.HTML5Props{
		Title:       "Arareader",
		Description: "Page to view quiz question",
		Language:    "en",
		Head: []g.Node{
			ar.CSS("/static/main.css"),
			ar.DeferJS("/static/main.js"),
			g.Raw("<script src='https://unpkg.com/htmx.org@1.9.12' integrity='sha384-ujb1lZYygJmzgSwoxRggbCHcjc0rB2XoQrxeTUQyRjrOnlCoYta87iKBWq3EsdM2' crossorigin='anonymous'></script>"),
			ar.Fonts(),
		},
		Body: []g.Node{ID("question-page"), htmx.Boost("true"),
			Nav(Class("nav"),
				ar.Hamburger(), // TODO(Amr Ojjeh): Pull up drawer
				ar.Icon(),
			),
			p.Excerpt,
			P(Class("prompt"),
				g.Text(p.Prompt),
			),
			p.InputMethod,
			Div(Class("question-ctrl"),
				A(Class("question-ctrl__first btn"), Href(p.PrevURL),
					g.Text("Previous"),
				),
				Button(Type("button"), Class("btn btn--primary"),
					g.Text("Submit"),
				),
				A(Class("question-ctrl__last btn"), Href(p.NextURL),
					g.Text("Next"),
				),
			),
		},
	})
}
