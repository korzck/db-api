package routes

import (
	"db-api/models"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)


func Storage(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("id"))
	storage, err := models.StorageGET(id)
	if err != nil {
		return ctx.NoContent(http.StatusBadRequest)
	}
	return ctx.JSON(http.StatusOK, storage)
}

func Storages(ctx echo.Context) error {
	method := ctx.Request().Method
	if method == "GET" {
		list, err := models.StoragesGET()
		if err != nil {
			return ctx.NoContent(http.StatusBadRequest)
		}
		ctx.JSON(http.StatusOK, list)
	}
	if method == "POST" {
		Storage := new(models.Storage)
		if err := ctx.Bind(Storage); err != nil {
			log.Fatalln(err.Error())
			return ctx.NoContent(http.StatusBadRequest)
		}
		err := models.StoragePOST(*Storage)
		if err != nil {
			return ctx.NoContent(http.StatusBadRequest)
		}
		return ctx.NoContent(http.StatusOK)
	}
	return ctx.NoContent(http.StatusOK)
}

func StorageDel(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("id"))
	models.StorageDELETE(id)
	return ctx.NoContent(http.StatusOK)
	
}