package models

type User struct {
	ID       uint `json:"id"`
	Email    string `gorm:"unique"`
	Password string `json:"-"`
	FullName string `json:"fullname"`
}
