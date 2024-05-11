package service

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHTTPRouteServeHTTP(t *testing.T) {
	tests := []struct {
		path string
		code int
		body []byte
	}{
		{
			path: "/",
			code: http.StatusOK,
			body: []byte("Submit"),
		},
		{
			path: "/static/main.css",
			code: http.StatusOK,
			body: []byte("@layer"),
		},
		{
			path: "/doesnotexist",
			code: http.StatusNotFound,
			body: []byte("404"),
		},
	}

	for _, tt := range tests {
		writer := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, tt.path, bytes.NewReader([]byte{}))
		HTTPRoute{}.ServeHTTP(writer, req)
		if writer.Code != tt.code {
			t.Errorf("incorrect status code (path: %s; expected: %d; actual: %d)", tt.path, tt.code, writer.Code)
		}
		if !bytes.Contains(writer.Body.Bytes(), tt.body) {
			t.Errorf("incorrect page (path: %s)", tt.path)
		}
	}
}
