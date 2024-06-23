/*
Copyright © 2024 Amr Ojjeh <amrojjeh@outlook.com>
*/

package components

import (
	"fmt"

	"github.com/amrojjeh/arareader/model"
	g "github.com/maragudk/gomponents"
	htmx "github.com/maragudk/gomponents-htmx"
	. "github.com/maragudk/gomponents/html"
)

// Assumes excerpt is valid
func Excerpt(e *model.Excerpt, highlightedRef int) g.Node {
	children := processNodes(e.Nodes, highlightedRef)
	return P(Class("max-h-96 mb-3 leading-relaxed text-3xl text-justify px-5 bg-gray-100 overflow-auto"), ID("excerpt"), Dir("rtl"), htmx.SwapOOB("true"),
		g.Group(children),
	)
}

func refNode(r *model.ReferenceNode, highlightedRef int) g.Node {
	children := processNodes(r.Nodes, highlightedRef)
	if r.ID == highlightedRef {
		return Span(Class("bg-opacity-50 bg-blue-400"), Data("selected-segment", ""),
			g.Group(children),
		)
	} else {
		return Span(
			g.Group(children),
		)
	}
}

func processNodes(nodes []model.ExcerptNode, highlightedRef int) []g.Node {
	acc := []g.Node{}
	for _, n := range nodes {
		switch n.(type) {
		case *model.ReferenceNode:
			acc = append(acc, refNode(n.(*model.ReferenceNode), highlightedRef))
		case *model.TextNode:
			text := n.(*model.TextNode).Text
			acc = append(acc, g.Text(text))
		default:
			panic(fmt.Sprintf("unexpected element of type %T", n))
		}
	}
	return acc
}
