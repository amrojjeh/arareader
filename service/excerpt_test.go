package service

import (
	"bytes"
	"testing"
	"text/template"

	"github.com/amrojjeh/arareader/arabic"
	"github.com/amrojjeh/arareader/must"
)

func TestPlain(t *testing.T) {
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
			buff := &bytes.Buffer{}
			template.Must(excerptTemplate().Parse(tt.excerpt)).Execute(buff, nil)
			plain := must.Get(Plain(buff))
			if plain != tt.expected {
				t.Errorf("expected: '%s'; actual: '%s'", tt.expected, plain)
			}
		})
	}
}
