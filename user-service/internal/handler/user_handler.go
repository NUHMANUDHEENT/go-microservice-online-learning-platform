package handler

import (
	"context"

	"github.com/nuhmanudheent/online-learning-platform/user-service/internal/service"
	pb "github.com/nuhmanudheent/online-learning-platform/user-service/proto"
)

type UserHandler struct {
	pb.UnimplementedUserServiceServer
	Service service.UserService
}

func NewUserHandler(svc service.UserService) *UserHandler {
	return &UserHandler{Service: svc}
}

func (h *UserHandler) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	err := h.Service.Register(req.Username, req.Password, req.Email)
	if err != nil {
		return nil, err
	}
	return &pb.RegisterResponse{Message: "User registered successfully"}, nil
}

func (h *UserHandler) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	token, err := h.Service.Login(req.Username, req.Password)
	if err != nil {
		return nil, err
	}
	return &pb.LoginResponse{Token: token}, nil
}

func (h *UserHandler) GetProfile(ctx context.Context, req *pb.ProfileRequest) (*pb.ProfileResponse, error) {
	user, err := h.Service.GetProfile(req.UserId)
	if err != nil {
		return nil, err
	}
	return &pb.ProfileResponse{
		UserId:   user.ID,
		Username: user.Username,
		Email:    user.Email,
	}, nil
}
