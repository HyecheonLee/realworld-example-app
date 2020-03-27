package handler

import (
	"github.com/hyecheonlee/realworld-example-app/model"
	"github.com/labstack/echo"
)

func (h *Handler) SignUp(c echo.Context) error {
	var u model.User
	req := &userRe
}
