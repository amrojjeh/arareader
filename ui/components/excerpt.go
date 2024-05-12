package components

import (
	"encoding/xml"
	"errors"
	"io"

	g "github.com/maragudk/gomponents"
	. "github.com/maragudk/gomponents/html"
)

// Assumes excerpt is valid
func Excerpt(r io.Reader, highlightedRef string) g.Node {
	decoder := xml.NewDecoder(r)
	node := excerptEl(decoder, highlightedRef)
	return node
}

func excerptEl(decoder *xml.Decoder, highlightedRef string) g.Node {
	// parsing the starting element <excerpt>
	token, err := decoder.Token()
	if err != nil {
		if errors.Is(err, io.EOF) {
			return P()
		}
	}

	nodes := []g.Node{}

	for {
		token, _ = decoder.Token()
		switch token.(type) {
		case xml.StartElement:
			start, _ := token.(xml.StartElement)
			nodes = append(nodes, refEl(decoder, start.Attr[0].Value, highlightedRef))
		case xml.EndElement:
			return P(Class("excerpt"), g.Group(nodes))
		case xml.CharData:
			data, _ := token.(xml.CharData)
			nodes = append(nodes, g.Text(string(data)))
		}
	}
}

func refEl(decoder *xml.Decoder, id, highlightedRef string) g.Node {
	// already parsed the starting element
	nodes := []g.Node{}
	for {
		token, _ := decoder.Token()
		switch token.(type) {
		case xml.StartElement:
			start, _ := token.(xml.StartElement)
			nodes = append(nodes, refEl(decoder, start.Attr[0].Value, highlightedRef))
		case xml.EndElement:
			if id == highlightedRef {
				return Span(Class("highlight"), g.Group(nodes))
			}
			return Span(g.Group(nodes))
		case xml.CharData:
			data, _ := token.(xml.CharData)
			nodes = append(nodes, g.Text(string(data)))
		}
	}
}
