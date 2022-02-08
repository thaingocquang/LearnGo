package main

import (
	"echo-mongo-api/configs"
	"echo-mongo-api/routes"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	// run database
	configs.ConnectDB()

	// routes
	routes.PlayerRoute(e)

	e.Logger.Fatal(e.Start(":1323"))
}