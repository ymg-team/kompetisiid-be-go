package models

type Berita struct {
	Id              int    `gorm:"primaryKey; column:id"`
	Title           string `gorm:"column:title"`
	Image           string `gorm:"image"`
	ImageCloudinary string `gorm:"image_cloudinary"`
	Content         string `gorm:"content"`
	CreatedAt       string `gorm:created_at`
	UpdatedAt       string `gorm:updated_at`
	UserId          int    `gorm:author`
	Status          string `gorm:status`
	IsDraft         int    `gorm:draft`
	Tags            string `gorm:"column:tag"`
	Username        string
	Draft           int `gorm:"column:draft"`
}
