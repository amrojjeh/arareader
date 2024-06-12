/*
Copyright © 2024 Amr Ojjeh <amrojjeh@outlook.com>
*/

package routes

// func (rh rootHandler) mustParseForm(r *http.Request) {
// 	err := r.ParseForm()
// 	if err != nil {
// 		panic(clientError(http.StatusBadRequest))
// 	}
// }

// func (rh rootHandler) fetchQuizSession(r *http.Request, quizID, studentID int) (model.QuizSession, bool) {
// 	quizSession, err := rh.queries.GetQuizSession(r.Context(), model.GetQuizSessionParams{
// 		StudentID: studentID,
// 		QuizID:    quizID,
// 	})
// 	if err != nil {
// 		if errors.Is(err, sql.ErrNoRows) {
// 			return model.QuizSession{}, false
// 		}
// 		err = fmt.Errorf("routes: could not retrieve quiz session. %w", err)
// 		panic(err)
// 	}
// 	return quizSession, true
// }

// func (qh quizHandler) applyVowelQuestions() {
// 	vowelQuestions := service.Filter(isVowelQuestionType, qh.questions)
// 	questionData := service.Map(model.MustParseQuestionData, vowelQuestions)
// 	refs := service.Map(excerptRef, questionData)
// 	qh.excerpt.UnpointRefs(refs)
// }

// func (qh quizHandler) applyVowelAnswers(r *http.Request) {
// 	vowelSessions := must.Get(qh.queries.ListQuestionSessionByType(r.Context(), model.ListQuestionSessionByTypeParams{
// 		QuizSessionID: qh.quizSession.ID,
// 		Type:          model.VowelQuestionType,
// 	}))
// 	ids := map[int]bool{}
// 	for _, vs := range vowelSessions {
// 		if vs.Status.IsSubmitted() {
// 			ids[vs.QuestionID] = true
// 		}
// 	}
// 	for _, q := range qh.questions {
// 		if ids[q.ID] {
// 			data := model.MustParseQuestionData(q)
// 			qh.excerpt.Ref(data.Reference).ReplaceWithText(data.Answer)
// 		}
// 	}
// }

// // Useful for map
// func excerptRef(qd model.QuestionData) int {
// 	return qd.Reference
// }

// func isVowelQuestionType(q model.Question) bool {
// 	return q.Type == model.VowelQuestionType
// }

// func (qh quizHandler) submitAnswer(r *http.Request, q model.Question, answer string, status model.QuestionStatus) {
// 	_, err := qh.queries.SubmitAnswer(r.Context(), model.SubmitAnswerParams{
// 		Answer:        answer,
// 		Status:        status,
// 		QuizSessionID: qh.quizSession.ID,
// 		QuestionID:    q.ID,
// 	})
// 	if err != nil {
// 		if errors.Is(err, sql.ErrNoRows) {
// 			qh.questionSession(r, q)
// 			// TEST(Amr Ojjeh): VERY IMPORTANT TO TEST THIS BRANCH
// 			qh.submitAnswer(r, q, answer, status)
// 		}
// 		panic(err)
// 	}
// }

// func (qh quizHandler) questionSession(r *http.Request, q model.Question) model.QuestionSession {
// 	questionSession, err := qh.queries.GetQuestionSession(r.Context(), model.GetQuestionSessionParams{
// 		QuizSessionID: qh.quizSession.ID,
// 		QuestionID:    q.ID,
// 	})
// 	if err != nil {
// 		if errors.Is(err, sql.ErrNoRows) {
// 			return must.Get(qh.queries.CreateQuestionSession(r.Context(), model.CreateQuestionSessionParams{
// 				QuizSessionID: qh.quizSession.ID,
// 				QuestionID:    q.ID,
// 				Answer:        "",
// 				Status:        model.UnattemptedQuestionStatus,
// 			}))
// 		}
// 		panic(err)
// 	}
// 	return questionSession
// }
