package service

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestShiftPath(t *testing.T) {
	tests := []struct {
		path         string
		headExpected string
		tailExpected string
	}{
		{
			path:         "/static/img/image.png",
			headExpected: "static",
			tailExpected: "/img/image.png",
		},
		{
			path:         "/img/image.png",
			headExpected: "img",
			tailExpected: "/image.png",
		},
		{
			path:         "/image.png",
			headExpected: "image.png",
			tailExpected: "/",
		},
		{
			path:         "/",
			headExpected: ".",
			tailExpected: "/",
		},
	}

	for _, tt := range tests {
		head, tail := shiftPath(tt.path)
		if head != tt.headExpected {
			t.Errorf("head was not matched (actual: '%s'; expected: '%s')", head, tt.headExpected)
		}
		if tail != tt.tailExpected {
			t.Errorf("tail was not matched (actual: '%s'; expected: '%s')", tail, tt.tailExpected)
		}
	}

	// testing panic
	defer func() { _ = recover() }()
	panicPath := "image.png"
	shiftPath(panicPath)
	t.Errorf("shiftPath did not panic (path: '%s')", panicPath)
}

func TestShiftURL(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "/static/img/image.png", bytes.NewReader([]byte{}))
	head := shiftURL(r)
	expectedHead, expectedTail := "static", "/img/image.png"
	if head != expectedHead {
		t.Errorf("incorrect head (expected: %s; actual: %s)", expectedHead, head)
	}
	if r.URL.Path != expectedTail {
		t.Errorf("incorrect r.URL.Path (expected: %s; actual: %s)", expectedTail, r.URL.Path)
	}
	if r.URL.RawPath != expectedTail {
		t.Errorf("incorrect r.URL.RawPath (expected: %s; actual: %s)", expectedTail, r.URL.RawPath)
	}
}
