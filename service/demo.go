package service

import (
	"bytes"
	"context"
	"database/sql"
	"errors"
	"text/template"

	"github.com/amrojjeh/arareader/arabic"
	"github.com/amrojjeh/arareader/model"
	"github.com/amrojjeh/arareader/must"
)

var excerpt = `<excerpt><ref id="1">{{bw "h*A by"}}<ref id="2">{{bw "tN"}}</ref></ref></excerpt>`

func DemoDB(ctx context.Context) *sql.DB {
	db := OpenDB(":memory:")
	Setup(ctx, db)
	q := model.New(db)
	teacher := must.Get(q.CreateTeacher(ctx, model.CreateTeacherParams{
		Email:        "smith@demo.com",
		Username:     "Professor Smith",
		PasswordHash: mustFromPlainPassword("password"),
	}))

	tmpl := template.Must(template.New("excerpt parser").Funcs(template.FuncMap{"bw": arabic.FromBuckwalter}).
		Parse(excerpt))
	buff := &bytes.Buffer{}
	tmpl.Execute(buff, nil)

	must.Get(q.CreateQuiz(ctx, model.CreateQuizParams{
		TeacherID: teacher.ID,
		Title:     "Quiz 1",
		Excerpt:   buff.Bytes(),
	}))

	return db
}

func mustFromPlainPassword(pass string) string {
	hash, err := FromPlainPassword(pass)
	if err != nil {
		panic(errors.Join(errors.New("hash could not be generated"), err))
	}
	return hash
}
