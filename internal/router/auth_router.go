package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"restapi-with-opentelemetry/internal/controller"
)

const authRoute = "/auth"

//NewAuthRouter V1
func NewAuthRouter(r *mux.Router, authController controller.IAuthController) {
	r.Handle(authRoute+"/login", otelhttp.NewHandler(http.HandlerFunc(authController.Login), "handler.auth.Login")).Methods(http.MethodPost)
	r.Handle(authRoute+"/signup", otelhttp.NewHandler(http.HandlerFunc(authController.Signup), "handler.auth.Signup")).Methods(http.MethodPost)
}
