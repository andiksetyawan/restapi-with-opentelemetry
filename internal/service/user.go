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

type IUserService interface {
	CreateUser(ctx context.Context, user *entity.User) (*entity.User, error)
	FindUserByID(ctx context.Context, ID *uint) (*entity.User, error)
}

type userService struct {
	repo repository.IUserRepository
}

//NewUserService create new instance user service
func NewUserService(repo repository.IUserRepository) IUserService {
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

//FindUserByID service
func (u *userService) FindUserByID(ctx context.Context, ID *uint) (*entity.User, error) {
	//service tracer
	_, span := otel.Tracer(config.ServiceName).Start(ctx, "service.user.FindUserByID")
	defer span.End()

	user, err := u.repo.FindUserByID(ctx, ID)
	if err != nil {
		return nil, err
	}

	return user, nil
}
