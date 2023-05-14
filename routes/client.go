package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Client(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "get client")
}

func Clients(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "get clients")
}