package routes

import (
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/amrojjeh/arareader/service"
)

func TestHTTPRouteServeHTTP(t *testing.T) {
	tests := []struct {
		name string
		path string
		code int
		body []byte
	}{
		{
			name: "Root",
			path: "/",
			code: http.StatusOK,
			body: []byte("Submit"),
		},
		{
			name: "Static",
			path: "/static/main.css",
			code: http.StatusOK,
			body: []byte("@layer"),
		},
		{
			name: "Does not exist",
			path: "/doesnotexist",
			code: http.StatusNotFound,
			body: []byte("404"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			writer := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, tt.path, bytes.NewReader([]byte{}))
			handler := NewRootHandler(service.DemoDB(context.Background()))
			handler.ServeHTTP(writer, req)
			if writer.Code != tt.code {
				t.Errorf("incorrect status code (expected: %d; actual: %d)", tt.code, writer.Code)
			}
			if !bytes.Contains(writer.Body.Bytes(), tt.body) {
				t.Errorf("incorrect page (path: %s)", tt.path)
			}
		})
	}
}
