package handler

import (
	"github.com/manarakozhamuratova/one-lab-task2/config"
	"github.com/manarakozhamuratova/one-lab-task2/internal/service"
	"github.com/manarakozhamuratova/one-lab-task2/transport/httpserver/middleware"
)

type Handler struct {
	srv *service.Service
	jwt *middleware.JWTAuth
}

func (h *Handler) JWT() *middleware.JWTAuth {
	return h.jwt
}

func NewHandler(conf *config.Config, srv *service.Service, jwt *middleware.JWTAuth) *Handler {
	return &Handler{
		srv: srv,
		jwt: jwt,
	}
}
