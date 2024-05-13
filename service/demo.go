package service

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"text/template"

	"github.com/amrojjeh/arareader/model"
	"github.com/amrojjeh/arareader/must"
)

func DemoDB(ctx context.Context) *sql.DB {
	var excerpt = `<excerpt>{{bw "<nmA Al>EmA"}}<ref id="1">{{bw "lu"}}</ref> {{bw "bAlnyA"}}<ref id="2">{{bw "ti"}}</ref>، {{bw "w<nmA lkl AmrY' mA nwY fmn kAnt hjrth <lY"}} <ref id="3">{{bw "Allh wrswlh fhjrth <lY Allh wrswlh، wmn kAnt hjrth ldnyA ySybhA، >w Amr>p ynkHhA fhjrt"}}<ref id="4">{{bw "hu"}}</ref> {{bw "<lY mA hAjr <lyh"}}</ref></excerpt>`

	db := OpenDB(":memory:")
	Setup(ctx, db)
	q := model.New(db)
	teacher := must.Get(q.CreateTeacher(ctx, model.CreateTeacherParams{
		Email:        "smith@demo.com",
		Username:     "Professor Smith",
		PasswordHash: must.Get(FromPlainPassword("password")),
	}))

	buff := &bytes.Buffer{}
	template.Must(excerptTemplate().Parse(excerpt)).Execute(buff, nil)

	quiz := must.Get(q.CreateQuiz(ctx, model.CreateQuizParams{
		TeacherID: teacher.ID,
		Title:     "Quiz 1",
		Excerpt:   buff.Bytes(),
	}))

	q.CreateQuestion(ctx, model.CreateQuestionParams{
		QuizID:   quiz.ID,
		Position: 0,
		Type:     string(VowelQuestionType),
		Data: must.Get(json.Marshal(VowelQuestionData{
			Reference: 1,
			Feedback:  "There's a damma because it's a raf'",
		})),
	})

	q.CreateQuestion(ctx, model.CreateQuestionParams{
		QuizID:   quiz.ID,
		Position: 1,
		Type:     string(VowelQuestionType),
		Data: must.Get(json.Marshal(VowelQuestionData{
			Reference: 2,
			Feedback:  "There's a kasra because it's a jarr'",
		})),
	})

	q.CreateQuestion(ctx, model.CreateQuestionParams{
		QuizID:   quiz.ID,
		Position: 2,
		Type:     string(VowelQuestionType),
		Data: must.Get(json.Marshal(VowelQuestionData{
			Reference: 4,
			Feedback:  "There's a damma because it's a raf'",
		})),
	})

	q.CreateQuestion(ctx, model.CreateQuestionParams{
		QuizID:   quiz.ID,
		Position: 4,
		Type:     string(ShortAnswerQuestionType),
		Data: must.Get(json.Marshal(ShortAnswerQuestionData{
			Reference: 1,
			Feedback:  "Possible translation: this is a house",
			Prompt:    "Translate the sentence",
		})),
	})

	class := must.Get(q.CreateClass(ctx, model.CreateClassParams{
		TeacherID: teacher.ID,
		Name:      "Class of 2024",
	}))

	q.AddQuizToClass(ctx, model.AddQuizToClassParams{
		QuizID:  quiz.ID,
		ClassID: class.ID,
	})

	s := must.Get(q.CreateStudent(ctx, model.CreateStudentParams{
		Name:    "Bob",
		ClassID: class.ID,
	}))

	quizSession := must.Get(q.CreateStudentQuizSession(ctx, model.CreateStudentQuizSessionParams{
		StudentID: s.ID,
		QuizID:    quiz.ID,
		Status:    string(UnsubmittedQuizStatus),
	}))

	questions := must.Get(q.ListQuestionsByQuiz(ctx, quiz.ID))
	for range questions {
		q.CreateStudentQuestionSession(ctx, model.CreateStudentQuestionSessionParams{
			StudentQuizSessionID: quizSession.ID,
			Status:               string(UnattemptedQuestionStatus),
		})
	}

	return db
}
