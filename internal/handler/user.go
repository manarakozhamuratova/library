package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/manarakozhamuratova/one-lab-task2/internal/model"
)

// CreateUser godoc
// @Summary      Создание пользователя
// @Description  Создание пользователя
// @ID           CreateOrder
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        rq   body      model.User  true  "Входящие данные"
// @Router       /user [post]
func (h Handler) CreateUser(c echo.Context) error {
	var req model.User
	if err := c.Bind(&req); err != nil {
		return err
	}
	// resp, err := h.srv.User.Create(c.Request().Context(), req)
	// if err != nil {
	// 	return echo.NewHTTPError(http.StatusBadRequest, err)
	// }
	return c.JSON(http.StatusOK, nil)
}
