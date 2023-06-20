package main

import (
	"ki-be/configs"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	// run database
	configs.ConnectDB()

	// routes

	e.Logger.Fatal(e.Start(":20224"))
}
