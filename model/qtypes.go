/*
Copyright © 2024 Amr Ojjeh <amrojjeh@outlook.com>
*/

package model

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

func (qs QuestionStatus) IsSubmitted() bool {
	return qs == CorrectQuestionStatus || qs == IncorrectQuestionStatus
}
