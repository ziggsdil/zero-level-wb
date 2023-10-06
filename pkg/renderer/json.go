package renderer

import (
	"encoding/json"
	"github.com/ziggsdil/zero-level-wb/pkg/models"
	"net/http"
)

type Renderer struct {
}

func (r Renderer) RenderJSON(w http.ResponseWriter, response interface{}) {
	var dataModel models.Message

	val, ok := response.([]byte)
	if !ok {
		return
	}
	err := json.Unmarshal(val, &dataModel)
	if err != nil {
		http.Error(w, "failed to unmarshal", http.StatusInternalServerError)
		return
	}

	encoder := json.NewEncoder(w)
	_ = encoder.Encode(dataModel)
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
