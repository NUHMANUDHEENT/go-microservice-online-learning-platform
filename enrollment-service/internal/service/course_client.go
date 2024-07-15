package service

import (
	"context"
	"log"

	coursepb "github.com/nuhmanudheent/online-learning-platform/enrollment-service/proto"
	"google.golang.org/grpc"
)

type Course struct {
	ID          string
	Title       string
	Description string
}

type CourseClient interface {
	GetCourse(courseID string) (*Course, error)
}

type courseClient struct {
	client coursepb.CourseServiceClient
}

func NewCourseClient(conn *grpc.ClientConn) CourseClient {
	return &courseClient{client: coursepb.NewCourseServiceClient(conn)}
}

func (c *courseClient) GetCourse(courseID string) (*Course, error) {
	req := &coursepb.GetCourseRequest{Id: courseID}
	res, err := c.client.GetCourse(context.Background(), req)
	if err != nil {
		log.Printf("Error calling GetCourse: %v", err)
		return nil, err
	}
	return &Course{
		ID:          res.Course.Id,
		Title:       res.Course.Title,
		Description: res.Course.Description,
	}, nil
}
