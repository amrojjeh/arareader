/*
Copyright © 2024 Amr Ojjeh <amrojjeh@outlook.com>
*/

package page

import (
	"github.com/amrojjeh/arareader/model"
	ar "github.com/amrojjeh/arareader/ui/components"
	g "github.com/maragudk/gomponents"
	c "github.com/maragudk/gomponents/components"
	. "github.com/maragudk/gomponents/html"
)

type QuestionParams struct {
	Excerpt          *model.Excerpt
	HighlightedRef   int
	QuizTitle        string
	Prompt           string
	QuestionNavProps ar.QuestionNavProps
	InputMethod      g.Node
}

// TODO(Amr Ojjeh): Make the nav and question into a grid so that centering can be done
func QuestionPage(p QuestionParams) g.Node {
	return c.HTML5(c.HTML5Props{
		Title:       "Arareader",
		Description: "TO DO",
		Language:    "en",
		Head: []g.Node{
			ar.CSS("/static/main.css"),
			ar.DeferJS("/static/main.js"),
			ar.Fonts(),
		},
		Body: []g.Node{ID("svowel"),
			Nav(
				ar.Hamburger(),
				ar.Icon(),
			),
		},
	})
}
