package main

import (
	"log"
	"net"

	application "rushplay/internal/application/usecase"
	database "rushplay/internal/infrastructure/database"
	repository "rushplay/internal/infrastructure/repository"

	"google.golang.org/grpc"

	"rushplay/api/generated/proto/userpb"
	"rushplay/internal/transport"
)

func main() {
	db, err := database.NewDatabase()

	if err != nil {
		log.Fatalf("Failed to Connect to database: %v", err)
	}

	userRepo := repository.NewUserRepository(db)
	userUseCase := application.NewUserUseCase(userRepo)
	userHandler := transport.NewUserHandler(userUseCase)

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	userpb.RegisterUserServiceServer(grpcServer, userHandler)

	log.Println("Rushplay gRPC server listening on port 50051")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
