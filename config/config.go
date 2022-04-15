package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func connectDB() *gorm.DB {

	db, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/contoh?charset=utf8mb4&parseTime=True&loc=Local"))

	if err != nil {
		panic(err)
	}
	return db
}
