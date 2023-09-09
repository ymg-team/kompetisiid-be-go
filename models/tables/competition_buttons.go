package models

type Kompetisi_btn struct {
	Id            int    `gorm:"primaryKey;column:id_kompetisi"`
	CompetitionId int    `gorm:"primaryKey;column:id_kompetisi"`
	UserId        int    `gorm:"primaryKey;column:id_user"`
	Like          string `gorm:"primaryKey;column:like"`
}
