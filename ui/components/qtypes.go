/*
Copyright © 2024 Amr Ojjeh <amrojjeh@outlook.com>
*/

package components

import (
	"github.com/amrojjeh/arareader/model"
	g "github.com/maragudk/gomponents"
	c "github.com/maragudk/gomponents/components"
	. "github.com/maragudk/gomponents/html"
)

func QuestionToInputMethod(question model.Question, session model.QuestionSession) g.Node {
	switch question.Type {
	case model.VowelQuestionType:
		if session.Status.IsSubmitted() {
			return VowelInputMethodSubmitted(question.Solution, session.Answer)
		}
		return VowelInputMethodUnsubmitted(question.Solution)
	case model.ShortAnswerQuestionType:
		if session.Status.IsSubmitted() {
			if session.Status == model.PendingQuestionStatus {
				return ShortAnswerInputMethodPending(session.Answer)
			}
			return ShortAnswerInputMethodSubmitted(question.Solution, session.Answer)
		}
		return ShortAnswerInputMethodUnsubmitted()
	}
	return P(
		g.Text("Question not implemented..."),
	)
}

func ShortAnswerInputMethodUnsubmitted() g.Node {
	return shortAnswerInput("", false)
}

func ShortAnswerInputMethodSubmitted(solution, answer string) g.Node {
	if solution == answer {
		return shortAnswerFeedback(answer, "/static/icons/circle-check-solid.svg")
	}
	return Div(Class("flex flex-col gap-3"),
		shortAnswerFeedback(answer, "/static/icons/circle-xmark-solid.svg"),
		shortAnswerFeedback(solution, "/static/icons/circle-check-solid.svg"),
	)
}

func shortAnswerFeedback(value, icon string) g.Node {
	return Div(Class("flex flex-row gap-2"),
		shortAnswerInput(value, true),
		shortAnswerIcon(icon),
	)
}

func shortAnswerIcon(src string) g.Node {
	return Img(Class("w-5"), Src(src))
}

func shortAnswerInput(value string, disabled bool) g.Node {
	return Input(Value(value),
		g.If(disabled, Disabled()),
		g.If(!disabled, g.Group([]g.Node{
			Data("input", ""),
			AutoFocus(),
		})),
		c.Classes{
			"shadow-md px-1 py-1": true,
			"bg-gray-200":         disabled,
		},
	)
}

func ShortAnswerInputMethodPending(answer string) g.Node {
	return Div(
		shortAnswerFeedback(answer, "/static/icons/circle-question-solid.svg"),
		P(Class("pt-1"),
			g.Text("Waiting for teacher's feedback..."),
		),
	)
}

func VowelInputMethodUnsubmitted(correct string) g.Node {
	buttons := []g.Node{}
	options, err := model.VowelQuestionOptions(correct)
	if err != nil {
		panic("generating options")
	}
	for _, o := range options {
		buttons = append(buttons, vowelButton(o, acceptingInput))
	}
	return vowelOptionsContainer(buttons)
}

func VowelInputMethodSubmitted(correct, chosen string) g.Node {
	buttons := []g.Node{}
	options, err := model.VowelQuestionOptions(correct)
	if err != nil {
		panic("generating options")
	}
	for _, o := range options {
		if o.Option == correct {
			buttons = append(buttons, vowelButton(o, correctStatus))
		} else if o.Option == chosen && chosen != correct {
			buttons = append(buttons, vowelButton(o, incorrectStatus))
		} else {
			buttons = append(buttons, vowelButton(o, neutralStatus))
		}
	}
	return vowelOptionsContainer(buttons)
}

func vowelOptionsContainer(buttons []g.Node) g.Node {
	return Div(Class("grid grid-cols-2 gap-2 mr-4"),
		g.Group(buttons),
	)
}

type inputStatus int

const (
	correctStatus = inputStatus(iota)
	incorrectStatus
	neutralStatus
	acceptingInput
)

func vowelButton(option model.VowelQuestionOption, status inputStatus) g.Node {
	return Button(Type("button"),
		g.If(status == acceptingInput, Data("substitute-button", option.Option)),
		g.If(status == correctStatus, Data("type", "primary")),
		Data("shortcut", option.Shortcut),
		c.Classes{
			"btn shadow py-2 text-2xl": true,
			"bg-red-300":               status == incorrectStatus,
		},
		g.Text(option.Option),
	)
}
