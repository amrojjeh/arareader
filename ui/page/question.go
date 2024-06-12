/*
Copyright © 2024 Amr Ojjeh <amrojjeh@outlook.com>
*/

package page

import (
	ar "github.com/amrojjeh/arareader/ui/components"
	g "github.com/maragudk/gomponents"
	c "github.com/maragudk/gomponents/components"
	. "github.com/maragudk/gomponents/html"
)

type QuestionParams struct {
	Excerpt     g.Node
	QuizTitle   string
	Prompt      string
	InputMethod g.Node
}

// TODO(Amr Ojjeh): Make the nav and question into a grid so that centering can be done
func QuestionPage(p QuestionParams) g.Node {
	return c.HTML5(c.HTML5Props{
		Title:       "Arareader",
		Description: "Page to view quiz question",
		Language:    "en",
		Head: []g.Node{
			ar.CSS("/static/main.css"),
			ar.DeferJS("/static/main.js"),
			ar.Fonts(),
		},
		Body: []g.Node{ID("question-page"),
			Nav(Class("nav"),
				ar.Hamburger(), // TODO(Amr Ojjeh): Pull up drawer
				ar.Icon(),
			),
			p.Excerpt,
		},
	})
}
