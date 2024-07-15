package service

import (
	"errors"
)

type User struct {
	ID       string
	Username string
	Password string
	Email    string
}

type UserRepository interface {
	CreateUser(username, password, email string) (*User, error)
	FindUserByUsername(username string) (*User, error)
	FindUserByID(userID string) (*User, error)
}

type UserService interface {
	Register(username, password, email string) error
	Login(username, password string) (string, error)
	GetProfile(userID string) (*User, error)
}

type userService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) Register(username, password, email string) error {
	_, err := s.repo.CreateUser(username, password, email)
	if err != nil {
		return err
	}
	return nil
}

func (s *userService) Login(username, password string) (string, error) {
	user, err := s.repo.FindUserByUsername(username)
	if err != nil || user.Password != password {
		return "", errors.New("invalid username or password")
	}
	token := "token-" + user.ID
	return token, nil
}

func (s *userService) GetProfile(userID string) (*User, error) {
	user, err := s.repo.FindUserByID(userID)
	if err != nil {
		return nil, err
	}
	return user, nil
}
	