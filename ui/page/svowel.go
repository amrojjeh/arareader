/*
Copyright © 2024 Amr Ojjeh <amrojjeh@outlook.com>
*/

package page

import (
	"github.com/amrojjeh/arareader/arabic"
	"github.com/amrojjeh/arareader/service"
	ar "github.com/amrojjeh/arareader/ui/components"
	g "github.com/maragudk/gomponents"
	c "github.com/maragudk/gomponents/components"
	. "github.com/maragudk/gomponents/html"
)

type SVowelParams struct {
	Excerpt          *service.Excerpt
	HighlightedRef   int
	QuizTitle        string
	QuestionNavProps ar.QuestionNavProps
}

// TODO(Amr Ojjeh): Make the nav and question into a grid so that centering can be done
func SVowel(p SVowelParams) g.Node {
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
			Header(
				ar.Icon(),
				H1(Class("title"),
					g.Text(p.QuizTitle),
				),
				Button(Type("button"), Class("username button"),
					g.Text("Amr Ojjeh")),
			),
			ar.QuestionNav(p.QuestionNavProps),
			ar.Excerpt(p.Excerpt, p.HighlightedRef),
			P(Class("instruction"),
				g.Text("Enter the correct vowel"),
			),
			FormEl(Class("stack"),
				Div(Class("svowel-options"),
					ar.VowelButton(arabic.FromBuckwalter("lo")),
					Div(Class("svowel-options__not-sukoon"),
						ar.VowelButton(arabic.FromBuckwalter("li")),
						ar.VowelButton(arabic.FromBuckwalter("la")),
						ar.VowelButton(arabic.FromBuckwalter("lu")),
						ar.VowelButton(arabic.FromBuckwalter("lK")),
						ar.VowelButton(arabic.FromBuckwalter("lF")),
						ar.VowelButton(arabic.FromBuckwalter("lN")),
					),
				),
				ar.SubmitButton("Submit"),
			),
		},
	})
}
