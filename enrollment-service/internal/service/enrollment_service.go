package service

import (
	"errors"
)

type Enrollment struct {
	UserID   string
	CourseID string
}

type EnrollmentRepository interface {
	EnrollUser(userID, courseID string) (*Enrollment, error)
	FindEnrollmentsByUser(userID string) ([]*Enrollment, error)
	FindEnrollmentsByCourse(courseID string) ([]*Enrollment, error)
}

type EnrollmentService interface {
	EnrollUser(userID, courseID string) error
	GetEnrollmentsByUser(userID string) ([]*Enrollment, error)
	GetEnrollmentsByCourse(courseID string) ([]*Enrollment, error)
}

type enrollmentService struct {
	repo         EnrollmentRepository
	userClient   UserClient
	courseClient CourseClient
}

func NewEnrollmentService(repo EnrollmentRepository, userClient UserClient, courseClient CourseClient) EnrollmentService {
	return &enrollmentService{
		repo:         repo,
		userClient:   userClient,
		courseClient: courseClient,
	}
}

func (s *enrollmentService) EnrollUser(userID, courseID string) error {
	_, err := s.userClient.GetUser(userID)
	if err != nil {
		return errors.New("user does not exist")
	}

	_, err = s.courseClient.GetCourse(courseID)
	if err != nil {
		return errors.New("course does not exist")
	}

	_, err = s.repo.EnrollUser(userID, courseID)
	if err != nil {
		return err
	}

	return nil
}

func (s *enrollmentService) GetEnrollmentsByUser(userID string) ([]*Enrollment, error) {
	return s.repo.FindEnrollmentsByUser(userID)
}

func (s *enrollmentService) GetEnrollmentsByCourse(courseID string) ([]*Enrollment, error) {
	return s.repo.FindEnrollmentsByCourse(courseID)
}
