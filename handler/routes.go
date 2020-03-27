package handler

import (
	"github.com/hyecheonlee/realworld-example-app/router/middleware"
	"github.com/hyecheonlee/realworld-example-app/utils"
	"github.com/labstack/echo"
)

func (h *Handler) Register(v1 *echo.Group) {
	jwtMiddleware := middleware.JWT(utils.JWTSecret)
	guestUsers := v1.Group("/users")

}
