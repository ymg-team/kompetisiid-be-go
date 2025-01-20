package repositories

import (
	"encoding/json"
	dataModels "ki-be/models/data"
	tableModels "ki-be/models/tables"
	storageDb "ki-be/storages/db"
	"ki-be/utils"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

type ParamsGetListCompetitions struct {
	Status         string //one of: all, waiting, approve, rejected, posted
	Condition      string //one of: all, active, ended
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
	Id             int
}

type ParamsGetCompetitionActions struct {
	CompetitionId int
}

type ParamsGetLatestCompetitionId struct {
	Status  string
	Draft   string
	Id_user int
	// Title   string
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

	// query by competition status
	// must be: 'all' , 'posted' , 'waiting' , 'approve' , 'rejected
	if params.Status != "" {
		if params.Status == "all" {
			query = query.Where("kompetisi.status IN (?)", []string{"posted", "waiting", "approve", "rejected"})
		} else {
			query = query.Where("kompetisi.status = ?", params.Status)
		}
	}

	// query by competition condition based on deadline time (used on dashboard)
	// must be: 'all', 'active', 'ended'
	if params.Condition != "" {
		if params.Condition == "active" {
			//get active competition, announcement date deadline > now
			query = query.Where("kompetisi.pengumuman >= CURTIME()")
		} else if params.Condition == "end" {
			// get ended competition, now > announcement date
			query = query.Where("CURTIME() > kompetisi.pengumuman")
		}
		// else show competition with all condition
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

	// query by manage / not
	if params.IsManage != "" {
		query = query.Where("kompetisi.manage = ?", params.IsManage)
	}

	return query
}

/**
 * function to generate query for list competitions based on params
 */
func QueryLatestCompetitionID(selectCols string, params ParamsGetLatestCompetitionId) *gorm.DB {
	db := storageDb.ConnectDB()

	query := db.Select(selectCols)

	// query by username
	if params.Id_user != 0 {
		query = query.Where("kompetisi.id_user = ?", params.Id_user)
	}

	// query by draft / not
	if params.Draft != "" {
		query = query.Where("kompetisi.draft = ?", params.Draft)
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

/*
* function to get list competition
 */
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
				Poster: utils.ImageCompetitionNormalizer(n.Poster, n.Poster_cloudinary),
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

/**
* function to get competition detail
 */
func GetCompetitionDetail(c echo.Context, params ParamsGetListCompetitions) []dataModels.CompetitionDataModel {
	dbData := []tableModels.Kompetisi{}

	query := QueryListCompetitions(`id_kompetisi,judul_kompetisi, kompetisi.sort, 
	kompetisi.poster, kompetisi.poster_cloudinary, kompetisi.poster_cloudinary, 
	draft, kompetisi.status, kompetisi.konten,
	kompetisi.total_hadiah, kompetisi.views, kompetisi.penyelenggara, 
	kompetisi.garansi, kompetisi.mediapartner, kompetisi.manage,
	kompetisi.created_at,kompetisi.updated_at, kompetisi.deadline, kompetisi.pengumuman,
	kompetisi.total_hadiah, kompetisi.hadiah,
	kompetisi.dataPengumuman,
	kompetisi.kontak,
	kompetisi.tag,
	kompetisi.sumber, kompetisi.ikuti,
	user.username, user.id_user, 
	main_kat.id_main_kat, main_kat.main_kat, 
	sub_kat.id_sub_kat, sub_kat,sub_kat`, params)
	query.Where("id_kompetisi = ?", params.Id).Limit(1).Offset(0).Find(&dbData)
	// query.Close()

	var normalizeData []dataModels.CompetitionDataModel

	if len(dbData) > 0 {

		for _, n := range dbData {

			// get total likes by competition id
			resultActions := []tableModels.Kompetisi_btn{}
			queryActions := QueryCompetitionActions("id", ParamsGetCompetitionActions{CompetitionId: n.Id})
			totalLikes := queryActions.Find(&resultActions).RowsAffected
			queryActions.Close()

			// JSON.parse announcements data
			dataAnnouncements := []dataModels.CompetitionAnnouncementModel{}
			json.Unmarshal([]byte(n.Announcements), &dataAnnouncements)

			// JSON.parse contacts data
			dataContacts := []dataModels.CompetitionContactModel{}
			json.Unmarshal([]byte(n.Contacts), &dataContacts)

			var newData = dataModels.CompetitionDataModel{
				Id:       utils.EncCompetitionId(n.Id),
				Title:    n.Title,
				Sort:     n.Sort,
				Poster:   utils.ImageCompetitionNormalizer(n.Poster, n.Poster_cloudinary),
				Status:   n.Status,
				Contacts: dataContacts,
				Content:  n.Content,
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
					Total:       n.PrizeTotal,
					Description: n.PrizeDescription,
				},
				AnnouncementAt: n.AnnouncementAt,
				Announcements:  dataAnnouncements,
				Organizer:      n.Organizer,
				CreatedAt:      n.CreatedAt,
				UpdatedAt:      n.UpdatedAt,
				DeadlineAt:     n.DeadlineAt,
				IsGuaranted:    n.IsGuaranted == "1",
				IsMediaPartner: n.IsMediaPartner == "1",
				IsManage:       n.IsManage == "1",
				Stats: dataModels.CompetitionStatsModel{
					Views: n.Views,
					Likes: int(totalLikes),
				},
				Tags: n.Tags,
				SourceLink: n.SourceLink,
				RegisterLink: n.RegisterLink,
			}

			normalizeData = append(normalizeData, newData)
		}

	}

	return normalizeData
}

/*
* function to get count competition
 */
func GetCountCompetitions(c echo.Context, params ParamsGetListCompetitions) int {
	resultData := []tableModels.Kompetisi{}

	query := QueryListCompetitions(`id_kompetisi, judul_kompetisi`, params)

	total := query.Find(&resultData).RowsAffected

	query.Close()

	return int(total)
}

func WriteCompetition(c echo.Context, data tableModels.Kompetisi) (error, *gorm.DB) {
	db := storageDb.ConnectDB()
	result := db.Omit("username", "avatar", "main_kat", "sub_kat").Create(data)
	db.Close()
	return result.Error, result
}

func UpdateCompetition(c echo.Context, data tableModels.Kompetisi, competitionId int) error {
	db := storageDb.ConnectDB()
	var kompetisi tableModels.Kompetisi
	result := db.Model(&kompetisi).Where("id_kompetisi = ?", competitionId).UpdateColumn(data)
	db.Close()
	return result.Error
}

func GetLatestCompetitionID(c echo.Context, params ParamsGetLatestCompetitionId) int {
	resultData := tableModels.Kompetisi{}

	query := QueryLatestCompetitionID(`id_kompetisi`, params)

	query.
		Limit(1).Order("id_kompetisi DESC").First(&resultData)

	query.Close()

	return resultData.Id
}

/**
* function to increment views of news by ud
 */
func IncrCompetitionViews(v echo.Context, competitionId int) {
	db := storageDb.ConnectDB()

	NewsData := tableModels.Kompetisi{}

	db.Model(&NewsData).Where("kompetisi.id_kompetisi = ?", competitionId).Update("Views", gorm.Expr("views + ?", 1))

	db.Close()
}
