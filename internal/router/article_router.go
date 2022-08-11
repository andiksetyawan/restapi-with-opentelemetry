package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"restapi-with-opentelemetry/internal/controller"
)

const articleRoute = "/article"

//NewArticleRouter V1
func NewArticleRouter(r *mux.Router, articleController controller.IArticleController) *mux.Router {
	r.Handle(articleRoute, otelhttp.NewHandler(http.HandlerFunc(articleController.Create), "handler.article.Create")).Methods(http.MethodPost)
	r.Handle(articleRoute+"/{slug}", otelhttp.NewHandler(http.HandlerFunc(articleController.GetBySlug), "handler.article.GetBySlug")).Methods(http.MethodGet)
	return r
}
