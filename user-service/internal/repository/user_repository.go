package repository

import (
	"errors"

	"github.com/google/uuid"
	"github.com/nuhmanudheent/online-learning-platform/user-service/internal/service"
)

type UserRepository interface {
	CreateUser(username, password, email string) (*service.User, error)
	FindUserByUsername(username string) (*service.User, error)
	FindUserByID(userID string) (*service.User, error)
}

type userRepository struct {
	users map[string]*service.User
}

func NewUserRepository() UserRepository {
	return &userRepository{users: make(map[string]*service.User)}
}

func (r *userRepository) CreateUser(username, password, email string) (*service.User, error) {
	for _, user := range r.users {
		if user.Username == username {
			return nil, errors.New("username already exists")
		}
	}
	userID := uuid.New().String()
	user := &service.User{
		ID:       userID,
		Username: username,
		Password: password, 
		Email:    email,
	}
	r.users[userID] = user
	return user, nil
}

func (r *userRepository) FindUserByUsername(username string) (*service.User, error) {
	for _, user := range r.users {
		if user.Username == username {
			return user, nil
		}
	}
	return nil, errors.New("user not found")
}

func (r *userRepository) FindUserByID(userID string) (*service.User, error) {
	user, exists := r.users[userID]
	if !exists {
		return nil, errors.New("user not found")
	}
	return user, nil
}
