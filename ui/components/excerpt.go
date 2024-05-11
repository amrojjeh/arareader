package components

import (
	g "github.com/maragudk/gomponents"
	. "github.com/maragudk/gomponents/html"
)

// FIXME(Amr Ojjeh): Change to service.SegmentedQuestion or something like that
func Excerpt(excerpt string) g.Node {
	return P(Class("excerpt"))
}
