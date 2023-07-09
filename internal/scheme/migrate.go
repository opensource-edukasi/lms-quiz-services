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
				id uuid NOT NULL PRIMARY KEY,
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
				id uuid NOT NULL PRIMARY KEY,
				quiz_id uuid NOT NULL,
				title varchar(45) NOT NULL,
				description varchar(255) NOT NULL,
				storage_id uuid,
				answer_id uuid,
				created_at timestamptz NOT NULL DEFAULT timezone('utc', NOW()),
				updated_at timestamp NOT NULL DEFAULT timezone('utc', NOW()),
				updated_by uuid 
			);
		`,
	},
	{
		Version:     4,
		Description: "Create options Table",
		Script: `
			CREATE TABLE options (
				id uuid NOT NULL PRIMARY KEY,
				question_id uuid NOT NULL,
				description varchar(255) NOT NULL,
				storage_id uuid,
				answer_id uuid,
				created_at timestamptz NOT NULL DEFAULT timezone('utc', NOW()),
				updated_at timestamp NOT NULL DEFAULT timezone('utc', NOW()),
				updated_by uuid 
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
