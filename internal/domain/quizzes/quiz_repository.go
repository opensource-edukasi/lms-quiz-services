package quizzes

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"lms-quiz-services/internal/pkg/app"
	"lms-quiz-services/internal/pkg/array"
	quizPb "lms-quiz-services/pb/quizzes"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type QuizRepository struct {
	db       *sql.DB
	tx       *sql.Tx
	pb       quizPb.Quiz
	pbAnswer quizPb.QuizAnswer
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
	query := fmt.Sprintf("DELETE FROM questions WHERE id IN (%s)", array.ConvertToWhereIn(1, ids))
	stmt, err := a.tx.PrepareContext(ctx, query)
	if err != nil {
		return status.Errorf(codes.Internal, "Prepare statement deleteQuestions: %v", err)
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, array.ConvertToAny(ids)...)
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

	var storageId sql.NullString
	if len(question.StorageId) > 0 {
		storageId.String = question.StorageId
		storageId.Valid = true
	}

	var answerId sql.NullString
	if len(question.AnswerId) > 0 {
		answerId.String = question.AnswerId
		answerId.Valid = true
	}

	question.UpdatedBy = ctx.Value(app.Ctx("user_id")).(string)
	err = stmt.QueryRowContext(ctx,
		a.pb.Id,
		question.Title,
		question.Description,
		storageId,
		answerId,
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
		UPDATE questions SET
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

	var storageId sql.NullString
	if len(question.StorageId) > 0 {
		storageId.String = question.StorageId
		storageId.Valid = true
	}

	var answerId sql.NullString
	if len(question.AnswerId) > 0 {
		answerId.String = question.AnswerId
		answerId.Valid = true
	}

	err = stmt.QueryRowContext(ctx,
		question.Title,
		question.Description,
		storageId,
		answerId,
		question.UpdatedBy,
		now,
		question.Id,
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
	query := fmt.Sprintf("DELETE FROM options WHERE id IN (%s)", array.ConvertToWhereIn(1, ids))
	stmt, err := a.tx.PrepareContext(ctx, query)
	if err != nil {
		return status.Errorf(codes.Internal, "Prepare statement deleteOptions: %v", err)
	}
	defer stmt.Close()
	_, err = stmt.ExecContext(ctx, array.ConvertToAny(ids)...)
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
	var storageId sql.NullString
	if len(option.StorageId) > 0 {
		storageId.String = option.StorageId
		storageId.Valid = true
	}

	err = stmt.QueryRowContext(ctx,
		quistionId,
		option.Description,
		storageId,
		option.UpdatedBy,
	).Scan(&option.Id, &option.UpdatedAt, &option.CreatedAt)
	if err != nil {
		return status.Errorf(codes.Internal, "Exec insert option: %v", err)
	}

	return nil

}

func (a *QuizRepository) UpdateOption(ctx context.Context, option *quizPb.Option) error {

	query := `
		UPDATE options SET
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
	var storageId sql.NullString
	if len(option.StorageId) > 0 {
		storageId.String = option.StorageId
		storageId.Valid = true
	}

	err = stmt.QueryRowContext(ctx,
		option.Description,
		storageId,
		option.UpdatedBy,
		now,
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

func (a *QuizRepository) FindQuizById(ctx context.Context) error {
	query := `
		SELECT id, subject_class_id, topic_subject_id, name, description, end_date, updated_at, updated_by, created_at
	 	FROM quizzes WHERE id = $1
	`
	stmt, err := a.tx.PrepareContext(ctx, query)
	if err != nil {
		return status.Errorf(codes.Internal, "Prepare statement FindQuizById: %v", err)
	}
	defer stmt.Close()

	err = stmt.QueryRowContext(ctx, a.pb.Id).Scan(
		&a.pb.Id,
		&a.pb.SubjectClassId,
		&a.pb.TopicSubjectId,
		&a.pb.Name,
		&a.pb.Description,
		&a.pb.EndDate,
		&a.pb.UpdatedAt,
		&a.pb.UpdatedBy,
		&a.pb.CreatedAt,
	)
	if err != nil {
		return status.Errorf(codes.Internal, "Query Row Context FindQuizById: %v", err)
	}

	err = a.GetQuestionByQuizId(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (a *QuizRepository) GetQuizAnswer(ctx context.Context) error {
	query := `
		SELECT sq.id, sq.score, 
			json_agg(
					jsonb_build_object(
						
						'question_id', saq.question_id,
						'is_correct', saq.is_correct
					)
				) AS answers
		FROM student_quizzes sq
		JOIN student_answer_quizzes saq ON saq.student_quiz_id = sq.id 
		WHERE sq.quiz_id = $1 AND sq.student_id = $2
		GROUP BY sq.id`

	stmt, err := a.db.PrepareContext(ctx, query)
	if err != nil {
		return status.Errorf(codes.Internal, "Prepare statement: %v", err)
	}
	defer stmt.Close()

	var jsonStr string
	err = stmt.QueryRowContext(ctx, a.pbAnswer.Quiz.Id, a.pbAnswer.StudentId).Scan(&a.pbAnswer.Id, &a.pbAnswer.Score, &jsonStr)
	if err != nil {
		return status.Errorf(codes.Internal, "Query Row Context: %v", err)
	}

	var answer []struct {
		IsCorrect  bool   `json:"is_correct"`
		QuestionId string `json:"question_id"`
	}

	err = json.Unmarshal([]byte(jsonStr), &answer)

	if err != nil {
		return status.Errorf(codes.Internal, "Unmarshal answers: %v", err)
	}

	for _, v := range answer {
		a.pbAnswer.QuestionAnswer = append(a.pbAnswer.QuestionAnswer, &quizPb.QuestionAnswer{
			Question:  &quizPb.Question{Id: v.QuestionId},
			IsCorrect: v.IsCorrect,
		})
	}

	return nil
}

func (a *QuizRepository) GetQuestionByQuizId(ctx context.Context) error {
	query := `
		SELECT questions.id, questions.title, questions.description, questions.storage_id, 
			questions.answer_id, questions.updated_at, questions.updated_by, questions.created_at,
			json_agg(DISTINCT jsonb_build_object(
				'id', options.id,
				'description', options.description,
				'storage_id', options.storage_id,
				'updated_at', options.updated_at,
				'updated_by', options.updated_by,
				'created_at', options.created_at
			)) as question_options	 	
		FROM questions JOIN options ON (questions.id = options.question_id) 
		WHERE questions.quiz_id = $1
		GROUP BY questions.id
	`
	stmt, err := a.tx.PrepareContext(ctx, query)
	if err != nil {
		return status.Errorf(codes.Internal, "Prepare statement GetQuestionByQuizId: %v", err)
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx, a.pb.Id)
	if err != nil {
		return status.Errorf(codes.Internal, "Query Context GetQuestionByQuizId: %v", err)
	}

	for rows.Next() {
		var question quizPb.Question
		var options string
		err := rows.Scan(
			&question.Id,
			&question.Title,
			&question.Description,
			&question.StorageId,
			&question.AnswerId,
			&question.UpdatedAt,
			&question.UpdatedBy,
			&question.CreatedAt,
			&options,
		)
		if err != nil {
			return status.Errorf(codes.Internal, "Scan GetQuestionByQuizId: %v", err)
		}

		optionStruct := []struct {
			Id          string
			Description string
			StorageId   string
			UpdatedAt   string
			UpdatedBy   string
			CreatedAt   string
		}{}
		err = json.Unmarshal([]byte(options), &optionStruct)
		if err != nil {
			return status.Errorf(codes.Internal, "unmarshal option GetQuestionByQuizId: %v", err)
		}

		for _, option := range optionStruct {
			question.Option = append(question.Option, &quizPb.Option{
				Id:          option.Id,
				Description: option.Description,
				StorageId:   option.StorageId,
				CreatedAt:   option.CreatedAt,
				UpdatedAt:   option.UpdatedAt,
				UpdatedBy:   option.UpdatedBy,
			})
		}

		a.pb.Question = append(a.pb.Question, &question)
	}

	if rows.Err() != nil {
		return status.Errorf(codes.Internal, "rows error on  GetQuestionByQuizId: %v", err)
	}

	return nil
}

func (a *QuizRepository) CalculateScore() {
	score := 0
	for _, question := range a.pb.Question {
		for _, answer := range a.pbAnswer.QuestionAnswer {
			if question.Id == answer.Question.Id {
				answer.Question = question
				if question.AnswerId == answer.AnswerId {
					score += 1
					answer.IsCorrect = true
				}
			}
		}
	}

	a.pbAnswer.Score = int32(score)
	return
}

func (a *QuizRepository) Answer(ctx context.Context) error {
	a.pbAnswer.StudentId = ctx.Value(app.Ctx("user_id")).(string)
	query := `
		INSERT INTO student_quizzes (quiz_id, student_id, score) VALUES ($1, $2, $3)
		RETURNING id, created_at
		`

	stmt, err := a.tx.PrepareContext(ctx, query)
	if err != nil {
		return status.Errorf(codes.Internal, "Prepare statement answer quiz: %v", err)
	}
	defer stmt.Close()

	err = stmt.QueryRowContext(ctx,
		a.pbAnswer.Quiz.Id,
		a.pbAnswer.StudentId,
		a.pbAnswer.Score,
	).Scan(&a.pbAnswer.Id, &a.pbAnswer.CreatedAt)

	if err != nil {
		return status.Errorf(codes.Internal, "Exec answer quiz: %v", err)
	}

	for _, questionAnswer := range a.pbAnswer.QuestionAnswer {
		if err := a.InsertQuestionAnswer(ctx, questionAnswer); err != nil {
			return err
		}
	}

	return nil
}

func (a *QuizRepository) InsertQuestionAnswer(ctx context.Context, questionAnswer *quizPb.QuestionAnswer) error {
	query := `
		INSERT INTO student_answer_quizzes (student_quiz_id, question_id, answer_id, is_correct) 
		VALUES ($1, $2, $3, $4)
		RETURNING created_at
		`

	stmt, err := a.tx.PrepareContext(ctx, query)
	if err != nil {
		return status.Errorf(codes.Internal, "Prepare statement InsertQuestionAnswer : %v", err)
	}
	defer stmt.Close()

	err = stmt.QueryRowContext(ctx,
		a.pbAnswer.Id,
		questionAnswer.Question.Id,
		questionAnswer.AnswerId,
		questionAnswer.IsCorrect,
	).Scan(&a.pbAnswer.CreatedAt)

	if err != nil {
		return status.Errorf(codes.Internal, "Exec answer quiz: %v", err)
	}

	return nil
}

func (a *QuizRepository) Delete(ctx context.Context) error {
	// Prepare the DELETE statement
	query := `DELETE FROM quizzes WHERE id = $1`

	stmt, err := a.tx.PrepareContext(ctx, query)
	if err != nil {
		return status.Errorf(codes.Internal, "Prepare statement delete quizzes: %v", err)
	}
	defer stmt.Close()

	result, err := stmt.ExecContext(ctx, a.pb.Id)
	if err != nil {
		return status.Errorf(codes.Internal, "Exec delete quizzes: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return status.Errorf(codes.Internal, "Error getting rows affected: %v", err)
	}

	if rowsAffected == 0 {
		return status.Errorf(codes.NotFound, "Quiz with ID %s not found", a.pb.Id)
	}

	err = a.tx.Commit()
	if err != nil {
		return status.Errorf(codes.Internal, "Error committing transaction: %v", err)
	}

	return nil
}
