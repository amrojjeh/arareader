/*
Copyright © 2024 Amr Ojjeh <amrojjeh@outlook.com>
*/

package components

import (
	"github.com/amrojjeh/arareader/arabic"
	g "github.com/maragudk/gomponents"
	. "github.com/maragudk/gomponents/html"
)

func VowelInputMethodUnsubmitted(letter string, submitURL string) g.Node {
	options := vowelOptions(letter)
	return FormEl(Class("stack"), Action(submitURL), Method("post"),
		Input(Type("hidden"), Name("ans"), DataAttr("vowel-form-answer", "")),
		Div(Class("svowel-options"),
			options[0],
			Div(Class("svowel-options__not-sukoon"),
				g.Group(options[1:]),
			),
		),
		SubmitButton("Submit"),
	)
}

// TODO(Amr Ojjeh): Write a disabled version of the above
func VowelInputMethodSubmitted(letter, correct, chosen string) g.Node {
	return Div()
}

func vowelButton(text string) g.Node {
	return button(Type("button"), DataAttr("substitute-button", text),
		g.Text(text),
	)
}

func vowelOptions(letter string) []g.Node {
	vowels := []string{"o", "i", "a", "u", "K", "F", "N"}
	buttons := []g.Node{}
	for _, v := range vowels {
		buttons = append(buttons, vowelButton(letter+arabic.FromBuckwalter(v)))
	}
	return buttons
}
