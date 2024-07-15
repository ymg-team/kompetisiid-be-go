package routes

import (
	"ki-be/handlers"

	"github.com/labstack/echo/v4"
)

func NewsRoute(e *echo.Echo) {
	// all routes relates to news comes here
	e.GET("/news", handlers.ListNews)
}
