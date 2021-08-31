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

func ValidateEmail(email string) (bool, error) {
	var user models.Users
	query := config.DB.Where("email = ?", email).Find(&user)
	if query.Error != nil {
		return false, query.Error
	} else if query.RowsAffected != 0 {
		return false, nil
	}
	return true, nil
}

func ValidatePhone(phone string) (bool, error) {
	var user models.Users
	query := config.DB.Where("phone = ?", phone).Find(&user)
	if query.Error != nil {
		return false, query.Error
	} else if query.RowsAffected != 0 {
		return false, nil
	}
	return true, nil
}

func ValidateInput(user *models.Users) (bool, string) {
	switch {
	case user.Name == "":
		return false, "name"
	case user.Email == "":
		return false, "email"
	case user.Phone == "":
		return false, "phone number"
	case user.Password == "":
		return false, "password"
	case user.Address == "":
		return false, "address"
	}

	return true, ""
}

func RegisterUser(user *models.Users) (interface{}, error) {
	query := config.DB.Create(user)
	if query.Error != nil {
		return nil, query.Error
	}
	voucher, err := ClaimUserVoucher(int(user.User_id), 1)
	if !voucher {
		return nil, err
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
