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
	Id       int
	Username string
	Keyword  string
	Tag      string
	Page     int
	Limit    int
	Status   string
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

	// query get by id news
	if params.Id != 0 {
		query = query.Where("berita.id = ?", params.Id)
	}

	// query by status
	if params.Status == "published" {
		query = query.Where("berita.draft = ?", "0")
	} else if params.Status == "draft" {
		query = query.Where("berita.draft = ?", "1")
	}

	return query
}

/**
* function to get list news
 */
func GetNews(c echo.Context, params ParamsGetListNews) []dataModels.NewsDataModel {
	dbData := []tableModels.Berita{}

	query := QueryListNews(`berita.id, berita.title, berita.image, berita.content, berita.created_at, berita.updated_at, berita.author,
		berita.image, berita.image_cloudinary, berita.tag,
		user.username`, params)

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
				Image:     utils.ImageNewsNormalizer(n.Image, n.ImageCloudinary),
				Content:   n.Content,
				CreatedAt: n.CreatedAt,
				UpdatedAt: n.UpdatedAt,
				Tags:      n.Tags,
				User: dataModels.UserModel{
					Username: n.Username,
				},
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

/**
* function to get deat
 */
func GetNewsDetail(c echo.Context, params ParamsGetListNews) []dataModels.NewsDataModel {
	dbData := []tableModels.Berita{}

	query := QueryListNews(`berita.id, berita.title, berita.image, berita.content, berita.created_at, berita.updated_at, berita.author,
		berita.image, berita.image_cloudinary, berita.tag,
		berita.views,
		user.username`, params)

	query.Limit(1).Offset(0).Order("id DESC").Find(&dbData)

	query.Close()

	var normalizeData []dataModels.NewsDataModel

	if len(dbData) > 0 {
		for _, n := range dbData {

			// generate url of image

			var newData = dataModels.NewsDataModel{
				Id:        utils.EncCompetitionId(n.Id),
				Title:     n.Title,
				Image:     utils.ImageNewsNormalizer(n.Image, n.ImageCloudinary),
				Content:   n.Content,
				CreatedAt: n.CreatedAt,
				UpdatedAt: n.UpdatedAt,
				Tags:      n.Tags,
				User: dataModels.UserModel{
					Username: n.Username,
				},
				Stats: dataModels.NewsStatsModel{
					Views: n.Views,
				},
			}

			normalizeData = append(normalizeData, newData)
		}
	}

	return normalizeData
}

/**
* function to increment views of news by ud
 */
func IncrNewsViews(v echo.Context, newsId int) {
	db := storageDb.ConnectDB()

	NewsData := tableModels.Berita{}

	db.Model(&NewsData).Where("berita.id = ?", newsId).Update("Views", gorm.Expr("views + ?", 1))

	db.Close()
}
