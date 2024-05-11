package service

import (
	"bytes"
	"context"
	"encoding/xml"
	"reflect"
	"testing"

	"github.com/amrojjeh/arareader/arabic"
	"github.com/amrojjeh/arareader/model"
	"github.com/amrojjeh/arareader/must"
)

func TestDemoDB(t *testing.T) {
	ctx := context.Background()
	db := DemoDB(ctx)
	q := model.New(db)
	quiz := must.Get(q.GetQuiz(ctx, 1))
	decoder := xml.NewDecoder(bytes.NewReader(quiz.Excerpt))

	token := must.Get(decoder.Token())
	assertStartElement(t, token, "excerpt")

	token = must.Get(decoder.Token())
	assertStartElement(t, token, "ref")
	assertAttribute(t, token.(xml.StartElement), "id", "1")

	token = must.Get(decoder.Token())
	assertCharData(t, token, arabic.FromBuckwalter("h*A by"))

	token = must.Get(decoder.Token())
	assertStartElement(t, token, "ref")
	assertAttribute(t, token.(xml.StartElement), "id", "2")

	token = must.Get(decoder.Token())
	assertCharData(t, token, arabic.FromBuckwalter("tN"))

	token = must.Get(decoder.Token())
	assertEndElement(t, token, "ref")

	token = must.Get(decoder.Token())
	assertEndElement(t, token, "ref")

	token = must.Get(decoder.Token())
	assertEndElement(t, token, "excerpt")
}

func assertCharData(t *testing.T, token xml.Token, expectedValue string) {
	t.Helper()
	charData, ok := token.(xml.CharData)
	if !ok {
		t.Error("expected chardata")
	}

	if string(charData) != expectedValue {
		t.Errorf("incorrect charData (expected: '%s'; actual: '%s')", expectedValue, string(charData))
	}
}

func assertAttribute(t *testing.T, start xml.StartElement, expectedName, expectedValue string) {
	t.Helper()
	for _, attr := range start.Attr {
		if attr.Name.Local == expectedName {
			if attr.Value != expectedValue {
				t.Errorf("attr has wrong value (name: '%s'; expected: '%s'; actual: '%s')",
					attr.Name.Local, expectedValue, attr.Value)
			} else {
				return
			}
		}
	}
	t.Errorf("attr not found (expected: '%s')", expectedName)
}

func assertStartElement(t *testing.T, token xml.Token, expectedName string) {
	t.Helper()
	start, ok := token.(xml.StartElement)
	if !ok {
		t.Errorf("expected start element, found: %v", reflect.TypeOf(token))
	}
	if start.Name.Local != expectedName {
		t.Errorf("tag name is incorrect (expected: '%s'; actual: '%s')", expectedName, start.Name.Local)
	}
}

func assertEndElement(t *testing.T, token xml.Token, expectedName string) {
	t.Helper()
	end, ok := token.(xml.EndElement)
	if !ok {
		t.Errorf("expected end element, found: %v", reflect.TypeOf(token))
	}
	if end.Name.Local != expectedName {
		t.Errorf("tag name is incorrect (expected: '%s'; actual: '%s')", expectedName, end.Name)
	}
}
