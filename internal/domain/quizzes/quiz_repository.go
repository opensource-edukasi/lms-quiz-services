package quizzes

import (
	"context"
	"database/sql"
	"lms-quiz-services/internal/pkg/app"
	"lms-quiz-services/internal/pkg/array"
	quizPb "lms-quiz-services/pb/quizzes"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type QuizRepository struct {
	db *sql.DB
	tx *sql.Tx
	pb quizPb.Quiz
}

func (a *QuizRepository) Update(ctx context.Context) error {

	query := `
		UPDATE quizzes SET
			name = $1,
			description = $2,
			updated_by = $3,
			updated_at = $4
		WHERE id = $5
		RETURNING subject_class_id, topic_subject_id, created_at
		`

	stmt, err := a.tx.PrepareContext(ctx, query)
	if err != nil {
		return status.Errorf(codes.Internal, "Prepare statement update quiz: %v", err)
	}
	defer stmt.Close()

	now := time.Now().UTC()
	a.pb.UpdatedBy = ctx.Value(app.Ctx("user_id")).(string)
	a.pb.UpdatedAt = now.String()

	err = stmt.QueryRowContext(ctx,
		a.pb.Name,
		a.pb.Description,
		a.pb.UpdatedBy,
		now,
		a.pb.Id,
	).Scan(&a.pb.SubjectClassId, &a.pb.TopicSubjectId, &a.pb.CreatedAt)

	if err != nil {
		return status.Errorf(codes.Internal, "Exec update quiz: %v", err)
	}

	extIds, err := a.getExtQuestions(ctx)
	if err != nil {
		return err
	}

	for _, question := range a.pb.Question {
		if len(question.Id) == 0 {
			if err := a.InsertQuestion(ctx, question); err != nil {
				return err
			}
		} else {
			if err := a.UpdateQuestion(ctx, question); err != nil {
				return err
			}

			extIds = array.RemoveByValue(extIds, question.Id)
		}
	}

	if len(extIds) > 0 {
		if err = a.deleteQuestions(ctx, extIds); err != nil {
			return err
		}
	}

	return nil
}

func (a *QuizRepository) deleteQuestions(ctx context.Context, ids []string) error {
	query := "DELETE FROM questions WHERE id IN (" + array.ConvertToWhereIn(ids) + ")"
	stmt, err := a.tx.PrepareContext(ctx, query)
	if err != nil {
		return status.Errorf(codes.Internal, "Prepare statement deleteQuestions: %v", err)
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, array.ConvertToAny(ids))
	if err != nil {
		return status.Errorf(codes.Internal, "exec context deleteQuestions: %v", err)
	}

	return nil
}

func (a *QuizRepository) getExtQuestions(ctx context.Context) ([]string, error) {
	var ids []string
	query := `SELECT id FROM questions WHERE quiz_id = $1`
	stmt, err := a.tx.PrepareContext(ctx, query)
	if err != nil {
		return ids, status.Errorf(codes.Internal, "Prepare statement getExtQuestions: %v", err)
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx, a.pb.Id)
	if err != nil {
		return ids, status.Errorf(codes.Internal, "Query Context getExtQuestions: %v", err)
	}

	for rows.Next() {
		var id string
		err = rows.Scan(&id)
		if err != nil {
			return ids, status.Errorf(codes.Internal, "Scan getExtQuestions: %v", err)
		}
		ids = append(ids, id)
	}

	if rows.Err() != nil {
		return ids, status.Errorf(codes.Internal, "Rows getExtQuestions: %v", err)
	}

	return ids, nil
}

func (a *QuizRepository) InsertQuestion(ctx context.Context, question *quizPb.Question) error {

	query := `
		INSERT INTO question (quiz_id, title, description, storage_id, answer_id, updated_by)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id, updated_at, created_at
		`

	stmt, err := a.tx.PrepareContext(ctx, query)
	if err != nil {
		return status.Errorf(codes.Internal, "Prepare statement insert question: %v", err)
	}
	defer stmt.Close()

	question.UpdatedBy = ctx.Value(app.Ctx("user_id")).(string)
	err = stmt.QueryRowContext(ctx,
		a.pb.Id,
		question.Title,
		question.Description,
		question.StorageId,
		question.AnswerId,
		question.UpdatedBy,
	).Scan(&question.Id, &question.UpdatedAt, &question.CreatedAt)
	if err != nil {
		return status.Errorf(codes.Internal, "Exec insert question: %v", err)
	}

	for _, option := range question.Option {
		if err := a.InsertOption(ctx, option, question.Id); err != nil {
			return err
		}
	}

	return nil

}

func (a *QuizRepository) UpdateQuestion(ctx context.Context, question *quizPb.Question) error {

	query := `
		UPDATE question SET
			title = $1,
			description = $2,
			storage_id = $3,
			answer_id = $4,
			updated_by = $5,
			updated_at = $6
		WHERE id = $7
		RETURNING created_at
		`

	stmt, err := a.tx.PrepareContext(ctx, query)
	if err != nil {
		return status.Errorf(codes.Internal, "Prepare statement update question: %v", err)
	}
	defer stmt.Close()

	now := time.Now().UTC()
	question.UpdatedBy = ctx.Value(app.Ctx("user_id")).(string)
	question.UpdatedAt = now.String()

	err = stmt.QueryRowContext(ctx,
		question.Title,
		question.Description,
		question.StorageId,
		question.AnswerId,
		question.Id,
		question.UpdatedBy,
		question.UpdatedAt,
	).Scan(&question.CreatedAt)
	if err != nil {
		return status.Errorf(codes.Internal, "Exec update question: %v", err)
	}

	extIds, err := a.getExtOptions(ctx, question.Id)
	if err != nil {
		return err
	}

	for _, option := range question.Option {
		if len(option.Id) == 0 {
			if err := a.InsertOption(ctx, option, question.Id); err != nil {
				return err
			}
		} else {
			if err := a.UpdateOption(ctx, option); err != nil {
				return err
			}

			extIds = array.RemoveByValue(extIds, option.Id)
		}
	}

	if len(extIds) > 0 {
		if err = a.deleteOptions(ctx, extIds); err != nil {
			return err
		}
	}

	return nil

}

func (a *QuizRepository) DeleteQuestion(ctx context.Context, id string) error {

	query := `DELETE FROM questions WHERE id = $1`

	stmt, err := a.tx.PrepareContext(ctx, query)
	if err != nil {
		return status.Errorf(codes.Internal, "Prepare statement delete question: %v", err)
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, id)
	if err != nil {
		return status.Errorf(codes.Internal, "Exec delete question: %v", err)
	}

	return nil

}

func (a *QuizRepository) deleteOptions(ctx context.Context, ids []string) error {
	query := "DELETE FROM options WHERE id IN (" + array.ConvertToWhereIn(ids) + ")"
	stmt, err := a.tx.PrepareContext(ctx, query)
	if err != nil {
		return status.Errorf(codes.Internal, "Prepare statement deleteOptions: %v", err)
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, array.ConvertToAny(ids))
	if err != nil {
		return status.Errorf(codes.Internal, "exec context deleteOptions: %v", err)
	}

	return nil
}

func (a *QuizRepository) getExtOptions(ctx context.Context, questionId string) ([]string, error) {
	var ids []string
	query := `SELECT id FROM options WHERE question_id = $1`
	stmt, err := a.tx.PrepareContext(ctx, query)
	if err != nil {
		return ids, status.Errorf(codes.Internal, "Prepare statement getExtOptions: %v", err)
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx, questionId)
	if err != nil {
		return ids, status.Errorf(codes.Internal, "Query Context getExtOptions: %v", err)
	}

	for rows.Next() {
		var id string
		err = rows.Scan(&id)
		if err != nil {
			return ids, status.Errorf(codes.Internal, "Scan getExtOptions: %v", err)
		}
		ids = append(ids, id)
	}

	if rows.Err() != nil {
		return ids, status.Errorf(codes.Internal, "Rows getExtOptions: %v", err)
	}

	return ids, nil
}

func (a *QuizRepository) InsertOption(ctx context.Context, option *quizPb.Option, quistionId string) error {

	query := `
		INSERT INTO options (question_id, description, storage_id, updated_by) 
		VALUES ($1, $2, $3, $4)
		RETURNING id, updated_at, created_at
		`

	stmt, err := a.tx.PrepareContext(ctx, query)
	if err != nil {
		return status.Errorf(codes.Internal, "Prepare statement insert option: %v", err)
	}
	defer stmt.Close()

	option.UpdatedBy = ctx.Value(app.Ctx("user_id")).(string)

	err = stmt.QueryRowContext(ctx,
		quistionId,
		option.Description,
		option.StorageId,
		option.UpdatedBy,
	).Scan(&option.Id, &option.UpdatedAt, &option.CreatedAt)
	if err != nil {
		return status.Errorf(codes.Internal, "Exec insert option: %v", err)
	}

	return nil

}

func (a *QuizRepository) UpdateOption(ctx context.Context, option *quizPb.Option) error {

	query := `
		UPDATE option SET
			description = $1,
			storage_id = $2,
			updated_by = $3,
			updated_at = $4
		WHERE id = $5
		RETURNING created_at
		`

	stmt, err := a.tx.PrepareContext(ctx, query)
	if err != nil {
		return status.Errorf(codes.Internal, "Prepare statement update option: %v", err)
	}
	defer stmt.Close()

	now := time.Now().UTC()
	option.UpdatedBy = ctx.Value(app.Ctx("user_id")).(string)
	option.UpdatedAt = now.String()

	err = stmt.QueryRowContext(ctx,
		option.Description,
		option.StorageId,
		option.Id,
	).Scan(&option.CreatedAt)
	if err != nil {
		return status.Errorf(codes.Internal, "Exec update option: %v", err)
	}

	return nil

}

func (a *QuizRepository) DeleteOption(ctx context.Context, id string) error {

	query := `DELETE FROM options WHERE id = $1`

	stmt, err := a.tx.PrepareContext(ctx, query)
	if err != nil {
		return status.Errorf(codes.Internal, "Prepare statement delete option: %v", err)
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, id)
	if err != nil {
		return status.Errorf(codes.Internal, "Exec delete option: %v", err)
	}

	return nil

}
