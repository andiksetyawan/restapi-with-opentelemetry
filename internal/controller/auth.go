package controller

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/rs/zerolog/log"
	"restapi-with-opentelemetry/internal/model"
	"restapi-with-opentelemetry/internal/service"
	"restapi-with-opentelemetry/pkg/web"
)

type IAuthController interface {
	Login(w http.ResponseWriter, r *http.Request)
	Signup(w http.ResponseWriter, r *http.Request)
	Authorize(next http.Handler) http.Handler
}

type authController struct {
	authService service.IAuthService
}

func NewAuthController(authService service.IAuthService) IAuthController {
	return &authController{authService: authService}
}

func (a *authController) Login(w http.ResponseWriter, r *http.Request) {
	var authLoginRequest model.AuthLoginRequest
	err := json.NewDecoder(r.Body).Decode(&authLoginRequest)
	if err != nil {
		log.Error().Err(err).Msg("failed to parsing body payload")
		web.RespondWithError(w, http.StatusInternalServerError, "failed to parsing body payload")
		return
	}
	log.Debug().Msg(authLoginRequest.Email)
	token, err := a.authService.Login(r.Context(), &authLoginRequest)
	if err != nil {
		log.Error().Err(err).Msg("failed to authorize user")
		web.RespondWithError(w, http.StatusUnauthorized, "failed to authorize user")
		return
	}
	web.RespondWithJSON(w, http.StatusOK, model.ApiResponse{
		Error:   false,
		Message: "OK",
		Data:    model.AuthLoginSignupResponse{Token: token},
	})
}

func (a *authController) Signup(w http.ResponseWriter, r *http.Request) {
	var authSignupRequest model.AuthSignupRequest
	err := json.NewDecoder(r.Body).Decode(&authSignupRequest)
	if err != nil {
		log.Error().Err(err).Msg("failed to parsing body payload")
		web.RespondWithError(w, http.StatusInternalServerError, "failed to parsing body payload")
		return
	}

	token, err := a.authService.SignUp(r.Context(), &authSignupRequest)
	if err != nil {
		log.Error().Err(err).Msg("failed to signup user")
		web.RespondWithError(w, http.StatusUnauthorized, "failed to signup user")
		return
	}
	web.RespondWithJSON(w, http.StatusOK, model.ApiResponse{
		Error:   false,
		Message: "OK",
		Data:    model.AuthLoginSignupResponse{Token: token},
	})
}

//Authorize this is authorize controller for middleware function
func (a *authController) Authorize(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenBearerString := r.Header.Get("Authorization")
		tokenString := strings.Replace(tokenBearerString, "Bearer ", "", 1)
		user, err := a.authService.Authorize(r.Context(), &tokenString)
		if err != nil {
			web.RespondWithError(w, http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized))
			return
		}

		ctx := context.WithValue(r.Context(), "user", user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
