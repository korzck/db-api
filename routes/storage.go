package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Storage(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "get storage")
}

func Storages(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "get storages")
}