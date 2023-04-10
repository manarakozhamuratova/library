package httpserver

import echoSwagger "github.com/swaggo/echo-swagger"

func (s *Server) SetupRoutes() {
	s.App.GET("/swagger/*", echoSwagger.WrapHandler)
	s.App.POST("/user", s.handler.CreateUser)
	s.App.POST("/auth", s.handler.Auth)
	s.App.PUT("/user", s.handler.UpdatePassword, s.handler.JWT().ValidateAuth)
}
