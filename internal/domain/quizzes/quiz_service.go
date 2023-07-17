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
	var err error
	quizRepo.tx, err = a.Db.BeginTx(ctx, nil)
	if err != nil {
		return &quizRepo.pb, err
	}

	// TODO: validate quizInput

	quizRepo.pb = quizPb.Quiz{
		Id:          in.Id,
		Description: in.Description,
		Name:        in.Name,
		EndDate:     in.EndDate,
	}

	for _, questionInput := range in.Question {
		// TODO: validate questionInput
		question := &quizPb.Question{
			Id:          questionInput.Id,
			Title:       questionInput.Title,
			Description: questionInput.Description,
			StorageId:   questionInput.StorageId,
			AnswerId:    questionInput.AnswerId,
		}

		for _, opt := range questionInput.Option {
			// TODO: validate optionInput
			question.Option = append(question.Option, &quizPb.Option{
				Id:          opt.Id,
				Description: opt.Description,
				StorageId:   opt.StorageId,
			})
		}

		quizRepo.pb.Question = append(quizRepo.pb.Question, question)
	}

	err = quizRepo.Update(ctx)

	if err != nil {
		return &quizRepo.pb, err
	}

	quizRepo.tx.Commit()

	return &quizRepo.pb, nil
}

func (a *QuizService) Answer(ctx context.Context, in *quizPb.QuizAnswerInput) (*quizPb.QuizAnswer, error) {
	return nil, nil
}
