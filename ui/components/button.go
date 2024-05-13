/*
Copyright © 2024 Amr Ojjeh <amrojjeh@outlook.com>
*/

package components

import (
	g "github.com/maragudk/gomponents"
	h "github.com/maragudk/gomponents/html"
)

func VowelButton(text string) g.Node {
	return button(h.Type("button"), h.DataAttr("substitute-button", text),
		g.Text(text),
	)
}

func SubmitButton(text string) g.Node {
	return button(h.Type("submit"),
		g.Text(text),
	)
}

func button(children ...g.Node) g.Node {
	return h.Button(h.Class("button"),
		g.Group(children),
	)
}
