// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: query.sql

package model

import (
	"context"
	"time"
)

const addQuizToClass = `-- name: AddQuizToClass :exec

INSERT INTO class_quiz (
    quiz_id, class_id
) VALUES (
    ?, ?
)
`

type AddQuizToClassParams struct {
	QuizID  int
	ClassID int
}

// **********
// CLASS_QUIZ TABLE
// **********
func (q *Queries) AddQuizToClass(ctx context.Context, arg AddQuizToClassParams) error {
	_, err := q.db.ExecContext(ctx, addQuizToClass, arg.QuizID, arg.ClassID)
	return err
}

const createClass = `-- name: CreateClass :one

INSERT INTO class (
    teacher_id, name, created, updated
) VALUES (
    ?, ?, datetime("now"), datetime("now")
) RETURNING id, teacher_id, name, created, updated
`

type CreateClassParams struct {
	TeacherID int
	Name      string
}

// **********
// CLASS TABLE
// **********
func (q *Queries) CreateClass(ctx context.Context, arg CreateClassParams) (Class, error) {
	row := q.db.QueryRowContext(ctx, createClass, arg.TeacherID, arg.Name)
	var i Class
	err := row.Scan(
		&i.ID,
		&i.TeacherID,
		&i.Name,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

const createQuestion = `-- name: CreateQuestion :one

INSERT INTO question (
    quiz_id, position, type, data, created, updated
) VALUES (
    ?, ?, ?, ?, datetime("now"), datetime("now")
) RETURNING id, quiz_id, position, type, data, created, updated
`

type CreateQuestionParams struct {
	QuizID   int
	Position int
	Type     QuestionType
	Data     []byte
}

// **********
// QUESTION TABLE
// **********
func (q *Queries) CreateQuestion(ctx context.Context, arg CreateQuestionParams) (Question, error) {
	row := q.db.QueryRowContext(ctx, createQuestion,
		arg.QuizID,
		arg.Position,
		arg.Type,
		arg.Data,
	)
	var i Question
	err := row.Scan(
		&i.ID,
		&i.QuizID,
		&i.Position,
		&i.Type,
		&i.Data,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

const createQuiz = `-- name: CreateQuiz :one

INSERT INTO quiz (
    teacher_id, title, excerpt, created, updated
) VALUES (
    ?, ?, ?, datetime("now"), datetime("now")
) RETURNING id, teacher_id, title, excerpt, created, updated
`

type CreateQuizParams struct {
	TeacherID int
	Title     string
	Excerpt   []byte
}

// **********
// QUIZ TABLE
// **********
func (q *Queries) CreateQuiz(ctx context.Context, arg CreateQuizParams) (Quiz, error) {
	row := q.db.QueryRowContext(ctx, createQuiz, arg.TeacherID, arg.Title, arg.Excerpt)
	var i Quiz
	err := row.Scan(
		&i.ID,
		&i.TeacherID,
		&i.Title,
		&i.Excerpt,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

const createStudent = `-- name: CreateStudent :one

INSERT INTO student (
    name, class_id, created, updated
) VALUES (
    ?, ?, datetime("now"), datetime("now")
) RETURNING id, name, class_id, created, updated
`

type CreateStudentParams struct {
	Name    string
	ClassID int
}

// **********
// STUDENT TABLE
// **********
func (q *Queries) CreateStudent(ctx context.Context, arg CreateStudentParams) (Student, error) {
	row := q.db.QueryRowContext(ctx, createStudent, arg.Name, arg.ClassID)
	var i Student
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.ClassID,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

const createStudentQuestionSession = `-- name: CreateStudentQuestionSession :one
INSERT INTO student_question_session (
    student_quiz_session_id, question_id, answer, status, created, updated
) VALUES (
    ?, ?, ?, ?, datetime("now"), datetime("now")
) RETURNING id, student_quiz_session_id, question_id, status, answer, created, updated
`

type CreateStudentQuestionSessionParams struct {
	StudentQuizSessionID int
	QuestionID           int
	Answer               string
	Status               QuestionStatus
}

func (q *Queries) CreateStudentQuestionSession(ctx context.Context, arg CreateStudentQuestionSessionParams) (StudentQuestionSession, error) {
	row := q.db.QueryRowContext(ctx, createStudentQuestionSession,
		arg.StudentQuizSessionID,
		arg.QuestionID,
		arg.Answer,
		arg.Status,
	)
	var i StudentQuestionSession
	err := row.Scan(
		&i.ID,
		&i.StudentQuizSessionID,
		&i.QuestionID,
		&i.Status,
		&i.Answer,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

const createStudentQuizSession = `-- name: CreateStudentQuizSession :one

INSERT INTO student_quiz_session (
    student_id, quiz_id, status, created, updated
) VALUES (
    ?, ?, ?, datetime("now"), datetime("now")
) RETURNING id, student_id, quiz_id, status, created, updated
`

type CreateStudentQuizSessionParams struct {
	StudentID int
	QuizID    int
	Status    QuizStatus
}

// **********
// STUDENT_QUIZ_SESSION TABLE
// **********
func (q *Queries) CreateStudentQuizSession(ctx context.Context, arg CreateStudentQuizSessionParams) (StudentQuizSession, error) {
	row := q.db.QueryRowContext(ctx, createStudentQuizSession, arg.StudentID, arg.QuizID, arg.Status)
	var i StudentQuizSession
	err := row.Scan(
		&i.ID,
		&i.StudentID,
		&i.QuizID,
		&i.Status,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

const createTeacher = `-- name: CreateTeacher :one


INSERT INTO teacher (
    email, username, password_hash, created, updated
) VALUES (
    ?, ?, ?, datetime("now"), datetime("now")
) RETURNING id, email, username, password_hash, created, updated
`

type CreateTeacherParams struct {
	Email        string
	Username     string
	PasswordHash string
}

// Copyright 2024 Amr Ojjeh <amrojjeh@outlook.com>
// **********
// TEACHER TABLE
// **********
func (q *Queries) CreateTeacher(ctx context.Context, arg CreateTeacherParams) (Teacher, error) {
	row := q.db.QueryRowContext(ctx, createTeacher, arg.Email, arg.Username, arg.PasswordHash)
	var i Teacher
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.Username,
		&i.PasswordHash,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

const deleteClass = `-- name: DeleteClass :exec
DELETE FROM class
WHERE id=?
`

func (q *Queries) DeleteClass(ctx context.Context, id int) error {
	_, err := q.db.ExecContext(ctx, deleteClass, id)
	return err
}

const deleteQuestion = `-- name: DeleteQuestion :exec
DELETE FROM question
WHERE id=?
`

func (q *Queries) DeleteQuestion(ctx context.Context, id int) error {
	_, err := q.db.ExecContext(ctx, deleteQuestion, id)
	return err
}

const deleteQuiz = `-- name: DeleteQuiz :exec
DELETE FROM quiz
WHERE id=?
`

func (q *Queries) DeleteQuiz(ctx context.Context, id int) error {
	_, err := q.db.ExecContext(ctx, deleteQuiz, id)
	return err
}

const deleteStudent = `-- name: DeleteStudent :exec
DELETE FROM student
WHERE id=?
`

func (q *Queries) DeleteStudent(ctx context.Context, id int) error {
	_, err := q.db.ExecContext(ctx, deleteStudent, id)
	return err
}

const deleteTeacher = `-- name: DeleteTeacher :exec
DELETE FROM teacher
WHERE email=?
`

func (q *Queries) DeleteTeacher(ctx context.Context, email string) error {
	_, err := q.db.ExecContext(ctx, deleteTeacher, email)
	return err
}

const getClass = `-- name: GetClass :one
SELECT id, teacher_id, name, created, updated FROM class
WHERE id=?
`

func (q *Queries) GetClass(ctx context.Context, id int) (Class, error) {
	row := q.db.QueryRowContext(ctx, getClass, id)
	var i Class
	err := row.Scan(
		&i.ID,
		&i.TeacherID,
		&i.Name,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

const getQuestion = `-- name: GetQuestion :one
SELECT id, quiz_id, position, type, data, created, updated FROM question
WHERE id=?
`

func (q *Queries) GetQuestion(ctx context.Context, id int) (Question, error) {
	row := q.db.QueryRowContext(ctx, getQuestion, id)
	var i Question
	err := row.Scan(
		&i.ID,
		&i.QuizID,
		&i.Position,
		&i.Type,
		&i.Data,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

const getQuiz = `-- name: GetQuiz :one
SELECT id, teacher_id, title, excerpt, created, updated FROM quiz
WHERE id=?
`

func (q *Queries) GetQuiz(ctx context.Context, id int) (Quiz, error) {
	row := q.db.QueryRowContext(ctx, getQuiz, id)
	var i Quiz
	err := row.Scan(
		&i.ID,
		&i.TeacherID,
		&i.Title,
		&i.Excerpt,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

const getTeacherByEmail = `-- name: GetTeacherByEmail :one
SELECT id, email, username, password_hash, created, updated FROM teacher
WHERE email=?
`

func (q *Queries) GetTeacherByEmail(ctx context.Context, email string) (Teacher, error) {
	row := q.db.QueryRowContext(ctx, getTeacherByEmail, email)
	var i Teacher
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.Username,
		&i.PasswordHash,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

const listClassesByTeacher = `-- name: ListClassesByTeacher :many
SELECT id, teacher_id, name, created, updated FROM class
WHERE teacher_id=?
ORDER BY created
`

func (q *Queries) ListClassesByTeacher(ctx context.Context, teacherID int) ([]Class, error) {
	rows, err := q.db.QueryContext(ctx, listClassesByTeacher, teacherID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Class
	for rows.Next() {
		var i Class
		if err := rows.Scan(
			&i.ID,
			&i.TeacherID,
			&i.Name,
			&i.Created,
			&i.Updated,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listQuestionsByQuiz = `-- name: ListQuestionsByQuiz :many
SELECT id, quiz_id, position, type, data, created, updated FROM question
WHERE quiz_id=?
ORDER BY position
`

func (q *Queries) ListQuestionsByQuiz(ctx context.Context, quizID int) ([]Question, error) {
	rows, err := q.db.QueryContext(ctx, listQuestionsByQuiz, quizID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Question
	for rows.Next() {
		var i Question
		if err := rows.Scan(
			&i.ID,
			&i.QuizID,
			&i.Position,
			&i.Type,
			&i.Data,
			&i.Created,
			&i.Updated,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listQuestionsByQuizAndType = `-- name: ListQuestionsByQuizAndType :many
SELECT id, quiz_id, position, type, data, created, updated FROM question
WHERE quiz_id=? AND type=?
ORDER BY position
`

type ListQuestionsByQuizAndTypeParams struct {
	QuizID int
	Type   QuestionType
}

func (q *Queries) ListQuestionsByQuizAndType(ctx context.Context, arg ListQuestionsByQuizAndTypeParams) ([]Question, error) {
	rows, err := q.db.QueryContext(ctx, listQuestionsByQuizAndType, arg.QuizID, arg.Type)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Question
	for rows.Next() {
		var i Question
		if err := rows.Scan(
			&i.ID,
			&i.QuizID,
			&i.Position,
			&i.Type,
			&i.Data,
			&i.Created,
			&i.Updated,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listQuizzesByClass = `-- name: ListQuizzesByClass :many
SELECT id, teacher_id, title, excerpt, q.created, q.updated, class_id, quiz_id, cq.created, cq.updated
FROM quiz AS q
INNER JOIN class_quiz AS cq ON cq.quiz_id=q.id
WHERE cq.class_id=?
`

type ListQuizzesByClassRow struct {
	ID        int
	TeacherID int
	Title     string
	Excerpt   []byte
	Created   time.Time
	Updated   time.Time
	ClassID   int
	QuizID    int
	Created_2 time.Time
	Updated_2 time.Time
}

func (q *Queries) ListQuizzesByClass(ctx context.Context, classID int) ([]ListQuizzesByClassRow, error) {
	rows, err := q.db.QueryContext(ctx, listQuizzesByClass, classID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListQuizzesByClassRow
	for rows.Next() {
		var i ListQuizzesByClassRow
		if err := rows.Scan(
			&i.ID,
			&i.TeacherID,
			&i.Title,
			&i.Excerpt,
			&i.Created,
			&i.Updated,
			&i.ClassID,
			&i.QuizID,
			&i.Created_2,
			&i.Updated_2,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listSegmentedQuestionsByQuiz = `-- name: ListSegmentedQuestionsByQuiz :many
SELECT id, quiz_id, position, type, data, created, updated FROM question
WHERE quiz_id=? AND segmented=TRUE
ORDER BY position
`

func (q *Queries) ListSegmentedQuestionsByQuiz(ctx context.Context, quizID int) ([]Question, error) {
	rows, err := q.db.QueryContext(ctx, listSegmentedQuestionsByQuiz, quizID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Question
	for rows.Next() {
		var i Question
		if err := rows.Scan(
			&i.ID,
			&i.QuizID,
			&i.Position,
			&i.Type,
			&i.Data,
			&i.Created,
			&i.Updated,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listStudentsByClass = `-- name: ListStudentsByClass :many
SELECT id, name, class_id, created, updated FROM student
WHERE class_id=?
`

func (q *Queries) ListStudentsByClass(ctx context.Context, classID int) ([]Student, error) {
	rows, err := q.db.QueryContext(ctx, listStudentsByClass, classID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Student
	for rows.Next() {
		var i Student
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.ClassID,
			&i.Created,
			&i.Updated,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listTeachers = `-- name: ListTeachers :many
SELECT id, email, username, password_hash, created, updated FROM teacher
WHERE username LIKE ? AND email LIKE ?
ORDER BY email
`

type ListTeachersParams struct {
	Username string
	Email    string
}

func (q *Queries) ListTeachers(ctx context.Context, arg ListTeachersParams) ([]Teacher, error) {
	rows, err := q.db.QueryContext(ctx, listTeachers, arg.Username, arg.Email)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Teacher
	for rows.Next() {
		var i Teacher
		if err := rows.Scan(
			&i.ID,
			&i.Email,
			&i.Username,
			&i.PasswordHash,
			&i.Created,
			&i.Updated,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const removeQuizFromClass = `-- name: RemoveQuizFromClass :exec
DELETE FROM class_quiz
WHERE quiz_id=? AND class_id=?
`

type RemoveQuizFromClassParams struct {
	QuizID  int
	ClassID int
}

func (q *Queries) RemoveQuizFromClass(ctx context.Context, arg RemoveQuizFromClassParams) error {
	_, err := q.db.ExecContext(ctx, removeQuizFromClass, arg.QuizID, arg.ClassID)
	return err
}

const submitAnswer = `-- name: SubmitAnswer :one

UPDATE student_question_session
SET answer=?, status=?, updated=datetime("now")
WHERE student_quiz_session_id=?
RETURNING id, student_quiz_session_id, question_id, status, answer, created, updated
`

type SubmitAnswerParams struct {
	Answer               string
	Status               QuestionStatus
	StudentQuizSessionID int
}

// **********
// STUDENT_QUESTION_SESSION TABLE
// **********
func (q *Queries) SubmitAnswer(ctx context.Context, arg SubmitAnswerParams) (StudentQuestionSession, error) {
	row := q.db.QueryRowContext(ctx, submitAnswer, arg.Answer, arg.Status, arg.StudentQuizSessionID)
	var i StudentQuestionSession
	err := row.Scan(
		&i.ID,
		&i.StudentQuizSessionID,
		&i.QuestionID,
		&i.Status,
		&i.Answer,
		&i.Created,
		&i.Updated,
	)
	return i, err
}
