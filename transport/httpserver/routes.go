package httpserver

import echoSwagger "github.com/swaggo/echo-swagger"

func (s *Server) SetupRoutes() {
	s.App.GET("/swagger/*", echoSwagger.WrapHandler)
	s.App.POST("/user", s.handler.CreateUser)
	s.App.POST("/auth", s.handler.Auth)
	s.App.PUT("/user", s.handler.UpdatePassword, s.handler.JWT().ValidateAuth)
	s.App.GET("/user", s.handler.GetUsersWithActiveBorrowedBooks)
	s.App.GET("/user/count_books", s.handler.GetUsersWithBorrowedBookCountByDate)
	s.App.POST("/book", s.handler.CreateBook)
	s.App.GET("/book/:id/take", s.handler.TakeABook, s.handler.JWT().ValidateAuth)
	s.App.GET("/book/:id/give", s.handler.GiveTheBook, s.handler.JWT().ValidateAuth)

}
