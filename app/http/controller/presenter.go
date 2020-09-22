package controller

import (
	"io"
	"net/http"
)

func RenderJSON(status int, w http.ResponseWriter, data []byte) {
	w.WriteHeader(status)
	io.WriteString(w, string(data))
}
