package controller

import (
	"encoding/json"
	"io"
	"net/http"
)

type payload struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

type payloadError struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

// RenderJSON is unable to take any struct array with declaring []intertace{} due to golang restriction
func RenderJSON(status string, w http.ResponseWriter, jsonStr string) {
	payloadStr := `{"status": "` + status + `", "data": ` + jsonStr + `}`

	var payload payload
	json.Unmarshal([]byte(payloadStr), &payload)

	json, _ := json.Marshal(payload)

	respond(http.StatusOK, w, string(json))
}

func RenderErrorJSON(status string, w http.ResponseWriter, errorMessage string) {
	payloadStr := `{"status": "` + status + `", "message": "` + errorMessage + `"}`

	var payloadError payloadError
	json.Unmarshal([]byte(payloadStr), &payloadError)

	json, _ := json.Marshal(payloadError)

	respond(http.StatusBadRequest, w, string(json))
}

func respond(httpStatus int, w http.ResponseWriter, body string) {
	w.WriteHeader(httpStatus)
	io.WriteString(w, body)
}
