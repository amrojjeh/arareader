package page

import (
	g "github.com/maragudk/gomponents"
	htmx "github.com/maragudk/gomponents-htmx"
	. "github.com/maragudk/gomponents/html"
)

type QuestionPageParams struct {
	QuestionParams
	Title            string
	Excerpt          g.Node
	SummaryURL       string
	SidebarQuestions []SidebarQuestion
}

type QuestionParams struct {
	Prompt      string
	InputMethod g.Node
	NextURL     string
	PrevURL     string
	SubmitURL   string
	Feedback    string
	InputError  string
}

func QuestionPage(p QuestionPageParams) g.Node {
	return base([]g.Node{Class("flex"), htmx.Boost("true"),
		Sidebar(false, p.SidebarQuestions, p.SummaryURL),
		Main(Class("h-svh flex flex-col"),
			navbar(p.Title),
			p.Excerpt,
			Question(p.QuestionParams),
		),
	})
}

func Question(p QuestionParams) g.Node {
	return Div(ID("target"), Class("flex flex-col flex-1"), htmx.Swap("show:none"), htmx.Target("#target"),
		P(Class("pl-2 text-lg"),
			g.Text(p.Prompt),
		),
		Div(Class("pl-2"),
			p.InputMethod,
			P(Class("text-red-500"),
				g.Text(p.InputError),
			),
		),
		g.If(p.Feedback != "",
			Div(Class("bg-green-300 min-h-32 m-4 p-1 drop-shadow"),
				Img(Class("w-6"), Src("/static/icons/message-solid.svg")),
				P(Class("mt-2"),
					g.Text(p.Feedback),
				),
			),
		),
		QuestionCtrl(false, p.PrevURL, p.NextURL, p.SubmitURL),
	)
}

func QuestionCtrl(oob bool, prevURL, nextURL, submitURL string) g.Node {
	inner := g.Group([]g.Node{ID("questionctrl"), g.If(oob, htmx.SwapOOB("true")),
		prev(prevURL),
		submitBtn(submitURL),
		next(nextURL),
	})

	class := Class("h-16 text-lg sticky bottom-0 left-0 flex flex-row mt-auto pt-2")

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
		return Button(Class("flex-1 rounded-tl-lg btn"), Data("type", "disabled"), Disabled(),
			g.Text("Previous"),
		)
	}
	return Button(Class("flex-1 rounded-tl-lg btn bg-blue-200 text-black"), htmx.Get(prevURL),
		Data("shortcut", "ctrl ArrowLeft"),
		g.Text("Previous"),
	)
}

func next(nextURL string) g.Node {
	if nextURL == "" {
		return Button(Class("flex-1 rounded-tr btn"), Data("type", "disabled"), Disabled(),
			g.Text("Next"),
		)
	}
	return Button(Class("flex-1 rounded-tr btn bg-blue-200 text-black"), htmx.Get(nextURL),
		Data("shortcut", "ctrl ArrowRight"),
		g.Text("Next"),
	)
}
