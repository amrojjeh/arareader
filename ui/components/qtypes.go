/*
Copyright © 2024 Amr Ojjeh <amrojjeh@outlook.com>
*/

package components

import (
	"github.com/amrojjeh/arareader/model"
	g "github.com/maragudk/gomponents"
	. "github.com/maragudk/gomponents/html"
)

type QuestionToInputMethodParams struct {
	Question        model.Question
	QuestionSession model.QuestionSession
}

func QuestionToInputMethod(p QuestionToInputMethodParams) g.Node {
	switch p.Question.Type {
	case model.VowelQuestionType:
		if p.QuestionSession.Status.IsSubmitted() {
			return VowelInputMethodSubmitted(p.Question.Solution, p.QuestionSession.Answer)
		}
		return VowelInputMethodUnsubmitted(p.Question.Solution)
	case model.ShortAnswerQuestionType:
		if p.QuestionSession.Status.IsSubmitted() {
			return ShortAnswerInputMethodSubmitted()
		}
		return ShortAnswerInputMethodUnsubmitted()
	}
	return P(
		g.Text("Question not implemented..."),
	)
}

func ShortAnswerInputMethodUnsubmitted() g.Node {
	return Input(Class("short-answer"), DataAttr("input", ""), AutoFocus())
}

func ShortAnswerInputMethodSubmitted() g.Node {
	return Input(Class("short-answer"))
}

func VowelInputMethodUnsubmitted(correct string) g.Node {
	btns := vowelOptions(correct)
	return g.Group([]g.Node{
		Div(Class("svowel-options"),
			g.Group(btns),
		),
	})
}

func VowelInputMethodSubmitted(correct, chosen string) g.Node {
	btns := vowelOptionsDisabled(correct, chosen)
	return Div(Class("svowel-options"),
		btns[0],
		g.Group(btns[1:]),
	)
}

func vowelOptions(correct string) []g.Node {
	buttons := []g.Node{}
	options, err := model.VowelQuestionOptions(correct)
	if err != nil {
		panic("generating options")
	}
	for _, o := range options {
		buttons = append(buttons, vowelButton(o))
	}
	return buttons
}

func vowelOptionsDisabled(correct, chosen string) []g.Node {
	buttons := []g.Node{}
	options, err := model.VowelQuestionOptions(correct)
	if err != nil {
		panic("generating options")
	}
	for _, o := range options {
		buttons = append(buttons, vowelButtonDisabled(o, correct, chosen))
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
