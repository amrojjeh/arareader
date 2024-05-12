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

	assertStartElement(t, decoder, "excerpt")
	assertCharData(t, decoder, arabic.FromBuckwalter("<nmA Al>EmA"))
	start := assertStartElement(t, decoder, "ref")
	assertAttribute(t, start, "id", "1")
	assertCharData(t, decoder, arabic.FromBuckwalter("lu"))
	assertEndElement(t, decoder, "ref")
	// goes on..., no need to test the whole thing

	questions := must.Get(q.ListQuestionsByQuiz(ctx, quiz.ID))

	if questions[0].Type != string(VowelQuestionType) {
		t.Error("first question should be a vowel question")
	}

	if questions[3].Type != string(ShortAnswerQuestionType) {
		t.Error("fourth question should be a short answer")
	}
}

func assertCharData(t *testing.T, decoder *xml.Decoder, expectedValue string) {
	t.Helper()
	token := must.Get(decoder.Token())
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

func assertStartElement(t *testing.T, decoder *xml.Decoder, expectedName string) xml.StartElement {
	t.Helper()
	token := must.Get(decoder.Token())
	start, ok := token.(xml.StartElement)
	if !ok {
		t.Errorf("expected start element, found: %v", reflect.TypeOf(token))
	}
	if start.Name.Local != expectedName {
		t.Errorf("tag name is incorrect (expected: '%s'; actual: '%s')", expectedName, start.Name.Local)
	}
	return start
}

func assertEndElement(t *testing.T, decoder *xml.Decoder, expectedName string) xml.EndElement {
	t.Helper()
	token := must.Get(decoder.Token())
	end, ok := token.(xml.EndElement)
	if !ok {
		t.Errorf("expected end element, found: %v", reflect.TypeOf(token))
	}
	if end.Name.Local != expectedName {
		t.Errorf("tag name is incorrect (expected: '%s'; actual: '%s')", expectedName, end.Name)
	}
	return end
}
