package page

import (
	g "github.com/maragudk/gomponents"
	htmx "github.com/maragudk/gomponents-htmx"
	. "github.com/maragudk/gomponents/html"
)

type SummaryParams struct {
	SidebarQuestions []SidebarQuestion
}

func SummaryPage(p SummaryParams) g.Node {
	return base([]g.Node{Class("flex flex-col"), htmx.Boost("true"),
		sidebar(p.SidebarQuestions),
		Main(Class("h-svh flex flex-col"),
			navbar(),
		),
	})
}
