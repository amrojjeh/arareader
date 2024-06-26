package page

import (
	"fmt"

	g "github.com/maragudk/gomponents"
	htmx "github.com/maragudk/gomponents-htmx"
	c "github.com/maragudk/gomponents/components"
	. "github.com/maragudk/gomponents/html"
)

type SummaryParams struct {
	SidebarQuestions []SidebarQuestion
	Progress         int
	RestartURL       string
}

func SummaryPage(p SummaryParams) g.Node {
	return base([]g.Node{Class("flex"), htmx.Boost("true"),
		Sidebar(false, p.SidebarQuestions, "#"),
		Main(Class("h-svh flex flex-col flex-1"),
			navbar(),
			Div(Class("flex flex-col items-center text-2xl font-bold mb-3"),
				g.Text(fmt.Sprintf("%d%% ", p.Progress)),
				g.If(p.Progress == 100, g.Text("\U0001F973")),
				Div(Class("border w-4/5 h-7 rounded-lg"),
					Div(
						c.Classes{
							"bg-green-500 h-full rounded-lg": true,
							"w-0":                            p.Progress <= 0,
						},
						Style(fmt.Sprintf("width:%d%%", p.Progress)),
					),
				),
			),
			Button(Class("bg-red-500 font-bold text-xl drop-shadow w-min p-2 text-white rounded mx-auto"), htmx.Delete(p.RestartURL), htmx.Target("closest main"), htmx.Confirm("Restarting will delete all your answers. Are you sure?"),
				g.Text("Restart"),
			),
		),
	})
}
