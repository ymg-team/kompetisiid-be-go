package repositories

import (
	dataModels "ki-be/models/data"
	tableModels "ki-be/models/tables"
	storageDb "ki-be/storages/db"
	"ki-be/utils"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

type ParamsGetListCompetitions struct {
	Status         string
	Limit          int
	Page           int
	Keyword        string
	Tag            string
	OrderBy        string
	IdMainCategory int
	IdSubCategory  int
	MainCategory   string
	SubCategory    string
	IsDraft        string
	IsGuaranted    string
	IsMediaPartner string
	IsManage       string
	Username       string
}

type ParamsGetCompetitionActions struct {
	CompetitionId int
}

/**
 * function to generate query for list competitions based on params
 */
func QueryListCompetitions(selectCols string, params ParamsGetListCompetitions) *gorm.DB {
	db := storageDb.ConnectDB()

	query := db.Select(selectCols).
		Joins("JOIN user ON user.id_user = kompetisi.id_user").
		Joins("JOIN main_kat ON main_kat.id_main_kat = kompetisi.id_main_kat").
		Joins("JOIN sub_kat ON sub_kat.id_sub_kat = kompetisi.id_sub_kat")

	// query search by title
	if params.Keyword != "" {
		query = query.Where("kompetisi.judul_kompetisi LIKE ?", "%"+params.Keyword+"%")
	}

	// query filter by tag
	if params.Tag != "" {
		query = query.Where("kompetisi.tag LIKE ?", "%"+params.Tag+"%")
	}

	// query by username
	if params.Username != "" {
		query = query.Where("user.username = ?", params.Username)
	}

	// query by draft / not
	if params.IsDraft != "" {
		query = query.Where("kompetisi.draft = ?", params.IsDraft)
	}

	// query by manage / not
	if params.IsManage != "" {
		query = query.Where("kompetisi.manage = ?", params.IsManage)
	}

	// query by competition status
	if params.Status != "" {
		if params.Status == "all" {
			query = query.Where("kompetisi.status IN (?)", []string{"posted", "waiting", "approve", "rejected"})
		} else if params.Status == "active" {
			query = query.Where("kompetisi.status = 'posted' AND kompetisi.pengumuman >= CURTIME()")
		} else {
			query = query.Where("kompetisi.status = ?", params.Status)
		}

	}

	// query by id main category
	if params.IdMainCategory != 0 {
		query = query.Where("kompetisi.id_main_kat = ?", params.IdMainCategory)
	}

	// query by main category
	if params.MainCategory != "" {
		query = query.Where("main_kat.main_kat = ?", params.MainCategory)
	}

	// query by id sub category
	if params.IdSubCategory != 0 {
		query = query.Where("kompetisi.id_sub_kat = ?", params.IdSubCategory)
	}

	// query by main category
	if params.SubCategory != "" {
		query = query.Where("sub_kat.sub_kat = ?", params.SubCategory)
	}

	// query by guaranted
	if params.IsGuaranted != "" {
		query = query.Where("kompetisi.garansi = ?", params.IsGuaranted)
	}

	// query by mediapartner
	if params.IsMediaPartner != "" {
		query = query.Where("kompetisi.mediapartner = ?", params.IsMediaPartner)
	}

	return query
}

/**
* function to query to kompetisi_btn table
 */
func QueryCompetitionActions(selectCols string, params ParamsGetCompetitionActions) *gorm.DB {
	db := storageDb.ConnectDB()

	query := db.Select(selectCols)

	if params.CompetitionId != 0 {
		query = query.Where("id_kompetisi=?", params.CompetitionId)
	}

	return query
}

func GetCompetitions(c echo.Context, params ParamsGetListCompetitions) []dataModels.CompetitionDataModel {
	resultData := []tableModels.Kompetisi{}

	query := QueryListCompetitions(`id_kompetisi,judul_kompetisi, kompetisi.sort, 
	kompetisi.poster, kompetisi.poster_cloudinary, kompetisi.poster_cloudinary, 
	draft, kompetisi.status,
	kompetisi.total_hadiah, kompetisi.views, kompetisi.penyelenggara, 
	kompetisi.garansi, kompetisi.mediapartner, kompetisi.manage,
	kompetisi.created_at,kompetisi.updated_at, kompetisi.deadline, kompetisi.pengumuman,
	kompetisi.total_hadiah,
	user.username, user.id_user, 
	main_kat.id_main_kat, main_kat.main_kat, 
	sub_kat.id_sub_kat, sub_kat,sub_kat`, params)

	query.
		Limit(params.Limit).Offset(params.Limit * (params.Page - 1)).Order("id_kompetisi DESC").Find(&resultData)

	query.Close()

	// start mapping and normalize data
	var competitionData []dataModels.CompetitionDataModel

	if len(resultData) > 0 {
		for _, n := range resultData {

			// get total likes by competition id
			resultActions := []tableModels.Kompetisi_btn{}
			queryActions := QueryCompetitionActions("id", ParamsGetCompetitionActions{CompetitionId: n.Id})
			totalLikes := queryActions.Find(&resultActions).RowsAffected
			queryActions.Close()

			var newData = dataModels.CompetitionDataModel{
				Id:     utils.EncCompetitionId(n.Id),
				Title:  n.Title,
				Sort:   n.Sort,
				Poster: utils.ImageNormalizer(n.Poster, n.Poster_cloudinary),
				Status: n.Status,
				User: dataModels.UserModel{
					Username: n.Username,
				},
				MainCategory: dataModels.MainCategoryModel{
					Id:   n.Id_main_cat,
					Name: n.Main_cat,
				},
				SubCategory: dataModels.SubCategoryModel{
					Id:   n.Id_sub_cat,
					Name: n.Sub_cat,
				},
				Draft: n.Draft == "1",
				Prize: dataModels.PrizeModel{
					Total: n.PrizeTotal,
				},
				Organizer:      n.Organizer,
				CreatedAt:      n.CreatedAt,
				UpdatedAt:      n.UpdatedAt,
				DeadlineAt:     n.DeadlineAt,
				AnnouncementAt: n.AnnouncementAt,
				IsGuaranted:    n.IsGuaranted == "1",
				IsMediaPartner: n.IsMediaPartner == "1",
				IsManage:       n.IsManage == "1",
				Stats: dataModels.CompetitionStatsModel{
					Views: n.Views,
					Likes: int(totalLikes),
				},
			}

			competitionData = append(competitionData, newData)
		}
	}

	return competitionData
}

func GetCountCompetitions(c echo.Context, params ParamsGetListCompetitions) int {
	resultData := []tableModels.Kompetisi{}

	query := QueryListCompetitions(`id_kompetisi, judul_kompetisi`, params)

	total := query.Find(&resultData).RowsAffected

	query.Close()

	return int(total)
}

func WriteCompetition(c echo.Context, data tableModels.Kompetisi) (error, int64) {
	db := storageDb.ConnectDB()
	result := db.Omit("username", "avatar", "main_kat", "sub_kat").Create(data)

	return result.Error, result.RowsAffected
}
