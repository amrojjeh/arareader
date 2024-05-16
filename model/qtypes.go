/*
Copyright © 2024 Amr Ojjeh <amrojjeh@outlook.com>
*/

package model

import (
	"encoding/json"
)

type QuestionData struct {
	Reference int
	Feedback  string
	Prompt    string
	Answer    string
}

type QuestionType string

const (
	VowelQuestionType          = QuestionType("Vowel")
	ShortAnswerQuestionType    = QuestionType("ShortAnswer")
	LongAnswerQuestionType     = QuestionType("LongAnswer")
	MultipleChoiceQuestionType = QuestionType("MultipleChoice")
	NumberQuestionType         = QuestionType("Number")
	PGNQuestionType            = QuestionType("PGN")
)

type QuizStatus string

const (
	UnsubmittedQuizStatus = QuizStatus("Unsubmitted")
	SubmittedQuizStatus   = QuizStatus("Submitted")
)

type QuestionStatus string

const (
	CorrectQuestionStatus     = QuestionStatus("Correct")
	IncorrectQuestionStatus   = QuestionStatus("Incorrect")
	PendingQuestionStatus     = QuestionStatus("Pending")
	UnsubmittedQuestionStatus = QuestionStatus("Unsubmitted")
	UnattemptedQuestionStatus = QuestionStatus("Unattempted")
)

func NewVowelQuestionData(ref *ReferenceNode, feedback string) *QuestionData {
	if ref.ID == 0 {
		panic("ref ID cannot be zero")
	}
	if !ref.IsLetterSegmented() {
		panic("reference must be letter segmented")
	}
	return &QuestionData{
		Reference: ref.ID,
		Feedback:  feedback,
		Prompt:    "Choose the correct vowel",
		Answer:    ref.Plain(),
	}
}

func NewSegmentedShortAnswerQuestionData(ref *ReferenceNode, feedback, prompt string) *QuestionData {
	if ref.ID == 0 {
		panic("ref ID cannot be zero")
	}
	return &QuestionData{
		Reference: ref.ID,
		Feedback:  feedback,
		Prompt:    prompt,
	}
}

func NewWholeShortAnswerQuestionData(feedback, prompt string) *QuestionData {
	return &QuestionData{
		Reference: 0,
		Feedback:  feedback,
		Prompt:    prompt,
	}
}

func ParseQuestionData(q Question) (QuestionData, error) {
	var data QuestionData
	err := json.Unmarshal(q.Data, &data)
	return data, err
}

func MustParseQuestionData(q Question) QuestionData {
	data, err := ParseQuestionData(q)
	if err != nil {
		panic(err)
	}
	return data
}

func ValidateQuestion(question QuestionData, answer string) bool {
	return question.Answer == answer
}
