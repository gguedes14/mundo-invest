package event

import (
	"encoding/json"
	"net/http"

	controllerEvent "github.com/gguedes14/mundo-invest/internal/controller/event"
	"github.com/gguedes14/mundo-invest/internal/dto"
	"github.com/gguedes14/mundo-invest/internal/handler"
)

func CrateEvent(w http.ResponseWriter, r *http.Request, h *handler.Handler) {
	var input dto.EventInput
	var clientID dto.ClientResponse
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "invalid input"})
		return
	}

	response, err := controllerEvent.CreateEvent(r.Context(), h.Controller, input, clientID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)

}
