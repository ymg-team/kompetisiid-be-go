package models

type Kompetisi struct {
	Id                int    `gorm:"primaryKey;column:id_kompetisi"`
	Title             string `gorm:"column:judul_kompetisi"`
	Sort              string `gorm:"column:sort"`
	Poster            string
	Poster_cloudinary string `gorm:"column:poster_cloudinary"`
	Id_user           int
	Username          string
	Avatar            string
	Id_main_cat       int    `gorm:"column:id_main_kat"`
	Main_cat          string `gorm:"column:main_kat"`
	Id_sub_cat        int    `gorm:"column:id_sub_kat"`
	Sub_cat           string `gorm:"column:sub_kat"`
	Draft             string `gorm:"column:draft"`
	Content           string `gorm:"column:konten"`
	Status            string `gorm:"column:status"`
	PrizeTotal        int    `gorm:"column:total_hadiah"`
	PrizeDescription  string `gorm:"column:hadiah"`
	Contact           string `gorm:"column:kontak"`
	Organizer         string `gorm:"column:penyelenggara"`
	AnnouncementAt    string `gorm:"column:pengumuman"`
	DeadlineAt        string `gorm:"column:deadline"`
	IsGuaranted       string `gorm:"column:garansi"`
	IsMediaPartner    string `gorm:"column:mediapartner"`
	IsManage          string `gorm:"column:manage"`
	SourceLink        string `gorm:"column:sumber"`
	RegisterLink      string `gorm:"column:ikuti"`
	Views             int    `gorm:"column:views"`
	Announcements     string `gorm:"column:dataPengumuman"`
	Tags              string `gorm:"column:tag"`
	CreatedAt         string `gorm:"column:created_at"`
	UpdatedAt         string `gorm:"column:updated_at"`
}
