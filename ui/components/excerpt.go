package components

import (
	"fmt"
	"reflect"

	"github.com/amrojjeh/arareader/service"
	g "github.com/maragudk/gomponents"
	. "github.com/maragudk/gomponents/html"
)

// Assumes excerpt is valid
func Excerpt(e *service.Excerpt, highlightedRef int) g.Node {
	children := processNodes(e.Nodes, highlightedRef)
	return P(Class("excerpt"),
		g.Group(children),
	)
}

func refNode(r *service.ReferenceNode, highlightedRef int) g.Node {
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

func processNodes(nodes []service.ExcerptNodes, highlightedRef int) []g.Node {
	acc := []g.Node{}
	for _, n := range nodes {
		switch n.(type) {
		case *service.ReferenceNode:
			acc = append(acc, refNode(n.(*service.ReferenceNode), highlightedRef))
		case *service.TextNode:
			text := n.(*service.TextNode).Text
			acc = append(acc, g.Text(text))
		default:
			panic(fmt.Sprintf("unexpected element of type %v", reflect.TypeOf(n)))
		}
	}
	return acc
}
