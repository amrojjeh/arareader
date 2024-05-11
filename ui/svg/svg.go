/*
Copyright © 2024 Amr Ojjeh <amrojjeh@outlook.com>
*/

package svg

import (
	g "github.com/maragudk/gomponents"
)

func SVG(children ...g.Node) g.Node {
	return g.El("svg", children...)
}

func Path(children ...g.Node) g.Node {
	return g.El("path", children...)
}

func Polygon(children ...g.Node) g.Node {
	return g.El("polygon", children...)
}

// Attributes

func ViewBox(viewbox string) g.Node {
	return g.Attr("viewBox", viewbox)
}

func Points(points string) g.Node {
	return g.Attr("points", points)
}

func Fill(fill string) g.Node {
	return g.Attr("fill", fill)
}

func Stroke(stroke string) g.Node {
	return g.Attr("stroke", stroke)
}

func StrokeWidth(width string) g.Node {
	return g.Attr("stroke-width", width)
}

func Transform(transform string) g.Node {
	return g.Attr("transform", transform)
}
