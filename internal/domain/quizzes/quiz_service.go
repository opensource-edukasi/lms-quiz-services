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
	var quizRepo QuizRepository
	var err error

	// TODO: validasi apakah user yang login mengambil kelas pada quiz ini

	// TODO: validate quizAnswerInput
	quizRepo.pbAnswer = quizPb.QuizAnswer{
		Quiz: &quizPb.Quiz{Id: in.QuizId},
	}

	for _, questionAnswerInput := range in.QuestionAnswer {
		// TODO: validate questionAnswerInput
		questionAnswer := &quizPb.QuestionAnswer{
			Question: &quizPb.Question{Id: questionAnswerInput.QuestionId},
			AnswerId: questionAnswerInput.AnswerId,
		}

		quizRepo.pbAnswer.QuestionAnswer = append(quizRepo.pbAnswer.QuestionAnswer, questionAnswer)
	}

	quizRepo.pb.Id = quizRepo.pbAnswer.Quiz.Id

	quizRepo.tx, err = a.Db.BeginTx(ctx, nil)
	if err != nil {
		return &quizRepo.pbAnswer, err
	}

	err = quizRepo.FindQuizById(ctx)
	if err != nil {
		return &quizRepo.pbAnswer, err
	}
	quizRepo.pbAnswer.Quiz = &quizRepo.pb

	quizRepo.CalculateScore()

	err = quizRepo.Answer(ctx)

	if err != nil {
		return &quizRepo.pbAnswer, err
	}

	quizRepo.tx.Commit()

	return &quizRepo.pbAnswer, nil
}
