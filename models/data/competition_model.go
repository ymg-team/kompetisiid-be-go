package models

type CompetitionStatsModel struct {
	Views int `json:"views"`
	Likes int `json:"likes"`
}

type CompetitionDataModel struct {
	Id             string                `json:"id, omitempty" validate:"required"`
	Title          string                `json:"title, omitempty" validate:"required"`
	Poster         ImageModel            `json:"poster, omitempty" validate:"required"`
	User           UserModel             `json:"user, omitempty" validate:"required"`
	MainCategory   MainCategoryModel     `json:"main_category"`
	SubCategory    SubCategoryModel      `json:"sub_category"`
	Draft          bool                  `json:"draft"`
	Status         string                `json:"status"`
	Prize          PrizeModel            `json:"prize"`
	Organizer      string                `json:"organizer"`
	CreatedAt      string                `json:"created_at"`
	UpdatedAt      string                `json:"updated_at"`
	DeadlineAt     string                `json:"deadline_at"`
	AnnouncementAt string                `json:"announcement_at"`
	IsGuaranted    bool                  `json:"is_guaranted"`
	IsMediaPartner bool                  `json:"is_mediapartner"`
	Stats          CompetitionStatsModel `json:"stats"`
}
