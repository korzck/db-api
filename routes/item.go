package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Item(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "get item")
}

func Items(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "get items")
}