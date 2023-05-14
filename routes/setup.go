package routes

import "github.com/labstack/echo/v4"

func Setup() *echo.Echo {
	e := echo.New()
	
	e.GET("/items", Items)
	return e
}