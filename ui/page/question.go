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
			g.Raw("<script src='https://unpkg.com/htmx.org@2.0.0'></script>"),
			g.Raw("<script src='https://unpkg.com/idiomorph@0.3.0'></script>"),
			ar.Fonts(),
		},
		Body: []g.Node{Class("flex flex-col"), htmx.Ext("morph"), htmx.Boost("true"),
			Nav(Class("p-2 mb-2 flex justify-between items-center"),
				ar.Hamburger(), // TODO(Amr Ojjeh): Pull up drawer
				ar.Icon(),
			),
			p.Excerpt,
			Div(ID("target"), Class("pl-2"), htmx.Select("#target"), htmx.Target("#target"), htmx.Swap("show:none"),
				P(Class("text-lg"),
					g.Text(p.Prompt),
				),
				p.InputMethod,
				g.If(p.Feedback != "",
					Div(Class("bg-green-300 min-h-32 mr-4 my-4 p-1 drop-shadow"),
						Img(Class("w-6"), Src("/static/icons/message-solid.svg")),
						P(Class("mt-2"),
							g.Text(p.Feedback),
						),
					),
				),
				questionCtrl(p.PrevURL, p.NextURL, p.SubmitURL),
			),
		},
	})
}

func questionCtrl(prevURL, nextURL, submitURL string) g.Node {
	inner := g.Group([]g.Node{
		prev(prevURL),
		submitBtn(submitURL),
		next(nextURL),
	})

	if submitURL == "" {
		return Div(Class("h-16 text-lg fixed w-screen bottom-0 left-0 flex flex-row"), inner)
	}
	return Form(Class("h-16 text-lg fixed w-screen bottom-0 left-0 flex flex-row"), Method("post"), Action(submitURL),
		Input(Type("hidden"), Name("ans"), DataAttr("form-answer", "")),
		inner)
}

func submitBtn(submitURL string) g.Node {
	return Button(Class("flex-1 btn"), ID("submit"),
		g.If(submitURL == "",
			g.Group([]g.Node{
				Type("button"),
				Data("type", "disabled"),
				Disabled(),
			}),
		),
		g.If(submitURL != "",
			g.Group([]g.Node{
				Type("submit"),
				Data("type", "primary"),
			}),
		),
		g.Text("Submit"),
	)
}

func prev(prevURL string) g.Node {
	if prevURL == "" {
		return A(Class("flex-1 rounded-tl-lg btn"), Data("type", "disabled"), Disabled(),
			g.Text("Previous"),
		)
	}
	return A(Class("flex-1 rounded-tl-lg btn bg-blue-200 text-black"), Href(prevURL), g.Text("Previous"))
}

func next(nextURL string) g.Node {
	if nextURL == "" {
		return A(Class("flex-1 rounded-tr btn"), Data("type", "disabled"), Disabled(),
			g.Text("Next"),
		)
	}
	return A(Class("flex-1 rounded-tr btn bg-blue-200 text-black"), Href(nextURL),
		g.Text("Next"),
	)
}
