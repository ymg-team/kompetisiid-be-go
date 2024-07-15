package main

import (
	"ki-be/routes"
	"ki-be/storages/db"
	"os"

	"github.com/labstack/echo/v4"
)

func main() {
	// global setup
	os.Setenv("TZ", "Asia/Jakarta")

	e := echo.New()

	// run database
	db.ConnectDB()

	// routes
	// /competition/*
	routes.CompetitionRoute(e)
	// /news/*
	routes.NewsRoute(e)

	e.Logger.Fatal(e.Start(":20224"))
}
