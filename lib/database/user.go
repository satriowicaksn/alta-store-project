package database

import (
	"alta-store-project/config"
	"alta-store-project/middlewares"
	"alta-store-project/models"
)

func GetDetailUsers(userId int) (interface{}, error) {
	var user models.Users
	query := config.DB.Find(&user, userId)
	if query.Error != nil {
		return nil, query.Error
	} else if query.RowsAffected == 0 {
		return false, nil
	}
	return user, nil
}

func LoginUsers(user *models.Users) (interface{}, error) {
	var err error
	if err = config.DB.Where("email = ? AND password = ?", user.Email, user.Password).First(user).Error; err != nil {
		return nil, err
	}
	user.Token, err = middlewares.CreateToken(int(user.User_id))
	if err != nil {
		return nil, err
	}
	if err := config.DB.Save(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
