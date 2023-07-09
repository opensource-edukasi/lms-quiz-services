package quizzes

import (
	"database/sql"
	quizPb "lms-quiz-services/pb/quizzes"
)

type QuizRepository struct {
	db *sql.DB
	pb quizPb.Quiz
}
