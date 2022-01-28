package config

import (
	"Final-ProjectBDS-Sanbercode-Golang-Batch-31/models"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDataBase() *gorm.DB {
	username := "root"
	password := ""
	host := "tcp(127.0.0.1:3306)"
	database := "ecom"

	dsn := fmt.Sprintf("%v:%v@%v/%v?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, database)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	db.AutoMigrate(&models.User{}, &models.Order{}, &models.Category{}, &models.Product{}, &models.OrderDetail{}, &models.Cart{}, &models.Confrimation{})

	return db
}