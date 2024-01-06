package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	payloadModels "ki-be/models/payload"
	responsesModels "ki-be/models/response"
	tableModels "ki-be/models/tables"
	"ki-be/repositories"
	"ki-be/utils"
	"net/http"
	"strconv"
	"strings"
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
		IsManage:       c.QueryParam("is_manage"),
		Username:       c.QueryParam("username"),
		Keyword:        c.QueryParam("keyword"),
		Tag:            c.QueryParam("tag"),
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

	// get query by id main category
	if c.QueryParam("id_main_category") != "" {
		number, _ := strconv.Atoi(c.QueryParam("id_main_category"))
		params.IdMainCategory = number
	}

	// get query by id main category
	if c.QueryParam("main_category") != "" {
		params.MainCategory = c.QueryParam("main_category")
	}

	// get query by id sub category
	if c.QueryParam("id_sub_category") != "" {
		number, _ := strconv.Atoi(c.QueryParam("id_sub_category"))
		params.IdSubCategory = number
	}

	// get query by sub category
	if c.QueryParam("sub_category") != "" {
		params.SubCategory = c.QueryParam("sub_category")
	}

	// get query by status
	if c.QueryParam("status") != "" {
		params.Status = c.QueryParam("status")
	} else {
		params.Status = "posted"
	}

	// get query by condition
	if c.QueryParam("condition") != "" {
		params.Condition = c.QueryParam("condition")
	} else {
		params.Condition = "active"
	}

	data := repositories.GetCompetitions(c, params)
	total := repositories.GetCountCompetitions(c, params)
	status := 204
	message := "Kompetisi tidak ditemukan"

	if data != nil {
		status = 200
		message = "Success"
	}

	return c.JSON(http.StatusOK, responsesModels.GlobalResponse{Status: status, Message: message, Data: &echo.Map{"competitions": data, "total": total}})
}

func AddCompetition(c echo.Context) error {
	req := c.Request()

	// userKey validation
	userKey := req.Header.Get("userKey")

	if userKey == "" {
		return c.JSON(http.StatusBadRequest, responsesModels.GlobalResponse{Status: http.StatusForbidden, Message: "Please login first", Data: nil})
	} else {
		// check is available user with userKey
		_, userData := repositories.GetUserByUserKey(userKey)

		if userData.Id < 1 {
			return c.JSON(http.StatusBadRequest, responsesModels.GlobalResponse{Status: http.StatusForbidden, Message: "Please login first", Data: nil})
		} else {
			// -- add competition

			// receive body
			var payload payloadModels.PayloadCompetition
			err := json.NewDecoder(req.Body).Decode(&payload)

			if err != nil {
				fmt.Println(err)
				return c.JSON(http.StatusBadRequest, responsesModels.GlobalResponse{Status: http.StatusBadRequest, Message: "Error parsing payload", Data: nil})
			} else {
				now := time.Now()

				// upload to cloudinary first
				uploadDir := "/kompetisi-id/competition/" + userData.Username + "/" + fmt.Sprintf("%d", now.Year())

				_, uploadResult := utils.UploadCloudinary(uploadDir, payload.Poster)

				poster := uploadResult

				posterString, _ := json.Marshal(&poster)

				currentTime := fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d", now.Year(), now.Month(), now.Day(),
					now.Hour(), now.Minute(), now.Second())

				var isGuaranteed string = "0"
				if payload.IsGuaranteed == true {
					isGuaranteed = "1"
				}
				var isMediaPartner string = "0"
				if payload.IsMediaPartner == true {
					isMediaPartner = "1"
				}
				var isDraft string = "0"
				if payload.Draft == true {
					isDraft = "1"
				}

				status := payload.Status
				if userData.Level == "user" {
					status = "waiting"
				}

				newData := tableModels.Kompetisi{
					Id_user:           userData.Id,
					Title:             payload.Title,
					Sort:              payload.Description,
					Poster:            "",
					Poster_cloudinary: string(posterString),
					Organizer:         payload.Organizer,
					DeadlineAt:        payload.DeadlineDate,
					AnnouncementAt:    payload.DeadlineDate,
					Id_main_cat:       payload.MainCat,
					Id_sub_cat:        payload.SubCat,
					Content:           payload.Content,
					PrizeTotal:        payload.PrizeTotal,
					PrizeDescription:  payload.PrizeDescription,
					Contact:           payload.Contacts,
					IsGuaranted:       isGuaranteed,
					IsMediaPartner:    isMediaPartner,
					IsManage:          "0",
					Draft:             isDraft,
					SourceLink:        payload.SourceLink,
					RegisterLink:      payload.RegisterLink,
					Announcements:     payload.Announcements,
					Tags:              payload.Tags,
					Status:            status,
					Views:             1,
					CreatedAt:         currentTime,
					UpdatedAt:         currentTime,
				}

				errInsert, _ := repositories.WriteCompetition(c, newData)

				if errInsert != nil {
					return c.JSON(http.StatusBadRequest, responsesModels.GlobalResponse{Status: http.StatusInternalServerError, Message: "Error insert ke DB", Data: nil})
				} else {

					// post to telegram channel
					if newData.Draft != "1" && newData.Status == "posted" {
						insertedId := repositories.GetLatestCompetitionID(c, repositories.ParamsGetLatestCompetitionId{
							Status:  newData.Status,
							Id_user: newData.Id_user,
							Draft:   newData.Draft,
						})
						encCompetitionId := utils.EncCompetitionId(insertedId)
						chatMessage := "#KompetisiBaru #Kompetisi\n" + newData.Title +
							"\nhttps://kompetisi.id/competition/" + encCompetitionId + "/regulations/" + strings.ToLower(strings.ReplaceAll(newData.Title, " ", "-"))
						repositories.TelegramSendMessage(chatMessage)
					}

					return c.JSON(http.StatusOK, responsesModels.GlobalResponse{Status: http.StatusOK, Message: "Sukses tambah kompetisi", Data: nil})
				}
			}
			// -- end of add competition
		}

	}

}

func UpdateCompetition(c echo.Context) error {
	req := c.Request()

	// userKey validation
	userKey := req.Header.Get("userKey")

	if userKey == "" {
		return c.JSON(http.StatusBadRequest, responsesModels.GlobalResponse{Status: http.StatusForbidden, Message: "Please login first", Data: nil})
	} else {
		// check is available user with userKey
		_, userData := repositories.GetUserByUserKey(userKey)

		if userData.Id < 1 {
			return c.JSON(http.StatusBadRequest, responsesModels.GlobalResponse{Status: http.StatusForbidden, Message: "Please login first", Data: nil})
		} else {
			// -- update competition
			//get influencer ids form query
			encCompetitionId := c.Param("competition_id") + "=="
			decCompetitionId := utils.DecCompetitionId(encCompetitionId)

			// receive body
			var payload payloadModels.PayloadCompetition
			err := json.NewDecoder(req.Body).Decode(&payload)

			if err != nil {
				fmt.Println(err)
				return c.JSON(http.StatusBadRequest, responsesModels.GlobalResponse{Status: http.StatusBadRequest, Message: "Error parsing payload", Data: nil})
			} else {
				now := time.Now()

				currentTime := fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d", now.Year(), now.Month(), now.Day(),
					now.Hour(), now.Minute(), now.Second())

				var isGuaranteed string = "0"
				if payload.IsGuaranteed == true {
					isGuaranteed = "1"
				}
				var isMediaPartner string = "0"
				if payload.IsMediaPartner == true {
					isMediaPartner = "1"
				}
				var isDraft string = "0"
				if payload.Draft == true {
					isDraft = "1"
				}

				status := payload.Status
				if userData.Level == "user" {
					status = "waiting"
				}

				newData := tableModels.Kompetisi{
					Title:            payload.Title,
					Sort:             payload.Description,
					Organizer:        payload.Organizer,
					DeadlineAt:       payload.DeadlineDate,
					AnnouncementAt:   payload.DeadlineDate,
					Id_main_cat:      payload.MainCat,
					Id_sub_cat:       payload.SubCat,
					Content:          payload.Content,
					PrizeTotal:       payload.PrizeTotal,
					PrizeDescription: payload.PrizeDescription,
					Contact:          payload.Contacts,
					IsGuaranted:      isGuaranteed,
					IsMediaPartner:   isMediaPartner,
					IsManage:         "0",
					Draft:            isDraft,
					SourceLink:       payload.SourceLink,
					RegisterLink:     payload.RegisterLink,
					Announcements:    payload.Announcements,
					Tags:             payload.Tags,
					Status:           status,
					UpdatedAt:        currentTime,
				}

				// WIP: handle change poster

				errUpdate := repositories.UpdateCompetition(c, newData, decCompetitionId)

				if errUpdate != nil {
					return c.JSON(http.StatusBadRequest, responsesModels.GlobalResponse{Status: http.StatusInternalServerError, Message: "Error insert ke DB", Data: nil})
				} else {
					if newData.Status == "posted" && newData.Draft != "1" {
						// send chat to telegram channel
						chatMessage := "#KompetisiUpdate #Kompetisi\n" + newData.Title +
							"\nhttps://kompetisi.id/competition/" + c.Param("competition_id") + "/regulations/" + strings.ToLower(strings.ReplaceAll(newData.Title, " ", "-"))
						repositories.TelegramSendMessage(chatMessage)
					}

					return c.JSON(http.StatusOK, responsesModels.GlobalResponse{Status: http.StatusOK, Message: "Sukses update kompetisi", Data: nil})
				}
			}
			// -- end of add competition
		}

	}

}
