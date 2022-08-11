// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package server

import (
	"github.com/google/wire"
	"github.com/gorilla/mux"
	"restapi-with-opentelemetry/internal/controller"
	"restapi-with-opentelemetry/internal/repository"
	"restapi-with-opentelemetry/internal/router"
	"restapi-with-opentelemetry/internal/service"
	"restapi-with-opentelemetry/internal/store"
)

// Injectors from wire.go:

func InitializedServerRouter() *mux.Router {
	db := store.NewSQLLite()
	iUserRepository := repository.NewUserRepository(db)
	iUserService := service.NewUserService(iUserRepository)
	iUserController := controller.NewUserController(iUserService)
	iAuthService := service.NewAuthService(iUserRepository)
	iAuthController := controller.NewAuthController(iAuthService)
	iArticleRepository := repository.NewArticleRepository(db)
	iArticleService := service.NewArticleService(iArticleRepository)
	iArticleController := controller.NewArticleController(iArticleService)
	muxRouter := router.NewRouter(iUserController, iAuthController, iArticleController)
	return muxRouter
}

// wire.go:

var userSet = wire.NewSet(repository.NewUserRepository, service.NewUserService, controller.NewUserController)

var authSet = wire.NewSet(service.NewAuthService, controller.NewAuthController)

var articleSet = wire.NewSet(repository.NewArticleRepository, service.NewArticleService, controller.NewArticleController)
