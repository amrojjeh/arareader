/*
Copyright © 2024 Amr Ojjeh <amrojjeh@outlook.com>
*/

package components

import (
	"unicode/utf8"

	"github.com/amrojjeh/arareader/arabic"
	"github.com/amrojjeh/arareader/model"
	g "github.com/maragudk/gomponents"
	. "github.com/maragudk/gomponents/html"
)

type QuestionToInputMethodParams struct {
	Question        model.Question
	QuestionSession model.QuestionSession
}

var vowels = []string{"u", "N", "a", "F", "i", "K", "o"}

func QuestionToInputMethod(p QuestionToInputMethodParams) func() g.Node {
	switch p.Question.Type {
	case model.VowelQuestionType:
		return func() g.Node {
			letter, _ := utf8.DecodeRuneInString(p.Question.Solution)
			if p.QuestionSession.Status.IsSubmitted() {
				return VowelInputMethodSubmitted(string(letter), p.Question.Solution, p.QuestionSession.Answer)
			}
			return VowelInputMethodUnsubmitted(string(letter))
		}
	}
	return func() g.Node {
		return P(
			g.Text("Question not implemented..."),
		)
	}
}

func VowelInputMethodUnsubmitted(letter string) g.Node {
	options := vowelOptions(letter)
	return g.Group([]g.Node{
		Div(Class("svowel-options"),
			g.Group(options),
		),
	})
}

func VowelInputMethodSubmitted(letter, correct, chosen string) g.Node {
	options := vowelOptionsDisabled(letter, correct, chosen)
	return Div(Class("svowel-options"),
		options[0],
		g.Group(options[1:]),
	)
}

func vowelOptions(letter string) []g.Node {
	buttons := []g.Node{}
	for _, v := range vowels {
		buttons = append(buttons, vowelButton(letter+arabic.FromBuckwalter(v)))
	}
	return buttons
}

func vowelOptionsDisabled(letter, correct, chosen string) []g.Node {
	buttons := []g.Node{}
	for _, v := range vowels {
		buttons = append(buttons, vowelButtonDisabled(letter+arabic.FromBuckwalter(v), correct, chosen))
	}
	return buttons
}

func vowelButton(text string) g.Node {
	return Button(Class("btn btn--shadow"), Type("button"), DataAttr("substitute-button", text),
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
	return Button(Class("btn"), Type("button"), Disabled(),
		g.Text(text),
	)
}

func vowelButtonCorrect(text string) g.Node {
	return Button(Class("btn btn--primary"), Type("button"), Disabled(),
		g.Text(text),
	)
}

func vowelButtonIncorrect(text string) g.Node {
	return Button(Class("btn btn--danger"), Type("button"), Disabled(),
		g.Text(text),
	)
}
