package handler

import (
	"context"

	"github.com/nuhmanudheent/online-learning-platform/enrollment-service/internal/service"
	pb "github.com/nuhmanudheent/online-learning-platform/enrollment-service/proto"
)

type EnrollmentHandler struct {
	pb.UnimplementedEnrollmentServiceServer
	Service service.EnrollmentService
}

func (h *EnrollmentHandler) EnrollUser(ctx context.Context, req *pb.EnrollRequest) (*pb.EnrollResponse, error) {
	err := h.Service.EnrollUser(req.UserId, req.CourseId)
	if err != nil {
		return nil, err
	}
	return &pb.EnrollResponse{Message: "User enrolled successfully"}, nil
}

func (h *EnrollmentHandler) GetEnrollmentsByUser(ctx context.Context, req *pb.UserIDRequest) (*pb.EnrollmentsResponse, error) {
	enrollments, err := h.Service.GetEnrollmentsByUser(req.UserId)
	if err != nil {
		return nil, err
	}

	pbEnrollments := make([]*pb.Enrollment, len(enrollments))
	for i, enrollment := range enrollments {
		pbEnrollments[i] = &pb.Enrollment{
			UserId:   enrollment.UserID,
			CourseId: enrollment.CourseID,
		}
	}

	return &pb.EnrollmentsResponse{Enrollments: pbEnrollments}, nil
}

func (h *EnrollmentHandler) GetEnrollmentsByCourse(ctx context.Context, req *pb.CourseIDRequest) (*pb.EnrollmentsResponse, error) {
	enrollments, err := h.Service.GetEnrollmentsByCourse(req.CourseId)
	if err != nil {
		return nil, err
	}

	pbEnrollments := make([]*pb.Enrollment, len(enrollments))
	for i, enrollment := range enrollments {
		pbEnrollments[i] = &pb.Enrollment{
			UserId:   enrollment.UserID,
			CourseId: enrollment.CourseID,
		}
	}

	return &pb.EnrollmentsResponse{Enrollments: pbEnrollments}, nil
}
