/*
Copyright © 2024 Amr Ojjeh <amrojjeh@outlook.com>
*/

package demo

import (
	"bytes"
	"context"
	"database/sql"
	"fmt"
	"text/template"

	"github.com/amrojjeh/arareader/arabic"
	"github.com/amrojjeh/arareader/model"
	"github.com/amrojjeh/arareader/must"
)

func Demo(ctx context.Context, db *sql.DB) {
	var excerptXML = `<excerpt>{{bw "<nmA Al>EmA"}}<ref id="1">{{bw "lu"}}</ref> <ref id="5">{{bw "bAlnyA"}}<ref id="2">{{bw "ti"}}</ref></ref>، {{bw "w<nmA lkl AmrY' mA nwY fmn kAnt hjrth <lY Allh wrswlh fhjrth <lY Allh wrswlh، wmn kAnt hjrth ldnyA ySybhA، >w Amr>p ynkHhA"}} <ref id="3">{{bw "fhjrt"}}<ref id="4">{{bw "hu"}}</ref> {{bw "<lY mA hAjr <lyh"}}</ref></excerpt>`

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
		Reference: 1,
		Type:      model.VowelQuestionType,
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
		Feedback:  "",
		Prompt:    "Translate the sentence",
		Solution:  "", // manually graded if empty
	})

	q.CreateQuestion(ctx, model.CreateQuestionParams{
		QuizID:    quiz.ID,
		Position:  3,
		Type:      model.ShortAnswerQuestionType,
		Reference: 5,
		Feedback:  "The word highlighted is the plural version, hence why it's rendered as 'surely actions are by their intentions.'",
		Prompt:    fmt.Sprintf("What is the meaning of the word %s?", arabic.FromBuckwalter("nyp")),
		Solution:  "intention",
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
