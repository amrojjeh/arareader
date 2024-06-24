/*
Copyright © 2024 Amr Ojjeh <amrojjeh@outlook.com>
*/

package page

import (
	"github.com/amrojjeh/arareader/model"
	ar "github.com/amrojjeh/arareader/ui/components"
	g "github.com/maragudk/gomponents"
	htmx "github.com/maragudk/gomponents-htmx"
	c "github.com/maragudk/gomponents/components"
	. "github.com/maragudk/gomponents/html"
)

type SidebarQuestion struct {
	Prompt string
	Status model.QuestionStatus
	URL    string
}

type QuestionParams struct {
	Excerpt          g.Node
	Prompt           string
	InputMethod      g.Node
	SidebarQuestions []SidebarQuestion
	NextURL          string
	PrevURL          string
	SubmitURL        string
	Feedback         string
}

func QuestionPage(p QuestionParams) g.Node {
	return c.HTML5(c.HTML5Props{
		Title:       "Arareader",
		Description: "Page to view quiz question",
		Language:    "en",
		Head: []g.Node{
			ar.CSS("/static/main.css"),
			ar.DeferJS("/static/main.js"),
			g.Raw("<script src='https://unpkg.com/htmx.org@2.0.0/dist/htmx.js' integrity='sha384-Xh+GLLi0SMFPwtHQjT72aPG19QvKB8grnyRbYBNIdHWc2NkCrz65jlU7YrzO6qRp' crossorigin='anonymous'></script>"),
			g.Raw("<script src='https://unpkg.com/idiomorph@0.3.0'></script>"),
			ar.Fonts(),
		},
		Body: []g.Node{Class("flex flex-col"), htmx.Ext("morph"), htmx.Boost("true"),
			Dialog(Data("sidebar", ""), Class("sidebar"), htmx.SwapOOB("true"), ID("sidebar"),
				Div(Class("flex flex-col justify-center p-5 gap-4"),
					Div(Class("flex gap-3 mb-2"),
						Img(Class("w-6 cursor-pointer"), Data("sidebar-toggle", ""), Src("/static/icons/xmark-solid.svg")),
						Button(Class("flex-1 bg-blue-200 p-1 cursor-pointer rounded-lg"), htmx.Post("/restart"), htmx.Confirm("Are you sure?"), htmx.Target("#target"), htmx.Select("#target"),
							g.Text("Restart"), // TEMP(Amr Ojjeh): Should be summary
						),
					),
					g.Group(g.Map(p.SidebarQuestions, func(s SidebarQuestion) g.Node {
						return Div(Class("flex gap-2 text-lg"), htmx.Select("#target"), htmx.Target("#target"),
							g.If(s.Status == model.CorrectQuestionStatus,
								Img(Class("w-5"), Src("/static/icons/circle-check-solid.svg")),
							),
							g.If(s.Status == model.IncorrectQuestionStatus,
								Img(Class("w-5"), Src("/static/icons/circle-xmark-solid.svg")),
							),
							g.If(s.Status == model.PendingQuestionStatus,
								Img(Class("w-5"), Src("/static/icons/circle-question-solid.svg")),
							),
							A(Class("underline truncate"), Href(s.URL),
								g.Text(s.Prompt),
							),
						)
					})),
				),
			),
			Main(Class("h-svh flex flex-col"),
				Nav(Class("p-2 mb-2 flex justify-between items-center "),
					Img(Class("cursor-pointer"), Width("20px"), Src("/static/icons/bars-solid.svg"),
						Data("sidebar-toggle", ""),
					),
					ar.Icon(),
				),
				p.Excerpt,
				Div(ID("target"), Class("flex flex-col flex-1"), htmx.Select("#target"), htmx.Target("#target"), htmx.Swap("show:none"),
					P(Class("pl-2 text-lg"),
						g.Text(p.Prompt),
					),
					Div(Class("pl-2"),
						p.InputMethod,
					),
					g.If(p.Feedback != "",
						Div(Class("bg-green-300 min-h-32 mr-4 my-4 pl-2 p-1 drop-shadow"),
							Img(Class("w-6"), Src("/static/icons/message-solid.svg")),
							P(Class("mt-2"),
								g.Text(p.Feedback),
							),
						),
					),
					questionCtrl(p.PrevURL, p.NextURL, p.SubmitURL),
				),
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

	class := Class("h-16 text-lg sticky w-screen bottom-0 left-0 flex flex-row mt-auto")

	if submitURL == "" {
		return Div(class,
			inner)
	}
	return Form(class, Method("post"), Action(submitURL),
		Input(Type("hidden"), Name("ans"), DataAttr("form-answer", "")),
		inner)
}

func submitBtn(submitURL string) g.Node {
	return Button(Class("flex-1 btn"), ID("submit"), Data("shortcut", "Enter"),
		g.If(submitURL == "",
			Type("button"),
		),
		g.If(submitURL != "",
			Type("submit"),
		),
		Data("type", "disabled"),
		Disabled(),
		g.Text("Submit"),
	)
}

func prev(prevURL string) g.Node {
	if prevURL == "" {
		return A(Class("flex-1 rounded-tl-lg btn"), Data("type", "disabled"), Disabled(),
			g.Text("Previous"),
		)
	}
	return A(Class("flex-1 rounded-tl-lg btn bg-blue-200 text-black"), Href(prevURL),
		Data("shortcut", "ctrl ArrowLeft"),
		g.Text("Previous"),
	)
}

func next(nextURL string) g.Node {
	if nextURL == "" {
		return A(Class("flex-1 rounded-tr btn"), Data("type", "disabled"), Disabled(),
			g.Text("Next"),
		)
	}
	return A(Class("flex-1 rounded-tr btn bg-blue-200 text-black"), Href(nextURL),
		Data("shortcut", "ctrl ArrowRight"),
		g.Text("Next"),
	)
}
