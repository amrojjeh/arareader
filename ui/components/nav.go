/*
Copyright © 2024 Amr Ojjeh <amrojjeh@outlook.com>
*/

package components

import (
	"github.com/amrojjeh/arareader/ui/svg"
	g "github.com/maragudk/gomponents"
	h "github.com/maragudk/gomponents/html"
)

func QuestionNav(children ...g.Node) g.Node {
	return h.Div(h.Class("question-nav"),
		h.Button(h.Type("button"), h.Class("question-nav--navigate"),
			prevIcon(true),
		),
		h.Button(h.Type("button"), h.Class("question-nav--navigate"),
			prevIcon(false),
		),
		h.P(h.Class("question-nav--prompt"),
			g.Text("Question 1 out of 20"),
		),
		h.Button(h.Type("button"), h.Class("question-nav--navigate"),
			nextIcon(false),
		),
		h.Button(h.Type("button"), h.Class("question-nav--navigate"),
			nextIcon(true),
		),
	)
}

func nextIcon(double bool) g.Node {
	if double {
		return svg.SVG(svg.ViewBox("0 0 100 100"), h.Height("1rem"), h.Width("1rem"),
			svg.Polygon(svg.Points("10,0 10,100 50,50"), svg.Fill("none"), svg.Stroke("black"), svg.StrokeWidth("10")),
			svg.Polygon(svg.Points("50,0 50,100 90,50"), svg.Fill("none"), svg.Stroke("black"), svg.StrokeWidth("10")),
		)
	}
	return svg.SVG(svg.ViewBox("0 0 60 100"), h.Height("1rem"), h.Width("1rem"),
		svg.Polygon(svg.Points("10,0 10,100 50,50"), svg.Fill("none"), svg.Stroke("black"), svg.StrokeWidth("10")),
	)
}

func prevIcon(double bool) g.Node {
	if double {
		return svg.SVG(svg.ViewBox("0 0 100 100"), h.Height("1rem"), h.Width("1rem"), svg.Transform("scale(-1,1)"),
			svg.Polygon(svg.Points("10,0 10,100 50,50"), svg.Fill("none"), svg.Stroke("black"), svg.StrokeWidth("10")),
			svg.Polygon(svg.Points("50,0 50,100 90,50"), svg.Fill("none"), svg.Stroke("black"), svg.StrokeWidth("10")),
		)
	}
	return svg.SVG(svg.ViewBox("0 0 60 100"), h.Height("1rem"), h.Width("1rem"), svg.Transform("scale(-1,1)"),
		svg.Polygon(svg.Points("10,0 10,100 50,50"), svg.Fill("none"), svg.Stroke("black"), svg.StrokeWidth("10")),
	)
}
