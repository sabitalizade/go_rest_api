package database

import (
	"goapi/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	connection, err := gorm.Open(mysql.Open("root:72248123s@tcp(localhost:3306)/dmgo?charset=utf8&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	DB = connection
	
	connection.AutoMigrate(&models.User{})
}