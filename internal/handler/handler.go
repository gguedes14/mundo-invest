package handler

import (
	"github.com/gguedes14/mundo-invest/internal/controller"
)

type Handler struct {
	Controller *controller.Controller
}

func NewHandler(c *controller.Controller) *Handler {
	return &Handler{Controller: c}
}
