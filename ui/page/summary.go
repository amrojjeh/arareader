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
}

func SummaryPage(p SummaryParams) g.Node {
	return base([]g.Node{Class("flex flex-col"), htmx.Boost("true"),
		Sidebar(false, p.SidebarQuestions, "#"),
		Main(Class("h-svh flex flex-col"),
			navbar(),
			Div(Class("flex flex-col items-center text-2xl font-bold"),
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
		),
	})
}
