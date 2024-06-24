/*
Copyright © 2024 Amr Ojjeh <amrojjeh@outlook.com>
*/

package demo

import (
	"bytes"
	"context"
	"database/sql"
	"fmt"
	"strings"
	"text/template"

	"github.com/amrojjeh/arareader/arabic"
	"github.com/amrojjeh/arareader/model"
	"github.com/amrojjeh/arareader/must"
)

func Demo(ctx context.Context, db *sql.DB) {
	excerptBytes, excerpt := fromTemplate(`<excerpt>{{bw "<nmA Al>EmA"}}<ref id="1">{{bw "lu"}}</ref> <ref id="5">{{bw "bAlnyA"}}<ref id="2">{{bw "ti"}}</ref></ref>، {{bw "w<nmA lkl AmrY' mA nwY fmn kAnt hjrth <lY Allh wrswlh fhjrth <lY Allh wrswlh، wmn kAnt hjrth ldnyA ySybhA، >w Amr>p ynkHhA"}} <ref id="3">{{bw "fhjrt"}}<ref id="4">{{bw "hu"}}</ref> {{bw "<lY mA hAjr <lyh"}}</ref></excerpt>`)

	q := model.New(db)
	teacher := must.Get(q.CreateTeacher(ctx, model.CreateTeacherParams{
		Email:        "smith@demo.com",
		Username:     "Professor Smith",
		PasswordHash: must.Get(model.FromPlainPassword("password")),
	}))

	quiz := must.Get(q.CreateQuiz(ctx, model.CreateQuizParams{
		TeacherID: teacher.ID,
		Title:     "Quiz 1",
		Excerpt:   excerptBytes,
	}))

	q.CreateQuestion(ctx, model.CreateQuestionParams{
		QuizID:    quiz.ID,
		Reference: 1,
		Position:  0,
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

	must.Get(q.CreateStudent(ctx, model.CreateStudentParams{
		Name:    "Bob",
		ClassID: class.ID,
	}))

	_, tempExcerpt := fromTemplate(`<excerpt>{{bw "Amna"}} <ref id="1">{{bw "Alrswlu"}}</ref> {{bw "bmA"}} <ref id="2">{{bw ">unzila"}}</ref> {{bw "Alyhi mn rbhi wAlm&mnwna klN Amna bAllhi wmlA}kthi wktbhi wrslhi lA nfrqu byna AHdK mno rslhi wqAlwA smEnA wATEnA gfrAnka rbnA wAlyka AlmSyru lA yklfu Allhu nfsA AlA wsEhA lhA mA ksbto wElyhA mA Aktsbto rbnA lA t&Ax*nA An nsynA Aw AxTAnA rbnA wlA tHmlo ElynA ASrA kmA Hmlthu ElY Al*yna mn qblnA rbnA wlA tHmlnA mA lA TAqpa lnA bhi wAEfu EnA wAgfr lnA wArHmnA Ant mwlAnA fAnSrnA ElY Alqwmi AlkAfryna"}}</excerpt>`)

	excerpt = genVowelRefs(tempExcerpt)

	excerptBuffer := &bytes.Buffer{}
	excerpt.Write(excerptBuffer)

	quiz = must.Get(q.CreateQuiz(ctx, model.CreateQuizParams{
		TeacherID: teacher.ID,
		Title:     "Surah al-Baqarah",
		Excerpt:   excerptBuffer.Bytes(),
	}))

	for _, n := range excerpt.Nodes {
		r, ok := n.(*model.ReferenceNode)
		if !ok {
			continue
		}
		genVowelQuestions(ctx, q, quiz.ID, r)
	}

	q.CreateQuestion(ctx, model.CreateQuestionParams{
		QuizID:    quiz.ID,
		Position:  0,
		Type:      model.ShortAnswerQuestionType,
		Reference: 1,
		Feedback:  "It's definite because of the ال at the beginning.",
		Prompt:    "Definite or indefinite?",
		Solution:  "definite",
	})

	q.CreateQuestion(ctx, model.CreateQuestionParams{
		QuizID:    quiz.ID,
		Position:  1,
		Type:      model.ShortAnswerQuestionType,
		Reference: 2,
		Feedback:  "The main giveaway is that it's in the form " + arabic.FromBuckwalter(">ufEila") + ".",
		Prompt:    "Active or passive?",
		Solution:  "passive",
	})
}

func genVowelQuestions(ctx context.Context, q *model.Queries, quizID int, r *model.ReferenceNode) {
	if lp, err := arabic.ParseLetterPack(r.Plain()); err == nil && lp.Vowel != 0 {
		q.CreateQuestion(ctx, model.CreateQuestionParams{
			QuizID:    quizID,
			Position:  0,
			Type:      model.VowelQuestionType,
			Reference: r.ID,
			Prompt:    "Choose the correct vowel",
			Solution:  r.Plain(),
		})
		return
	}
	for _, n := range r.Nodes {
		r, ok := n.(*model.ReferenceNode)
		if !ok {
			continue
		}
		genVowelQuestions(ctx, q, quizID, r)
	}
}

// TEMP(Amr Ojjeh): This is basically just a prototype
// Does not handle the case where a short vowel reference already exists
func genVowelRefs(e *model.ReferenceNode) *model.ReferenceNode {
	excerpt := &model.ReferenceNode{}
	refCounter := e.AvailableID() // the next available reference
	for _, n := range e.Nodes {
		switch typed := n.(type) {
		case *model.ReferenceNode:
			var newRef *model.ReferenceNode
			refCounter, newRef = genVowelRefsRef(typed, refCounter)
			excerpt.Nodes = append(excerpt.Nodes, newRef)
		case *model.TextNode:
			var nodes []model.ExcerptNode
			refCounter, nodes = genVowelRefsText(typed, refCounter)
			excerpt.Nodes = append(excerpt.Nodes, nodes...)
		}
	}
	return excerpt
}

func genVowelRefsRef(r *model.ReferenceNode, counter int) (int, *model.ReferenceNode) {
	ref := &model.ReferenceNode{
		ID:    r.ID,
		Nodes: []model.ExcerptNode{},
	}
	for _, n := range r.Nodes {
		switch typed := n.(type) {
		case *model.ReferenceNode:
			var newRef *model.ReferenceNode
			counter, newRef = genVowelRefsRef(typed, counter)
			ref.Nodes = append(ref.Nodes, newRef)
		case *model.TextNode:
			var nodes []model.ExcerptNode
			counter, nodes = genVowelRefsText(typed, counter)
			ref.Nodes = append(ref.Nodes, nodes...)
		}
	}
	return counter, ref
}

func genVowelRefsText(t *model.TextNode, counter int) (int, []model.ExcerptNode) {
	nodes := []model.ExcerptNode{}
	words := strings.Split(t.Text, " ")
	for i, w := range words {
		lps := arabic.LetterPacks(w)
		if lps[len(lps)-1].Vowel != 0 {
			nodes = append(nodes, &model.TextNode{
				Text: arabic.LetterPacksToString(lps[:len(lps)-1]),
			})
			nodes = append(nodes, &model.ReferenceNode{
				ID: counter,
				Nodes: []model.ExcerptNode{&model.TextNode{
					Text: lps[len(lps)-1].String(),
				}},
			})
			if i != len(words)-1 {
				nodes = append(nodes, &model.TextNode{
					Text: " ",
				})
			}

			counter++
			continue
		}
		if i != len(words)-1 {
			nodes = append(nodes, &model.TextNode{
				Text: w + " ",
			})
		} else {
			nodes = append(nodes, &model.TextNode{
				Text: w,
			})
		}
	}
	return counter, nodes
}
func fromTemplate(s string) ([]byte, *model.ReferenceNode) {

	buff := &bytes.Buffer{}
	template.Must(model.ExcerptTemplate().Parse(s)).Execute(buff, nil)
	excerpt, _ := model.ExcerptFromXML(bytes.NewReader(buff.Bytes()))
	return buff.Bytes(), excerpt
}
