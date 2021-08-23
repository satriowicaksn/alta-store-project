package models

type Users struct {
	User_id  uint   `gorm:"primaryKey"`
	Name     string `gorm:"size:100" json: "name" form: "name"`
	Email    string `json: "email" form:"email"`
	Password string `json: "password" form:"password"`
	Token    string `json: "token" form:"token"`
}
