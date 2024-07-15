package handler

import (
	"context"

	"github.com/nuhmanudheent/online-learning-platform/course-service/internal/service"
	pb "github.com/nuhmanudheent/online-learning-platform/course-service/proto"
)

type CourseHandler struct {
	pb.UnimplementedCourseServiceServer
	Service service.CourseService
}

func NewCourseHandler(svc service.CourseService) *CourseHandler {
	return &CourseHandler{Service: svc}
}

func (h *CourseHandler) CreateCourse(ctx context.Context, req *pb.CreateCourseRequest) (*pb.CreateCourseResponse, error) {
	course, err := h.Service.CreateCourse(req.Title, req.Description)
	if err != nil {
		return nil, err
	}
	return &pb.CreateCourseResponse{
		Course: &pb.Course{
			Id:          course.ID,
			Title:       course.Title,
			Description: course.Description,
		},
	}, nil
}

func (h *CourseHandler) GetCourse(ctx context.Context, req *pb.GetCourseRequest) (*pb.GetCourseResponse, error) {
	course, err := h.Service.GetCourse(req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.GetCourseResponse{
		Course: &pb.Course{
			Id:          course.ID,
			Title:       course.Title,
			Description: course.Description,
		},
	}, nil
}

func (h *CourseHandler) ListCourses(ctx context.Context, req *pb.ListCoursesRequest) (*pb.ListCoursesResponse, error) {
	courses, err := h.Service.ListCourses()
	if err != nil {
		return nil, err
	}
	var pbCourses []*pb.Course
	for _, course := range courses {
		pbCourses = append(pbCourses, &pb.Course{
			Id:          course.ID,
			Title:       course.Title,
			Description: course.Description,
		})
	}
	return &pb.ListCoursesResponse{Courses: pbCourses}, nil
}
