package controller

import (
	"encoding/json"
	"net/http"

	"github.com/rs/zerolog/log"
	"restapi-with-opentelemetry/internal/entity"
	"restapi-with-opentelemetry/internal/model"
	"restapi-with-opentelemetry/internal/service"
	"restapi-with-opentelemetry/pkg/web"
)

type IUserController interface {
	Create(w http.ResponseWriter, r *http.Request)
}

type userController struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) IUserController {
	return &userController{userService: userService}
}

func (u *userController) Create(w http.ResponseWriter, r *http.Request) {
	var userCreateRequest model.UserCreateRequest
	err := json.NewDecoder(r.Body).Decode(&userCreateRequest)
	if err != nil {
		log.Error().Err(err).Msg("failed to parsing body payload")
		web.RespondWithError(w, http.StatusInternalServerError, "failed to parsing body payload")
		return
	}
	newUser := entity.User{Name: userCreateRequest.Name}
	createdUser, err := u.userService.CreateUser(r.Context(), &newUser)
	if err != nil {
		log.Error().Err(err).Msg("failed to create user")
		web.RespondWithError(w, http.StatusInternalServerError, "failed to create user")
		return
	}
	web.RespondWithJSON(w, http.StatusOK, model.ApiResponse{
		Error:   false,
		Message: "OK",
		Data: model.UserCreateResponse{
			ID:        createdUser.ID,
			Name:      createdUser.Name,
			CreatedAt: createdUser.CreatedAt,
		},
	})
}
