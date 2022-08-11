package controller

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
	"restapi-with-opentelemetry/internal/entity"
	"restapi-with-opentelemetry/internal/model"
	"restapi-with-opentelemetry/internal/service"
	"restapi-with-opentelemetry/pkg/web"
)

type IArticleController interface {
	Create(w http.ResponseWriter, r *http.Request)
	GetBySlug(w http.ResponseWriter, r *http.Request)
}

type articleController struct {
	articleService service.IArticleService
}

func NewArticleController(userService service.IArticleService) IArticleController {
	return &articleController{articleService: userService}
}

func (a articleController) Create(w http.ResponseWriter, r *http.Request) {
	var articleCreateRequest model.ArticleCreateRequest
	err := json.NewDecoder(r.Body).Decode(&articleCreateRequest)
	if err != nil {
		log.Error().Err(err).Msg("failed to parsing body payload")
		web.RespondWithError(w, http.StatusInternalServerError, "failed to parsing body payload")
		return
	}

	user, ok := r.Context().Value("user").(*entity.User)
	if !ok {
		log.Error().Err(err).Msg(http.StatusText(http.StatusInternalServerError))
		web.RespondWithError(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	newUser := entity.Article{UserID: user.ID, Title: articleCreateRequest.Title, Content: articleCreateRequest.Content}
	_, err = a.articleService.CreateArticle(r.Context(), &newUser)
	if err != nil {
		log.Error().Err(err).Msg("failed to create new article")
		web.RespondWithError(w, http.StatusInternalServerError, "failed to create new article")
		return
	}
	web.RespondWithJSON(w, http.StatusOK, model.ApiResponse{
		Error:   false,
		Message: "OK",
		Data:    nil,
	})
}

func (a articleController) GetBySlug(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	slug := vars["slug"]

	article, err := a.articleService.GetBySlugArticle(r.Context(), &slug)
	if err != nil {
		log.Error().Err(err).Msg("failed to get article")
		web.RespondWithError(w, http.StatusInternalServerError, "failed to get article")
		return
	}

	web.RespondWithJSON(w, http.StatusOK, model.ApiResponse{
		Error:   false,
		Message: "OK",
		Data:    article,
	})
}
