package service

import (
	"encoding/xml"
	"errors"
	"io"
	"strings"
	"text/template"

	"github.com/amrojjeh/arareader/arabic"
)

// TODO(Amr Ojjeh): Write a IsValid function
// - cannot contain non-Arabic characters
// - the root must be excerpt
// - elements must either be excerpt or ref
// - all refs have exactly one attribute, which is the id
// - the id is a number

func Plain(r io.Reader) (string, error) {
	builder := strings.Builder{}
	decoder := xml.NewDecoder(r)
	for {
		token, err := decoder.Token()
		if err != nil {
			if errors.Is(err, io.EOF) {
				return builder.String(), nil
			}
			return "", err
		}
		data, ok := token.(xml.CharData)
		if ok {
			builder.Write(data)
		}
	}
}

func excerptTemplate() *template.Template {
	return template.New("excerpt parser").Funcs(template.FuncMap{"bw": arabic.FromBuckwalter})
}
