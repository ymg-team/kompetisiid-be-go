package handlers

import (
	"context"
	responses "ki-be/models/response"
	"ki-be/repositories"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

func ListCompetition(c echo.Context) error {
	_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var params = repositories.ParamsGetListCompetitions{
		Status:         c.QueryParam("status"),
		IsDraft:        c.QueryParam("is_draft"),
		IsGuaranted:    c.QueryParam("is_guaranted"),
		IsMediaPartner: c.QueryParam("is_mediapartner"),
		Username:       c.QueryParam("username"),
		Keyword:        c.QueryParam("keyword"),
	}

	// get query page
	if c.QueryParam("page") != "" {
		pageNumber, _ := strconv.Atoi(c.QueryParam("page"))
		params.Page = pageNumber
	} else {
		params.Page = 1
	}

	// get query page
	if c.QueryParam("limit") != "" {
		limitNumber, _ := strconv.Atoi(c.QueryParam("limit"))
		params.Limit = limitNumber
	} else {
		params.Limit = 9
	}

	// get query by main category
	if c.QueryParam("id_main_category") != "" {
		number, _ := strconv.Atoi(c.QueryParam("id_main_category"))
		params.IdMainCategory = number
	}

	// get query by sub category
	if c.QueryParam("id_sub_category") != "" {
		number, _ := strconv.Atoi(c.QueryParam("id_sub_category"))
		params.IdSubCategory = number
	}

	// get query by status
	if c.QueryParam("status") != "" {
		params.Status = c.QueryParam("status")
	} else {
		params.Status = "posted"
	}

	data := repositories.GetCompetitions(c, params)
	total := repositories.GetCountCompetitions(c, params)

	return c.JSON(http.StatusOK, responses.GlobalResponse{Status: http.StatusOK, Message: "success", Data: &echo.Map{"competitions": data, "total": total}})
}
