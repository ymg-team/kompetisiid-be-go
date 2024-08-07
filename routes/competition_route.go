package routes

import (
	"ki-be/handlers"

	"github.com/labstack/echo/v4"
)

func CompetitionRoute(e *echo.Echo) {
	// all routes relates to competitions comes here
	e.GET("/competitions", handlers.ListCompetition)
	e.POST("/competitions", handlers.AddCompetition)
	e.PUT("/competitions/:competition_id", handlers.UpdateCompetition)
	e.GET("/competitions/detail", handlers.DetailCompetition)
}
