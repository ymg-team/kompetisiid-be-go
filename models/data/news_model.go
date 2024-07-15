package models

type NewsStatsModel struct {
	Views int `json:"views"`
	Likes int `json:"likes"`
}

type NewsDataModel struct {
	Id        string         `json:"id, omitempty" validate:"required"`
	Title     string         `json:"title, omitempty" validate:"required"`
	Image     ImageModel     `json:"image, omitempty"`
	Content   string         `json:"content, omitempty"`
	CreatedAt string         `json:"created_at"`
	UpdatedAt string         `json:"updated_at"`
	User      UserModel      `json:"author, omitempty" validate:"required"`
	Stats     NewsStatsModel `json:"stats"`
	IsDraft   bool           `json:"is_boolean" bson:"draft"`
	Tags      string         `json:"tags, omitempty"`
}
