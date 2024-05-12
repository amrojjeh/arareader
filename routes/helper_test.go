package routes

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestShiftPath(t *testing.T) {
	tests := []struct {
		name         string
		path         string
		headExpected string
		tailExpected string
	}{
		{
			name:         "Shift 0",
			path:         "/static/img/image.png",
			headExpected: "static",
			tailExpected: "/img/image.png",
		},
		{
			name:         "Shift 1",
			path:         "/img/image.png",
			headExpected: "img",
			tailExpected: "/image.png",
		},
		{
			name:         "Shift 2",
			path:         "/image.png",
			headExpected: "image.png",
			tailExpected: "/",
		},
		{
			name:         "Shift 3",
			path:         "/",
			headExpected: ".",
			tailExpected: "/",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head, tail := shiftPath(tt.path)
			if head != tt.headExpected {
				t.Errorf("head was not matched (actual: '%s'; expected: '%s')", head, tt.headExpected)
			}
			if tail != tt.tailExpected {
				t.Errorf("tail was not matched (actual: '%s'; expected: '%s')", tail, tt.tailExpected)
			}
		})
	}

	// testing panic
	t.Run("Relative path", func(t *testing.T) {
		defer func() { _ = recover() }()
		panicPath := "image.png"
		shiftPath(panicPath)
		t.Errorf("shiftPath did not panic (path: '%s')", panicPath)
	})
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
