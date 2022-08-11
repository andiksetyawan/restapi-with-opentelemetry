package repository

//repository helper
//helper for usecase layer
//TODO with transaction

import (
	"context"

	"gorm.io/gorm"
	"restapi-with-opentelemetry/internal/entity"
)

type IArticleRepository interface {
	CreateArticle(ctx context.Context, article *entity.Article) (*entity.Article, error)
	GetBySlugArticle(ctx context.Context, slug *string) (*entity.Article, error)
}

type articleRepository struct {
	Db *gorm.DB
}

func NewArticleRepository(db *gorm.DB) IArticleRepository {
	return &articleRepository{
		Db: db,
	}
}

func (u *articleRepository) CreateArticle(ctx context.Context, article *entity.Article) (*entity.Article, error) {
	if err := u.Db.WithContext(ctx).Create(article).Error; err != nil {
		return nil, err
	}
	return article, nil
}

func (u *articleRepository) GetBySlugArticle(ctx context.Context, slug *string) (*entity.Article, error) {
	var article entity.Article
	if err := u.Db.WithContext(ctx).Joins("User").Where(&entity.Article{Slug: *slug}).First(&article).Error; err != nil {
		return nil, err
	}
	return &article, nil
}
