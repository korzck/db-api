package main

import (
	"db-api/models"

	"github.com/labstack/echo/v4"
)

func main() {
	models.Setup()
	e := echo.New()
	e.Start(":8080")
}