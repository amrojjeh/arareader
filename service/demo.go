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
	var excerptXML = `<excerpt>{{bw "<nmA Al>EmA"}}<ref id="1">{{bw "lu"}}</ref> {{bw "bAlnyA"}}<ref id="2">{{bw "ti"}}</ref>، {{bw "w<nmA lkl AmrY' mA nwY fmn kAnt hjrth <lY"}} <ref id="3">{{bw "Allh wrswlh fhjrth <lY Allh wrswlh، wmn kAnt hjrth ldnyA ySybhA، >w Amr>p ynkHhA fhjrt"}}<ref id="4">{{bw "hu"}}</ref> {{bw "<lY mA hAjr <lyh"}}</ref></excerpt>`

	db := MustOpenDB(":memory:")
	MustSetup(ctx, db)
	q := model.New(db)
	teacher := must.Get(q.CreateTeacher(ctx, model.CreateTeacherParams{
		Email:        "smith@demo.com",
		Username:     "Professor Smith",
		PasswordHash: must.Get(FromPlainPassword("password")),
	}))

	buff := &bytes.Buffer{}
	template.Must(model.ExcerptTemplate().Parse(excerptXML)).Execute(buff, nil)
	excerpt, _ := model.ExcerptFromXML(bytes.NewReader(buff.Bytes()))

	quiz := must.Get(q.CreateQuiz(ctx, model.CreateQuizParams{
		TeacherID: teacher.ID,
		Title:     "Quiz 1",
		Excerpt:   buff.Bytes(),
	}))

	q.CreateQuestion(ctx, model.CreateQuestionParams{
		QuizID:   quiz.ID,
		Position: 0,
		Type:     model.VowelQuestionType,
		Data:     must.Get(json.Marshal(model.NewVowelQuestionData(excerpt.Ref(1), "There's a damma because it's a raf'"))),
	})

	q.CreateQuestion(ctx, model.CreateQuestionParams{
		QuizID:   quiz.ID,
		Position: 1,
		Type:     model.VowelQuestionType,
		Data:     must.Get(json.Marshal(model.NewVowelQuestionData(excerpt.Ref(2), "There's a kasra because it's a jarr'"))),
	})

	q.CreateQuestion(ctx, model.CreateQuestionParams{
		QuizID:   quiz.ID,
		Position: 2,
		Type:     model.VowelQuestionType,
		Data:     must.Get(json.Marshal(model.NewVowelQuestionData(excerpt.Ref(4), "There's a damma because it's a raf'"))),
	})

	q.CreateQuestion(ctx, model.CreateQuestionParams{
		QuizID:   quiz.ID,
		Position: 3,
		Type:     model.ShortAnswerQuestionType,
		Data:     must.Get(json.Marshal(model.NewSegmentedShortAnswerQuestionData(excerpt.Ref(4), "Possible translation: this is a house", "Translate the sentence"))),
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

	q.CreateStudentQuizSession(ctx, model.CreateStudentQuizSessionParams{
		StudentID: s.ID,
		QuizID:    quiz.ID,
		Status:    model.UnsubmittedQuizStatus,
	})
	return db
}
