package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gorilla/mux/otelmux"
	"restapi-with-opentelemetry/config"
	"restapi-with-opentelemetry/internal/controller"
)

func NewRouter(userController controller.IUserController, authController controller.IAuthController) *mux.Router {
	r := mux.NewRouter()

	//otelmux middleware, tracing parentspan with path route
	r.Use(otelmux.Middleware(config.ServiceName))
	r.Use(mux.CORSMethodMiddleware(r))

	//TODO add handler: Liveness, Readiness k8s
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("i'm a live"))
	})

	v1 := r.PathPrefix("/api/v1").Subrouter()
	{
		NewAuthRouter(v1, authController)
		privateV1 := v1.PathPrefix("/").Subrouter()
		{
			privateV1.Use(authController.Authorize)
			NewUserRouter(privateV1, userController)
		}
	}

	return r
}
