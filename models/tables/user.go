package models

type User struct {
	Id            int    `gorm:"primaryKey;column:id_user"`
	Username      string `gorm:"column:username"`
	Password      string `gorm:"column:password"`
	Fullname      string `gorm:"column:fullname"`
	Email         string `gorm:"column:email"`
	Province      string `gorm:"column:provinsi"`
	City          string `gorm:"column:kota"`
	Address       string `gorm:"column:alamat"`
	JoinDate      string `gorm:"column:tgl_gabung"`
	LastLoginTime string `gorm:"column:last_login"`
	Status        string `gorm:"column:status"`
	Level         string `gorm:"column:level"`
	Gender        string `gorm:"column:gender"`
	Moto          string `gorm:"column:moto"`
	IsVerified    string `gorm:"column:is_verified"`
	UpdatedAt     string `gorm:"column:updated_at"`
	Avatar        string `gorm:"column:avatar"`
	UserKey       string `gorm:"column:user_key"`
	Phone         string `gorm:"column:phone"`
}
