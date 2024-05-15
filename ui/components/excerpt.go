package components

import (
	"fmt"
	"reflect"

	"github.com/amrojjeh/arareader/model"
	g "github.com/maragudk/gomponents"
	. "github.com/maragudk/gomponents/html"
)

// Assumes excerpt is valid
func Excerpt(e *model.Excerpt, highlightedRef int) g.Node {
	children := processNodes(e.Nodes, highlightedRef)
	return P(Class("excerpt"),
		g.Group(children),
	)
}

func refNode(r *model.ReferenceNode, highlightedRef int) g.Node {
	children := processNodes(r.Nodes, highlightedRef)
	if r.ID == highlightedRef {
		return Span(Class("highlight"), DataAttr("selected-segment", ""),
			g.Group(children),
		)
	} else {
		return Span(
			g.Group(children),
		)
	}
}

func processNodes(nodes []model.ExcerptNodes, highlightedRef int) []g.Node {
	acc := []g.Node{}
	for _, n := range nodes {
		switch n.(type) {
		case *model.ReferenceNode:
			acc = append(acc, refNode(n.(*model.ReferenceNode), highlightedRef))
		case *model.TextNode:
			text := n.(*model.TextNode).Text
			acc = append(acc, g.Text(text))
		default:
			panic(fmt.Sprintf("unexpected element of type %v", reflect.TypeOf(n)))
		}
	}
	return acc
}
