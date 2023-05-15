package routes

import "github.com/labstack/echo/v4"

func Setup() *echo.Echo {
	e := echo.New()
	
	e.GET("/items", Items)

	e.GET("/couriers", Couriers)
	e.GET("/couriers/:id", Courier)
	e.POST("/couriers", Couriers)
	e.POST("/couriers/delete/:id", CourierDel)


	e.GET("/clients", Clients)
	e.GET("/clients/:id", Client)
	e.POST("/clients", Clients)
	e.POST("/clients/delete/:id", ClientDel)

	e.GET("/items", Items)
	e.GET("/items/:id", Item)
	e.POST("/items", Items)
	e.POST("/items/delete/:id", ItemDel)


	e.GET("/storages", Storages)
	e.GET("/storages/:id", Storage)
	e.POST("/storages", Storages)
	e.POST("/storages/delete/:id", StorageDel)

	return e
}