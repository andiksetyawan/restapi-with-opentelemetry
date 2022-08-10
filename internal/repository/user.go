package repository

//repository helper
//helper for usecase layer
//TODO with transaction

import (
	"context"

	"gorm.io/gorm"
	"restapi-with-opentelemetry/internal/entity"
)

type IUserRepository interface {
	CreateUser(ctx context.Context, user *entity.User) (*entity.User, error)
	FindUserByEmail(ctx context.Context, email *string) (*entity.User, error)
	FindUserByID(ctx context.Context, ID *uint) (*entity.User, error)
}

type userRepository struct {
	Db *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &userRepository{
		Db: db,
	}
}

func (u *userRepository) CreateUser(ctx context.Context, user *entity.User) (*entity.User, error) {
	if err := u.Db.WithContext(ctx).Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (u *userRepository) FindUserByEmail(ctx context.Context, email *string) (*entity.User, error) {
	var user entity.User
	if err := u.Db.WithContext(ctx).Where(&entity.User{Email: *email}).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *userRepository) FindUserByID(ctx context.Context, ID *uint) (*entity.User, error) {
	var user entity.User
	if err := u.Db.WithContext(ctx).Where(&entity.User{ID: *ID}).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
