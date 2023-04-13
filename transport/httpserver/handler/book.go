package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/manarakozhamuratova/one-lab-task2/internal/model"
)

// CreateBook godoc
// @Summary      Создание книги
// @Description  Создание книги
// @Tags         book
// @Accept       json
// @Produce      json
// @Param        rq   body      model.Book  true  "Входящие данные"
// @Success	     200  {}  uint
// @Router       /book [post]
func (h *Handler) CreateBook(c echo.Context) error {
	var req model.Book
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	resp, err := h.srv.Book.Create(c.Request().Context(), req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, resp)
}

// TakeABook godoc
// @Summary      Взять книгу
// @Description  Взять книгу
// @Tags         book
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Book ID"
// @Success	     200  {}  uint
// @Router       /book/{id}/borrow [post]
// @Security 	 ApiKeyAuth
func (h *Handler) TakeABook(c echo.Context) error {
	bookID, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	userId, ok := c.Request().Context().Value(model.ContextUsername).(uint)
	if !ok {
		return echo.NewHTTPError(http.StatusUnauthorized, nil)
	}
	err := h.srv.Book.TakeABook(c.Request().Context(), model.BookOperation{
		UserID: userId,
		BookID: uint(bookID),
	})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, nil)
}

// GiveTheBook godoc
// @Summary      Вернуть книгу
// @Description  Вернуть книгу
// @Tags         book
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Book ID"
// @Success	     200  {}  uint
// @Router       /book/{id}/return [post]
// @Security 	 ApiKeyAuth
func (h *Handler) GiveTheBook(c echo.Context) error {
	bookID, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	userId, ok := c.Request().Context().Value(model.ContextUsername).(uint)
	if !ok {
		return echo.NewHTTPError(http.StatusUnauthorized, nil)
	}
	err := h.srv.Book.GiveTheBook(c.Request().Context(), model.BookOperation{
		UserID: userId,
		BookID: uint(bookID),
	})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, nil)
}

// CreateTransaction godoc
// @Summary      Купить книгу
// @Description  Купить книгу
// @Tags         book
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Book ID"
// @Success	     200  {}  uint
// @Router       /book/{id}/buy [post]
// @Security 	 ApiKeyAuth
func (h *Handler) BuyABook(c echo.Context) error {
	bookID, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	userId, ok := c.Request().Context().Value(model.ContextUsername).(uint)
	if !ok {
		return echo.NewHTTPError(http.StatusUnauthorized, nil)
	}
	err := h.srv.Book.BuyABook(c.Request().Context(), model.Transaction{
		UserID: userId,
		BookID: uint(bookID),
	})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, nil)
}
