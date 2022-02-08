package routes

import (
	"echo-mongo-api/controllers" 

	"github.com/labstack/echo/v4"
)

// PlayerRoute manage all the player-related routes
func PlayerRoute(e *echo.Echo) {
	e.POST("/player", controllers.CreatePlayer)
	e.GET("/player/:playerId", controllers.GetAPlayer)
	e.PUT("/player/:playerId", controllers.UpdatePlayer)
	e.GET("/players", controllers.GetAllPlayers)
}
