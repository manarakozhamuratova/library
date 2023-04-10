package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/manarakozhamuratova/one-lab-task2/internal/model"
)

// CreateUser godoc
// @Summary      Создание пользователя
// @Description  Создание пользователя
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        rq   body      model.User  true  "Входящие данные"
// @Success	     200  {object}  model.CreateResp
// @Router       /user [post]
func (h *Handler) CreateUser(c echo.Context) error {
	var req model.User
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	resp, err := h.srv.User.Create(c.Request().Context(), req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, resp)
}

// AuthUser godoc
// @Summary      Аутентификация пользователя
// @Description  Аутентификация пользователя
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        rq   body      model.AuthUser  true  "Входящие данные"
// @Success	     200  {object}  model.AuthResponse
// @Router       /auth [post]
func (h *Handler) Auth(c echo.Context) error {
	var req model.AuthUser
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	err := h.srv.User.Auth(c.Request().Context(), req)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}

	token, err := h.jwt.GenerateJWT(req.Username)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}
	return c.JSON(http.StatusOK, model.AuthResponse{
		Token: token,
	})
}

// UpdatePassword godoc
// @Summary      Обновление пароля пользователя
// @Description  Обновление пароля пользователя
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        rq   body      model.UpdatePassword  true  "Входящие данные"
// @Router       /user [put]
// @Security 	 ApiKeyAuth
func (h *Handler) UpdatePassword(c echo.Context) error {
	var req model.UpdatePassword
	if err := c.Bind(&req); err != nil {
		echo.NewHTTPError(http.StatusBadRequest, err)
	}
	username, ok := c.Request().Context().Value(model.ContextUsername).(string)
	if !ok {
		return echo.NewHTTPError(http.StatusUnauthorized, nil)
	}
	req.Username = username
	if err := h.srv.User.UpdatePassword(c.Request().Context(), req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, nil)
}
