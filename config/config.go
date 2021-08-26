package config

import (
	"alta-store-project/models"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	config := map[string]string{
		"DB_Username": "satrio",
		"DB_Password": "satrio12345",
		"DB_Port":     "3306",
		"DB_Host":     "127.0.0.1",
		"DB_Name":     "alta_store",
	}

	connectionString :=
		fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
			config["DB_Username"],
			config["DB_Password"],
			config["DB_Host"],
			config["DB_Port"],
			config["DB_Name"])

	var err error
	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}

}

func InitMigrate() {
	DB.AutoMigrate(&models.Product_categories{})
	DB.AutoMigrate(&models.Product{})
	DB.AutoMigrate(&models.Cart{})
	DB.AutoMigrate(&models.Cart_item{})
	DB.AutoMigrate(&models.Payment{})
	DB.AutoMigrate(&models.Payment_item{})
	DB.AutoMigrate(&models.Users{})
	DB.AutoMigrate(&models.Payment_method{})
}
