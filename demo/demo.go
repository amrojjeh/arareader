/*
Copyright © 2024 Amr Ojjeh <amrojjeh@outlook.com>
*/

package demo

import (
	"bytes"
	"context"
	"database/sql"
	"text/template"

	"github.com/amrojjeh/arareader/model"
	"github.com/amrojjeh/arareader/must"
)

func Demo(ctx context.Context, db *sql.DB) {
	var excerptXML = `<excerpt>{{bw "<nmA Al>EmA"}}<ref id="1">{{bw "lu"}}</ref> {{bw "bAlnyA"}}<ref id="2">{{bw "ti"}}</ref>، {{bw "w<nmA lkl AmrY' mA nwY fmn kAnt hjrth <lY"}} <ref id="3">{{bw "Allh wrswlh fhjrth <lY Allh wrswlh، wmn kAnt hjrth ldnyA ySybhA، >w Amr>p ynkHhA fhjrt"}}<ref id="4">{{bw "hu"}}</ref> {{bw "<lY mA hAjr <lyh"}}</ref></excerpt>`

	q := model.New(db)
	teacher := must.Get(q.CreateTeacher(ctx, model.CreateTeacherParams{
		Email:        "smith@demo.com",
		Username:     "Professor Smith",
		PasswordHash: must.Get(model.FromPlainPassword("password")),
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
		QuizID:    quiz.ID,
		Position:  0,
		Type:      model.VowelQuestionType,
		Reference: 1,
		Feedback:  "There's a damma because it's a raf'",
		Prompt:    "Choose the correct vowel",
		Solution:  excerpt.Ref(1).Plain(),
	})

	q.CreateQuestion(ctx, model.CreateQuestionParams{
		QuizID:    quiz.ID,
		Position:  1,
		Type:      model.VowelQuestionType,
		Reference: 2,
		Feedback:  "There's a kasra because it's a jarr'",
		Prompt:    "Choose the correct vowel",
		Solution:  excerpt.Ref(2).Plain(),
	})

	q.CreateQuestion(ctx, model.CreateQuestionParams{
		QuizID:    quiz.ID,
		Position:  2,
		Type:      model.VowelQuestionType,
		Reference: 4,
		Feedback:  "There's a damma because it's a raf'",
		Prompt:    "Choose the correct vowel",
		Solution:  excerpt.Ref(4).Plain(),
	})

	q.CreateQuestion(ctx, model.CreateQuestionParams{
		QuizID:    quiz.ID,
		Position:  3,
		Type:      model.ShortAnswerQuestionType,
		Reference: 3,
		Feedback:  "Possible translation: this is a house",
		Prompt:    "Translate the sentence",
		Solution:  "<MANUAL>",
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

	q.CreateQuizSession(ctx, model.CreateQuizSessionParams{
		StudentID: s.ID,
		QuizID:    quiz.ID,
		Status:    model.UnsubmittedQuizStatus,
	})
}
