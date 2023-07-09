package quizzes

import (
	"context"
	"database/sql"
	quizPb "lms-quiz-services/pb/quizzes"

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
			description = $2
		WHERE id = $3
		`

	stmt, err := a.db.PrepareContext(ctx, query)
	if err != nil {
		return status.Errorf(codes.Internal, "Prepare statement update quiz: %v", err)
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx,
		a.pb.GetName(),
		a.pb.GetDescription(),
		a.pb.GetId(),
	)
	if err != nil {
		return status.Errorf(codes.Internal, "Exec update employee: %v", err)
	}

	return nil

}
