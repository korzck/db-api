package routes

import (
	"db-api/models"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func Courier(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("id"))
	courier, err := models.CourierGET(id)
	if err != nil {
		return ctx.NoContent(http.StatusBadRequest)
	}
	return ctx.JSON(http.StatusOK, courier)
}

func Couriers(ctx echo.Context) error {
	method := ctx.Request().Method
	if method == "GET" {
		list, err := models.CouriersGET()
		if err != nil {
			return ctx.NoContent(http.StatusBadRequest)
		}
		ctx.JSON(http.StatusOK, list)
	}
	if method == "POST" {
		courier := new(models.Courier)
		if err := ctx.Bind(courier); err != nil {
			log.Fatalln(err.Error())
			return ctx.NoContent(http.StatusBadRequest)
		}
		err := models.CourierPOST(*courier)
		if err != nil {
			return ctx.NoContent(http.StatusBadRequest)
		}
		return ctx.NoContent(http.StatusOK)
	}
	return ctx.NoContent(http.StatusOK)
}

func CourierDel(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("id"))
	models.CourierDELETE(id)
	return ctx.NoContent(http.StatusOK)
	
}