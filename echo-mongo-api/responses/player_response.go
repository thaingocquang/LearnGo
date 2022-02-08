package responses

import (
	"github.com/labstack/echo/v4"
)

// PlayerResponse describe API’s response
type PlayerResponse struct {
	Status  int       `json:"status"`
	Message string    `json:"message"`
	Data    *echo.Map `json:"data"`
}
