package routes

import (
	"net/http"

	"github.com/amrojjeh/arareader/ui/page"
)

func (rs rootResource) summaryGet(w http.ResponseWriter, r *http.Request) {
	studentID, found := studentIDFromRequest(r)
	if !found {
		panic("getting student id")
	}

	quiz := rs.quiz(r)
	quizSession := rs.createQuizSession(r, studentID, quiz.ID)
	questionSessions := rs.questionSessions(r, quizSession.ID)
	questions := rs.questions(r, quiz.ID)

	page.SummaryPage(page.SummaryParams{
		SidebarQuestions: sidebar(false, quiz.ID, questionSessions, questions),
	}).Render(w)
}
