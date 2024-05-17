/*
Copyright © 2024 Amr Ojjeh <amrojjeh@outlook.com>
*/

package model

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"
	"text/template"

	"github.com/amrojjeh/arareader/arabic"
)

type ExcerptNodes interface {
	Write(enc *xml.Encoder) error
	Plain() string
	Unpoint()
}

type Excerpt struct {
	Nodes []ExcerptNodes
}

type ReferenceNotFoundError struct {
	ID int
}

func (r ReferenceNotFoundError) Error() string {
	return fmt.Sprintf("ReferenceNotFound: could not find reference (id: %d)", r.ID)
}

// TODO(Amr Ojjeh): Support colors
func ExcerptFromXML(r io.Reader) (*Excerpt, error) {
	excerpt := &Excerpt{}
	decoder := xml.NewDecoder(r)
	root, err := decoder.Token()
	if err != nil {
		if errors.Is(err, io.EOF) {
			return excerpt, nil
		}
		return &Excerpt{}, nil
	}

	rootStart, _ := root.(xml.StartElement)
	if rootStart.Name.Local != "excerpt" {
		return &Excerpt{}, fmt.Errorf("xml document must begin with <excerpt> (found: %s)", rootStart.Name.Local)
	}

	if err = excerptEl(excerpt, decoder); err != nil {
		return &Excerpt{}, err
	}
	return excerpt, nil
}

func excerptEl(e *Excerpt, decoder *xml.Decoder) error {
	// root tag is already parsed
	for {
		token, err := decoder.Token()
		if err != nil {
			return err
		}
		switch token.(type) {
		case xml.StartElement:
			start, _ := token.(xml.StartElement)
			r, err := refStartEl(start)
			if err != nil {
				return err
			}
			refEl(r, decoder)
			e.Nodes = append(e.Nodes, r)
		case xml.CharData:
			data, _ := token.(xml.CharData)
			textNode := &TextNode{string(data)}
			e.Nodes = append(e.Nodes, textNode)
		case xml.EndElement:
			return nil
		}
	}
}

func refStartEl(start xml.StartElement) (r *ReferenceNode, err error) {
	if start.Name.Local != "ref" {
		return r, fmt.Errorf("element not recognized (el: %s)", start.Name.Local)
	}
	r = &ReferenceNode{ID: -1}
	for _, attr := range start.Attr {
		if attr.Name.Local == "id" {
			r.ID, err = strconv.Atoi(attr.Value)
			if err != nil {
				return r, errors.Join(errors.New("could not convert id to integer"), err)
			}
			break
		}
	}
	if r.ID == -1 {
		return r, errors.New("ref must have an id")
	}
	if r.ID == 0 {
		return r, errors.New("ref id must be greater than 0")
	}

	return r, nil
}

func refEl(r *ReferenceNode, decoder *xml.Decoder) error {
	// root tag is already parsed
	for {
		token, err := decoder.Token()
		if err != nil {
			return err
		}

		switch token.(type) {
		case xml.StartElement:
			start, _ := token.(xml.StartElement)
			subr, err := refStartEl(start)
			if err != nil {
				return err
			}
			refEl(subr, decoder)
			r.Nodes = append(r.Nodes, subr)
		case xml.CharData:
			data, _ := token.(xml.CharData)
			textNode := &TextNode{string(data)}
			r.Nodes = append(r.Nodes, textNode)
		case xml.EndElement:
			return nil
		}
	}
}

func (e *Excerpt) Write(w io.Writer) error {
	enc := xml.NewEncoder(w)
	start, end := e.tags()
	if err := enc.EncodeToken(start); err != nil {
		return err
	}

	for _, n := range e.Nodes {
		if err := n.Write(enc); err != nil {
			return err
		}
	}
	if err := enc.EncodeToken(end); err != nil {
		return err
	}
	return enc.Flush()
}

func (e *Excerpt) tags() (xml.StartElement, xml.EndElement) {
	return xml.StartElement{
			Name: xml.Name{
				Local: "excerpt",
			},
		}, xml.EndElement{
			Name: xml.Name{Local: "excerpt"},
		}
}

func (e *Excerpt) Plain() string {
	builder := strings.Builder{}
	for _, n := range e.Nodes {
		builder.WriteString(n.Plain())
	}
	return builder.String()
}

func (e *Excerpt) Ref(targetRef int) *ReferenceNode {
	for _, n := range e.Nodes {
		ref, ok := n.(*ReferenceNode)
		if !ok {
			continue
		}
		if target := ref.Ref(targetRef); target != nil {
			return target
		}
	}
	return nil
}

func (e *Excerpt) UnpointRef(targetRef int) error {
	r := e.Ref(targetRef)
	if r == nil {
		return ReferenceNotFoundError{targetRef}
	}
	r.Unpoint()
	return nil
}

func (e *Excerpt) UnpointRefs(refs []int) error {
	for _, r := range refs {
		if err := e.UnpointRef(r); err != nil {
			return err
		}
	}
	return nil
}

type ReferenceNode struct {
	ID    int
	Nodes []ExcerptNodes
}

func (r *ReferenceNode) Write(enc *xml.Encoder) error {
	start, end := r.tags()
	if err := enc.EncodeToken(start); err != nil {
		return err
	}

	for _, n := range r.Nodes {
		if err := n.Write(enc); err != nil {
			return err
		}
	}

	if err := enc.EncodeToken(end); err != nil {
		return err
	}
	return nil
}

func (r *ReferenceNode) tags() (xml.StartElement, xml.EndElement) {
	return xml.StartElement{
			Name: xml.Name{
				Local: "ref",
			},
			Attr: []xml.Attr{
				{
					Name: xml.Name{
						Local: "id",
					},
					Value: strconv.Itoa(r.ID),
				},
			},
		}, xml.EndElement{
			Name: xml.Name{Local: "ref"},
		}
}

func (r *ReferenceNode) Plain() string {
	builder := strings.Builder{}
	for _, n := range r.Nodes {
		builder.WriteString(n.Plain())
	}
	return builder.String()
}

func (r *ReferenceNode) Unpoint() {
	for _, n := range r.Nodes {
		n.Unpoint()
	}
}

func (r *ReferenceNode) Ref(targetRef int) *ReferenceNode {
	if r.ID == targetRef {
		return r
	}
	for _, n := range r.Nodes {
		ref, ok := n.(*ReferenceNode)
		if !ok {
			continue
		}
		if target := ref.Ref(targetRef); target != nil {
			return target
		}
	}
	return nil
}

func (r *ReferenceNode) Replace(other *ReferenceNode) {
	*r = *other
}

func (r *ReferenceNode) ReplaceWithText(text string) {
	node := &TextNode{
		Text: text,
	}
	r.Nodes = []ExcerptNodes{node}
}

func (r *ReferenceNode) IsLetterSegmented() bool {
	_, err := arabic.ParseLetterPack(r.Plain())
	return err == nil
}

type TextNode struct {
	Text string
}

func (c *TextNode) Write(enc *xml.Encoder) error {
	return enc.EncodeToken(xml.CharData(c.Text))
}

func (c *TextNode) Plain() string {
	return c.Text
}

func (c *TextNode) Unpoint() {
	c.Text = arabic.Unpointed(c.Text)
}

func ExcerptTemplate() *template.Template {
	return template.New("excerpt parser").Funcs(template.FuncMap{"bw": arabic.FromBuckwalter})
}
