/*
Copyright © 2024 Amr Ojjeh <amrojjeh@outlook.com>
*/

package page

import (
	"io"

	"github.com/amrojjeh/arareader/arabic"
	ar "github.com/amrojjeh/arareader/ui/components"
	g "github.com/maragudk/gomponents"
	c "github.com/maragudk/gomponents/components"
	. "github.com/maragudk/gomponents/html"
)

type SVowelParams struct {
	Excerpt        io.Reader
	HighlightedRef string
}

// TODO(Amr Ojjeh): Fix the multiple choice at the bottom of the page
func SVowel(p SVowelParams) g.Node {
	return c.HTML5(c.HTML5Props{
		Title:       "Arareader",
		Description: "TO DO",
		Language:    "en",
		Head: []g.Node{
			ar.CSS("/static/main.css"),
			ar.Fonts(),
		},
		Body: []g.Node{ID("svowel"),
			Header(
				ar.Icon(),
				H1(Class("title"),
					g.Text(arabic.FromBuckwalter("Hdyv 1")),
				),
				Button(Type("button"), Class("username button"),
					g.Text("Amr Ojjeh")),
			),
			ar.QuestionNav(),
			ar.Excerpt(p.Excerpt, p.HighlightedRef),
			P(Class("instruction"),
				g.Text("Enter the correct vowel"),
			),
			FormEl(Class("stack"),
				Div(Class("svowel-options"),
					ar.DefaultButton(arabic.FromBuckwalter("lo")),
					Div(Class("svowel-options--not-sukoon"),
						ar.DefaultButton(arabic.FromBuckwalter("li")),
						ar.DefaultButton(arabic.FromBuckwalter("la")),
						ar.DefaultButton(arabic.FromBuckwalter("lu")),
						ar.DefaultButton(arabic.FromBuckwalter("lK")),
						ar.DefaultButton(arabic.FromBuckwalter("lF")),
						ar.DefaultButton(arabic.FromBuckwalter("lN")),
					),
				),
				ar.SubmitButton("Submit"),
			),
		},
	})
}
