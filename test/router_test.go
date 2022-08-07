package test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"restapi-with-opentelemetry/internal/controller"
	"restapi-with-opentelemetry/internal/model"
	"restapi-with-opentelemetry/internal/repository"
	"restapi-with-opentelemetry/internal/router"
	"restapi-with-opentelemetry/internal/service"
	"restapi-with-opentelemetry/internal/store"
)

func setupRouterTest() *mux.Router {
	db := store.NewPostgreeSQLDbTest().Connect()
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userController := controller.NewUserController(userService)
	router := router.NewRouter(userController)
	return router
}

//TestCreateUser integration test
func TestCreateUserSuccess(t *testing.T) {
	//TODO TestTable

	name := "John Due Test"
	r := setupRouterTest()
	w := httptest.NewRecorder()

	reqPayload := strings.NewReader("{\"name\":\"" + name + "\"}")
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/user", reqPayload)
	req.Header.Set("Content-Type", "application/json")

	r.ServeHTTP(w, req)

	var response model.ApiResponse
	json.NewDecoder(w.Body).Decode(&response)
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "OK", response.Message)

	createdUser := response.Data.(map[string]interface{})
	assert.Equal(t, name, createdUser["name"])
}
