package models

type PayloadCompetition struct {
	Id_user          int    `json:"id_user,omitempty"`
	Title            string `json:"title,omitempty"`
	Description      string `json:"description,omitempty"`
	Organizer        string `json:"organizer,omitempty"`
	Poster           string `json:"poster,omitempty"`
	DeadlineDate     string `json:"deadline_date,omitempty"`
	AnnouncementDate string `json:"announcement_date,omitempty"`
	MainCat          int    `json:"main_cat,omitempty"`
	SubCat           int    `json:"sub_cat,omitempty"`
	PrizeTotal       int    `json:"prize_total,omitempty"`
	PrizeDescription string `json:"prize_description,omitempty"`
	Content          string `json:"content,omitempty"`
	Contacts         string `json:"contacts,omitempty"`
	IsGuaranteed     bool   `json:"is_guaranteed,omitempty"`
	IsMediaPartner   bool   `json:"is_mediapartner,omitempty"`
	IsManage         bool   `json:"is_manage,omitempty"`
	Draft            bool   `json:"draft,omitempty"`
	SourceLink       string `json:"source_link,omitempty"`
	RegisterLink     string `json:"register_link,omitempty"`
	Tags             string `json:"tags,omitempty"`
	Announcements    string `json:"announcements,omitempty"`
	Status           string `json:"status,omitempty"`
}
