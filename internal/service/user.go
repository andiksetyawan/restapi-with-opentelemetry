package service

//usecase or bussines logic layer
//TODO : add metrics counter method

import (
	"context"

	"go.opentelemetry.io/otel"
	"restapi-with-opentelemetry/config"
	"restapi-with-opentelemetry/internal/entity"
	"restapi-with-opentelemetry/internal/repository"
)

type UserService interface {
	CreateUser(ctx context.Context, user *entity.User) (*entity.User, error)
}

type userService struct {
	repo repository.IUserRepository
}

//NewUserService create new instance user service
func NewUserService(repo repository.IUserRepository) UserService {
	return &userService{repo: repo}
}

//CreateUser service
func (u *userService) CreateUser(ctx context.Context, user *entity.User) (*entity.User, error) {
	//service tracer
	_, span := otel.Tracer(config.ServiceName).Start(ctx, "service.user.CreateUser")
	defer span.End()

	createdUser, err := u.repo.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	return createdUser, nil
}
