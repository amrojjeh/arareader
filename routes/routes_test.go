/*
Copyright © 2024 Amr Ojjeh <amrojjeh@outlook.com>
*/

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
			code: http.StatusNotFound,
			body: []byte("404"),
		},
		{
			name: "Static",
			path: "/static/main.css",
			code: http.StatusOK,
			body: []byte("body"),
		},
		{
			name: "Does not exist",
			path: "/doesnotexist",
			code: http.StatusNotFound,
			body: []byte("404"),
		},
		{
			name: "Question",
			path: "/quiz/1/question/0",
			code: http.StatusOK,
			body: []byte("html"),
		},
		{
			name: "Negative question",
			path: "/quiz/1/question/-1",
			code: http.StatusNotFound,
		},
		{
			name: "Too many questions",
			path: "/quiz/1/question/4",
			code: http.StatusBadRequest,
		},
		{
			name: "Quiz does not exist",
			path: "/quiz/0/question/2",
			code: http.StatusNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			writer := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, tt.path, bytes.NewReader([]byte{}))
			handler := Routes(service.DemoDB(context.Background()))
			handler.ServeHTTP(writer, req)
			if writer.Code != tt.code {
				t.Errorf("incorrect status code (expected: %d; actual: %d)", tt.code, writer.Code)
			}
			if !bytes.Contains(writer.Body.Bytes(), tt.body) {
				t.Errorf("incorrect page")
			}
		})
	}
}
