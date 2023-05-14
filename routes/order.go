package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Order(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "get order")
}

func Orders(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "get orders")
}