package repository

import (
	"errors"

	"github.com/google/uuid"
	"github.com/nuhmanudheent/online-learning-platform/course-service/internal/service"
)

type CourseRepository interface {
	CreateCourse(title, description string) (*service.Course, error)
	FindCourseByID(courseID string) (*service.Course, error)
	ListCourses() ([]*service.Course, error)
}

type courseRepository struct {
	courses map[string]*service.Course
}

func NewCourseRepository() CourseRepository {
	return &courseRepository{courses: make(map[string]*service.Course)}
}

func (r *courseRepository) CreateCourse(title, description string) (*service.Course, error) {
	courseID := uuid.New().String()
	course := &service.Course{
		ID:          courseID,
		Title:       title,
		Description: description,
	}
	r.courses[courseID] = course
	return course, nil
}

func (r *courseRepository) FindCourseByID(courseID string) (*service.Course, error) {
	course, exists := r.courses[courseID]
	if !exists {
		return nil, errors.New("course not found")
	}
	return course, nil
}

func (r *courseRepository) ListCourses() ([]*service.Course, error) {
	var courses []*service.Course
	for _, course := range r.courses {
		courses = append(courses, course)
	}
	return courses, nil
}
