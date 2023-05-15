package routes

import (
	"db-api/models"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func Order(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("id"))
	order, err := models.OrderGET(id)
	if err != nil {
		return ctx.NoContent(http.StatusBadRequest)
	}
	return ctx.JSON(http.StatusOK, order)
}

func Orders(ctx echo.Context) error {
	method := ctx.Request().Method
	if method == "GET" {
		list, err := models.OrdersGET()
		if err != nil {
			return ctx.NoContent(http.StatusBadRequest)
		}
		ctx.JSON(http.StatusOK, list)
	}
	if method == "POST" {
		order := new(models.Order)
		if err := ctx.Bind(order); err != nil {
			log.Fatalln(err.Error())
			return ctx.NoContent(http.StatusBadRequest)
		}
		err := models.OrderPOST(*order)
		if err != nil {
			return ctx.NoContent(http.StatusBadRequest)
		}
		return ctx.NoContent(http.StatusOK)
	}
	return ctx.NoContent(http.StatusOK)
}

func OrderDel(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("id"))
	models.OrderDELETE(id)
	return ctx.NoContent(http.StatusOK)
	
}