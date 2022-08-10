//go:build wireinject
// +build wireinject

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

var userSet = wire.NewSet(
	repository.NewUserRepository,
	service.NewUserService,
	controller.NewUserController,
)

func InitializedServerRouter() *mux.Router {
	wire.Build(
		store.NewSQLLite,
		userSet,
		router.NewRouter,
	)
	return nil
}
