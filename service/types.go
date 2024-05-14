package service

type QuestionType string

const (
	VowelQuestionType          = QuestionType("Vowel")
	ShortAnswerQuestionType    = QuestionType("ShortAnswer")
	LongAnswerQuestionType     = QuestionType("LongAnswer")
	MultipleChoiceQuestionType = QuestionType("MultipleChoice")
	NumberQuestionType         = QuestionType("Number")
	PGNQuestionType            = QuestionType("PGN")
)

type QuestionData interface {
	RefID() int
}

type VowelQuestionData struct {
	Reference int // should never be 0. TODO(Amr Ojjeh): Enforce that somehow?
	Feedback  string
}

func (vd VowelQuestionData) RefID() int {
	return vd.Reference
}

type ShortAnswerQuestionData struct {
	Reference int // 0 = whole
	Feedback  string
	Prompt    string
}

func (sd ShortAnswerQuestionData) RefID() int {
	return sd.Reference
}

// TODO(Amr Ojjeh): Move these to models and modify sqlc.yaml to convert to these types
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
