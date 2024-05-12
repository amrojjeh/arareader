-- Copyright 2024 Amr Ojjeh <amrojjeh@outlook.com>

-- **********
-- TEACHER TABLE
-- **********

-- name: CreateTeacher :one
INSERT INTO teacher (
    email, username, password_hash, created, updated
) VALUES (
    ?, ?, ?, datetime("now"), datetime("now")
) RETURNING *;

-- name: GetTeacherByEmail :one
SELECT * FROM teacher
WHERE email=?;

-- name: ListTeachers :many
SELECT * FROM teacher
WHERE username LIKE ? AND email LIKE ?
ORDER BY email;

-- name: DeleteTeacher :exec
DELETE FROM teacher
WHERE email=?;

-- **********
-- QUIZ TABLE
-- **********

-- name: CreateQuiz :one
INSERT INTO quiz (
    teacher_id, title, excerpt, created, updated
) VALUES (
    ?, ?, ?, datetime("now"), datetime("now")
) RETURNING *;

-- name: GetQuiz :one
SELECT * FROM quiz
WHERE id=?;

-- name: ListQuizzesByClass :many
SELECT *
FROM quiz AS q
INNER JOIN class_quiz AS cq ON cq.quiz_id=q.id
WHERE cq.class_id=?;

-- name: DeleteQuiz :exec
DELETE FROM quiz
WHERE id=?;

-- **********
-- QUESTION TABLE
-- **********

-- name: CreateQuestion :one
INSERT INTO question (
    quiz_id, position, type, data, created, updated
) VALUES (
    ?, ?, ?, ?, datetime("now"), datetime("now")
) RETURNING *;

-- name: GetQuestion :one
SELECT * FROM question
WHERE id=?;

-- name: ListQuestionsByQuizAndType :many
SELECT * FROM question
WHERE quiz_id=? AND type=?
ORDER BY position;

-- name: ListQuestionsByQuiz :many
SELECT * FROM question
WHERE quiz_id=?
ORDER BY position;

-- name: ListSegmentedQuestionsByQuiz :many
SELECT * FROM question
WHERE quiz_id=? AND segmented=TRUE
ORDER BY position;


-- name: DeleteQuestion :exec
DELETE FROM question
WHERE id=?;

-- **********
-- CLASS TABLE
-- **********

-- name: CreateClass :one
INSERT INTO class (
    teacher_id, name, created, updated
) VALUES (
    ?, ?, datetime("now"), datetime("now")
) RETURNING *;

-- name: GetClass :one
SELECT * FROM class
WHERE id=?;

-- name: ListClassesByTeacher :many
SELECT * FROM class
WHERE teacher_id=?
ORDER BY created;

-- name: DeleteClass :exec
DELETE FROM class
WHERE id=?;

-- **********
-- CLASS_QUIZ TABLE
-- **********

-- name: AddQuizToClass :exec
INSERT INTO class_quiz (
    quiz_id, class_id
) VALUES (
    ?, ?
);

-- name: RemoveQuizFromClass :exec
DELETE FROM class_quiz
WHERE quiz_id=? AND class_id=?;


-- **********
-- STUDENT TABLE
-- **********

-- name: CreateStudent :one
INSERT INTO student (
    name, class_id, created, updated
) VALUES (
    ?, ?, datetime("now"), datetime("now")
) RETURNING *;

-- name: ListStudentsByClass :many
SELECT * FROM student
WHERE class_id=?;

-- name: DeleteStudent :exec
DELETE FROM student
WHERE id=?;

-- **********
-- STUDENT_QUIZ_SESSION TABLE
-- **********

-- name: CreateStudentQuizSession :one
INSERT INTO student_quiz_session (
    student_id, quiz_id, status, created, updated
) VALUES (
    ?, ?, ?, datetime("now"), datetime("now")
) RETURNING *;

-- **********
-- STUDENT_QUESTION_SESSION TABLE
-- **********

-- name: CreateStudentQuestionSession :one
INSERT INTO student_question_session (
    student_quiz_session_id, status, created, updated
) VALUES (
    ?, ?, datetime("now"), datetime("now")
) RETURNING *;
