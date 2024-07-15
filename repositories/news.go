package repositories

import (
	dataModels "ki-be/models/data"
	tableModels "ki-be/models/tables"
	storageDb "ki-be/storages/db"
	"ki-be/utils"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

type ParamsGetListNews struct {
	Username string
	Keyword  string
	Tag      string
	Page     int
	Limit    int
}

/**
 * function to generate query list news based on params
 */
func QueryListNews(selectCols string, params ParamsGetListNews) *gorm.DB {
	db := storageDb.ConnectDB()

	query := db.Select(selectCols).
		Joins("JOIN user ON user.id_user = berita.author")

	//query search by title
	if params.Keyword != "" {
		query = query.Where("berita.title LIKE ?", "%"+params.Keyword+"%")
	}

	// query get by tag
	if params.Tag != "" {
		query = query.Where("berita.tag LIKE ?", "%"+params.Tag+"%")
	}

	// query get by username
	if params.Username != "" {
		query = query.Where("berita.author = ?", params.Username)
	}

	return query
}

/**
* function to get list news
 */
func GetNews(c echo.Context, params ParamsGetListNews) []dataModels.NewsDataModel {
	dbData := []tableModels.Berita{}

	query := QueryListNews(`id, title`, params)

	query.Limit(params.Limit).Offset(params.Limit * (params.Page - 1)).Order("id DESC").Find(&dbData)

	query.Close()

	// normalize data
	var normalizeData []dataModels.NewsDataModel

	if len(dbData) > 0 {
		for _, n := range dbData {

			// generate url of image

			var newData = dataModels.NewsDataModel{
				Id:        utils.EncCompetitionId(n.Id),
				Title:     n.Title,
				Image:     "",
				Content:   n.Content,
				CreatedAt: n.CreatedAt,
				UpdatedAt: n.UpdatedAt,
				// User: userdata
				// Stats: stats
				// IsDraft: n.IsDraft == "1"
			}

			normalizeData = append(normalizeData, newData)
		}
	}

	return normalizeData
}

/**
* function to get count news
 */
func GetCountNews(c echo.Context, params ParamsGetListNews) int {
	resultData := []tableModels.Berita{}

	query := QueryListNews(`id, title`, params)

	total := query.Find(&resultData).RowsAffected

	query.Close()

	return int(total)
}
