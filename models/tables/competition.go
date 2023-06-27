package models

type Kompetisi struct {
	Id             int    `gorm:"primaryKey;column:id_kompetisi"`
	Title          string `gorm:"column:judul_kompetisi"`
	Poster         string
	Id_user        int
	Username       string
	Avatar         string
	Id_main_cat    string `gorm:"column:id_main_kat"`
	Main_cat       string `gorm:"column:main_kat"`
	Id_sub_cat     string `gorm:"column:id_sub_kat"`
	Sub_cat        string `gorm:"column:sub_kat"`
	Draft          string `gorm:"column:draft"`
	Status         string `gorm:"column:status"`
	PrizeTotal     int    `gorm:"column:total_hadiah"`
	Organizer      string `gorm:"column:penyelenggara"`
	CreatedAt      string `gorm:"column:created_at"`
	UpdatedAt      string `gorm:"column:updated_at"`
	AnnouncementAt string `gorm:"column:pengumuman"`
	DeadlineAt     string `gorm:"column:deadline"`
	IsGuaranted    string `gorm:"column:garansi"`
	IsMediaPartner string `gorm:"column:mediapartner"`
	Views          int    `gorm:"column:views"`
}
