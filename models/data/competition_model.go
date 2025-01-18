package models

type CompetitionStatsModel struct {
	Views int `json:"views"`
	Likes int `json:"likes"`
}

type CompetitionAnnouncementModel struct {
	Data string `json:"data"`
	By   string `json:"by"`
	Date string `json:"date"`
}

type CompetitionContactModel struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

type CompetitionDataModel struct {
	Id               string                         `json:"id, omitempty" validate:"required"`
	Title            string                         `json:"title, omitempty" validate:"required"`
	Sort             string                         `json:"sort"`
	Poster           ImageModel                     `json:"poster, omitempty" validate:"required"`
	PosterCloudinary ImageModel                     `json:"poster_cloudinary, omitempty" validate:"required"`
	User             UserModel                      `json:"user, omitempty" validate:"required"`
	MainCategory     MainCategoryModel              `json:"main_category"`
	SubCategory      SubCategoryModel               `json:"sub_category"`
	Draft            bool                           `json:"draft"`
	Status           string                         `json:"status"`
	Prize            PrizeModel                     `json:"prize"`
	Organizer        string                         `json:"organizer"`
	CreatedAt        string                         `json:"created_at"`
	UpdatedAt        string                         `json:"updated_at"`
	DeadlineAt       string                         `json:"deadline_at"`
	AnnouncementAt   string                         `json:"announcement_at"`
	Announcements    []CompetitionAnnouncementModel `json:"announcements"`
	IsGuaranted      bool                           `json:"is_guaranted"`
	IsMediaPartner   bool                           `json:"is_mediapartner"`
	IsManage         bool                           `json:"is_manage"`
	Stats            CompetitionStatsModel          `json:"stats"`
	Content          string                         `json:"content"`
	Contacts         []CompetitionContactModel      `json:"contacts"`
	Tags             string                         `json:"tags"`
	RegisterLink     string                         `json:"register_link"`
	SourceLink       string                         `json:"source_link"`
}
