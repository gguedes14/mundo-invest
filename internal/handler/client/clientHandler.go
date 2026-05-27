package client

import (
	"encoding/json"
	"net/http"

	controllerclient "github.com/gguedes14/mundo-invest/internal/controller/client"
	"github.com/gguedes14/mundo-invest/internal/dto"
	"github.com/gguedes14/mundo-invest/internal/handler"
)

func CreateClient(w http.ResponseWriter, r *http.Request, h *handler.Handler) {
	var input dto.ClientInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "invalid input"})
		return
	}

	response, err := controllerclient.CreateClient(r.Context(), h.Controller, input)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}
