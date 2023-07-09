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
	var quizRepo QuizRepository
	quizRepo.pb = quizPb.Quiz{
		Id:          in.Id,
		Description: in.Description,
		Name:        in.Name,
		EndDate:     in.EndDate,
	}
	return &quizRepo.pb, nil
}
