package models

type Users struct {
	User_id  uint   `json:"user_id" gorm:"primaryKey"`
	Name     string `gorm:"size:100" json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Phone    string `json:"phone"`
	Password string `json:"password" form:"password"`
	Address  string `json:"address" form:"address"`
	Token    string `json:"token" form:"token"`
}
