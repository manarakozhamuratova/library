package handler

import (
	"github.com/manarakozhamuratova/one-lab-task2/config"
	"github.com/manarakozhamuratova/one-lab-task2/internal/service"
)

type Handler struct {
	srv *service.Service
}

func NewHandler(conf *config.Config, srv *service.Service) *Handler {
	return &Handler{
		srv: srv,
	}
}
