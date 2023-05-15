package main

import (
	"db-api/models"
	"db-api/routes"
)

func main() {
	models.Setup()
	e := routes.Setup()
	
	e.Start(":8080")
}