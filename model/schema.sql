-- Copyright © 2024 Amr Ojjeh <amrojjeh@outlook.com>

CREATE TABLE IF NOT EXISTS teacher (
    id INTEGER NOT NULL PRIMARY KEY,
    email TEXT NOT NULL,
    username TEXT NOT NULL,
    password_hash TEXT NOT NULL,
    created DATETIME NOT NULL,
    updated DATETIME NOT NULL,

    CONSTRAINT teacher_email_uc UNIQUE (email)
);

CREATE TABLE IF NOT EXISTS quiz (
    id INTEGER NOT NULL PRIMARY KEY,
    teacher_id INTEGER NOT NULL,
    title TEXT NOT NULL,
    excerpt XML NOT NULL,
    created DATETIME NOT NULL,
    updated DATETIME NOT NULL,

    FOREIGN KEY (teacher_id)
        REFERENCES teacher(id)
        ON DELETE CASCADE
        ON UPDATE CASCADE
);

CREATE TABLE IF NOT EXISTS question (
    id INTEGER NOT NULL PRIMARY KEY,
    quiz_id INTEGER NOT NULL,
    position INTEGER NOT NULL,
    type TEXT NOT NULL,
    data JSON NOT NULL,
    created DATETIME NOT NULL,
    updated DATETIME NOT NULL,

    FOREIGN KEY (quiz_id)
        REFERENCES quiz(id)
        ON DELETE SET NULL
        ON UPDATE CASCADE
);

CREATE TABLE IF NOT EXISTS class (
    id INTEGER NOT NULL PRIMARY KEY,
    teacher_id INTEGER NOT NULL,
    name TEXT NOT NULL,
    created DATETIME NOT NULL,
    updated DATETIME NOT NULL,
    FOREIGN KEY (teacher_id)
        REFERENCES teacher(id)
        ON DELETE CASCADE
        ON UPDATE CASCADe
);

CREATE TABLE IF NOT EXISTS class_quiz (
    class_id INTEGER NOT NULL,
    quiz_id INTEGER NOT NULL,
    created DATETIME NOT NULL,
    updated DATETIME NOT NULL,
    PRIMARY KEY (class_id, quiz_id),
    FOREIGN KEY (class_id)
        REFERENCES class(id)
        ON DELETE CASCADE
        ON UPDATE CASCADE,
    FOREIGN KEY (quiz_id)
        REFERENCES quiz(id)
        ON DELETE CASCADE
        ON UPDATE CASCADE
);

CREATE TABLE IF NOT EXISTS student (
    id INTEGER NOT NULL PRIMARY KEY,
    name TEXT NOT NULL,
    class_id INTEGER NOT NULL,
    created DATETIME NOT NULL,
    updated DATETIME NOT NULL,
    FOREIGN KEY (class_id)
        REFERENCES class(id)
        ON DELETE CASCADE
        ON UPDATE CASCADE
);

CREATE TABLE IF NOT EXISTS student_quiz_session (
    id INTEGER NOT NULL PRIMARY KEY,
    student_id INTEGER NOT NULL,
    quiz_id INTEGER NOT NULL,
    status TEXT NOT NULL, -- "submitted" | "unsubmitted"
    created DATETIME NOT NULL,
    updated DATETIME NOT NULL,
    FOREIGN KEY (student_id)
        REFERENCES student(id)
        ON DELETE CASCADE
        ON UPDATE CASCADE,
    FOREIGN KEY (quiz_id)
        REFERENCES quiz(id)
        ON DELETE CASCADE
        ON UPDATE CASCADE
);

CREATE TABLE IF NOT EXISTS student_question_session (
    id INTEGER NOT NULL PRIMARY KEY,
    student_quiz_session_id INTEGER NOT NULL,
    question_id INTEGER NOT NULL,
    status TEXT NOT NULL, -- "correct" | "incorrect" | "pending" | "unsubmitted" | "unattempted"
    answer TEXT NOT NULL,
    created DATETIME NOT NULL,
    updated DATETIME NOT NULL,
    FOREIGN KEY (student_quiz_session_id)
        REFERENCES student_quiz_session(id)
        ON DELETE CASCADE
        ON UPDATE CASCADE,
    FOREIGN KEY (question_id)
        REFERENCES question(id)
        ON DELETE CASCADE
        ON UPDATE CASCADE
);


-- For: https://github.com/alexedwards/scs/tree/master/sqlite3store
-- Should be "session" :(
CREATE TABLE sessions (
	token TEXT PRIMARY KEY,
	data BLOB NOT NULL,
	expiry REAL NOT NULL
);

CREATE INDEX sessions_expiry_idx ON sessions(expiry);
