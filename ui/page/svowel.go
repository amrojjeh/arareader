/*
Copyright © 2024 Amr Ojjeh <amrojjeh@outlook.com>
*/

package page

import (
	"github.com/amrojjeh/arareader/arabic"
	na "github.com/amrojjeh/arareader/ui/components"
	g "github.com/maragudk/gomponents"
	c "github.com/maragudk/gomponents/components"
	. "github.com/maragudk/gomponents/html"
)

type SVowelParams struct {
	Excerpt g.Node
	Prompt  string
}

func SVowel(p SVowelParams) g.Node {
	return c.HTML5(c.HTML5Props{
		Title:       "NahwApp",
		Description: "TO DO",
		Language:    "en",
		Head: []g.Node{
			na.CSS("/static/main.css"),
			na.Fonts(),
		},
		Body: []g.Node{ID("svowel"),
			Header(
				na.Icon(),
				H1(Class("title"),
					g.Text(arabic.FromBuckwalter("Hdyv 1")),
				),
				Button(Type("button"), Class("username button"),
					g.Text("Amr Ojjeh")),
			),
			na.QuestionNav(),
			p.Excerpt,
			P(Class("instruction"),
				g.Text("Enter the correct vowel"),
			),
			FormEl(Class("stack"),
				Div(Class("svowel-options"),
					na.Button(arabic.FromBuckwalter("lo")),
					Div(Class("svowel-options--not-sukoon"),
						na.Button(arabic.FromBuckwalter("li")),
						na.Button(arabic.FromBuckwalter("la")),
						na.Button(arabic.FromBuckwalter("lu")),
						na.Button(arabic.FromBuckwalter("lK")),
						na.Button(arabic.FromBuckwalter("lF")),
						na.Button(arabic.FromBuckwalter("lN")),
					),
				),
				na.SubmitButton("Submit"),
			),
		},
	})
}
