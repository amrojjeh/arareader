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

type VowelQuestionData struct {
	Reference int
	Feedback  string
}

type ShortAnswerQuestionData struct {
	Reference int // 0 = whole
	Feedback  string
	Prompt    string
}
