package httpserver

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/manarakozhamuratova/one-lab-task2/config"
	"github.com/manarakozhamuratova/one-lab-task2/transport/httpserver/handler"
)

type Server struct {
	cfg     *config.Config
	handler *handler.Handler
	App     *echo.Echo
}

func NewServer(cfg *config.Config, handler *handler.Handler) *Server {
	return &Server{cfg: cfg, handler: handler}
}

func (s *Server) StartHTTPServer(ctx context.Context) error {
	s.App = s.BuildEngine()

	s.SetupRoutes()

	go func() {
		if err := s.App.Start(fmt.Sprintf("%s", s.cfg.ServerAddress)); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen:%+s\n", err)
		}
	}()
	<-ctx.Done()

	ctxShutDown, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		cancel()
	}()
	if err := s.App.Shutdown(ctxShutDown); err != nil {
		log.Fatalf("server Shutdown Failed:%+s", err)
	}
	log.Print("server exited properly")
	return nil
}

func (s *Server) BuildEngine() *echo.Echo {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{"*"},
	}))
	e.Use(middleware.RequestID())
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			log.Println("request, status", v.URI, v.Status)
			return nil
		},
	}))
	return e
}
