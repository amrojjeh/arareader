/*
Copyright © 2024 Amr Ojjeh <amrojjeh@outlook.com>
*/

package components

import (
	g "github.com/maragudk/gomponents"
	. "github.com/maragudk/gomponents/html"
)

func Hamburger() g.Node {
	return Img(Class("cursor-pointer"), Width("20px"), Src("/static/icons/bars-solid.svg"))
}
