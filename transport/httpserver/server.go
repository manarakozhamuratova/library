package httpserver

import (
	"github.com/labstack/echo/v4"
	"github.com/manarakozhamuratova/one-lab-task2/config"
	"github.com/manarakozhamuratova/one-lab-task2/internal/handler"
)

type Server struct {
	cfg     *config.Config
	handler *handler.Handler
	App     *echo.Echo
}

func NewServer(cfg *config.Config, handler *handler.Handler) *Server {
	return &Server{cfg: cfg, handler: handler}
}
