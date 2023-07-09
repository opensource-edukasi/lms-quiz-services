package main

import (
	"context"
	"log"
	"net"
	"os"
	"time"

	_ "github.com/lib/pq"
	"google.golang.org/grpc"

	"lms-quiz-services/internal/config"
	"lms-quiz-services/internal/pkg/db/postgres"
	"lms-quiz-services/internal/pkg/db/redis"
	"lms-quiz-services/internal/route"
)

const defaultPort = "8000"

func main() {
	// lookup and setup env
	if _, ok := os.LookupEnv("PORT"); !ok {
		config.Setup(".env")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	// init log
	log := log.New(os.Stdout, "LMS : ", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)

	// create postgres database connection
	db, err := postgres.Open()
	if err != nil {
		log.Fatalf("connecting to db: %v", err)
		return
	}
	log.Print("connecting to postgresql database")

	defer db.Close()

	// create redis cache connection
	cache, err := redis.NewCache(context.Background(), os.Getenv("REDIS_ADDRESS"), os.Getenv("REDIS_PASSWORD"), 24*time.Hour)
	if err != nil {
		log.Fatalf("cannot create redis connection: %v", err)
		return
	}
	log.Print("connecting to redis cache")

	// listen tcp port
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	// routing grpc services
	route.GrpcRoute(grpcServer, db, log, cache)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
		return
	}
	log.Print("serve grpc on port: " + port)
}
