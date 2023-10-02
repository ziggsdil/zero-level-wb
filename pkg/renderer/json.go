package renderer

import (
	"encoding/json"
	"net/http"
)

type Renderer struct {
}

func (r Renderer) RenderJSON(w http.ResponseWriter, response interface{}) {
	w.WriteHeader(http.StatusOK)
	encoder := json.NewEncoder(w)
	_ = encoder.Encode(response)
}

func (r Renderer) RenderOK(w http.ResponseWriter) {
	w.WriteHeader(http.StatusOK)
}

func (r Renderer) RenderError(w http.ResponseWriter, err error) {
	status := http.StatusInternalServerError

	w.WriteHeader(status)
	encoder := json.NewEncoder(w)
	_ = encoder.Encode(ErrorResponse{
		Error: err.Error(),
	})
}

type ErrorResponse struct {
	Error string `json:"error"`
}
