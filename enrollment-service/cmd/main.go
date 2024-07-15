package main

import (
	"log"
	"net"

	"github.com/nuhmanudheent/online-learning-platform/enrollment-service/internal/handler"
	"github.com/nuhmanudheent/online-learning-platform/enrollment-service/internal/repository"
	"github.com/nuhmanudheent/online-learning-platform/enrollment-service/internal/service"
	pb "github.com/nuhmanudheent/online-learning-platform/enrollment-service/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	userServiceAddress   = "user-service:50051"
	courseServiceAddress = "course-service:50052" 
)

func main() {
	lis, err := net.Listen("tcp", ":50053")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()

	connUser, err := grpc.Dial(userServiceAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to user service: %v", err)
	}
	defer connUser.Close()

	connCourse, err := grpc.Dial(courseServiceAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to course service: %v", err)
	}
	defer connCourse.Close()

	userClient := service.NewUserClient(connUser)
	courseClient := service.NewCourseClient(connCourse)
	repo := repository.NewEnrollmentRepository()
	enrollmentService := service.NewEnrollmentService(repo, userClient, courseClient)

	handler := handler.EnrollmentHandler{
		Service: enrollmentService,
	}

	pb.RegisterEnrollmentServiceServer(s, &handler)
	reflection.Register(s)

	log.Println("Enrollment Service is running on port 50053...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
