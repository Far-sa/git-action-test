package handler

import (
	"net/http"
	"user-service/port"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userSvc port.UserService
}

func NewUserHandler(userSvc port.UserService) *UserHandler {
	return &UserHandler{userSvc: userSvc}
}

func (h *UserHandler) Register(c echo.Context) error {
	username := c.FormValue("username")
	email := c.FormValue("email")
	password := c.FormValue("password")

	if err := h.userSvc.RegisterUser(username, email, password); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.NoContent(http.StatusCreated)
}
