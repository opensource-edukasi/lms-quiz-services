package quizzes

import (
	"context"
	"database/sql"
	"lms-quiz-services/internal/pkg/app"
	"lms-quiz-services/internal/pkg/db/redis"
	quizPb "lms-quiz-services/pb/quizzes"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type QuizService struct {
	Db    *sql.DB
	Cache *redis.Cache
	Log   *log.Logger
}

func (a *QuizService) GetResultQuiz(ctx context.Context, in *quizPb.GetResultQuizInput) (*quizPb.QuizAnswer, error) {
	var quizRepo QuizRepository
	var err error

	quizRepo.db = a.Db
	quizRepo.pbAnswer = quizPb.QuizAnswer{
		StudentId: in.StudentId,
		Quiz:      &quizPb.Quiz{Id: in.QuizId},
	}

	err = quizRepo.GetQuizAnswer(ctx)
	if err != nil {
		return nil, err
	}

	return &quizRepo.pbAnswer, nil
}

func (a *QuizService) Get(ctx context.Context, in *quizPb.Id) (*quizPb.Quiz, error) {
	var quizRepo QuizRepository
	var err error
	quizRepo.pb.Id = in.Id

	quizRepo.tx, err = a.Db.BeginTx(ctx, nil)
	if err != nil {
		a.Log.Println("Error beginning DB Transaction: ", err)
		return nil, err
	}

	err = quizRepo.FindQuizById(ctx)
	if err != nil {
		return nil, err
	}
	quizRepo.tx.Commit()
	return &quizRepo.pb, nil
}

func (a *QuizService) ListScoresByQuizId(ctx context.Context, in *quizPb.Id) (*quizPb.QuizScoreList, error) {
	if len(in.Id) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "quiz id is required")
	}

	var quizRepo QuizRepository
	quizRepo.db = a.Db
	quizRepo.Log = a.Log

	return quizRepo.ListScoresByQuizId(ctx, in.Id)
}

func (a *QuizService) Update(ctx context.Context, in *quizPb.QuizUpdateInput) (*quizPb.Quiz, error) {
	var quizRepo QuizRepository
	var err error
	quizRepo.tx, err = a.Db.BeginTx(ctx, nil)
	if err != nil {
		a.Log.Println("Error beginning DB Transaction: ", err)
		return &quizRepo.pb, err
	}

	// Validate quiz input
	if len(in.Id) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "id is required")
	}
	if len(in.Name) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "name is required")
	}
	if len(in.Description) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "description is required")
	}

	quizRepo.pb = quizPb.Quiz{
		Id:          in.Id,
		Description: in.Description,
		Name:        in.Name,
		EndDate:     in.EndDate,
	}

	for _, questionInput := range in.Question {
		// Validate question input
		if len(questionInput.Title) == 0 {
			return nil, status.Errorf(codes.InvalidArgument, "question title is required")
		}
		if len(questionInput.Option) < 2 {
			return nil, status.Errorf(codes.InvalidArgument, "each question must have at least 2 options")
		}

		question := &quizPb.Question{
			Id:          questionInput.Id,
			Title:       questionInput.Title,
			Description: questionInput.Description,
			StorageId:   questionInput.StorageId,
			AnswerId:    questionInput.AnswerId,
		}

		for _, opt := range questionInput.Option {
			// Validate option input
			if len(opt.Description) == 0 {
				return nil, status.Errorf(codes.InvalidArgument, "option description is required")
			}
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

	// Validate quiz answer input
	if len(in.QuizId) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "quiz_id is required")
	}

	quizRepo.pbAnswer = quizPb.QuizAnswer{
		Quiz: &quizPb.Quiz{Id: in.QuizId},
	}

	for _, questionAnswerInput := range in.QuestionAnswer {
		// Validate question answer input
		if len(questionAnswerInput.QuestionId) == 0 {
			return nil, status.Errorf(codes.InvalidArgument, "question_id is required")
		}
		if len(questionAnswerInput.AnswerId) == 0 {
			return nil, status.Errorf(codes.InvalidArgument, "answer_id is required")
		}

		questionAnswer := &quizPb.QuestionAnswer{
			Question: &quizPb.Question{Id: questionAnswerInput.QuestionId},
			AnswerId: questionAnswerInput.AnswerId,
		}

		quizRepo.pbAnswer.QuestionAnswer = append(quizRepo.pbAnswer.QuestionAnswer, questionAnswer)
	}

	quizRepo.pb.Id = quizRepo.pbAnswer.Quiz.Id

	quizRepo.tx, err = a.Db.BeginTx(ctx, nil)
	if err != nil {
		a.Log.Println("Error beginning DB Transaction: ", err)
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

func (a *QuizService) Delete(ctx context.Context, in *quizPb.Id) (*quizPb.BoolMessage, error) {
	var quizRepo QuizRepository

	var err error

	quizRepo.pb.Id = in.Id
	quizRepo.db = a.Db

	err = quizRepo.Delete(ctx)
	if err != nil {
		return &quizPb.BoolMessage{IsTrue: false}, err
	}
	return &quizPb.BoolMessage{IsTrue: true}, nil
}

func (a *QuizService) Create(ctx context.Context, in *quizPb.QuizCreateInput) (*quizPb.Quiz, error) {
	var quizRepo QuizRepository
	var err error
	quizRepo.Log = a.Log
	quizRepo.tx, err = a.Db.BeginTx(ctx, nil)
	if err != nil {
		a.Log.Println("Error beginning DB Transaction: ", err)
		return &quizRepo.pb, err
	}

	// Validate quiz input
	if len(in.SubjectClassId) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "subject_class_id is required")
	}
	if len(in.TopicSubjectId) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "topic_subject_id is required")
	}
	if len(in.Name) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "name is required")
	}
	if len(in.Description) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "description is required")
	}
	if len(in.EndDate) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "end_date is required")
	}

	quizRepo.pb = quizPb.Quiz{
		Description:    in.Description,
		Name:           in.Name,
		SubjectClassId: in.SubjectClassId,
		TopicSubjectId: in.TopicSubjectId,
		EndDate:        in.EndDate,
	}

	for _, questionInput := range in.Question {
		// Validate question input
		if len(questionInput.Title) == 0 {
			return nil, status.Errorf(codes.InvalidArgument, "question title is required")
		}
		if len(questionInput.Option) < 2 {
			return nil, status.Errorf(codes.InvalidArgument, "each question must have at least 2 options")
		}

		question := &quizPb.Question{
			Title:       questionInput.Title,
			Description: questionInput.Description,
			StorageId:   questionInput.StorageId,
		}

		for _, opt := range questionInput.Option {
			// Validate option input
			if len(opt.Description) == 0 {
				return nil, status.Errorf(codes.InvalidArgument, "option description is required")
			}
			question.Option = append(question.Option, &quizPb.Option{
				Description: opt.Description,
				StorageId:   opt.StorageId,
			})
		}

		quizRepo.pb.Question = append(quizRepo.pb.Question, question)
	}

	err = quizRepo.Create(ctx)

	if err != nil {
		return &quizRepo.pb, err
	}

	quizRepo.tx.Commit()

	// Call post service to create a post with type QUIZ
	go func() {
		postClient := &postServiceClient{Log: a.Log}
		userID := ctx.Value(app.Ctx("user_id")).(string)
		postClient.createPost(ctx, in.SubjectClassId, in.TopicSubjectId, quizRepo.pb.Id, in.Name, userID)
	}()

	return &quizRepo.pb, nil
}
