package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Courier(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "get courier")
}

func Couriers(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "get couriers")
}