package web

import (
	"encoding/json"
	"net/http"

	"restapi-with-opentelemetry/internal/model"
)

func RespondWithError(w http.ResponseWriter, code int, message string) {
	RespondWithJSON(w, code, model.ApiResponse{
		Error:   true,
		Message: message,
		Data:    nil,
	})
}

func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
