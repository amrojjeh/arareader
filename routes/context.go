/*
Copyright © 2024 Amr Ojjeh <amrojjeh@outlook.com>
*/

package routes

import "net/http"

type contextKey string

var (
	studentIDKey     = contextKey("studentID")
	quizIDKey        = contextKey("quizID")
	questionPosKey   = contextKey("questionPos")
	questionErrorKey = contextKey("questionErrorKey")
)

func studentIDFromRequest(r *http.Request) (int, bool) {
	if r.Context().Value(studentIDKey) == nil {
		return 0, false
	}
	return r.Context().Value(studentIDKey).(int), true
}

func questionPositionFromRequest(r *http.Request) (int, bool) {
	if r.Context().Value(questionPosKey) == nil {
		return 0, false
	}
	return r.Context().Value(questionPosKey).(int), true
}

func quizIDFromRequest(r *http.Request) (int, bool) {
	if r.Context().Value(quizIDKey) == nil {
		return 0, false
	}
	return r.Context().Value(quizIDKey).(int), true
}

func questionErrorFromRequest(r *http.Request) (string, bool) {
	if r.Context().Value(questionErrorKey) == nil {
		return "", false
	}
	return r.Context().Value(questionErrorKey).(string), true
}
