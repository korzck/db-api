package routes

import (
	"db-api/models"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// func Item(ctx echo.Context) error {
// 	return ctx.String(http.StatusOK, "get item")
// }

// func Items(ctx echo.Context) error {
// 	return ctx.String(http.StatusOK, "get items")
// }

func Item(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("id"))
	Item, err := models.ItemGET(id)
	if err != nil {
		return ctx.NoContent(http.StatusBadRequest)
	}
	return ctx.JSON(http.StatusOK, Item)
}

func Items(ctx echo.Context) error {
	method := ctx.Request().Method
	if method == "GET" {
		list, err := models.ItemsGET()
		if err != nil {
			return ctx.NoContent(http.StatusBadRequest)
		}
		ctx.JSON(http.StatusOK, list)
	}
	if method == "POST" {
		Item := new(models.Item)
		if err := ctx.Bind(Item); err != nil {
			log.Fatalln(err.Error())
			return ctx.NoContent(http.StatusBadRequest)
		}
		err := models.ItemPOST(*Item)
		if err != nil {
			return ctx.NoContent(http.StatusBadRequest)
		}
		return ctx.NoContent(http.StatusOK)
	}
	return ctx.NoContent(http.StatusOK)
}

func ItemDel(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("id"))
	models.ItemDELETE(id)
	return ctx.NoContent(http.StatusOK)
	
}