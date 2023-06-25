package repositories

import (
	dataModels "ki-be/models/data"
	models "ki-be/models/data"
	tableModels "ki-be/models/tables"
	storageDb "ki-be/storages/db"
	"ki-be/utils"

	"github.com/labstack/echo/v4"
)

type ParamsGetListCompetitions struct {
	Status         string
	Limit          int
	Page           int
	Keyword        string
	OrderBy        string
	IdMainCategory int
	IdSubCategory  int
	Draft          string
}

func GetCompetitions(c echo.Context, params ParamsGetListCompetitions) []dataModels.CompetitionDataModel {
	db := storageDb.ConnectDB()

	resultData := []tableModels.Kompetisi{}

	query := db.Select(`id_kompetisi,judul_kompetisi, poster, draft, kompetisi.status,
	kompetisi.total_hadiah, kompetisi.views, kompetisi.penyelenggara, 
	kompetisi.garansi,
	kompetisi.created_at,kompetisi.updated_at, kompetisi.deadline, kompetisi.pengumuman,
	kompetisi.total_hadiah,
	user.username, user.id_user, 
	main_kat.id_main_kat, main_kat.main_kat, 
	sub_kat.id_sub_kat, sub_kat,sub_kat`).
		Joins("JOIN user ON user.id_user = kompetisi.id_user").
		Joins("JOIN main_kat ON main_kat.id_main_kat = kompetisi.id_main_kat").
		Joins("JOIN sub_kat ON sub_kat.id_main_kat = kompetisi.id_sub_kat")

	// query by draft / not
	if params.Draft != "" {
		query = query.Where("kompetisi.draft = ?", params.Draft)
	}

	// query by competition status
	if params.Status != "" {
		if params.Status == "all" {
			query = query.Where("kompetisi.status IN (?)", []string{"posted", "waiting", "approve", "rejected"})
		} else if params.Status == "active" {
			query = query.Where("kompetisi.status = 'posted' AND kompetisi.pengumuman >= CURDATE()")
		} else {
			query = query.Where("kompetisi.status = ?", params.Status)
		}

	}

	// query by main category
	if params.IdMainCategory != 0 {
		query = query.Where("kompetisi.id_main_kat = ?", params.IdMainCategory)
	}

	query.
		Limit(params.Limit).Offset(params.Limit * (params.Page - 1)).Order("id_kompetisi desc").Find(&resultData)

	// start mapping and normalize data
	var competitionData []dataModels.CompetitionDataModel

	if len(resultData) > 0 {
		for _, n := range resultData {

			var newData = dataModels.CompetitionDataModel{
				Id:     utils.EncCompetitionId(n.Id),
				Title:  n.Title,
				Poster: utils.ImageNormalizer(n.Poster),
				Status: n.Status,
				User: models.UserModel{
					Username: n.Username,
				},
				MainCategory: models.MainCategoryModel{
					Id:   n.Id_main_cat,
					Name: n.Main_cat,
				},
				SubCategory: models.SubCategoryModel{
					Id:   n.Id_sub_cat,
					Name: n.Sub_cat,
				},
				Draft: n.Draft == "1",
				Prize: models.PrizeModel{
					Total: n.PrizeTotal,
				},
				Organizer:      n.Organizer,
				CreatedAt:      n.CreatedAt,
				UpdatedAt:      n.UpdatedAt,
				DeadlineAt:     n.DeadlineAt,
				AnnouncementAt: n.AnnouncementAt,
			}

			competitionData = append(competitionData, newData)
		}
	}

	return competitionData
}
