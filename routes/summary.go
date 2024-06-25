package routes

import (
	"fmt"
	"net/http"

	"github.com/amrojjeh/arareader/model"
	"github.com/amrojjeh/arareader/ui/page"
)

func (rs rootResource) summaryGet(w http.ResponseWriter, r *http.Request) {
	var quizID, studentID int
	var found bool
	if studentID, found = studentIDFromRequest(r); !found {
		panic("getting student id")
	}
	if quizID, found = quizIDFromRequest(r); !found {
		panic("getting quiz id")
	}

	q := model.New(rs.db)

	quizSession, err := q.GetQuizSession(r.Context(), model.GetQuizSessionParams{
		StudentID: studentID,
		QuizID:    quizID,
	})
	if err != nil {
		// TODO(Amr Ojjeh): Just create the quiz session
		http.Redirect(w, r, questionURLUnsafe(quizID, 0), http.StatusSeeOther)
		return
	}

	questionSessions, err := q.ListQuestionSessions(r.Context(), quizSession.ID)
	if err != nil {
		err = fmt.Errorf("retrieving quiz sessions: %v", err)
		panic(err)
	}

	questions, err := q.ListQuestionsByQuiz(r.Context(), quizID)
	if err != nil {
		err = fmt.Errorf("retrieving questions: %v", err)
		panic(err)
	}

	page.SummaryPage(page.SummaryParams{
		SidebarQuestions: sidebar(false, quizID, questionSessions, questions),
	}).Render(w)
}
