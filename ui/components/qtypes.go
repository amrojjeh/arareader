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
		Button(Class("button"), Type("submit"),
			g.Text("Submit"),
		),
	)
}

func VowelInputMethodSubmitted(letter, correct, chosen string) g.Node {
	options := vowelOptionsDisabled(letter, correct, chosen)
	return Div(Class("svowel-options"),
		options[0],
		Div(Class("svowel-options__not-sukoon"),
			g.Group(options[1:]),
		),
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

func vowelOptionsDisabled(letter, correct, chosen string) []g.Node {
	vowels := []string{"o", "i", "a", "u", "K", "F", "N"}
	buttons := []g.Node{}
	for _, v := range vowels {
		buttons = append(buttons, vowelButtonDisabled(letter+arabic.FromBuckwalter(v), correct, chosen))
	}
	return buttons
}

func vowelButton(text string) g.Node {
	return Button(Class("button"), Type("button"), DataAttr("substitute-button", text),
		g.Text(text),
	)
}

func vowelButtonDisabled(text, correct, chosen string) g.Node {
	if text == correct {
		return vowelButtonCorrect(text)
	}
	if text == chosen {
		return vowelButtonIncorrect(text)
	}
	return Button(Class("button button--disabled"), Type("button"),
		g.Text(text),
	)
}

func vowelButtonCorrect(text string) g.Node {
	return Button(Class("primary button button--disabled"), Type("button"), DataAttr("substitute-button", text),
		g.Text(text),
	)
}

func vowelButtonIncorrect(text string) g.Node {
	return Button(Class("danger button button--disabled"), Type("button"),
		g.Text(text),
	)
}
