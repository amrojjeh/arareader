/*
Copyright © 2024 Amr Ojjeh <amrojjeh@outlook.com>
*/

package components

import (
	"fmt"

	"github.com/amrojjeh/arareader/ui/svg"
	g "github.com/maragudk/gomponents"
	. "github.com/maragudk/gomponents/html"
)

func Hamburger() g.Node {
	return Img(Width("20px"), Src("/static/icons/bars-solid.svg"))
}

type QuestionNavProps struct {
	CurrentQuestion int
	TotalQuestions  int
	SkipForwardURL  string
	SkipBackwardURL string
	NextURL         string
	PrevURL         string
}

func QuestionNav(p QuestionNavProps) g.Node {
	return Div(Class("question-nav"),
		g.If(p.SkipBackwardURL != "",
			A(Class("question-nav__navigate"), Href(p.SkipBackwardURL),
				prevIcon(true),
			),
		),
		g.If(p.PrevURL != "",
			A(Class("question-nav__navigate"), Href(p.PrevURL),
				prevIcon(false),
			),
		),
		P(Class("question-nav__prompt"),
			g.Text(fmt.Sprintf("Question %d out of %d", p.CurrentQuestion, p.TotalQuestions)),
		),
		g.If(p.NextURL != "",
			A(Class("question-nav__navigate"), Href(p.NextURL),
				nextIcon(false),
			),
		),
		g.If(p.SkipForwardURL != "",
			A(Class("question-nav__navigate"), Href(p.SkipForwardURL),
				nextIcon(true),
			),
		),
	)
}

func nextIcon(double bool) g.Node {
	if double {
		return svg.SVG(svg.ViewBox("0 0 100 100"), Height("1rem"), Width("1rem"),
			svg.Polygon(svg.Points("10,0 10,100 50,50"), svg.Fill("none"), svg.Stroke("black"), svg.StrokeWidth("10")),
			svg.Polygon(svg.Points("50,0 50,100 90,50"), svg.Fill("none"), svg.Stroke("black"), svg.StrokeWidth("10")),
		)
	}
	return svg.SVG(svg.ViewBox("0 0 60 100"), Height("1rem"), Width("1rem"),
		svg.Polygon(svg.Points("10,0 10,100 50,50"), svg.Fill("none"), svg.Stroke("black"), svg.StrokeWidth("10")),
	)
}

func prevIcon(double bool) g.Node {
	if double {
		return svg.SVG(svg.ViewBox("0 0 100 100"), Height("1rem"), Width("1rem"), svg.Transform("scale(-1,1)"),
			svg.Polygon(svg.Points("10,0 10,100 50,50"), svg.Fill("none"), svg.Stroke("black"), svg.StrokeWidth("10")),
			svg.Polygon(svg.Points("50,0 50,100 90,50"), svg.Fill("none"), svg.Stroke("black"), svg.StrokeWidth("10")),
		)
	}
	return svg.SVG(svg.ViewBox("0 0 60 100"), Height("1rem"), Width("1rem"), svg.Transform("scale(-1,1)"),
		svg.Polygon(svg.Points("10,0 10,100 50,50"), svg.Fill("none"), svg.Stroke("black"), svg.StrokeWidth("10")),
	)
}
