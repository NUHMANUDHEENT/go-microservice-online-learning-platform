package main

import (
	"log"
	"net"

	"github.com/nuhmanudheent/online-learning-platform/course-service/internal/handler"
	"github.com/nuhmanudheent/online-learning-platform/course-service/internal/repository"
	"github.com/nuhmanudheent/online-learning-platform/course-service/internal/service"
	pb "github.com/nuhmanudheent/online-learning-platform/course-service/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	repo := repository.NewCourseRepository()
	courseService := service.NewCourseService(repo)
	courseHandler := handler.NewCourseHandler(courseService)

	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterCourseServiceServer(s, courseHandler)

	reflection.Register(s)

	log.Printf("Server listening on port 50052")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
