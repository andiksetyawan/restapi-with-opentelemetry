package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"restapi-with-opentelemetry/internal/controller"
)

const userRoute = "/user"

//NewUserRouter V1
func NewUserRouter(r *mux.Router, userController controller.IUserController) {
	r.Handle(userRoute, otelhttp.NewHandler(http.HandlerFunc(userController.Create), "handler.user.Create")).Methods(http.MethodPost)
}
