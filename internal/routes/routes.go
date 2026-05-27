package routes

import (
	"net/http"

	"github.com/gguedes14/mundo-invest/internal/handler"
	handlerclient "github.com/gguedes14/mundo-invest/internal/handler/client"
	"github.com/go-chi/chi/v5"
)

func ClientRoutes(r chi.Router, h *handler.Handler) {
	r.Post("/clientes", func(w http.ResponseWriter, r *http.Request) {
		handlerclient.CreateClient(w, r, h)
	})
}
