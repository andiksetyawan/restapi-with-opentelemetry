package router

import (
	"github.com/gorilla/mux"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gorilla/mux/otelmux"
	"restapi-with-opentelemetry/config"
	"restapi-with-opentelemetry/internal/controller"
)

func NewRouter(userController controller.IUserController) *mux.Router {
	r := mux.NewRouter()

	//otelmux middleware, tracing parentspan with path route
	r.Use(otelmux.Middleware(config.ServiceName))

	v1 := r.PathPrefix("/api/v1").Subrouter()
	{
		NewUserRouter(v1, userController)
		//NewOrderRouter(v1, orderController)
	}

	return r
}
