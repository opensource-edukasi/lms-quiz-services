package route

import (
	"database/sql"
	"log"

	"google.golang.org/grpc"

	quizDomain "lms-quiz-services/internal/domain/quizzes"
	"lms-quiz-services/internal/pkg/db/redis"
	quizPb "lms-quiz-services/pb/quizzes"
)

// GrpcRoute func
func GrpcRoute(grpcServer *grpc.Server, db *sql.DB, log *log.Logger, cache *redis.Cache) {
	quizServer := quizDomain.QuizService{Db: db, Cache: cache, Log: log}
	quizPb.RegisterQuizzesServer(grpcServer, &quizServer)
}
