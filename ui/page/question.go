/*
Copyright © 2024 Amr Ojjeh <amrojjeh@outlook.com>
*/

package page

import (
	ar "github.com/amrojjeh/arareader/ui/components"
	g "github.com/maragudk/gomponents"
	htmx "github.com/maragudk/gomponents-htmx"
	c "github.com/maragudk/gomponents/components"
	. "github.com/maragudk/gomponents/html"
)

type QuestionParams struct {
	Excerpt     g.Node
	Prompt      string
	InputMethod g.Node
	NextURL     string
	PrevURL     string
	SubmitURL   string
	Feedback    string
}

func QuestionPage(p QuestionParams) g.Node {
	return c.HTML5(c.HTML5Props{
		Title:       "Arareader",
		Description: "Page to view quiz question",
		Language:    "en",
		Head: []g.Node{
			ar.CSS("/static/main.css"),
			ar.DeferJS("/static/main.js"),
			g.Raw("<script src='https://unpkg.com/htmx.org@1.9.12' integrity='sha384-ujb1lZYygJmzgSwoxRggbCHcjc0rB2XoQrxeTUQyRjrOnlCoYta87iKBWq3EsdM2' crossorigin='anonymous'></script>"),
			ar.Fonts(),
		},
		Body: []g.Node{ID("question-page"), htmx.Boost("true"),
			Nav(Class("nav"),
				ar.Hamburger(), // TODO(Amr Ojjeh): Pull up drawer
				ar.Icon(),
			),
			p.Excerpt,
			P(Class("prompt"),
				g.Text(p.Prompt),
			),
			p.InputMethod,
			g.If(p.Feedback != "",
				Div(Class("callout"),
					Img(Class("callout__icon"), Src("/static/icons/message-solid.svg")),
					P(Class("callout__text"),
						g.Text(p.Feedback),
					),
				),
			),
			questionCtrl(p.PrevURL, p.NextURL, p.SubmitURL),
		},
	})
}

func questionCtrl(prevURL, nextURL, submitURL string) g.Node {
	inner := g.Group([]g.Node{
		prev(prevURL),
		Button(g.If(submitURL == "", Type("button")), g.If(submitURL != "", Type("submit")), c.Classes{
			"btn":           true,
			"btn--disabled": submitURL == "",
			"btn--primary":  submitURL != "",
		},
			g.Text("Submit"),
		),
		next(nextURL),
	})

	if submitURL == "" {
		return Div(Class("question-ctrl"), inner)
	}
	return Form(Class("question-ctrl"), Method("post"), Action(submitURL),
		Input(Type("hidden"), DataAttr("form-answer", "")),
		inner)
}

func prev(prevURL string) g.Node {
	if prevURL == "" {
		return A(Class("question-ctrl__first btn btn--disabled"), Disabled(),
			g.Text("Previous"),
		)
	}
	return A(Class("question-ctrl__first btn btn--secondary"), Href(prevURL), g.Text("Previous"))
}

func next(nextURL string) g.Node {
	if nextURL == "" {
		return A(Class("question-ctrl__last btn btn--disabled"), Disabled(),
			g.Text("Next"),
		)
	}
	return A(Class("question-ctrl__last btn btn--secondary"), Href(nextURL),
		g.Text("Next"),
	)
}
