package handlers

import (
	"context"
	responsesModels "ki-be/models/response"
	"ki-be/repositories"
	"ki-be/utils"
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
		Status:   c.QueryParam("status"),
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

func DetailNews(c echo.Context) error {
	_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// get query id
	if c.QueryParam("id") != "" {
		encCompetitionId := c.QueryParam("id") + "=="
		decId := utils.DecCompetitionId(encCompetitionId)

		if decId != 0 {
			var params = repositories.ParamsGetListNews{
				Id: decId,
			}

			data := repositories.GetNewsDetail(c, params)

			if len(data) > 0 {
				// increment views
				repositories.IncrNewsViews(c, decId)

				return c.JSON(http.StatusOK, responsesModels.GlobalResponse{Status: 200, Message: "Success", Data: &echo.Map{"news": data}})
			} else {
				return c.JSON(http.StatusOK, responsesModels.GlobalResponse{Status: 204, Message: "Berita tidak ditemukan"})
			}
		}

	}

	// return not found result
	return c.JSON(http.StatusOK, responsesModels.GlobalResponse{Status: 204, Message: "Berita tidak ditemukan"})
}
