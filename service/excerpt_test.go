package service

import (
	"bytes"
	"errors"
	"reflect"
	"testing"
	"text/template"

	"github.com/amrojjeh/arareader/arabic"
	"github.com/amrojjeh/arareader/must"
)

func TestExcerptPlain(t *testing.T) {
	tests := []struct {
		name     string
		excerpt  string
		expected string
	}{
		{
			name:     "Basic",
			excerpt:  `<excerpt>{{bw "h*A baytN"}}</excerpt>`,
			expected: arabic.FromBuckwalter("h*A baytN"),
		},
		{
			name:     "With references",
			excerpt:  `<excerpt>{{bw "h*A"}} {{bw "bay"}}<ref id="1">{{bw "tN"}}</ref> {{bw "wh*A"}} <ref id="2">{{bw "$y'"}}</ref></excerpt>`,
			expected: arabic.FromBuckwalter("h*A baytN wh*A $y'"),
		},
		{
			name:     "With nested references",
			excerpt:  `<excerpt>{{bw "h*A"}} {{bw "bay"}}<ref id="1">{{bw "tN"}}</ref> {{bw "wh*A"}} <ref id="2">{{bw "$y"}}<ref id="3">{{bw "'"}}</ref></ref></excerpt>`,
			expected: arabic.FromBuckwalter("h*A baytN wh*A $y'"),
		},
		{
			name:     "Empty Excerpt",
			excerpt:  "<excerpt></excerpt>",
			expected: "",
		},
		{
			name:     "Empty",
			excerpt:  "",
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := parseExcerpt(t, tt.excerpt)
			plain := e.Plain()
			if plain != tt.expected {
				t.Errorf("expected: '%s'; actual: '%s'", tt.expected, plain)
			}
		})
	}
}

func TestExcerptUnpointRef(t *testing.T) {
	tests := []struct {
		name     string
		excerpt  string
		ref      int
		expected string
	}{
		{
			name:     "Basic",
			excerpt:  `<excerpt>{{bw "h*A"}} {{bw "bay"}}<ref id="1">{{bw "tN"}}</ref> {{bw "wh*A"}} <ref id="2">{{bw "$y'"}}</ref></excerpt>`,
			ref:      1,
			expected: arabic.FromBuckwalter("h*A bayt wh*A $y'"),
		},
		{
			name:     "Nested references",
			excerpt:  `<excerpt>{{bw "h*A"}} {{bw "bay"}}<ref id="1">{{bw "tN"}}</ref> {{bw "wh*A"}} <ref id="2">{{bw "$y"}}<ref id="3">{{bw "'N"}}</ref></ref></excerpt>`,
			ref:      3,
			expected: arabic.FromBuckwalter("h*A baytN wh*A $y'"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := parseExcerpt(t, tt.excerpt)
			err := e.UnpointRef(tt.ref)
			if err != nil {
				t.Error(err)
			}
			plain := e.Plain()
			if plain != tt.expected {
				t.Errorf("expected: '%s'; actual: '%s'", tt.expected, plain)
			}
		})
	}

	// testing error
	t.Run("Reference not found", func(t *testing.T) {
		e := parseExcerpt(t, `<excerpt>{{bw "h*A"}} {{bw "bay"}}<ref id="1">{{bw "tN"}}</ref> {{bw "wh*A"}} <ref id="2">{{bw "$y'"}}</ref></excerpt>`)
		err := e.UnpointRef(25)
		var concreteErr ReferenceNotFoundError
		if !errors.As(err, &concreteErr) {
			t.Errorf("expected ReferenceNotFoundError; got %v", reflect.TypeOf(err))
		}
		if concreteErr.ID != 25 {
			t.Errorf("expected err with id 25, found %d", concreteErr.ID)
		}
	})
}

func TestExcerptWrite(t *testing.T) {
	tests := []struct {
		name    string
		excerpt string
	}{
		{
			name:    "Basic",
			excerpt: `<excerpt>{{bw "h*A baytN"}}</excerpt>`,
		},
		{
			name:    "With references",
			excerpt: `<excerpt>{{bw "h*A"}} {{bw "bay"}}<ref id="1">{{bw "tN"}}</ref> {{bw "wh*A"}} <ref id="2">{{bw "$y'"}}</ref></excerpt>`,
		},
		{
			name:    "With nested references",
			excerpt: `<excerpt>{{bw "h*A"}} {{bw "bay"}}<ref id="1">{{bw "tN"}}</ref> {{bw "wh*A"}} <ref id="2">{{bw "$y"}}<ref id="3">{{bw "'"}}</ref></ref></excerpt>`,
		},
		{
			name:    "Empty Excerpt",
			excerpt: "<excerpt></excerpt>",
		},
		{
			name:    "Empty",
			excerpt: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			expected := &bytes.Buffer{}
			template.Must(excerptTemplate().Parse(tt.excerpt)).Execute(expected, nil)
			e := must.Get(ExcerptFromXML(expected))
			actual := &bytes.Buffer{}
			e.Write(actual)
			if actual.String() != expected.String() {
				t.Errorf("expected: '%s'; actual: '%s'", expected.String(), actual.String())
			}
		})
	}
}

func parseExcerpt(t *testing.T, excerpt string) *Excerpt {
	t.Helper()
	buff := &bytes.Buffer{}
	template.Must(excerptTemplate().Parse(excerpt)).Execute(buff, nil)
	return must.Get(ExcerptFromXML(buff))
}
