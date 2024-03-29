package database

import (
	"github.com/sabitalizade/go-resttAPI/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	connection, err := gorm.Open(mysql.Open("root:root@tcp(localhost:3306)/dmgo?charset=utf8&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	DB = connection
	
	connection.AutoMigrate(&models.User{})
}
