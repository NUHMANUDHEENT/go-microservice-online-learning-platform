package main

import (
	"log"
	"net"

	"github.com/nuhmanudheent/online-learning-platform/user-service/internal/handler"
	"github.com/nuhmanudheent/online-learning-platform/user-service/internal/repository"
	"github.com/nuhmanudheent/online-learning-platform/user-service/internal/service"
	pb "github.com/nuhmanudheent/online-learning-platform/user-service/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	repo := repository.NewUserRepository()
	userService := service.NewUserService(repo)
	userHandler := handler.NewUserHandler(userService)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, userHandler)

	reflection.Register(s)

	log.Printf("Server listening on port 50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
