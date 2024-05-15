package model

import "encoding/json"

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

func NewVowelQuestionData(ref int, feedback string) *QuestionData {
	if ref == 0 {
		panic("Vowel question data cannot be zero")
	}
	return &QuestionData{
		Reference: ref,
		Feedback:  feedback,
		Prompt:    "Choose the correct vowel",
	}
}

func NewShortAnswerQuestionData(ref int, feedback, prompt string) *QuestionData {
	return &QuestionData{
		Reference: ref,
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
