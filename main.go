package main

import (
	"fmt"
	"intro-golang/datastore/user"
	"intro-golang/entities"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func connectDB() *gorm.DB {

	dsn := "root:@tcp(127.0.0.1:3306)/book?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println(err)
	}
	return db
}

func main() {
	// Mymenu := []string{"1. Login", "2. Register", "3. Lis Buku", "4. Keluar"}
	// fmt.Println("Menu Utama")
	// for _, v := range Mymenu {
	// 	fmt.Println(v)
	// }
	// var nomor int
	// fmt.Print("Masukkan Angka :")
	// fmt.Scanf("%d", &nomor)
	// fmt.Println(menu2(nomor))
	conn := connectDB()
	// user := entities.Users(conn)
	conn.AutoMigrate(&entities.Users{})
	user := user.UserDB{Db: conn}
	var nama, email, password, hp string
	fmt.Scanf("%s", &nama)
	fmt.Scanf("%s", &email)
	fmt.Scanf("%s", &password)
	fmt.Scanf("%s", &hp)
	adduser, err := user.Register(entities.Users{Nama: nama, Email: email, Password: password, Hp: hp})
	// Create
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(adduser)
}
