package route

import (
	"database/sql"
	"log"

	"google.golang.org/grpc"

	"lms-quiz-services/internal/pkg/db/redis"
)

// GrpcRoute func
func GrpcRoute(grpcServer *grpc.Server, db *sql.DB, log *log.Logger, cache *redis.Cache) {
	//authServer := service.Auth{Db: db, Cache: cache}
	//users.RegisterAuthServiceServer(grpcServer, &authServer)
}
