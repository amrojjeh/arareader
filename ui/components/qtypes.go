/*
Copyright © 2024 Amr Ojjeh <amrojjeh@outlook.com>
*/

package components

import (
	"github.com/amrojjeh/arareader/arabic"
	g "github.com/maragudk/gomponents"
	. "github.com/maragudk/gomponents/html"
)

func VowelInputMethod(submitURL string) g.Node {
	return FormEl(Class("stack"), Action(submitURL), Method("post"),
		Input(Type("hidden"), Name("ans"), DataAttr("vowel-form-answer", "")),
		Div(Class("svowel-options"),
			vowelButton(arabic.FromBuckwalter("lo")),
			Div(Class("svowel-options__not-sukoon"),
				vowelButton(arabic.FromBuckwalter("li")),
				vowelButton(arabic.FromBuckwalter("la")),
				vowelButton(arabic.FromBuckwalter("lu")),
				vowelButton(arabic.FromBuckwalter("lK")),
				vowelButton(arabic.FromBuckwalter("lF")),
				vowelButton(arabic.FromBuckwalter("lN")),
			),
		),
		SubmitButton("Submit"),
	)
}

func vowelButton(text string) g.Node {
	return button(Type("button"), DataAttr("substitute-button", text),
		g.Text(text),
	)
}
