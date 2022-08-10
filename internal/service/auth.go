package service

import (
	"context"
	"errors"

	"go.opentelemetry.io/otel"
	"restapi-with-opentelemetry/config"
	"restapi-with-opentelemetry/internal/entity"
	"restapi-with-opentelemetry/internal/model"
	"restapi-with-opentelemetry/internal/repository"
	"restapi-with-opentelemetry/pkg/helper"
	"restapi-with-opentelemetry/pkg/token"
)

type IAuthService interface {
	Login(ctx context.Context, authRequest *model.AuthLoginRequest) (string, error)
	SignUp(ctx context.Context, reqUser *model.AuthSignupRequest) (string, error)
	Authorize(ctx context.Context, jwtToken *string) (*entity.User, error)
}

type authService struct {
	userRepository repository.IUserRepository
}

func NewAuthService(userRepository repository.IUserRepository) IAuthService {
	return &authService{userRepository: userRepository}
}

func (a *authService) Login(ctx context.Context, authRequest *model.AuthLoginRequest) (string, error) {
	//service tracer
	_, span := otel.Tracer(config.ServiceName).Start(ctx, "service.auth.Login")
	defer span.End()

	user, err := a.userRepository.FindUserByEmail(ctx, &authRequest.Email)
	if err != nil {
		return "", err
	}

	//compare password and hash
	if err := token.HashIsMatch(authRequest.Password, user.Password); err != nil {
		return "", err
	}

	//generate token
	jwtToken, err := token.BuildJwtToken(helper.UintToString(user.ID), user.Email)
	if err != nil {
		return "", err
	}
	return jwtToken, nil
}

func (a *authService) SignUp(ctx context.Context, reqUser *model.AuthSignupRequest) (string, error) {
	//service tracer
	_, span := otel.Tracer(config.ServiceName).Start(ctx, "service.auth.Signup")
	defer span.End()

	hashPassword, err := token.HashBuilder(reqUser.Password)
	if err != nil {
		return "", err
	}

	newUser := entity.User{
		Name:     reqUser.Name,
		Email:    reqUser.Email,
		Password: hashPassword,
	}
	createdUser, err := a.userRepository.CreateUser(ctx, &newUser)
	if err != nil {
		return "", err
	}

	jwtToken, err := token.BuildJwtToken(helper.UintToString(createdUser.ID), createdUser.Email)
	if err != nil {
		return "", err
	}

	return jwtToken, nil
}

func (a *authService) Authorize(ctx context.Context, jwtToken *string) (*entity.User, error) {
	_, span := otel.Tracer(config.ServiceName).Start(ctx, "service.auth.Authorize")
	defer span.End()

	//TODO apakah akses helper harus melalui injeksi dari controller dgn menambah instance di struct authService?.
	claims, err := token.JwtTokenIsValid(*jwtToken)
	if err != nil {
		return nil, err
	}

	//TODO claim set to model struct
	IDString, ok := claims["id"].(string)
	if !ok || IDString == "" {
		return nil, errors.New("claim not valid")
	}

	uintID, err := helper.StringToUint(IDString)
	if err != nil {
		return nil, err
	}

	user, err := a.userRepository.FindUserByID(ctx, &uintID)
	if err != nil {
		return nil, err
	}

	//TODO create specific model for return
	user.Password = ""
	return user, nil

}
