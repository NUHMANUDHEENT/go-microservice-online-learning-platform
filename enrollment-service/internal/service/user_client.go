package service

import (
	"context"
	"log"

	userpb "github.com/nuhmanudheent/online-learning-platform/enrollment-service/proto"
	"google.golang.org/grpc"
)

type User struct {
	ID       string
	Username string
	Email    string
}

type UserClient interface {
	GetUser(userID string) (*User, error)
}

type userClient struct {
	client userpb.UserServiceClient
}

func NewUserClient(conn *grpc.ClientConn) UserClient {
	return &userClient{client: userpb.NewUserServiceClient(conn)}
}

func (c *userClient) GetUser(userID string) (*User, error) {
	req := &userpb.ProfileRequest{UserId: userID}
	res, err := c.client.GetProfile(context.Background(), req)
	if err != nil {
		log.Printf("Error calling GetProfile: %v", err)
		return nil, err
	}
	return &User{
		ID:       res.UserId,
		Username: res.Username,
		Email:    res.Email,
	}, nil
}
