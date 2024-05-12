package service

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"text/template"

	"github.com/amrojjeh/arareader/arabic"
	"github.com/amrojjeh/arareader/model"
	"github.com/amrojjeh/arareader/must"
)

// TODO(Amr Ojjeh): Add a second longer excerpt
var excerpt = `<excerpt><ref id="1">{{bw "h*A by"}}<ref id="2">{{bw "tN"}}</ref></ref></excerpt>`

func DemoDB(ctx context.Context) *sql.DB {
	db := OpenDB(":memory:")
	Setup(ctx, db)
	q := model.New(db)
	teacher := must.Get(q.CreateTeacher(ctx, model.CreateTeacherParams{
		Email:        "smith@demo.com",
		Username:     "Professor Smith",
		PasswordHash: must.Get(FromPlainPassword("password")),
	}))

	tmpl := template.Must(template.New("excerpt parser").Funcs(template.FuncMap{"bw": arabic.FromBuckwalter}).
		Parse(excerpt))
	buff := &bytes.Buffer{}
	tmpl.Execute(buff, nil)

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
			Reference: 2,
			Feedback:  "There's a damma because it's a predicate'",
		})),
	})

	q.CreateQuestion(ctx, model.CreateQuestionParams{
		QuizID:   quiz.ID,
		Position: 1,
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

	must.Get(q.CreateStudent(ctx, model.CreateStudentParams{
		Name:    "Bob",
		ClassID: class.ID,
	}))

	return db
}
