package handlers

import (
	"context"
	responsesModels "ki-be/models/response"
	"ki-be/repositories"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

func ListNews(c echo.Context) error {
	_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var params = repositories.ParamsGetListNews{
		Username: c.QueryParam("username"),
		Keyword:  c.QueryParam("keyword"),
		Tag:      c.QueryParam("tag"),
	}

	// get query page
	if c.QueryParam("page") != "" {
		pageNumber, _ := strconv.Atoi(c.QueryParam("page"))
		params.Page = pageNumber
	} else {
		params.Page = 1
	}

	// get query limit
	if c.QueryParam("limit") != "" {
		limitNumber, _ := strconv.Atoi(c.QueryParam("limit"))
		params.Limit = limitNumber
	} else {
		params.Limit = 9
	}

	data := repositories.GetNews(c, params)
	total := repositories.GetCountNews(c, params)
	status := 204
	message := "Berita tidak ditemukan"

	if data != nil {
		status = 200
		message = "Success"
	}

	return c.JSON(http.StatusOK, responsesModels.GlobalResponse{Status: status, Message: message, Data: &echo.Map{"news": data, "total": total}})
}
