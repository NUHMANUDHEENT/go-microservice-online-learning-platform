package repository

import (
	"errors"

	"github.com/nuhmanudheent/online-learning-platform/enrollment-service/internal/service"
)

type EnrollmentRepository interface {
	EnrollUser(userID, courseID string) (*service.Enrollment, error)
	FindEnrollmentsByUser(userID string) ([]*service.Enrollment, error)
	FindEnrollmentsByCourse(courseID string) ([]*service.Enrollment, error)
}

type enrollmentRepository struct {
	enrollments map[string]*service.Enrollment
}

func NewEnrollmentRepository() EnrollmentRepository {
	return &enrollmentRepository{enrollments: make(map[string]*service.Enrollment)}
}

func (r *enrollmentRepository) EnrollUser(userID, courseID string) (*service.Enrollment, error) {
	for _, enrollment := range r.enrollments {
		if enrollment.UserID == userID && enrollment.CourseID == courseID {
			return nil, errors.New("user already enrolled in this course")
		}
	}
	enrollment := &service.Enrollment{
		UserID:   userID,
		CourseID: courseID,
	}
	r.enrollments[userID+courseID] = enrollment
	return enrollment, nil
}

func (r *enrollmentRepository) FindEnrollmentsByUser(userID string) ([]*service.Enrollment, error) {
	var enrollments []*service.Enrollment
	for _, enrollment := range r.enrollments {
		if enrollment.UserID == userID {
			enrollments = append(enrollments, enrollment)
		}
	}
	return enrollments, nil
}

func (r *enrollmentRepository) FindEnrollmentsByCourse(courseID string) ([]*service.Enrollment, error) {
	var enrollments []*service.Enrollment
	for _, enrollment := range r.enrollments {
		if enrollment.CourseID == courseID {
			enrollments = append(enrollments, enrollment)
		}
	}
	return enrollments, nil
}
