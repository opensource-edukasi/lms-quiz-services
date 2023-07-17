package scheme

import (
	"database/sql"

	"github.com/GuiaBolso/darwin"
)

var migrations = []darwin.Migration{
	{
		Version:     1,
		Description: "Create uuid extension",
		Script:      `CREATE EXTENSION "uuid-ossp";`,
	},
	{
		Version:     2,
		Description: "Create quizzes Table",
		Script: `
			CREATE TABLE quizzes (
				id uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4 (),
				subject_class_id uuid NOT NULL,
				topic_subject_id uuid NOT NULL,
				name varchar(45) NOT NULL,
				description varchar(255) NOT NULL,
				end_date timestamptz NOT NULL,
				created_at timestamptz NOT NULL DEFAULT timezone('utc', NOW()),
				updated_at timestamp NOT NULL DEFAULT timezone('utc', NOW()),
				updated_by uuid 
			);
		`,
	},
	{
		Version:     3,
		Description: "Create questions Table",
		Script: `
			CREATE TABLE questions (
				id uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4 (),
				quiz_id uuid NOT NULL,
				title varchar(45) NOT NULL,
				description varchar(255) NOT NULL,
				storage_id uuid,
				answer_id uuid,
				created_at timestamptz NOT NULL DEFAULT timezone('utc', NOW()),
				updated_at timestamp NOT NULL DEFAULT timezone('utc', NOW()),
				updated_by uuid,
				CONSTRAINT fk_questions_to_quizzes FOREIGN KEY(quiz_id) REFERENCES quizzes(id) ON DELETE CASCADE 
			);
		`,
	},
	{
		Version:     4,
		Description: "Create options Table",
		Script: `
			CREATE TABLE options (
				id uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4 (),
				question_id uuid NOT NULL,
				description varchar(255) NOT NULL,
				storage_id uuid,
				created_at timestamptz NOT NULL DEFAULT timezone('utc', NOW()),
				updated_at timestamp NOT NULL DEFAULT timezone('utc', NOW()),
				updated_by uuid,
				CONSTRAINT fk_options_to_questions FOREIGN KEY(question_id) REFERENCES questions(id) ON DELETE CASCADE 
			);
		`,
	},
	{
		Version:     5,
		Description: "Create student_quizzes Table",
		Script: `
			CREATE TABLE student_quizzes (
				id uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4 (),
				quiz_id uuid NOT NULL,
				student_id uuid NOT NULL,
				score smallint NOT NULL,
				created_at timestamptz NOT NULL DEFAULT timezone('utc', NOW()),
				UNIQUE(quiz_id, student_id),
				CONSTRAINT fk_student_quizzes_to_quizzes FOREIGN KEY(quiz_id) REFERENCES quizzes(id) ON DELETE CASCADE
			);
		`,
	},
	{
		Version:     6,
		Description: "Create student_answer_quizzes Table",
		Script: `
			CREATE TABLE student_answer_quizzes (
				id uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4 (),
				student_quiz_id uuid NOT NULL,
				question_id uuid NOT NULL,
				answer_id uuid,
				is_correct boolean NOT NULL DEFAULT false,
				created_at timestamptz NOT NULL DEFAULT timezone('utc', NOW()),
				UNIQUE(student_quiz_id, question_id),
				CONSTRAINT fk_student_answer_quizzes_to_student_quizzes FOREIGN KEY(student_quiz_id) REFERENCES student_quizzes(id) ON DELETE CASCADE
			);
		`,
	},
}

// Migrate attempts to bring the schema for db up to date with the migrations
// defined in this package.
func Migrate(db *sql.DB) error {
	driver := darwin.NewGenericDriver(db, darwin.PostgresDialect{})

	d := darwin.New(driver, migrations, nil)

	return d.Migrate()
}
