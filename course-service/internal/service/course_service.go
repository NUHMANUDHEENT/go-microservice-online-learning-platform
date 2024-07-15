package service

type Course struct {
	ID          string
	Title       string
	Description string
}

type CourseRepository interface {
	CreateCourse(title, description string) (*Course, error)
	FindCourseByID(courseID string) (*Course, error)
	ListCourses() ([]*Course, error)
}

type CourseService interface {
	CreateCourse(title, description string) (*Course, error)
	GetCourse(courseID string) (*Course, error)
	ListCourses() ([]*Course, error)
}

type courseService struct {
	repo CourseRepository
}

func NewCourseService(repo CourseRepository) CourseService {
	return &courseService{repo: repo}
}

func (s *courseService) CreateCourse(title, description string) (*Course, error) {
	return s.repo.CreateCourse(title, description)
}

func (s *courseService) GetCourse(courseID string) (*Course, error) {
	return s.repo.FindCourseByID(courseID)
}

func (s *courseService) ListCourses() ([]*Course, error) {
	return s.repo.ListCourses()
}
