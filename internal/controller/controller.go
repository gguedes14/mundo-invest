package controller

import "github.com/gguedes14/mundo-invest/internal/service"

type Controller struct {
	Service *service.Service
}

func NewController(s *service.Service) *Controller {
	return &Controller{Service: s}
}
