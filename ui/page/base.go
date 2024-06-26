package page

import (
	"fmt"

	"github.com/amrojjeh/arareader/model"
	g "github.com/maragudk/gomponents"
	htmx "github.com/maragudk/gomponents-htmx"
	c "github.com/maragudk/gomponents/components"
	. "github.com/maragudk/gomponents/html"
)

func base(body []g.Node) g.Node {
	return c.HTML5(c.HTML5Props{
		Title:       "Arareader",
		Description: "Page to view quiz question",
		Language:    "en",
		Head: []g.Node{
			css("/static/main.css"),
			deferJS("/static/main.js"),
			g.Raw("<script src='https://unpkg.com/htmx.org@2.0.0/dist/htmx.js' integrity='sha384-Xh+GLLi0SMFPwtHQjT72aPG19QvKB8grnyRbYBNIdHWc2NkCrz65jlU7YrzO6qRp' crossorigin='anonymous'></script>"),
			fonts(),
		},
		Body: body,
	})
}

func navbar(title string) g.Node {
	return Nav(Class("p-2 mb-2 flex justify-between items-center"),
		Img(Class("lg:invisible cursor-pointer"), Width("20px"), Src("/static/icons/bars-solid.svg"),
			Data("sidebar-toggle", ""),
		),
		P(Class("basis-1/2 overflow-x-auto font-medium text-xl text-nowrap overflow-hidden"),
			g.Text(title),
		),
		icon(),
	)
}

type SidebarQuestion struct {
	Prompt   string
	Selected bool
	Status   model.QuestionStatus
	URL      string
	Target   bool
}

func Sidebar(oob bool, qs []SidebarQuestion, summaryURL string) g.Node {
	return Dialog(Data("sidebar", ""), Class("sidebar lg:shrink-0 lg:static lg:block"), g.If(oob, htmx.SwapOOB("true")), ID("sidebar"),
		Div(Class("flex flex-col justify-center p-5 gap-4"),
			Div(Class("flex gap-3 mb-2"),
				Img(Class("w-6 cursor-pointer lg:hidden"), Data("sidebar-toggle", ""), Src("/static/icons/xmark-solid.svg")),
				A(Class("flex-1 bg-blue-200 text-center p-1 cursor-pointer rounded-lg"), Href(summaryURL),
					g.Text("Summary"),
				),
			),
			g.Group(g.Map(qs, func(s SidebarQuestion) g.Node {
				return Div(Class("flex gap-2 text-lg"), g.If(s.Target, htmx.Target("#target")), g.If(!s.Target, htmx.Target("closest body")),
					g.If(s.Status == model.CorrectQuestionStatus,
						Img(Class("w-5"), Src("/static/icons/circle-check-solid.svg")),
					),
					g.If(s.Status == model.IncorrectQuestionStatus,
						Img(Class("w-5"), Src("/static/icons/circle-xmark-solid.svg")),
					),
					g.If(s.Status == model.PendingQuestionStatus,
						Img(Class("w-5"), Src("/static/icons/circle-question-solid.svg")),
					),

					g.If(s.Selected, P(Class("truncate font-bold"), Data("selected-question", "true"),
						g.Text(s.Prompt),
					)),
					g.If(!s.Selected, Button(Class("truncate underline cursor"), htmx.Get(s.URL), g.If(!s.Target, htmx.PushURL("true")),
						g.Text(s.Prompt),
					)),
				)
			})),
		),
	)
}

// TODO(Amr Ojjeh): Change icon
func icon() g.Node {
	return Div(Class("arareader-icon"),
		g.Text("AR"),
	)
}

func css(path string) g.Node {
	return Link(Rel("stylesheet"), Href(path))
}

func deferJS(path string) g.Node {
	return Script(Src(path), Defer())
}

func roboto() string {
	return "family=Roboto:ital,wght@0,100;0,300;0,400;0,500;0,700;0,900;1,100;1,300;1,400;1,500;1,700;1,900&"
}

func notoNaskh() string {
	return "family=Noto+Naskh+Arabic:wght@400..700&"
}

func permanentMarker() string {
	return "family=Permanent+Marker&"
}

func fonts() g.Node {
	fonts := fmt.Sprintf("https://fonts.googleapis.com/css2?%vdisplay=swap", notoNaskh()+roboto()+permanentMarker())
	return g.Group([]g.Node{
		Link(Rel("preconnect"), Href("https://fonts.googleapis.com")),
		Link(Rel("preconnect"), Href("https://fonts.gstatic.com"), g.Attr("crossorigin")),
		Link(Href(fonts), Rel("stylesheet")),
	})
}
