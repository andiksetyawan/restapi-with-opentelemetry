package service

//usecase or bussines logic layer
//TODO : add metrics counter method

import (
	"context"

	"github.com/gosimple/slug"
	"go.opentelemetry.io/otel"
	"restapi-with-opentelemetry/config"
	"restapi-with-opentelemetry/internal/entity"
	"restapi-with-opentelemetry/internal/repository"
)

type IArticleService interface {
	CreateArticle(ctx context.Context, article *entity.Article) (*entity.Article, error)
	GetBySlugArticle(ctx context.Context, slug *string) (*entity.Article, error)
}

type articleService struct {
	repo repository.IArticleRepository
}

//NewArticleService create new instance article service
func NewArticleService(repo repository.IArticleRepository) IArticleService {
	return &articleService{repo: repo}
}

func (a *articleService) CreateArticle(ctx context.Context, article *entity.Article) (*entity.Article, error) {
	_, span := otel.Tracer(config.ServiceName).Start(ctx, "service.article.CreateArticle")
	defer span.End()

	article.Slug = slug.Make(article.Title)
	createdUser, err := a.repo.CreateArticle(ctx, article)
	if err != nil {
		return nil, err
	}

	return createdUser, nil
}

func (a *articleService) GetBySlugArticle(ctx context.Context, slug *string) (*entity.Article, error) {
	_, span := otel.Tracer(config.ServiceName).Start(ctx, "service.article.GetBySlugArticle")
	defer span.End()

	article, err := a.repo.GetBySlugArticle(ctx, slug)
	if err != nil {
		return nil, err
	}
	return article, nil
}
