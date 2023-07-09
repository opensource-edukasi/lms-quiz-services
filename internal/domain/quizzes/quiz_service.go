package quizzes

import (
	"context"
	"database/sql"
	"lms-quiz-services/internal/pkg/db/redis"
	quizPb "lms-quiz-services/pb/quizzes"
)

type QuizService struct {
	Db    *sql.DB
	Cache *redis.Cache
}

func (a *QuizService) Update(ctx context.Context, in *quizPb.QuizUpdateInput) (*quizPb.Quiz, error) {
	return &quizPb.Quiz{}, nil
}
