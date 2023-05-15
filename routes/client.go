package routes

import (
	"db-api/models"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)


func Client(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("id"))
	client, err := models.ClientGET(id)
	if err != nil {
		return ctx.NoContent(http.StatusBadRequest)
	}
	return ctx.JSON(http.StatusOK, client)
}

func Clients(ctx echo.Context) error {
	method := ctx.Request().Method
	if method == "GET" {
		list, err := models.ClientsGET()
		if err != nil {
			return ctx.NoContent(http.StatusBadRequest)
		}
		ctx.JSON(http.StatusOK, list)
	}
	if method == "POST" {
		client := new(models.Client)
		if err := ctx.Bind(client); err != nil {
			log.Fatalln(err.Error())
			return ctx.NoContent(http.StatusBadRequest)
		}
		err := models.ClientPOST(*client)
		if err != nil {
			return ctx.NoContent(http.StatusBadRequest)
		}
		return ctx.NoContent(http.StatusOK)
	}
	return ctx.NoContent(http.StatusOK)
}

func ClientDel(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("id"))
	models.ClientDELETE(id)
	return ctx.NoContent(http.StatusOK)
	
}